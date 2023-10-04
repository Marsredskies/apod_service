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

//import "github.com/kelseyhightower/envconfig"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	env := envconfig.MustGetConfig()
	log.Println("env_parsed")
	dbClient := database.RequireNewDBClient(ctx, env.DB)
	log.Println("db_client_initialized")

	nasa := nasa.MustInitClient(env.APOD, dbClient)
	go nasa.DoJobFetchAndSaveImages()

	api := api.MustInitNewAPI(env.DB)
	api.StartServer(env.APOD.ApiPort)

	exit := make(chan (os.Signal), 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	err := api.Shutdown(ctx)
	if err != nil {
		log.Printf("failed to shutdown the server: %v", err)
	}

	log.Println("shutting down")
}
