package main

import (
	"github.com/skarakasoglu/g-case-challenge/api"
	"github.com/skarakasoglu/g-case-challenge/config"
	"github.com/skarakasoglu/g-case-challenge/inmem"
	"github.com/skarakasoglu/g-case-challenge/mongodb"
	"github.com/skarakasoglu/g-case-challenge/record"
	rediscl "github.com/skarakasoglu/g-case-challenge/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfig := config.ReadFromEnvironmentVariables()
	dbConfig := appConfig.Database

	conn := mongodb.NewConnection(dbConfig.ConnectionString, dbConfig.DefaultDatabaseName)
	conn.Connect()
	defer conn.Disconnect()

	recordDao := record.Dao{Db: conn}
	recordService := record.Service{Dao: recordDao}
	recordController := record.Controller{Repository: recordService}

	redisCl := rediscl.NewClient(appConfig.RedisConnectionString)

	inMemoryDao := inmem.Dao{Db: redisCl}
	inMemoryService := inmem.Service{Dao: inMemoryDao}
	inMemoryController := inmem.Controller{ Repository: inMemoryService}

	endpoints := []api.Endpoint{
		{ Path: "/records", Handler: recordController},
		{ Path: "/in-memory", Handler: inMemoryController},
	}
	go func() {
		err := api.Start(appConfig.Api.Address, endpoints...)
		if err != nil {
			log.Fatalf("error on serving HTTP: %v", err)
		}
	}()

	log.Println("Getir Case Challenge API is now running. Press CTRL + C to interrupt.")

	signalHandler := make(chan os.Signal)
	signal.Notify(signalHandler, os.Interrupt, syscall.SIGUSR1)
	receivedSignal := <-signalHandler

	log.Printf("API received %v signal. Gracefully shutting down the application.", receivedSignal)
}
