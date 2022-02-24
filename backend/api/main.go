package main

import (
	"backend/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"

	"google.golang.org/api/option"
)

type config struct {
	port int
}

type AppStatus struct {
	Status     string `json:"status"`
	Enviroment string `json:"enviroment"`
	Version    string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {

	// load the .env file

	err := godotenv.Load("D:/PROJECTS/GigStateCMS/backend/api/.env")

	if err != nil {
		log.Fatal("cannot looad .env file ", err)
	}
	// assign the log file
	// create log file for current run

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var cfg config
	// pre-set cfg's port number to 4000
	cfg.port = 4000
	// redirect loggers stroud to the log file we created above
	logger := log.New(file, "", log.Ldate|log.Ltime)

	// connect to firebase
	opt := option.WithCredentialsFile(os.Getenv("FIREBASECLIENTCREDFILE"))
	fireapp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Fatalln(err)
	}

	// get firestore client
	client, err := fireapp.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)

		logger.Fatalln(err)

	}
	defer func() {
		client.Close()
	}()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(client),
	}

	// create a custom http server

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		// no time-outs... for now.
	}

	log.Printf("Starting server on port:%d\n", cfg.port)
	logger.Printf("Starting server on port:%d\n", cfg.port)
	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
		logger.Println(err)
	}
}
