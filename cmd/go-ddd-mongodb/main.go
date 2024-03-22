package main

import (
	"context"
	"fmt"
	"github.com/tutorial_ddd_go-mongodb/pkg/domain"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tutorial_ddd_go-mongodb/pkg/application"
	"github.com/tutorial_ddd_go-mongodb/pkg/infrastructure"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://root:toor@localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("users")
	repository := &infrastructure.MongoUserRepository{Collection: collection}
	service := &application.UserService{Repository: repository}

	// Create a user with the name "John Doe"
	user := domain.User{Name: "John Doe"}
	id, err := service.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created user with ID:", id)

	// Add a delay
	time.Sleep(2 * time.Second)

	userget, err := service.GetUser(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User:", userget.Name)
}
