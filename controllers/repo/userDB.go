package repo

import (
	"context"
	"errors"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.EnvMongoCollection1())

func CreateUserDB(userDb models.User, a context.Context) (*mongo.InsertOneResult, error) {
	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Email:    userDb.Email,
		Password: userDb.Password,
		Name:     userDb.Name,
		Age:      userDb.Age,
		Orgs:     userDb.Orgs,
		About:    userDb.About,
		Gender:   userDb.Gender,
	}

	result, err := userCollection.InsertOne(a, newUser)
	if err != nil {
		return nil, errors.New("email already exists")
	} else {
		return result, nil
	}
}
