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
		log.Fatalf("Mongo client creation failed: %v", err)
	}

	err = c.Connect(context.Background())
	if err != nil {
		log.Fatalf("Mongo connection failed: %v", err)
	}
	defer c.Disconnect(context.Background())

	// List all databases.
	ds, err := c.ListDatabaseNames(context.Background(), nil)
	if err != nil {
		log.Fatalf("Listing database failed: %v", err)
	}
	fmt.Printf("Available dbs: %v\n", ds)

	// Insert a document.
	collection := c.Database("users").Collection("info")
	res, err := collection.InsertOne(context.Background(), map[string]string{"name": "siyou"})
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}
	fmt.Printf("Insert response: %v\n", res)

	// Find a document.
	filter := bson.NewDocument(bson.EC.String("name", "siyou"))
	fmt.Printf("Filter: %v\n", filter)

	cur, err := collection.Find(context.Background(), filter)
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
		fmt.Printf("document: %v\n", d)
	}
	if cur.Err() != nil {
		log.Fatalf("Cur has an error: %v", cur.Err())
	}

	// Delete a document.
	dRes, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	fmt.Printf("Delete response: %v\n", *dRes)
}
