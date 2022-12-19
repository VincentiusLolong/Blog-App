package repo

import (
	"context"
	"errors"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fmt"

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
	var user *models.User
	err := userCollection.FindOne(a, bson.M{"email": userDb.Email}).Decode(&user)
	str := fmt.Sprintf("Email (%v) Already Registered", user.Email)
	if err == nil {
		return nil, errors.New(str)
	}
	result, err := userCollection.InsertOne(a, newUser)
	if err != nil {
		return nil, errors.New("err: 505")
	}
	return result, nil
}

func SignInDB(userLogin models.Login, a context.Context) (*models.User, error) {
	var user *models.User

	err := userCollection.FindOne(a, bson.M{
		"email": userLogin.Email}).Decode(&user)

	if err != nil {
		return nil, errors.New("cant find the account")
	} else {
		return user, nil
	}
}

func UserPorfileDB(a context.Context, user models.User, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(a, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return errors.New("cant find the account")
	} else {
		return nil
	}
}

func UserEdit(a context.Context, str string, bsondata primitive.M) (*mongo.UpdateResult, error) {
	objId, _ := primitive.ObjectIDFromHex(str)
	filter := bson.D{{Key: "id", Value: objId}}
	update := bson.D{{Key: "$set", Value: bsondata}}

	result, err := userCollection.UpdateOne(a, filter, update)
	if err != nil {
		return nil, errors.New("cant edit")
	} else {
		return result, nil
	}
}
