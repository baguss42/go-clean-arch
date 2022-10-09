package app

import (
	"context"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"log"
	"net/http"
)

type Engine struct {
	Router   Router
	Database *database.Database
	Ctx      context.Context
}

func NewEngine(router Router) *Engine {
	return &Engine{
		Router: router,
		Ctx:    context.Background(),
	}
}

type Router interface {
	GET(path string, fn func(w http.ResponseWriter, r *http.Request))
	POST(path string, fn func(w http.ResponseWriter, r *http.Request))
	Serve(port string) error
}

func (e *Engine) Shutdown() {
	log.Println("shutdown ...")
	_ = e.Database.Close()
}
