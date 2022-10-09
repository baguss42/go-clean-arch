package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouter struct {
	router *mux.Router
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{
		router: mux.NewRouter(),
	}
}

func (r *MuxRouter) GET(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(path, fn).Methods(http.MethodGet)
}

func (r *MuxRouter) POST(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(path, fn).Methods(http.MethodPost)
}

func (r *MuxRouter) Serve(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r.router)
}
