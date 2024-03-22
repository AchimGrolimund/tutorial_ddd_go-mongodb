package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/tutorial_ddd_go-mongodb/pkg/domain"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func (r *MongoUserRepository) GetUser(id string) (domain.User, error) {
	var user domain.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *MongoUserRepository) CreateUser(user domain.User) (string, error) {
	result, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
