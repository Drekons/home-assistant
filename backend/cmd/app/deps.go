package app

import (
	"context"
	"github.com/Drekons/home-assistant/backend/config"
	"github.com/Drekons/home-assistant/backend/internal/database"
	"github.com/Drekons/home-assistant/backend/internal/interfaces"
	"github.com/Drekons/home-assistant/backend/internal/repository"
	"log"
)

type Deps struct {
	db   *database.DB
	cfg  *config.Config
	repo interfaces.Repository
}

func NewDeps(ctx context.Context) *Deps {
	deps := &Deps{}
	var err error

	// Load configuration
	deps.cfg, err = config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize DB connection
	deps.db, err = database.NewDB(ctx, deps.cfg.MongoDB.URI, deps.cfg.MongoDB.Database)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	deps.repo, err = repository.NewRepository(ctx, deps.db)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	return deps
}

func (d *Deps) DB() *database.DB {
	return d.db
}

func (d *Deps) Config() *config.Config {
	return d.cfg
}

func (d *Deps) Repo() interfaces.Repository {
	return d.repo
}
