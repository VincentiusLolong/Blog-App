package routes

import (
	"fiber-mongo-api/controllers" //add this

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// ============       ADMIN adn User     =================
	// find, editm delete by id (one)
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId/:orgs/:about", controllers.EditAUser) // both admin and user can use this but admin has all access to all user not the user one
	app.Delete("/user/:userId", controllers.DeleteAUser)

	//=============         only admin       =================
	// find, edit, delete many by name (many)
	app.Get("/alluser/:name/:age", controllers.Alluserbyage)

}
