package web

import (
	"github.com/baguss42/go-clean-arch/app"
	"github.com/baguss42/go-clean-arch/controller"
	"github.com/baguss42/go-clean-arch/infrastructure/router"
	"github.com/baguss42/go-clean-arch/infrastructure/router/middleware"
	_interface "github.com/baguss42/go-clean-arch/interface"
	"log"
	"net/http"
)

func Boot(engine *app.Engine) {
	r := SetupRouter(engine)

	log.Printf("HTTP starting at %s", engine.Environment.HTTPPort)
	if err := r.Serve(engine.Environment.HTTPPort); err != nil {
		log.Fatal("could not serve http", err)
	}
}

// SetupRouter setup routing for application
func SetupRouter(engine *app.Engine) _interface.Router {
	middlewares := []middleware.Filter{
		middleware.Trace(),
		middleware.Log(),
	}

	ctrl := controller.NewController(engine.Context, engine.Service)

	r := router.NewMuxRouter()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.GET("/product", middleware.Apply(middleware.WithFilters(ctrl.Product.All, middlewares...)))
	r.POST("/product/create", middleware.Apply(middleware.WithFilters(ctrl.Product.Create, middlewares...)))
	r.GET("/product/read", middleware.Apply(middleware.WithFilters(ctrl.Product.Read, middlewares...)))
	r.POST("/product/update", middleware.Apply(middleware.WithFilters(ctrl.Product.Update, middlewares...)))
	r.POST("/product/delete", middleware.Apply(middleware.WithFilters(ctrl.Product.Delete, middlewares...)))

	return r
}
