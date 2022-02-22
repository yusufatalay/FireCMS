package main

import (
	"backend/models"
	"context"
	"log"

	firebase "firebase.google.com/go"

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
	client models.Client
}

func main() {

	// connect to firebase
	opt := option.WithCredentialsFile("../gigstateprototip-404f9-firebase-adminsdk-b6y1w-6d35065726.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// get firestore client
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	client.Collection("Gig Sites").Doc().Create()
	defer func() {
		client.Close()
	}()

}
