package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Message struct {
	Name       string
	Id         int       `bson:"_id"`
	IsDisabled bool      `bson:"is_disabled"`
	Date       time.Time `bson:"date,omitempty"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file in ")
	}
	MongoURILoc := os.Getenv("MONGODB_URI_LOC")
	if MongoURILoc == "" {
		log.Fatal("You must set you 'MongoURI' environment variable ")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoURILoc))
	if err != nil {
		log.Println(err)
	}
	coll := client.Database("test").Collection("messages")
	newMessage := Message{Id: 0, Name: "hello", IsDisabled: false}
	// Insert a Data
	_, err = coll.InsertOne(ctx, newMessage)
	log.Printf("%+v\n", err)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No Documents: ", err)
			return
		}
	}

	filter := bson.D{{"name", "hello"}}
	var result Message
	coll.FindOne(ctx, filter).Decode(&result)
	fmt.Printf("%+v", result)
	// TODO InsertOne InsertMany

}
