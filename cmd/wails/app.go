package main

import (
	"context"
	"fmt"
	"os"
)

// App application struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	home, err := os.UserHomeDir()

	if err != nil {
		panic("cannot get home dir")
	}

	home = home + "/ScottishGlen"

	err = os.MkdirAll(home, os.ModePerm)

	if err != nil {
		panic("cannot setup app dir")
	}

	ctx := context.WithValue(context.TODO(), "HomeDir", home)

	return &App{
		ctx: ctx,
	}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	//
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
