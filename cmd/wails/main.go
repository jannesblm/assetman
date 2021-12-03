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
	"log"
	"os"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	// Create an instance of the app structure
	app := NewApp()
	io := fs.NewService(app.ctx)

	createUsers := false

	if _, err := os.Stat(io.GetDatabasePath()); errors.Is(err, os.ErrNotExist) {
		createUsers = true
	}

	db, err := sqlite.Connect(io.GetDatabasePath())
	err = db.MigrateAll()

	if err != nil {
		panic(err)
	}

	ar := sqlite.NewAssetRepository(db)
	mr := sqlite.NewManufacturerRepository(db)
	ur := sqlite.NewUserRepository(db)
	rr := sqlite.NewReportRepository(db)

	if createUsers {
		ur.Save(storage.User{
			Name:     "admin",
			Password: []byte("admin"),
		})

		ur.Save(storage.User{
			Name:     "reporter",
			Password: []byte("reporter"),
		})
	}

	auth := auth.NewService(ur)
	vs := vulnerability.NewService(os.Getenv("NVD_API_URL"), os.Getenv("NVD_API_KEY"))

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

	err = wails.Run(opts)

	if err != nil {
		log.Fatal(err)
	}
}
