package repo

import (
	"context"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoDataset interface {
	//user
	CreateUserDB(userDb models.User, a context.Context) (*mongo.InsertOneResult, error)
	SignInDB(userLogin models.Login, a context.Context) (*models.User, error)
	UserEdit(a context.Context, str string, bsondata primitive.M) (*mongo.UpdateResult, error)
	UserPorfileDB(a context.Context, user models.User, id string) error
	// user content
	CreateContent(addcontents models.AllContents, str string, a context.Context) (*mongo.InsertOneResult, error)
	FindContent(a context.Context, str string) ([]models.AllContents, error)
	ContentEdit(a context.Context, str string, cid string, bsondata primitive.M) (*mongo.UpdateResult, error)
	ContentDelete(a context.Context, str, content_id string) (*mongo.DeleteResult, error)
}
