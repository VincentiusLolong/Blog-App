package routes

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/controllers" //add this
	"fiber-mongo-api/middleware"
	"fiber-mongo-api/services"

	"github.com/gofiber/fiber/v2"
)

var (
	db   configs.MongoDB         = configs.New()
	serv services.SerivceDataset = services.NewSerivce(db)
	con  controllers.Controller  = controllers.Control(serv)
)

func UserRoute(a *fiber.App) {

	// ============       ADMIN adn User     =================
	// find, editm delete by id (one)
	UserRoute := a.Group("/api/V1")
	UserRoute.Post("/user/register", con.CreateUser)
	UserRoute.Post("/user/signin", con.SignIn)
	secured := UserRoute.Group("/secured").Use(middleware.Auth())
	secured.Post("user/logout", con.Logout)
	secured.Delete("user/delete", con.DeleteMyAccount)
	secured.Get("user/Get", con.GetMyAccountProfile)
	secured.Put("user/edit/:orgs/:about", con.EditMyPorfile)
	secured.Post("user/addcontent", con.AddContent)
	secured.Get("user/findcontent", con.FindContent)
	secured.Put("user/editcontent/:content_id", con.EditContent)
	secured.Delete("user/deletecontent/:content_id", con.DeleteContent)
	secured.Post("user/addcomment", con.AddComment)

	// //=============         Public        =================
	// // find, edit, delete many by name (many)
	// PublicRoute := app.Group("/api/V2")
	// PublicRoute.Get("/users/:name/:age", con.Alluserbyage)
}
