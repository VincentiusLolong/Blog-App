package controllers

import (
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (ct *controller) AddComment(c *fiber.Ctx) error {
	str := fmt.Sprintf("%v", c.Locals("id"))
	var addcomment models.Comments
	a, b := contectx()
	defer b()

	if err := c.BodyParser(&addcomment); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error()}})
	}

	if validationErr := validate.Struct(&addcomment); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"data": validationErr.Error()}})
	}

	result, err := ct.service.CreateComment(a, addcomment, str)
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
