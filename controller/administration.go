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
// @Param       	userInput body models.UserManagementRequest true "User Input"
// @Success		  	200 {object} models.UserManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_usermanagement/ [post]
func GetUserManagements(c *fiber.Ctx) error {
	userInput := models.UserManagementRequest{}

	if parErr := c.BodyParser(&userInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	umModel := []models.UserManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_usermanagement").Find(&umModel, userInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(umModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    umModel,
	})

}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	rolesInput body models.RolesManagementRequest true "Roles Input"
// @Success		  	200 {object} models.RolesManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_rolesmanagement/ [post]
func GetRolesManagements(c *fiber.Ctx) error {
	rolesInput := models.RolesManagementRequest{}

	if parErr := c.BodyParser(&rolesInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	rmModel := []models.RolesManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_rolemanagement").Find(&rmModel, rolesInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(rmModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    rmModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	hierarchyInput body models.HierarchyRequest true "Hierarchy Input"
// @Success		  	200 {object} models.HierarchyResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_hierarchy/ [post]
func GetHierarchy(c *fiber.Ctx) error {
	hierarchyInput := models.HierarchyRequest{}

	if parErr := c.BodyParser(&hierarchyInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	hModel := []models.HierarchyResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_hierarchy").Find(&hModel, hierarchyInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(hModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    hModel,
	})
}
