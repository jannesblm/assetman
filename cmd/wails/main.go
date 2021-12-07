package main

import (
	"embed"
	"errors"
	"github.com/cmp307/assetman/pkg/auth"
	"github.com/cmp307/assetman/pkg/fs"
	"github.com/cmp307/assetman/pkg/storage"
	"github.com/cmp307/assetman/pkg/storage/sqlite"
	"github.com/cmp307/assetman/pkg/vulnerability"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
	"log"
	"os"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	godotenv.Load()

	// Create an instance of the app structure
	app := NewApp()
	io := fs.NewService(app.ctx)

	createUsers := false

	if _, err := os.Stat(io.GetDatabasePath()); errors.Is(err, os.ErrNotExist) {
		createUsers = true
	}

	// Initialise dummy db while we initialise the proper connection
	// in another thread
	gd, _ := gorm.Open(tests.DummyDialector{}, nil)
	db := &sqlite.DB{DB: gd}

	ar := sqlite.NewAssetRepository(db)
	mr := sqlite.NewManufacturerRepository(db)
	ur := sqlite.NewUserRepository(db)
	rr := sqlite.NewReportRepository(db)

	auth := auth.NewService(ur)

	// Initialise pre-execution callbacks manually to prevent calls to an uninitialised db
	db.InitialiseCallbacks()

	go func(db2 *sqlite.DB) {
		db, err := sqlite.Connect(io.GetDatabasePath(), auth)

		if err != nil {
			panic(err)
		}

		*db2 = *db
		err = db2.MigrateAll()

		if err != nil {
			panic(err)
		}

		if createUsers {
			ur.Save(storage.User{
				Name:     "admin",
				Group:    "admin",
				Password: []byte("admin"),
			})

			ur.Save(storage.User{
				Name:     "reporter",
				Group:    "reporter",
				Password: []byte("reporter"),
			})
		}

		db.Initialized = true
	}(db)

	vs := vulnerability.NewService(os.Getenv("NVD_API_URL"), os.Getenv("NVD_API_KEY"))
	vs.WatchResults()

	// Create application with options
	opts := &options.App{
		Title:             "AssetMan",
		Width:             1250,
		Height:            700,
		MinWidth:          1250,
		MinHeight:         700,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app,
			auth,
			io,
			ar,
			mr,
			rr,
			vs,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	}

	err := wails.Run(opts)

	if err != nil {
		log.Fatal(err)
	}
}
