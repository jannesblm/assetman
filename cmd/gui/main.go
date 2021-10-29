package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	app := fx.New(
		fx.Invoke(InitGui),
	)

	app.Run()
}
