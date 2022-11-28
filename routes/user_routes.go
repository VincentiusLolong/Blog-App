package routes

import (
	"fiber-mongo-api/controllers" //add this

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// ============       ADMIN adn User     =================
	// find, editm delete by id (one)
	UserRoute := app.Group("/api/V1")
	UserRoute.Post("/user", controllers.CreateUser)
	UserRoute.Get("/users/:userId/:email/:pass", controllers.SignIn)
	UserRoute.Put("/user/:userId/:orgs/:about", controllers.EditAUser) // both admin and user can use this but admin has all access to all user not the user one
	UserRoute.Delete("/user/:userId", controllers.DeleteAUser)

	//=============         only admin       =================
	// find, edit, delete many by name (many)
	AdminRoute := app.Group("/api/V2")
	AdminRoute.Get("/users/:name/:age", controllers.Alluserbyage)
}
