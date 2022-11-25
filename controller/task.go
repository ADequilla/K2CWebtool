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
// @Param       	taskInput body models.TaskRequest true "Task Input"
// @Success		  	200 {object} models.TaskResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_task/ [post]
func GetTask(c *fiber.Ctx) error {
	taskInput := models.TaskRequest{}

	if parErr := c.BodyParser(&taskInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	taskModel := []models.TaskResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_task").Find(&taskModel, taskInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(taskModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    taskModel,
	})

}
