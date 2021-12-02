package main

import (
	"embed"
	"github.com/cmp307/assetman/pkg/backup"
	"github.com/cmp307/assetman/pkg/storage/sqlite"
	_ "github.com/cmp307/assetman/pkg/storage/sqlite"
	_ "github.com/cmp307/assetman/pkg/vulnerability"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	"path/filepath"
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

	path := filepath.FromSlash(app.ctx.Value("HomeDir").(string) + "/assets.db")
	db, err := sqlite.Connect(path)

	ar := sqlite.NewAssetRepository(db)
	mr := sqlite.NewManufacturerRepository(db)

	bak := backup.NewService(app.ctx)

	// Create application with options
	opts := &options.App{
		Title:             "AssetMan",
		Width:             1250,
		Height:            488,
		MinWidth:          720,
		MinHeight:         570,
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
			ar,
			mr,
			bak,
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

	// Initialise vulnerability service and configure Fx container
	//vs := vulnerability.NewService(os.Getenv("NVD_API_URL"), os.Getenv("NVD_API_KEY"))
}
