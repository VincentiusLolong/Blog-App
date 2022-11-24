package repo

import (
	"context"
	"fiber-mongo-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// import "github.com/gofiber/fiber"

type Repo interface {
	CreateUserDB(newUser models.User, a context.Context) (*mongo.InsertOneResult, error)
}
