package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func AppRoutes2(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Administration
	app.Post("/select_usermanagement/", controller.SelectUserManagementbyID)
	app.Post("/select_heirarchy/", controller.SelectHeirarchybyID)

	//Enrollment

	//Monitoring
	app.Post("/select_useddevice/", controller.SelectUseddevicebyID)

	//Utilities
	app.Post("/select_feestructure/", controller.SelectFeeStructurebyID)
	app.Post("/select_paramconfig/", controller.SelectParamConfigbyID)
	app.Post("/select_atmloc/", controller.SelectAtmLocbyID)
	///CustomerService

	///Reports

	///Task

}
