package repo

import (
	"context"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dataset interface {
	CreateUserDB(userDb models.User, a context.Context) (*mongo.InsertOneResult, error)
	SignInDB(userLogin models.Login, a context.Context) (*models.User, error)
	UserEdit(a context.Context, str string, bsondata primitive.M) (*mongo.UpdateResult, error)
	UserPorfileDB(a context.Context, user models.User, id string) error
}
