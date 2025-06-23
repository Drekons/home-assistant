package app

import (
	"context"
	"github.com/Drekons/home-assistant/backend/internal/database"
	"log"
	"os"
	"os/signal"
)

type App struct {
	Deps *Deps
}

func NewApp(ctx context.Context) *App {
	app := &App{
		Deps: NewDeps(ctx),
	}

	app.gracefulShutdown(ctx)

	return app
}

func (a *App) gracefulShutdown(ctx context.Context) {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	go func() {
		<-gracefulShutdown
		a.Shutdown(ctx)
		log.Println("Server is shutting down...")
	}()
}

func (a *App) Shutdown(ctx context.Context) {
	defer func(db *database.DB, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			log.Fatalf("Failed to close DB connection: %v", err)
		}
	}(a.Deps.DB(), ctx)
}
