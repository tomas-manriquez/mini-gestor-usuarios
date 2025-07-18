package main

import (
	"context"
	"fmt"
	"time"

	"go-template/models"
	"go-template/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Initialize MongoDB connection
	services.InitMongo()

	// Get collection
	coll := services.GetUserCollection()

	// Create sample users
	users := []interface{}{
		models.User{
			ID:        "", // Mongo will auto-generate
			Name:      "Alice Smith",
			Email:     "alice@example.com",
			BirthDate: time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Now(),
		},
		models.User{
			ID:        "",
			Name:      "Bob Johnson",
			Email:     "bob@example.com",
			BirthDate: time.Date(1985, 12, 2, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Now(),
		},
		models.User{
			ID:        "",
			Name:      "Charlie Brown",
			Email:     "charlie@example.com",
			BirthDate: time.Date(2000, 1, 20, 0, 0, 0, 0, time.UTC),
			CreatedAt: time.Now(),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert users
	res, err := coll.InsertMany(ctx, users)
	if err != nil {
		panic(err)
	}

	// Print inserted IDs
	fmt.Println("Inserted user IDs:")
	for _, id := range res.InsertedIDs {
		oid := id.(primitive.ObjectID)
		fmt.Println(oid.Hex())
	}
}
