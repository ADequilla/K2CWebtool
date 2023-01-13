package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Login(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Login/Authenticated/Logout
	app.Post("/Login/", controller.Login)
	app.Post("/Authenticate/", controller.Authenticated)
	app.Post("/Logout/", controller.Logout)
}
