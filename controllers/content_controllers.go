package controllers

import (
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (ct *controller) AddContent(c *fiber.Ctx) error {
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

	result, err := ct.service.CreateContent(addcontents, str, a)
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

func (ct *controller) DeleteContent(c *fiber.Ctx) error {
	content_id := c.Params("content_id")
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))
	defer b()

	data, err := ct.service.ContentDelete(a, str, content_id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: &fiber.Map{
			"data": data}})
}

func (ct *controller) EditContent(c *fiber.Ctx) error {
	content_id := c.Params("content_id")
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))
	defer b()

	var editcontent models.AllContents

	if err := c.BodyParser(&editcontent); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": err.Error()}})
	}

	primitive, err := parsejson(editcontent)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data: &fiber.Map{
					"DataNull": err.Error()}})
	}

	result, errs := ct.service.ContentEdit(a, str, content_id, primitive)
	if errs != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: &fiber.Map{
			"data": result}})
}

func (ct *controller) FindContent(c *fiber.Ctx) error {
	a, b := contectx()
	str := fmt.Sprintf("%v", c.Locals("id"))
	defer b()

	data, err := ct.service.FindContent(a, str)
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
