package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SlelectbyCid(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Administration
	app.Post("/select_usermanagement/", controller.SelectUserManagementbyID)
	app.Post("/select_rolemanagement/", controller.SelectRolesManagementbyID)
	app.Post("/select_heirarchy/", controller.SelectHeirarchybyID)

	//Enrollment

	//Monitoring
	app.Post("/select_remittancesent/", controller.SelectRemittanceSent)
	app.Post("/select_remittanceclaimed/", controller.SelectRemittanceClaimed)
	app.Post("/select_remittancecancelled/", controller.SelectRemittanceCancelled)
	app.Post("/select_remittancepending/", controller.SelectRemittancePending)
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
	app.Post("/select_commission/", controller.SelectCommissionbyID)
	app.Post("/select_banklist/", controller.SelectBanklistbyID)
	app.Post("/select_partnerlist/", controller.SelectPartnerlistbyID)
	app.Post("/select_splashscreen/", controller.SelectSplashScreenbyID)
	///CustomerService
	app.Post("/select_broadcastsms/", controller.SelectBroadcastSmsbyID)
	app.Post("/select_concerntype/", controller.SelectConcernTypebyID)
	app.Post("/select_csrhotline/", controller.SelectCsrHotlinebyID)

	///Reports

	///Task

}
