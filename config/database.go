package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {

	datauri, ok := os.LookupEnv("DB_URI")

	if !ok {
		log.Fatalln("can'nt find the database uri")

	}

	mongoDbOpion := options.Client().ApplyURI(datauri)
	mongoClient, err := mongo.Connect(context.Background(), mongoDbOpion)
	if err != nil {
		log.Fatalln("error while connecting the mongodb database")
	}
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("error While connecting MONGO DATABASE..! %v ->", err)
	} else {
		log.Println("Connected to MONGO DATABASE ")
	}
	databasename, ok := os.LookupEnv("DATABASE_NAME")

	if !ok {
		log.Fatalln("can'nt find the database name")

	}

	DB = mongoClient.Database(databasename)

}
