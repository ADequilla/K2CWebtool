package controller

import (
	"net/http"
	"webtool-api/middleware"
	"webtool-api/models"

	"github.com/gofiber/fiber/v2"
)

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	webreportInput body models.WebReportRequest true "WebReport Input"
// @Success		  	200 {object} models.WebReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_webreport/ [post]
func GetWebReport(c *fiber.Ctx) error {
	webreportInput := models.WebReportRequest{}

	if parErr := c.BodyParser(&webreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	webreportModel := []models.WebReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_web_report").Find(&webreportModel, webreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(webreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    webreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	transreportInput body models.TransactionReportRequest true "TransactionReport Input"
// @Success		  	200 {object} models.TransactionReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_transreport/ [post]
func GetTransactionReport(c *fiber.Ctx) error {
	transreportInput := models.TransactionReportRequest{}

	if parErr := c.BodyParser(&transreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	transreportModel := []models.TransactionReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_transaction_report").Find(&transreportModel, transreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(transreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    transreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	remittancesentreportInput body models.RemittanceSentReportRequest true "RemittanceSentReport Input"
// @Success		  	200 {object} models.RemittanceSentReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_remittancesentreport/ [post]
func GetRemittanceSentReport(c *fiber.Ctx) error {
	remittancesentreportInput := models.RemittanceSentReportRequest{}

	if parErr := c.BodyParser(&remittancesentreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	remittancesentreportModel := []models.RemittanceSentReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_remittance_sent_report").Find(&remittancesentreportModel, remittancesentreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(remittancesentreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    remittancesentreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	remittanceclaimedreportInput body models.RemittanceClaimedReportRequest true "RemittanceClaimedReport Input"
// @Success		  	200 {object} models.RemittanceClaimedReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_remittanceclaimedreport/ [post]
func GetRemittanceClaimedReport(c *fiber.Ctx) error {
	remittanceclaimedreportInput := models.RemittanceClaimedReportRequest{}

	if parErr := c.BodyParser(&remittanceclaimedreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	remittanceclaimedreportModel := []models.RemittanceClaimedReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_remittance_claimed_report").Find(&remittanceclaimedreportModel, remittanceclaimedreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(remittanceclaimedreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    remittanceclaimedreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	remittancecancelledreportInput body models.RemittanceCancelledReportRequest true "RemittanceCancelledReport Input"
// @Success		  	200 {object} models.RemittanceCancelledReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_remittancecancelledreport/ [post]
func GetRemittanceCancelledReport(c *fiber.Ctx) error {
	remittancecancelledreportInput := models.RemittanceCancelledReportRequest{}

	if parErr := c.BodyParser(&remittancecancelledreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	remittancecancelledreportModel := []models.RemittanceCancelledReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_remittance_cancelled_report").Find(&remittancecancelledreportModel, remittancecancelledreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(remittancecancelledreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    remittancecancelledreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	activityhistoryreportInput body models.ActivityHistoryReportRequest true "ActivityHistoryReport Input"
// @Success		  	200 {object} models.ActivityHistoryReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_activityhistoryreport/ [post]
func GetActivityHistoryReport(c *fiber.Ctx) error {
	activityhistoryreportInput := models.ActivityHistoryReportRequest{}

	if parErr := c.BodyParser(&activityhistoryreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	activityhistoryreportModel := []models.ActivityHistoryReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_activity_history_report").Find(&activityhistoryreportModel, activityhistoryreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(activityhistoryreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    activityhistoryreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	mpinchangereportInput body models.MpinChangeReportRequest true "MpinChangeReport Input"
// @Success		  	200 {object} models.MpinChangeReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_mpinchangereport/ [post]
func GetMpinChangeReport(c *fiber.Ctx) error {
	mpinchangereportInput := models.MpinChangeReportRequest{}

	if parErr := c.BodyParser(&mpinchangereportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	mpinchangereportModel := []models.MpinChangeReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_mpin_change_report").Find(&mpinchangereportModel, mpinchangereportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(mpinchangereportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    mpinchangereportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	smsactivationreportInput body models.SmsActivationReportRequest true "SmsActivationReport Input"
// @Success		  	200 {object} models.SmsActivationReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_smsactivationreport/ [post]
func GetSmsActivationReport(c *fiber.Ctx) error {
	smsactivationreportInput := models.SmsActivationReportRequest{}

	if parErr := c.BodyParser(&smsactivationreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	smsactivationreportModel := []models.SmsActivationReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_sms_activation_report").Find(&smsactivationreportModel, smsactivationreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(smsactivationreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    smsactivationreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	loginlogoutreportInput body models.LoginLogoutReportRequest true "LoginLogoutReport Input"
// @Success		  	200 {object} models.LoginLogoutReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_loginlogoutreport/ [post]
func GetLoginLogoutReport(c *fiber.Ctx) error {
	loginlogoutreportInput := models.LoginLogoutReportRequest{}

	if parErr := c.BodyParser(&loginlogoutreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	loginlogoutreportModel := []models.LoginLogoutReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_login_logout_report").Find(&loginlogoutreportModel, loginlogoutreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(loginlogoutreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    loginlogoutreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	useractivityreportInput body models.UserActivityReportRequest true "UserActivityReport Input"
// @Success		  	200 {object} models.UserActivityReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_useractivityreport/ [post]
func GetUserActivityReport(c *fiber.Ctx) error {
	useractivityreportInput := models.UserActivityReportRequest{}

	if parErr := c.BodyParser(&useractivityreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	useractivityreportModel := []models.UserActivityReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_user_activity_report").Find(&useractivityreportModel, useractivityreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(useractivityreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    useractivityreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	transactionsuspiciousInput body models.TransactionSuspiciousReportRequest true "TransactionSuspiciousReport Input"
// @Success		  	200 {object} models.TransactionSuspiciousReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_transactionsuspiciousreport/ [post]
func GetTransactionSuspiciousReport(c *fiber.Ctx) error {
	transactionsuspiciousreportInput := models.TransactionSuspiciousReportRequest{}

	if parErr := c.BodyParser(&transactionsuspiciousreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	transactionsuspiciousreportModel := []models.TransactionSuspiciousReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_transaction_suspicious_report").Find(&transactionsuspiciousreportModel, transactionsuspiciousreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(transactionsuspiciousreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    transactionsuspiciousreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	transactionvalidInput body models.TransactionValidReportRequest true "TransactionValidReport Input"
// @Success		  	200 {object} models.TransactionValidReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_transactionvalidreport/ [post]
func GetTransactionValidReport(c *fiber.Ctx) error {
	transactionvalidreportInput := models.TransactionValidReportRequest{}

	if parErr := c.BodyParser(&transactionvalidreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	transactionvalidreportModel := []models.TransactionValidReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_transaction_valid_report").Find(&transactionvalidreportModel, transactionvalidreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(transactionvalidreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    transactionvalidreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	csdashboardInput body models.CsDashboardReportRequest true "CsDashboardReport Input"
// @Success		  	200 {object} models.CsDashboardReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_csdashboardreport/ [post]
func GetCsDashboardReport(c *fiber.Ctx) error {
	csdashboardreportInput := models.CsDashboardReportRequest{}

	if parErr := c.BodyParser(&csdashboardreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	csdashboardreportModel := []models.CsDashboardReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_cs_dashboard_report").Find(&csdashboardreportModel, csdashboardreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(csdashboardreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    csdashboardreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	reconccmInput body models.ReconCcmReportRequest true "ReconCcmReport Input"
// @Success		  	200 {object} models.ReconCcmReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_reconccmreport/ [post]
func GetReconCcmReport(c *fiber.Ctx) error {
	reconccmreportInput := models.ReconCcmReportRequest{}

	if parErr := c.BodyParser(&reconccmreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	reconccmreportModel := []models.ReconCcmReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_recon_ccm_report").Find(&reconccmreportModel, reconccmreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(reconccmreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    reconccmreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	smslogInput body models.SmsLogReportRequest true "SmsLogReport Input"
// @Success		  	200 {object} models.SmsLogReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_smslogreport/ [post]
func GetSmsLogReport(c *fiber.Ctx) error {
	smslogreportInput := models.SmsLogReportRequest{}

	if parErr := c.BodyParser(&smslogreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	smslogreportModel := []models.SmsLogReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_sms_log_report").Find(&smslogreportModel, smslogreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(smslogreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    smslogreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	accountstatusInput body models.AccountStatusReportRequest true "AccountStatusReport Input"
// @Success		  	200 {object} models.AccountStatusReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_accountstatusreport/ [post]
func GetAccountStatusReport(c *fiber.Ctx) error {
	accountstatusreportInput := models.AccountStatusReportRequest{}

	if parErr := c.BodyParser(&accountstatusreportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	accountstatusreportModel := []models.AccountStatusReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_account_status_report").Find(&accountstatusreportModel, accountstatusreportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(accountstatusreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    accountstatusreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	enableagentInput body models.EnableAgentReportRequest true "EnableAgentReport Input"
// @Success		  	200 {object} models.EnableAgentReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_enableagentreport/ [post]
func GetEnableAgentReport(c *fiber.Ctx) error {
	enableagentInput := models.EnableAgentReportRequest{}

	if parErr := c.BodyParser(&enableagentInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	enableagentreportModel := []models.EnableAgentReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_enable_agent_report").Find(&enableagentreportModel, enableagentInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(enableagentreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    enableagentreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	igatereconInput body models.IgateReconReportRequest true "IgateReconReport Input"
// @Success		  	200 {object} models.IgateReconReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_igatereconreport/ [post]
func GetIgateReconReport(c *fiber.Ctx) error {
	igatereconInput := models.IgateReconReportRequest{}

	if parErr := c.BodyParser(&igatereconInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	igatereconreportModel := []models.IgateReconReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_igate_recon_report").Find(&igatereconreportModel, igatereconInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(igatereconreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    igatereconreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	activatedmerchantInput body models.ActivatedMerchantReportRequest true "ActivatedMerchantReport Input"
// @Success		  	200 {object} models.ActivatedMerchantReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_activatedmerchantreport/ [post]
func GetActivatedMerchantReport(c *fiber.Ctx) error {
	activatedmerchantInput := models.ActivatedMerchantReportRequest{}

	if parErr := c.BodyParser(&activatedmerchantInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	activatedmerchantreportModel := []models.ActivatedMerchantReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_activated_merchant_report").Find(&activatedmerchantreportModel, activatedmerchantInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(activatedmerchantreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    activatedmerchantreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	deactivatedmerchantInput body models.DeactivatedMerchantReportRequest true "DeactivatedMerchantReport Input"
// @Success		  	200 {object} models.DeactivatedMerchantReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_deactivatedmerchantreport/ [post]
func GetDeactivatedMerchantReport(c *fiber.Ctx) error {
	deactivatedmerchantInput := models.DeactivatedMerchantReportRequest{}

	if parErr := c.BodyParser(&deactivatedmerchantInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	deactivatedmerchantreportModel := []models.DeactivatedMerchantReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_deactivated_merchant_report").Find(&deactivatedmerchantreportModel, deactivatedmerchantInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(deactivatedmerchantreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    deactivatedmerchantreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	useddeviceInput body models.UsedDeviceReportRequest true "UsedDeviceReport Input"
// @Success		  	200 {object} models.UsedDeviceReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_useddevicereport/ [post]
func GetUsedDeviceReport(c *fiber.Ctx) error {
	useddeviceInput := models.UsedDeviceReportRequest{}

	if parErr := c.BodyParser(&useddeviceInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	useddevicereportModel := []models.UsedDeviceReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_used_device_report").Find(&useddevicereportModel, useddeviceInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(useddevicereportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    useddevicereportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	failedenrollmentInput body models.FailedEnrollmentReportRequest true "FailedEnrollmentReport Input"
// @Success		  	200 {object} models.FailedEnrollmentReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_failedenrollmentreport/ [post]
func GetFailedEnrollmentReport(c *fiber.Ctx) error {
	failedenrollmentInput := models.FailedEnrollmentReportRequest{}

	if parErr := c.BodyParser(&failedenrollmentInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	failedenrollmentreportModel := []models.FailedEnrollmentReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_failed_enrollment_report").Find(&failedenrollmentreportModel, failedenrollmentInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(failedenrollmentreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    failedenrollmentreportModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	registeredclientInput body models.RegisteredClientReportRequest true "RegisteredClientReport Input"
// @Success		  	200 {object} models.RegisteredClientReportResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_registeredclientReportreport/ [post]
func GetRegisteredClientReport(c *fiber.Ctx) error {
	registeredclientReportInput := models.RegisteredClientReportRequest{}

	if parErr := c.BodyParser(&registeredclientReportInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	registeredclientreportModel := []models.RegisteredClientReportResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_registered_client_report").Find(&registeredclientreportModel, registeredclientReportInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(registeredclientreportModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    registeredclientreportModel,
	})

}
