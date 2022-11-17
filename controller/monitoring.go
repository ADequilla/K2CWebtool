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
// @Param       	clientprofileInput body models.ClientProfileRequest true "Clientprofile Input"
// @Success		  	200 {object} models.ClientProfileResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_clientprofile/ [post]
func GetClientProfile(c *fiber.Ctx) error {
	clientprofileInput := models.ClientProfileRequest{}

	if parErr := c.BodyParser(&clientprofileInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	cfModel := []models.ClientProfileResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_clientprofile").Find(&cfModel, clientprofileInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    cfModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	remittanceLogInput body models.RemittanceLogRequest true "RemittanceLog Input"
// @Success		  	200 {object} models.RemittanceLogResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_remittancelog/ [post]
func GetRemittanceLog(c *fiber.Ctx) error {
	remittancelogInput := models.RemittanceLogRequest{}

	if parErr := c.BodyParser(&remittancelogInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	rlModel := []models.RemittanceLogResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_remittancelog").Find(&rlModel, remittancelogInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    rlModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Success		  	200 {object} models.RemittanceStatusResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_remittancestatus/ [post]
func GetRemittanceStatus(c *fiber.Ctx) error {
	rsModel := []models.RemittanceStatusResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.view_remittance_status").Scan(&rsModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    rsModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	smsLogInput body models.SmsLogRequest true "SmsLog Input"
// @Success		  	200 {object} models.SmsLogResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_smslog/ [post]
func GetSmsLog(c *fiber.Ctx) error {
	smslogInput := models.SmsLogRequest{}

	if parErr := c.BodyParser(&smslogInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	slModel := []models.SmsLogResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_smslog").Find(&slModel, smslogInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    slModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	transLogInput body models.TransLogRequest true "TransLog Input"
// @Success		  	200 {object} models.TransLogResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_transactionlog/ [post]
func GetTransLog(c *fiber.Ctx) error {
	translogInput := models.TransLogRequest{}

	if parErr := c.BodyParser(&translogInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	translogModel := []models.TransLogResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_translog").Find(&translogModel, translogInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    translogModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	usedDeviceInput body models.UsedDeviceRequest true "UsedDevice Input"
// @Success		  	200 {object} models.UsedDeviceResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_listuseddevice/ [post]
func GetUsedDevice(c *fiber.Ctx) error {
	useddeviceInput := models.UsedDeviceRequest{}

	if parErr := c.BodyParser(&useddeviceInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	useddeviceModel := []models.UsedDeviceResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_listuseddevice").Find(&useddeviceModel, useddeviceInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    useddeviceModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	failedEnrollmentInput body models.FailedEnrollmentRequest true "FailedEnrollment Input"
// @Success		  	200 {object} models.FailedEnrollmentResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_failedenrollment/ [post]
func GetFailedEnrollment(c *fiber.Ctx) error {
	failedenrollmentInput := models.FailedEnrollmentRequest{}

	if parErr := c.BodyParser(&failedenrollmentInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	failedenrollmentModel := []models.FailedEnrollmentResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_failedenrollment").Find(&failedenrollmentModel, failedenrollmentInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    failedenrollmentModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	listofAgentInput body models.ListofAgentRequest true "ListofAgent Input"
// @Success		  	200 {object} models.ListofAgentResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_listofagent/ [post]
func GetListofAgent(c *fiber.Ctx) error {
	listofagentInput := models.ListofAgentRequest{}

	if parErr := c.BodyParser(&listofagentInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	listofagentModel := []models.ListofAgentResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_listofagent").Find(&listofagentModel, listofagentInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    listofagentModel,
	})
}
