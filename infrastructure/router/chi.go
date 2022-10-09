package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiRouter struct {
	router *chi.Mux
}

func NewChiRouter() *ChiRouter {
	return &ChiRouter{
		router: chi.NewRouter(),
	}
}

func (r *ChiRouter) GET(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.router.Get(path, fn)
}

func (r *ChiRouter) POST(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.router.Post(path, fn)
}

func (r *ChiRouter) Serve(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r.router)
}
