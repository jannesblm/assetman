package main

import (
	"assetman/pkg/storage"
	"assetman/pkg/storage/sqlite"
	"assetman/pkg/vulnerability"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	repo := sqlite.NewRepository()
	vuln := vulnerability.NewService(os.Getenv("NVD_API_URL"), os.Getenv("NVD_API_KEY"))

	app := fx.New(
		fx.Provide(
			func() storage.AssetRepository { return repo },
			func() vulnerability.Service { return vuln },
		),

		fx.Invoke(InitGui),
	)

	app.Run()
}
