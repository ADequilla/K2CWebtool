package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func AppRoutes3(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Administration
	app.Post("/edit_usermanagement/", controller.EditUserManagement)
	app.Post("/edit_hierarchy/", controller.EditHierarchy)

	//Enrollment

	//Monitoring
	app.Post("/edit_useddevice/", controller.EditListUseddevice)
	//Utilities
	app.Post("/edit_feestructure/", controller.EditFeeStructure)
	app.Post("/edit_paramconfig/", controller.EditParamConfig)
	app.Post("/edit_atmloc/", controller.EditAtmLoc)
	///CustomerService

	///Reports

	///Task

}
