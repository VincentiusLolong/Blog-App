package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("CONTENTCOLLECTION"))
var validate = validator.New()
var Store = session.New(session.Config{
	Expiration:     15 * time.Minute,
	CookieHTTPOnly: true,
})

func contectx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}
