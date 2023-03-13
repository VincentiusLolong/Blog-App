package controllers

import (
	"context"
	"encoding/json"
	"fiber-mongo-api/models"
	"fiber-mongo-api/services"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller interface {
	//Account
	CreateUser(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	GetMyAccountProfile(c *fiber.Ctx) error
	DeleteMyAccount(c *fiber.Ctx) error
	EditMyPorfile(c *fiber.Ctx) error

	//Content
	AddContent(c *fiber.Ctx) error
	DeleteContent(c *fiber.Ctx) error
	EditContent(c *fiber.Ctx) error
	FindContent(c *fiber.Ctx) error

	//Comments
	AddComment(c *fiber.Ctx) error
	EditComment(c *fiber.Ctx) error
}

type controller struct {
	service services.SerivceDataset
}

func Control(serv services.SerivceDataset) Controller {
	return &controller{
		service: serv,
	}
}

var validate = validator.New()
var Store = session.New(session.Config{
	Expiration:     168 * time.Hour,
	CookieHTTPOnly: true,
})

func contectx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

func parsejson[test models.UserPorfile | models.AllContents | models.Comments](edit test) (primitive.M, error) {
	data := make(map[string]interface{})
	userJson, err := json.Marshal(edit)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(userJson, &data); err != nil {
		return nil, err
	}

	if _, ok := data["content_id"]; !ok {
		delete(data, "user_id")
	} else {
		delete(data, "content_id")
		delete(data, "user_id")
		delete(data, "comments_id")
	}
	return bson.M(data), nil
}
