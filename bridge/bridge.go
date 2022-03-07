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

const DatasourceName = "htickets-cb4d0"
const CollectionName = "ProfileVisitors"
const required_File = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"

func Firestore_Object() *firestore.Client {

	_, err := os.Stat("credentials/" + required_File)
	if os.IsExist(err) {
		fmt.Println("Error to read credentials file", err.Error())
		return &firestore.Client{}
	}

	firestore_credentials := "credentials/" + required_File
	firestore_connect, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: DatasourceName}, option.WithCredentialsFile(firestore_credentials))
	if err != nil {
		fmt.Println("Error creating firestore credentials", err.Error())
		return &firestore.Client{}
	}

	client_obj, err := firestore_connect.Firestore(context.Background())
	if err != nil {
		fmt.Println("Error creating firestore client", err.Error())
		return &firestore.Client{}
	}

	return client_obj
}

func GetUserLogin(email, password string, client *firestore.Client) map[string]interface{} {

	var profile map[string]interface{} = nil
	query := client.Collection(CollectionName).Where("email", "==", email).Documents(context.Background())
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
