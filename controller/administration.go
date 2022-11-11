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

	userinput := c.Params("userInput")
	fmt.Println("UserInput:", userinput)

	umModel := []models.UserManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.usermanagement(?)", userInput.SearchName).Scan(&umModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
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

	rolesinput := c.Params("rolesInput")
	fmt.Println("RolesInput:", rolesinput)

	rmModel := []models.RolesManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.rolesmanagement(?)", rolesInput.SearchRole).Scan(&rmModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(rmModel)
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

	hierarchyinput := c.Params("hierarchyInput")
	fmt.Println("HierarchyInput:", hierarchyinput)

	hModel := []models.HierarchyResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.hierarchy(?)", hierarchyInput.SearchHierarchy).Scan(&hModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(hModel)
}
