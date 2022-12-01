package secure

import (
	"fiber-mongo-api/models"

	"golang.org/x/crypto/bcrypt"
)

// var userCollection *mongo.Collection = configs.GetCollection(configs.AllEnv("CONTENTCOLLECTION"))
// var validate = validator.New()

func HashPassword(u *models.User, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func CheckPassword(u *models.User, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

// func GenearateToken(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var request models.TokenRequest
// 	if err := c.BodyParser(&request); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(
// 			responses.UserResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data: &fiber.Map{
// 					"data": err.Error()}})
// 	}
// 	if validationErr := validate.Struct(&request); validationErr != nil {
// 		return c.Status(http.StatusBadRequest).JSON(
// 			responses.UserResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data: &fiber.Map{
// 					"data": validationErr.Error()}})
// 	}

// 	var user models.User
// 	if err := userCollection.FindOne(ctx, bson.M{"email": request.Email}).Decode(&user); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(
// 			responses.UserResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data: &fiber.Map{
// 					"data": err.Error()}})
// 	}

// 	if CredentialError := CheckPassword(&user, request.Password); CredentialError != nil {
// 		return c.Status(http.StatusBadRequest).JSON(
// 			responses.UserResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data: &fiber.Map{
// 					"data": CredentialError.Error()}})
// 	}

// 	tokenstring, err := GenerateJWT(user.Email, user.Name)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(
// 			responses.UserResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data: &fiber.Map{
// 					"data": err.Error()}})
// 	}
// 	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
// 		Status:  http.StatusCreated,
// 		Message: "success",
// 		Data: &fiber.Map{
// 			"data": tokenstring}})
// }
