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
// @Param       	csrhotlineInput body models.CsrHotlineRequest true "CsrHotline Input"
// @Success		  	200 {object} models.CsrHotlineResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_csrhotline/ [post]
func GetCsrHotline(c *fiber.Ctx) error {
	csrhotlineInput := models.CsrHotlineRequest{}

	if parErr := c.BodyParser(&csrhotlineInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	csrhotlineModel := []models.CsrHotlineResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_csrhotline").Find(&csrhotlineModel, csrhotlineInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(csrhotlineModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    csrhotlineModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	corncerntypeInput body models.ConcernTypeRequest true "ConcernType Input"
// @Success		  	200 {object} models.ConcernTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_concerntype/ [post]
func GetConcernType(c *fiber.Ctx) error {
	concerntypeInput := models.ConcernTypeRequest{}

	if parErr := c.BodyParser(&concerntypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	concerntypeModel := []models.ConcernTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_concerntype").Find(&concerntypeModel, concerntypeInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(concerntypeModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    concerntypeModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	csrdashboardInput body models.CSRDashboardRequest true "CSRDashboard Input"
// @Success		  	200 {object} models.CSRDashboardResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_csrdashboard/ [post]
func GetCSRDashboard(c *fiber.Ctx) error {
	csrdashboardInput := models.CSRDashboardRequest{}

	if parErr := c.BodyParser(&csrdashboardInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	csrdashboardModel := []models.CSRDashboardResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_cs_ticket").Find(&csrdashboardModel, csrdashboardInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(csrdashboardModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    csrdashboardModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	broadcastsmsInput body models.BroadcastSmsRequest true "BroadcastSms Input"
// @Success		  	200 {object} models.BroadcastSmsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_broadcastsms/ [post]
func GetBroadcastSms(c *fiber.Ctx) error {
	broadcastsmsInput := models.BroadcastSmsRequest{}

	if parErr := c.BodyParser(&broadcastsmsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	broadcastsmsModel := []models.BroadcastSmsResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_broadcastsms").Find(&broadcastsmsModel, broadcastsmsInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(broadcastsmsModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    broadcastsmsModel,
	})
}
