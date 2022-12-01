package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/controllers/repo"
	"fiber-mongo-api/controllers/secure"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("CONTENTCOLLECTION"))
var validate = validator.New()

func contectx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

func CreateUser(c *fiber.Ctx) error {
	a, b := contectx()
	var user models.User
	defer b()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": validationErr.Error()}})
	}

	errs := secure.HashPassword(&user, user.Password)

	if errs != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": errs.Error()}})
	}

	res, err := repo.CreateUserDB(user, a)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: &fiber.Map{
			"data": res}})
}

func SignIn(c *fiber.Ctx) error {
	a, b := contectx()
	defer b()

	var signin models.Login
	if err := c.BodyParser(&signin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}
	if validationErr := validate.Struct(&signin); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": validationErr.Error()}})
	}

	res, err := repo.SignInDB(signin, a)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error()}})
	}

	if CredentialError := secure.CheckPassword(res, signin.Password); CredentialError != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": CredentialError.Error()}})
	}

	tokenstring, err := secure.GenerateJWT(res.Email, res.Name)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &fiber.Map{
			"data":  res.Name,
			"token": tokenstring}})
}

func GetAUser(c *fiber.Ctx) error {
	a, b := contectx()
	userId := c.Params("userId")
	var user models.User
	defer b()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(a, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func DeleteAUser(c *fiber.Ctx) error {
	a, b := contectx()
	userId := c.Params("userId")
	defer b()

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "id", Value: objId}}
	result, err := userCollection.DeleteOne(a, filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

func EditAUser(c *fiber.Ctx) error {
	a, b := contectx()
	orgs, about, userId := c.Params("orgs"), c.Params("about"), c.Params("userId")
	defer b()

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "id", Value: objId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "orgs", Value: orgs}, {Key: "about", Value: about}}}}

	result, err := userCollection.UpdateOne(a, filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

func Alluserbyage(c *fiber.Ctx) error {
	a, b := contectx()
	name, age := c.Params("name"), c.Params("age")
	var users []models.User
	defer b()

	results, err := userCollection.Find(a, bson.M{"orgs": name, "age": age})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(a)
	for results.Next(a) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}},
	)
}
