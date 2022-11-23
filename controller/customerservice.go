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
