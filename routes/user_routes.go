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
	secured.Post("user/logout", controllers.Logout)
	secured.Delete("user/delete", controllers.DeleteMyAccount)
	secured.Get("user/Get", controllers.GetMyAccountProfile)
	secured.Put("user/edit/:orgs/:about", controllers.EditMyPorfile)
	secured.Post("user/addcontent", controllers.AddContent)
	secured.Get("user/findcontent", controllers.FindContent)

	// //=============         Public        =================
	// // find, edit, delete many by name (many)
	PublicRoute := app.Group("/api/V2")
	PublicRoute.Get("/users/:name/:age", controllers.Alluserbyage)
}
