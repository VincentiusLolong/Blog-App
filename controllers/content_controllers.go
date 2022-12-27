package controllers

import (
	"fiber-mongo-api/controllers/repo"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddContent(c *fiber.Ctx) error {
	var addcontents models.AllContents
	str := fmt.Sprintf("%v", c.Locals("id"))
	a, b := contectx()
	defer b()

	if err := c.BodyParser(&addcontents); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}
	if validationErr := validate.Struct(&addcontents); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": validationErr.Error()}})
	}

	result, err := repo.CreateContent(addcontents, str, a)
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
			"data": result}})
}

// func DeleteContent(c *fiber.Ctx) error {
// 	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
// 		Status:  http.StatusCreated,
// 		Message: "success",
// 		Data: &fiber.Map{
// 			"data": ""}})
// }

// func EditContent(c *fiber.Ctx) error {
// 	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
// 		Status:  http.StatusCreated,
// 		Message: "success",
// 		Data: &fiber.Map{
// 			"data": ""}})
// }

func FindContent(c *fiber.Ctx) error {
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))
	defer b()

	data, err := repo.FindContent(a, str)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "can't find content",
			Data: &fiber.Map{
				"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: &fiber.Map{
			"data": data}})
}
