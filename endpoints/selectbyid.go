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
	app.Post("/select_banknews/", controller.SelectBankNewsbyID)
	app.Post("/select_productservices/", controller.SelectProductServicesbyID)
	app.Post("/select_servicedowntime/", controller.SelectServiceDowntimebyID)
	app.Post("/select_insti/", controller.SelectInstibyID)
	app.Post("/select_branch/", controller.SelectBranchbyID)
	app.Post("/select_unit/", controller.SelectUnitbyID)
	app.Post("/select_center/", controller.SelectCenterbyID)
	app.Post("/select_provider/", controller.SelectProviderbyID)
	app.Post("/select_producttype/", controller.SelectProductTypebyID)
	app.Post("/select_productcategory/", controller.SelectProductCategorybyID)
	app.Post("/select_billerproduct/", controller.SelectBillerProductbyID)
	app.Post("/select_loadproduct/", controller.SelectLoadProductbyID)
	///CustomerService

	///Reports

	///Task

}
