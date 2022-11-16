package controller

import (
	"fmt"
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

	csrhotlineinput := c.Params("csrhotlineInput")
	fmt.Println("CsrHotlineInput:", csrhotlineinput)

	csrhotlineModel := []models.CsrHotlineResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.csrhotline(?)", csrhotlineInput.SearchCsrhotline).Scan(&csrhotlineModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
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

	concerntypeinput := c.Params("corncerntypeInput")
	fmt.Println("ConcernTypeInput:", concerntypeinput)

	concerntypeModel := []models.ConcernTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.concerntype(?)", concerntypeInput.SearchConcerntype).Scan(&concerntypeModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    concerntypeModel,
	})
}
