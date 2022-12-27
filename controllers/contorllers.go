package controllers

import (
	"context"
	"encoding/json"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("THECOLLECTION"))

// var contentCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("PXCOLLECTIONS"))
var validate = validator.New()
var Store = session.New(session.Config{
	Expiration:     168 * time.Hour,
	CookieHTTPOnly: true,
})

func contectx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

func ParseJson[test models.UserPorfile | models.AllContents](edit test) (primitive.M, error) {
	data := make(map[string]interface{})
	userJson, err := json.Marshal(edit)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(userJson, &data); err != nil {
		return nil, err
	}
	delete(data, "content_id")
	delete(data, "user_id")
	return bson.M(data), nil
}
