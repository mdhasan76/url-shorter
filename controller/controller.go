package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://go_netflix:lW28H9FzpWhDA2gJ@cluster0.mj4ed9j.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func ServerMainFn() {
	clientOption := options.Client().ApplyURI(connectionString)

	// connection to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		// log.Fatal(err)
		log.Println("Mongodb throw error", err)
	}

	fmt.Println("Mongodb Connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// Collection instance
	fmt.Println("Collection instance is ready")
}

// func getAllDocuments() []primitive.M {
// 	result, err := collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var urlDataLists []primitive.M

// 	for result.Next(context.Background()) {
// 		var urlData primitive.M
// 		err := result.Decode(&urlData)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		urlDataLists = append(urlDataLists, urlData)
// 	}

//		defer result.Close(context.Background())
//		return urlDataLists
//	}
// func getAllDocuments() []bson.M {
// 	cursor, err := collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		log.Fatal("Find error:", err)
// 	}
// 	defer cursor.Close(context.Background()) // Close cursor properly

// 	var urlDataLists []bson.M

// 	for cursor.Next(context.Background()) {
// 		var urlData bson.M
// 		if err := cursor.Decode(&urlData); err != nil {
// 			log.Println("Error decoding document:", err)
// 			continue
// 		}
// 		urlDataLists = append(urlDataLists, urlData)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		log.Println("Cursor iteration error:", err)
// 	}

// 	return urlDataLists
// }

func getAllDocuments() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Got error on find document", err)
	}
	var allDocuments []primitive.M
	for cursor.Next(context.Background()) {
		var urlData bson.M
		err := cursor.Decode(&urlData)
		if err != nil {
			log.Fatal(err)
		}
		allDocuments = append(allDocuments, urlData)
	}
	defer cursor.Close(context.Background())
	return allDocuments
}

// func getAllMovies() []primitive.M {
// 	cur, err := collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var movies []primitive.M
// 	for cur.Next(context.Background()) {
// 		var movie bson.M
// 		err := cur.Decode(&movie)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		movies = append(movies, movie)
// 	}
// 	defer cur.Close(context.Background())
// 	return movies
// }

func GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := getAllDocuments()
	// fmt.Println(result, "here is the result")
	// jsonData := map[string]string{"name": "Hasan Mia", "email": "mdhasan76@gmail.com"}
	json.NewEncoder(w).Encode(result)
}
