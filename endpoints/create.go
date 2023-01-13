package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Create(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Administration
	app.Post("/create_usermanagement/", controller.CreateUserManagements)
	app.Post("/create_heirarchy/", controller.CreateHeirarchy)

	//Enrollment

	//Monitoring

	//Utilities
	app.Post("/create_feestructure/", controller.CreateFeeStructure)
	app.Post("/create_paramconfig/", controller.CreateParamConfig)
	app.Post("/create_atmloc/", controller.CreateAtmLoc)
	app.Post("/create_banknews/", controller.CreateBankNews)
	app.Post("/create_productservices/", controller.CreateProductServices)
	app.Post("/create_servicedowntime/", controller.CreateServiceDowntime)
	app.Post("/create_insti/", controller.CreateInsti)
	app.Post("/create_branch/", controller.CreateBranch)
	app.Post("/create_unit/", controller.CreateUnit)
	app.Post("/create_center/", controller.CreateCenter)
	app.Post("/create_provider/", controller.CreateProvider)
	app.Post("/create_product_type/", controller.CreateProductType)
	app.Post("/create_productcategory/", controller.CreateProductCategory)
	app.Post("/create_billerproduct/", controller.CreateBillerProduct)
	app.Post("/create_loadproduct/", controller.CreateLoadProduct)
	app.Post("/create_commission/", controller.CreateCommission)
	app.Post("/create_bank/", controller.CreateBankList)
	app.Post("/create_partner/", controller.CreatePartnerList)
	app.Post("/create_splashscreen/", controller.CreateSplashScreen)

	///CustomerService
	app.Post("/create_broadcastsms/", controller.CreateBroadcastSms)
	app.Post("/create_concerntype/", controller.CreateConcernType)
	app.Post("/create_csrhotline/", controller.CreateCsrHotline)

	///Reports

	///Task
}
