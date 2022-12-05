package routes

import (
	"fiber-mongo-api/controllers" //add this
	"fiber-mongo-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// ============       ADMIN adn User     =================
	// find, editm delete by id (one)
	UserRoute := app.Group("/api/V1")
	UserRoute.Post("/user/register", controllers.CreateUser)
	UserRoute.Post("/user/signin", controllers.SignIn)
	secured := UserRoute.Group("/secured").Use(middleware.Auth())
	secured.Get("/ping", controllers.Ping)
	secured.Post("/logout", controllers.Logout)
	// UserRoute.Get("/users/:userId/:email/:pass", controllers.SignIn)
	// UserRoute.Put("/user/:userId/:orgs/:about", controllers.EditAUser) // both admin and user can use this but admin has all access to all user not the user one
	// UserRoute.Delete("/user/:userId", controllers.DeleteAUser)

	// //=============         Public        =================
	// // find, edit, delete many by name (many)
	PublicRoute := app.Group("/api/V2")
	PublicRoute.Get("/users/:name/:age", controllers.Alluserbyage)
}
