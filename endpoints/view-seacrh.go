package endpoints

import (
	"webtool-api/controller"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func ViewSearch(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	//Login/Authenticated/Logout
	app.Post("/Login/", controller.Login)
	app.Post("/Authenticate/", controller.Authenticated)
	app.Post("/Logout/", controller.Logout)

	//Administration
	app.Post("/get_usermanagement/", controller.GetUserManagements)
	app.Post("/get_rolesmanagement/", controller.GetRolesManagements)
	app.Post("/get_hierarchy/", controller.GetHierarchy)

	//Enrollment
	app.Post("/get_listforregistration/", controller.GetListforRegistration)

	//Monitoring
	app.Post("/get_clientprofile/", controller.GetClientProfile)
	app.Post("/get_remittancelog/", controller.GetRemittanceLog)
	app.Post("/get_remittancestatus/", controller.GetRemittanceStatus)
	app.Post("/get_smslog/", controller.GetSmsLog)
	app.Post("/get_transactionLog/", controller.GetTransLog)
	app.Post("/get_listuseddevice/", controller.GetUsedDevice)
	app.Post("/get_failedenrollment/", controller.GetFailedEnrollment)
	app.Post("/get_listofagent/", controller.GetListofAgent)
	app.Post("/get_slfrequest/", controller.GetSlfRequest)
	app.Post("/get_operationdashboard/", controller.GetOperationDashboard)
	app.Post("/get_arp/", controller.GetAuthorResetPassword)
	app.Post("/get_agentdashboard/", controller.GetAgentDashboard)
	app.Post("/get_transconfirmation/", controller.GetTransConfirmation)

	//Utilities
	app.Post("/get_feestructure/", controller.GetFeeStructure)
	app.Post("/get_paramconfig/", controller.GetParamConfig)
	app.Post("/get_atmloc/", controller.GetAtmLoc)
	app.Post("/get_banknews/", controller.GetBankNews)
	app.Post("/get_productservices/", controller.GetProductServices)
	app.Post("/get_servicedowntime/", controller.GetServiceDowntime)
	app.Post("/get_insti/", controller.GetInstitution)
	app.Post("/get_branch/", controller.GetBranch)
	app.Post("/get_unit/", controller.GetUnit)
	app.Post("/get_center/", controller.GetCenter)
	app.Post("/get_provider/", controller.GetProvider)
	app.Post("/get_producttype/", controller.GetProductType)
	app.Post("/get_productcategory/", controller.GetProductCategory)
	app.Post("/get_billerproduct/", controller.GetBillerProduct)
	app.Post("/get_loadproduct/", controller.GetLoadProduct)
	app.Post("/get_commission/", controller.GetCommission)
	app.Post("/get_banklist/", controller.GetBankList)
	app.Post("/get_partnerlist/", controller.GetPartnerList)
	app.Post("/get_splashscreen/", controller.GetSplashScreen)

	///CustomerService
	app.Post("/get_csrhotline/", controller.GetCsrHotline)
	app.Post("/get_concerntype/", controller.GetConcernType)
	app.Post("/get_csrdashboard/", controller.GetCSRDashboard)
	app.Post("/get_broadcastsms/", controller.GetBroadcastSms)

	///Reports
	app.Post("/get_webreport/", controller.GetWebReport)
	app.Post("/get_transreport/", controller.GetTransactionReport)
	app.Post("/get_remittancesentreport/", controller.GetRemittanceSentReport)
	app.Post("/get_remittanceclaimedreport/", controller.GetRemittanceClaimedReport)
	app.Post("/get_remittancecancelledreport/", controller.GetRemittanceCancelledReport)
	app.Post("/get_activityhistoryreport/", controller.GetActivityHistoryReport)
	app.Post("/get_mpinchangereport/", controller.GetMpinChangeReport)
	app.Post("/get_smsactivationreport/", controller.GetSmsActivationReport)
	app.Post("/get_loginlogoutreport/", controller.GetLoginLogoutReport)
	app.Post("/get_useractivityreport/", controller.GetUserActivityReport)
	app.Post("/get_transactionsuspiciousreport/", controller.GetTransactionSuspiciousReport)
	app.Post("/get_registeredclientreport/", controller.GetRegisteredClientReport)
	app.Post("/get_transactionvalidreport/", controller.GetTransactionValidReport)
	app.Post("/get_csdashboardreport/", controller.GetCsDashboardReport)
	app.Post("/get_reconccmreport/", controller.GetReconCcmReport)
	app.Post("/get_smslogreport/", controller.GetSmsLogReport)
	app.Post("/get_accountstatusreport/", controller.GetAccountStatusReport)
	app.Post("/get_enableagentreport/", controller.GetEnableAgentReport)
	app.Post("/get_igatereconreport/", controller.GetIgateReconReport)
	app.Post("/get_activatedmerchantreport/", controller.GetActivatedMerchantReport)
	app.Post("/get_deactivatedmerchantreport/", controller.GetDeactivatedMerchantReport)
	app.Post("/get_useddevicereport/", controller.GetUsedDeviceReport)
	app.Post("/get_failedenrollmentreport/", controller.GetFailedEnrollmentReport)

	///Task
	app.Post("/get_task/", controller.GetTask)
}
