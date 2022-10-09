package middleware

import (
	"log"
	"net/http"
	"time"
)

type HandleWithError func(writer http.ResponseWriter, request *http.Request) (int, error)

type Filter func(HandleWithError) HandleWithError

func WithFilters(handle HandleWithError, middlewares ...Filter) HandleWithError {
	for _, filter := range middlewares {
		handle = filter(handle)
	}
	return handle
}

// Apply should be applied with WithFilters
func Apply(handle HandleWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = handle(w, r)
	}
}

func Trace() Filter {
	return func(handle HandleWithError) HandleWithError {
		return func(writer http.ResponseWriter, request *http.Request) (int, error) {
			start := time.Now()
			code, err := handle(writer, request)
			end := time.Since(start)

			log.Println("time: ", end.String())
			return code, err
		}
	}
}

func Log() Filter {
	return func(handle HandleWithError) HandleWithError {
		return func(writer http.ResponseWriter, request *http.Request) (int, error) {
			i, err := handle(writer, request)
			log.Printf("code: %v, error: %v", i, err)
			return i, err
		}
	}
}
