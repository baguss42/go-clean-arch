package main

import (
	"context"
	"github.com/baguss42/go-clean-arch/app"
	"github.com/baguss42/go-clean-arch/app/grpc"
	"github.com/baguss42/go-clean-arch/app/web"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"github.com/baguss42/go-clean-arch/infrastructure/environment"
	"github.com/baguss42/go-clean-arch/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env, err := environment.Load("env", ".env")
	if err != nil {
		log.Fatal("could not load environment ", err)
	}

	db := SetupDatabase(env)
	svc := service.NewService(db)
	ctx, cancel := context.WithCancel(context.Background())

	engine := app.NewEngine(ctx, db, svc, env)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		web.Boot(engine)
	}()

	go func() {
		grpc.Boot(engine)
	}()

	<-c
	log.Println("shutdown ...")
	engine.Shutdown(cancel)
}

// SetupDatabase setup for application database
func SetupDatabase(env *environment.Environment) *database.Database {
	return database.NewPostgresql(database.Config{
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
