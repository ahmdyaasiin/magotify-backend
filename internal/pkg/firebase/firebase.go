package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
	"log"
)

var firebaseClient *db.Client
var authClient *auth.Client
var ctx context.Context

func init() {

	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://try-firebase-889e3-default-rtdb.asia-southeast1.firebasedatabase.app",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("./key.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	firebaseClient, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
}

func CreateCustomToken(nim string) (string, error) {
	token, err := authClient.CustomToken(ctx, nim)
	if err != nil {

		return token, err
	}

	return token, nil
}

func DecodeCustomToken(idToken string) (string, error) {
	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}
