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
	app.Post("/edit_banknews/", controller.EditBankNews)
	app.Post("/edit_productservices/", controller.EditProductServices)
	app.Post("/edit_servicedowntime/", controller.EditServiceDowntime)
	app.Post("/edit_insti/", controller.EditInsti)
	app.Post("/edit_branch/", controller.EditBranch)
	app.Post("/edit_unit/", controller.EditUnit)
	app.Post("/edit_center/", controller.EditCenter)
	///CustomerService

	///Reports

	///Task

}
