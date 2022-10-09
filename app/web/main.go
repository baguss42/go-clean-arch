package web

import (
	"github.com/baguss42/go-clean-arch/app"
	"github.com/baguss42/go-clean-arch/controller"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"github.com/baguss42/go-clean-arch/infrastructure/router"
	"github.com/baguss42/go-clean-arch/infrastructure/router/middleware"
	"github.com/baguss42/go-clean-arch/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Boot() {
	engine := app.NewEngine(router.NewMuxRouter())
	log.Println("starting server ...")

	engine.Database = database.NewPostgres(database.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "root",
		DBName:   "go_clean",
		SSLMode:  "disable",
	})

	RegisterEndpoints(engine)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := engine.Router.Serve("8080"); err != nil {
			log.Fatal("could not serve http", err)
		}
	}()

	<-c
	engine.Shutdown()
}

func RegisterEndpoints(engine *app.Engine) {
	product := controller.Product{
		Service: service.NewProductService(engine.Database),
		Ctx:     engine.Ctx,
	}

	middlewares := []middleware.Filter{
		middleware.Trace(),
		middleware.Log(),
	}

	engine.Router.GET("/", Index)
	engine.Router.POST("/create", middleware.Apply(middleware.WithFilters(product.Create, middlewares...)))
	engine.Router.GET("/read", middleware.Apply(middleware.WithFilters(product.Read, middlewares...)))
	engine.Router.POST("/update", middleware.Apply(middleware.WithFilters(product.Update, middlewares...)))
	engine.Router.POST("/delete", middleware.Apply(middleware.WithFilters(product.Delete, middlewares...)))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
