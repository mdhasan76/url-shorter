package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://go_netflix:lW28H9FzpWhDA2gJ@cluster0.mj4ed9j.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "url-shorter"
const collectionName = "url-encoded"

var collection *mongo.Collection

func Server() {
	clientConnection := options.Client().ApplyURI(connectionString)

	// connection to mongodb
	client, err := mongo.Connect(clientConnection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb Connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// Collection instance
	fmt.Println("Collection instance is ready")
}

func getAllDocuments() []primitive.M {
	result, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var urlDataLists []primitive.M

	for result.Next(context.Background()) {
		var urlData primitive.M
		err := result.Decode(&urlData)
		if err != nil {
			log.Fatal(err)
		}
		urlDataLists = append(urlDataLists, urlData)
	}

	defer result.Close(context.Background())
	return urlDataLists
}
func GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := getAllDocuments()
	json.NewEncoder(w).Encode(result)
}
