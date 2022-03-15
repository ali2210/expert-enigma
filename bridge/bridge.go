package bridge

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// database credentials

func Firestore_Object() *firestore.Client {

	// test whether file is accessible in a directory
	_, err := os.Stat("credentials/" + "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json")
	if os.IsExist(err) {
		return &firestore.Client{}
	}

	// get database credentials path
	firestore_credentials := "credentials/" + "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"

	// init database conections
	firestore_connect, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: "htickets-cb4d0"}, option.WithCredentialsFile(firestore_credentials))
	if err != nil {
		return &firestore.Client{}
	}

	// connect with databse server
	client_obj, err := firestore_connect.Firestore(context.Background())
	if err != nil {
		return &firestore.Client{}
	}

	return client_obj
}

func GetUserLogin(email, password string, client *firestore.Client) map[string]interface{} {

	// get user login credentials
	var profile map[string]interface{} = nil

	// if client is authenticated then return the user information
	query := client.Collection("ProfileVisitors").Where("email", "==", email).Documents(context.Background())
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		snaps, err := doc.Ref.Get(context.Background())
		if err != nil {
			fmt.Println("Error getting user credentials: ", err.Error())
			return map[string]interface{}{}
		}
		profile = snaps.Data()
	}

	return profile
}
