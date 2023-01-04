package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Delete(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Administration

	app.Post("/drop_hierarchy/", controller.DropHierarchy)

	//Enrollment

	//Monitoring

	//Utilities
	app.Post("/drop_feestructure/", controller.DropFeeStructure)
	app.Post("/drop_atmloc/", controller.DropAtmLoc)
	app.Post("/drop_banknews/", controller.DropBankNews)
	app.Post("/drop_productservices/", controller.DropProductServices)
	app.Post("/drop_servicedowntime/", controller.DropServiceDowntime)
	app.Post("/drop_insti/", controller.DropInstitution)
	app.Post("/drop_branch/", controller.DropBranch)
	app.Post("/drop_unit/", controller.DropUnit)
	app.Post("/drop_center/", controller.DropCenter)
	app.Post("/drop_provider/", controller.DropProvider)
	app.Post("/drop_producttype/", controller.DropProductType)
	app.Post("/drop_productcategory/", controller.DropProductCategory)
	app.Post("/drop_billerproduct/", controller.DropBillerProduct)
	app.Post("/drop_loadproduct/", controller.DropLoadProduct)
	app.Post("/drop_commission/", controller.DropCommission)
	app.Post("/drop_banklist/", controller.DropBanklist)
	app.Post("/drop_partnerlist/", controller.DropPartnerList)
	app.Post("/drop_splashscreen/", controller.DropSplashScreen)

	///CustomerService
	app.Post("/drop_broadcastsms/", controller.DropBroadcastSms)
	app.Post("/drop_concerntype/", controller.DropConcernType)
	app.Post("/drop_csrhotline/", controller.DropCsrHotline)

	///Reports

	///Task

}
