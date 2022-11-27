package repo

import (
	"context"
	"errors"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("CONTENTCOLLECTION"))

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

func SignInDB(userId, email, password string, a context.Context) (*models.User, error) {
	var user *models.User
	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(a, bson.M{
		"id":       objId,
		"email":    email,
		"password": password}).Decode(&user)

	if err != nil {
		return nil, errors.New("cant find the account")
	} else {
		return user, nil
	}
}
