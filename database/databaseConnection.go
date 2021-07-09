package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	clientOption := options.Client().ApplyURI("mongodb+srv://karan:maher7505@Spring.crnpa.mongodb.net/movies?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	collection := client.Database("movies").Collection("movie")
	return collection

}
