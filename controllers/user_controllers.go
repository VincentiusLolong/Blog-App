package controllers

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/controllers/repo"
	"fiber-mongo-api/controllers/secure"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	var allsession, sessionerr = Store.Get(c)
	if sessionerr != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": sessionerr.Error()}})
	}
	// loggedInCookie := c.Cookies("logged_in")
	// if loggedInCookie != "" {
	// 	return c.Status(http.StatusOK).JSON(responses.UserResponse{
	// 		Status:  http.StatusBadRequest,
	// 		Message: "you're already login",
	// 		Data: &fiber.Map{
	// 			"statusLogin": "already"}})
	// }

	var signin models.Login
	if err := c.BodyParser(&signin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": err.Error()}})
	}
	if validationErr := validate.Struct(&signin); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"ValidateErr": validationErr.Error()}})
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

	rt, err := secure.GenerateJWT(&models.RefreshDataToken{
		Id: res.User_Id,
	}, 43200)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}

	t, err := secure.GenerateJWT(&models.GetDataToken{
		Id:    res.User_Id,
		Email: res.Email,
		Name:  res.Name,
	}, 15)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}

	RedisSet := configs.RedisSet(res.User_Id.Hex(), rt)
	if RedisSet != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": RedisSet.Error()}})
	}

	// c.Cookie(&fiber.Cookie{
	// 	Name:    "logged_in",
	// 	Value:   t,
	// 	Expires: time.Now().Add(time.Minute * 15),
	// })

	allsession.Set("logged_in", t)

	if err := allsession.Save(); err != nil {
		panic(err)
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &fiber.Map{
			"Access Token": t,
			"name":         res.Name}})
}

func Logout(c *fiber.Ctx) error {
	var allsession, sessionerr = Store.Get(c)
	if sessionerr != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": sessionerr.Error()}})
	}
	str := fmt.Sprintf("%v", c.Locals("id"))

	if err := allsession.Destroy(); err != nil {
		panic(err)
	}
	RedisDel := configs.RedisDelete(str)
	if RedisDel != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": RedisDel.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &fiber.Map{
			"token": "expired"}})
}

func GetMyAccountProfile(c *fiber.Ctx) error {
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))

	defer b()

	var user models.User

	objId, _ := primitive.ObjectIDFromHex(str)
	err := userCollection.FindOne(a, bson.M{"user_id": objId}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": user}})
}

func DeleteMyAccount(c *fiber.Ctx) error {
	var allsession, sessionerr = Store.Get(c)
	if sessionerr != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": sessionerr.Error()}})
	}
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))
	defer b()

	objId, _ := primitive.ObjectIDFromHex(str)
	filter := bson.D{{Key: "id", Value: objId}}

	result, err := userCollection.DeleteOne(a, filter)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if err := allsession.Destroy(); err != nil {
		panic(err)
	}

	RedisDel := configs.RedisDelete(str)
	if RedisDel != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": RedisDel.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success", Data: &fiber.Map{"data": result}})
}

func EditMyPorfile(c *fiber.Ctx) error {
	var edit models.UserPorfile
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))
	defer b()

	if err := c.BodyParser(&edit); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": err.Error()}})
	}

	data, err := ParseJson(edit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	result, errs := repo.UserEdit(a, str, data)
	if errs != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": result}})
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
