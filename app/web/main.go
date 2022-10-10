package web

import (
	"context"
	"github.com/baguss42/go-clean-arch/app"
	"github.com/baguss42/go-clean-arch/controller"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"github.com/baguss42/go-clean-arch/infrastructure/environment"
	"github.com/baguss42/go-clean-arch/infrastructure/router"
	"github.com/baguss42/go-clean-arch/infrastructure/router/middleware"
	_interface "github.com/baguss42/go-clean-arch/interface"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Boot() {
	env, err := environment.Load("env", "./.env")
	if err != nil {
		log.Fatal("could not load environment ", err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	db := SetupDatabase(env)
	ctrl := controller.NewController(ctx, db)
	r := SetupRouter(ctrl)

	engine := app.NewEngine(r, db, env, ctx, cancel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	log.Println("starting server ...")
	go func() {
		if err = engine.Start(); err != nil {
			log.Fatal("could not serve http", err)
		}
	}()

	<-c
	engine.Shutdown()
}

// SetupDatabase setup for application database
func SetupDatabase(env *environment.Environment) *database.Database {
	return database.NewPostgres(database.Config{
		Host:        env.DBHost,
		Port:        env.DBPort,
		Username:    env.DBUsername,
		Password:    env.DBPassword,
		DBName:      env.DBName,
		SSLMode:     env.DBSslMode,
		MaxOpenConn: env.DBMaxOpenConn,
		MaxLifeTime: env.DBMaxLifeTime,
		MaxIdleTime: env.DBMaxIdleTime,
		MaxIdleConn: env.DBMaxIdleConn,
	})
}

// SetupRouter setup routing for application
func SetupRouter(ctrl *controller.Controller) _interface.Router {
	middlewares := []middleware.Filter{
		middleware.Trace(),
		middleware.Log(),
	}

	r := router.NewMuxRouter()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.POST("/create", middleware.Apply(middleware.WithFilters(ctrl.Product.Create, middlewares...)))
	r.GET("/read", middleware.Apply(middleware.WithFilters(ctrl.Product.Read, middlewares...)))
	r.POST("/update", middleware.Apply(middleware.WithFilters(ctrl.Product.Update, middlewares...)))
	r.POST("/delete", middleware.Apply(middleware.WithFilters(ctrl.Product.Delete, middlewares...)))

	return r
}
