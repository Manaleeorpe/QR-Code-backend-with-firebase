package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var FirebaseDB *db.Client

func ConnectFirebase() {
	// Initialize Firebase app with service account
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}
	FirebaseApp = app

	// Initialize Firebase Realtime Database client (change URL to your project DB URL)
	databaseURL := "https://qr-code-2bd00-default-rtdb.firebaseio.com/"
	client, err := app.DatabaseWithURL(ctx, databaseURL)
	if err != nil {
		log.Fatalf("error initializing firebase database: %v", err)
	}
	FirebaseDB = client

	log.Println("Connected to Firebase!")
}
