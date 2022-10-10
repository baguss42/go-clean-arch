package app

import (
	"context"
	"github.com/baguss42/go-clean-arch/controller"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"github.com/baguss42/go-clean-arch/infrastructure/environment"
	_interface "github.com/baguss42/go-clean-arch/interface"
	"log"
)

type Engine struct {
	Router      _interface.Router
	Database    *database.Database
	Environment *environment.Environment
	Ctx         context.Context
	Controller  controller.Controller

	cancel context.CancelFunc
}

func NewEngine(router _interface.Router,
	db *database.Database,
	env *environment.Environment,
	ctx context.Context,
	cancel context.CancelFunc) *Engine {
	return &Engine{
		Router:      router,
		Database:    db,
		Environment: env,
		Ctx:         ctx,
		cancel:      cancel,
	}
}

func (e *Engine) Start() error {
	return e.Router.Serve(e.Environment.AppPort)
}

func (e *Engine) Shutdown() {
	log.Println("shutdown ...")
	e.cancel()
	_ = e.Database.Close()
}
