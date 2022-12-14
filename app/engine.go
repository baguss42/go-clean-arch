package app

import (
	"context"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"github.com/baguss42/go-clean-arch/infrastructure/environment"
	"github.com/baguss42/go-clean-arch/service"
	"log"
)

type Engine struct {
	Context     context.Context
	Database    *database.Database
	Environment *environment.Environment
	Service     *service.Service
}

func NewEngine(
	ctx context.Context,
	db *database.Database,
	svc *service.Service,
	env *environment.Environment,
) *Engine {
	return &Engine{
		Context:     ctx,
		Database:    db,
		Service:     svc,
		Environment: env,
	}
}

func (e *Engine) Shutdown(cancel context.CancelFunc) {
	log.Println("shutdown ...")
	cancel()
	_ = e.Database.Close()
}
