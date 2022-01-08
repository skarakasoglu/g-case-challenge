package main

import (
	"flag"
	"getir-assignment/api"
	"getir-assignment/inmem"
	"getir-assignment/mongodb"
	"getir-assignment/record"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	dbUsername string
	dbPassword string
	dbHost string
	dbName string
)

func init() {
	flag.StringVar(&dbUsername, "db-username", "", "database login username")
	flag.StringVar(&dbPassword, "db-password", "", "database login password")
	flag.StringVar(&dbHost, "db-host", "", "database host address")
	flag.StringVar(&dbName, "db-name", "", "default database name")
	flag.Parse()
}

func main() {
	conn := mongodb.NewConnection(dbUsername, dbPassword, dbHost, dbName)
	conn.Connect()
	defer conn.Disconnect()

	recordDao := record.Dao{Db: conn}
	recordService := record.Service{Dao: recordDao}
	recordController := record.Controller{Repository: recordService}

	inMemoryController := inmem.Controller{}

	endpoints := []api.Endpoint{
		{ Path: "/records", Handler: recordController},
		{ Path: "/in-memory", Handler: inMemoryController},
	}
	go func() {
		err := api.Start(":8080", endpoints...)
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
