package services

import (
	"context"
	"errors"
	"fiber-mongo-api/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *services) CreateUserDB(userDb models.User, a context.Context) (*mongo.InsertOneResult, error) {
	newUser := models.User{
		User_Id:  primitive.NewObjectID(),
		Email:    userDb.Email,
		Password: userDb.Password,
		Name:     userDb.Name,
		Age:      userDb.Age,
		Orgs:     userDb.Orgs,
		About:    userDb.About,
		Gender:   userDb.Gender,
	}
	var user *models.User
	err := c.monggose.UserCollection().FindOne(a, bson.M{"email": userDb.Email}).Decode(&user.Email)
	str := fmt.Sprintf("Email (%v) Already Registered", user.Email)
	if err == nil {
		return nil, errors.New(str)
	}
	result, err := c.monggose.UserCollection().InsertOne(a, newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *services) SignInDB(userLogin models.Login, a context.Context) (*models.User, error) {
	var user *models.User

	err := c.monggose.UserCollection().FindOne(a, bson.M{
		"email": userLogin.Email}).Decode(&user)

	return user, err
}

func (c *services) UserPorfileDB(a context.Context, id string) (*models.User, error) {
	var user *models.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := c.monggose.UserCollection().FindOne(a, bson.M{"user_id": objId}).Decode(&user)
	return user, err
}

func (c *services) UserEdit(a context.Context, str string, bsondata primitive.M) (*mongo.UpdateResult, error) {
	objId, _ := primitive.ObjectIDFromHex(str)
	filter := bson.D{{Key: "user_id", Value: objId}}
	update := bson.D{{Key: "$set", Value: bsondata}}

	result, err := c.monggose.UserCollection().UpdateOne(a, filter, update)

	return result, err
}

func (c *services) DeleteAccount(a context.Context, id string) (*mongo.DeleteResult, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "user_id", Value: objId}}

	result, err := c.monggose.UserCollection().DeleteOne(a, filter)
	return result, err
}
