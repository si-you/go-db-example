package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var (
	mongoAddr = flag.String("mongo_addr", "localhost:27017", "Redis address.")
)

func main() {
	flag.Parse()

	c, err := mongo.NewClient(fmt.Sprintf("mongodb://%s", *mongoAddr))
	if err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}

	collection := c.Database("users").Collection("info")
	res, err := collection.InsertOne(context.Background(), map[string]string{"name": "siyou"})
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}
	fmt.Printf("Insert response: %v", res)

	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatalf("Full fetch failed: %v", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		d := bson.NewDocument()
		err := cur.Decode(d)
		if err != nil {
			log.Fatalf("Decoding failed: %v", err)
		}
		fmt.Printf("document: %v", d)
	}

	if cur.Err() != nil {
		log.Fatalf("Cur has an error: %v", cur.Err())
	}
}
