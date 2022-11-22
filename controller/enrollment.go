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
// @Param       	listforregistrationInput body models.ListforRegistrationRequest true "ListforRegistration Input"
// @Success		  	200 {object} models.ListforRegistrationResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_listforregistration/ [post]
func GetListforRegistration(c *fiber.Ctx) error {
	listforregistrationInput := models.ListforRegistrationRequest{}

	if parErr := c.BodyParser(&listforregistrationInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	listforregistrationModel := []models.ListforRegistrationResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_clientregistration").Find(&listforregistrationModel, listforregistrationInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    listforregistrationModel,
	})
}
