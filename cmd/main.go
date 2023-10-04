package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Marsredskies/apod_service/cmd/apod_service/api"
	"github.com/Marsredskies/apod_service/cmd/apod_service/database"
	"github.com/Marsredskies/apod_service/cmd/apod_service/nasa"
	"github.com/Marsredskies/apod_service/envconfig"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	env := envconfig.MustGetConfig()
	dbClient := database.RequireNewDBClient(ctx, env.DB)

	database.MustApplyMigrations(ctx, env.DB)

	nasa := nasa.MustInitClient(env.APOD, dbClient)
	go nasa.DoJobFetchAndSaveImages()

	api := api.MustInitNewAPI(env.DB)
	api.StartServer(env.APOD.ApiPort)

	exit := make(chan (os.Signal), 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	err := api.Shutdown(ctxShutDown)
	if err != nil {
		log.Printf("failed to shutdown the server: %v", err)
	}

	log.Println("shutting down")
}
