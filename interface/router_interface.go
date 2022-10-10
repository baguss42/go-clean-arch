package _interface

import "net/http"

type Router interface {
	GET(path string, fn func(w http.ResponseWriter, r *http.Request))
	POST(path string, fn func(w http.ResponseWriter, r *http.Request))
	Serve(port string) error
}
