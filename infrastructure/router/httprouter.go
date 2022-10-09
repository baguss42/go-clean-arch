package router

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HttpRouter struct {
	router *httprouter.Router
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{
		router: httprouter.New(),
	}
}

func (r *HttpRouter) GET(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.router.GET(path, wrapHandler(fn))
}

func (r *HttpRouter) POST(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	r.router.POST(path, wrapHandler(fn))
}

func (r *HttpRouter) Serve(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r.router)
}

func wrapHandler(fn func(w http.ResponseWriter, r *http.Request)) httprouter.Handle {
	return func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(req.Context(), "params", ps)
		fn(rw, req.WithContext(ctx))
	}
}

func HttpGetParams(r *http.Request, key string) string {
	params, ok := r.Context().Value("params").(httprouter.Params)
	if !ok {
		return ""
	}

	for i := range params {
		if params[i].Key == key {
			return params[i].Value
		}
	}
	return ""
}
