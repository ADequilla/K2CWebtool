package controller

import (
	"net/http"
	"time"
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
// @Param       	allUserInput body models.AllUserManagementRequest true "AllUser Input"
// @Success		  	200 {object} models.AllUserManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_usermanagement/ [post]
func SelectUserManagementbyID(c *fiber.Ctx) error {
	alluserInput := models.AllUserManagementRequest{}

	if parErr := c.BodyParser(&alluserInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	alluserModel := []models.AllUserManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_usermanagement(?)", alluserInput.Get_id).Scan(&alluserModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(alluserModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    alluserModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editUserInput body models.EditUserManagementRequest true "EditUser Input"
// @Success		  	200 {object} models.EditUserManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_usermanagement/ [post]
func EditUserManagement(c *fiber.Ctx) error {
	edituserInput := models.EditUserManagementRequest{}

	if parErr := c.BodyParser(&edituserInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	edituserModel := models.EditUserManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_usermanagement(?,?,?,?,?,?,?,?,?,?,?,?)", edituserInput.Get_user_id, edituserInput.Get_user_name, edituserInput.Get_given_name, edituserInput.Get_middle_name, edituserInput.Get_last_name, edituserInput.Get_user_email, edituserInput.Get_user_phone, edituserInput.Get_user_status, edituserInput.Get_check_status, edituserInput.Get_role_name, edituserInput.Get_inst_desc, edituserInput.Get_branch_desc).Scan(&edituserModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Updated Successfully",
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	createUserInput body models.CreateUserManagementRequest true "CreateUser Input"
// @Success		  	200 {object} models.CreateUserManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_usermanagement/ [post]
func CreateUserManagements(c *fiber.Ctx) error {
	var createInput models.CreateUserManagementRequest
	var err error

	if parErr := c.BodyParser(&createInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "401",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	if createInput.Set_user_password == "" {
		createInput.Set_user_password = "fdsap2023"
	}

	if createInput.Set_created_at == "" {
		createInput.Set_created_at = time.Now().Format("2014-02-04 18:05:00")
	}

	createInput.Set_user_password, err = middleware.HashPassword(createInput.Set_user_password)
	if err != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "HashPassword Error",
		})
	}

	createModel := models.CreateUserManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_usermanagement(?,?,?,?,?,?,?,?,?,?,?,?,?,?)", createInput.Set_user_name, createInput.Set_user_password, createInput.Set_user_email, createInput.Set_user_phone, createInput.Set_created_at, createInput.Set_created_by, createInput.Set_user_status, createInput.Set_inst_desc, createInput.Set_branch_desc, createInput.Set_given_name, createInput.Set_middle_name, createInput.Set_last_name, createInput.Set_check_status, createInput.Set_role).Scan(&createModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Created Successfully",
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
// @Param       	allRoleManagementInput body models.AllRoleManagementRequest true "AllRoleManagement Input"
// @Success		  	200 {object} models.AllRoleManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_rolemanagement/ [post]
func SelectRolesManagementbyID(c *fiber.Ctx) error {
	allrolemanagementInput := models.AllRoleManagementRequest{}

	if parErr := c.BodyParser(&allrolemanagementInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allrolemanagementModel := []models.AllRoleManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_role(?)", allrolemanagementInput.Get_id).Scan(&allrolemanagementModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allrolemanagementModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allrolemanagementModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editRoleManagementInput body models.EditRoleManagementRequest true "EditRoleManagement Input"
// @Success		  	200 {object} models.EditUserManagementResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_rolemanagement/ [post]
func EditRoleManagement(c *fiber.Ctx) error {
	editrolemanagementInput := models.EditRoleManagementRequest{}

	if parErr := c.BodyParser(&editrolemanagementInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editrolemanagementModel := models.EditRoleManagementResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_role(?,?,?,?,?,?)", editrolemanagementInput.Get_role_id, editrolemanagementInput.Get_role_name, editrolemanagementInput.Get_role_desc, editrolemanagementInput.Get_menu_comp_name, editrolemanagementInput.Get_menu_comp_desc, editrolemanagementInput.Get_menu_desc).Scan(&editrolemanagementModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Updated Successfully",
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

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allHeirarchyInput body models.AllHeirarchyRequest true "AllHeirarchy Input"
// @Success		  	200 {object} models.AllHeirarchyResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_heirarchy/ [post]
func SelectHeirarchybyID(c *fiber.Ctx) error {
	hierarchyInput := models.AllHeirarchyRequest{}

	if parErr := c.BodyParser(&hierarchyInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allheirarchyModel := []models.AllHeirarchyResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_heirarchy(?)", hierarchyInput.Get_id).Scan(&allheirarchyModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allheirarchyModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allheirarchyModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	edithierarchyInput body models.EditHierarchyRequest true "EditHierarchy Input"
// @Success		  	200 {object} models.EditHeirarchyResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_hierarchy/ [post]
func EditHierarchy(c *fiber.Ctx) error {
	edithierarchyInput := models.EditHierarchyRequest{}

	if parErr := c.BodyParser(&edithierarchyInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	edithierarchyModel := models.EditHeirarchyResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_hierarchy(?,?,?,?,?,?,?)", edithierarchyInput.Get_hierarchy_id, edithierarchyInput.Get_branch_code, edithierarchyInput.Get_branch_desc, edithierarchyInput.Get_unit_code, edithierarchyInput.Get_unit_desc, edithierarchyInput.Get_center_code, edithierarchyInput.Get_center_desc).Scan(&edithierarchyModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Updated Successfully",
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	drophierarchyInput body models.DropHierarchyRequest true "DropHierarchy Input"
// @Success		  	200 {object} models.DropHeirarchyResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_hierarchy/ [post]
func DropHierarchy(c *fiber.Ctx) error {
	drophierarchyInput := models.DropHierarchyRequest{}

	if parErr := c.BodyParser(&drophierarchyInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	drophierarchyModel := models.DropHeirarchyResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_hierarchy(?,?)", drophierarchyInput.Drop_id, drophierarchyInput.Drop_hierarchy).Scan(&drophierarchyModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Updated Successfully",
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	createHeirarchyInput body models.CreateHeirarchyRequest true "CreateHeirarchy Input"
// @Success		  	200 {object} models.CreateHeirarchyResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_heirarchy/ [post]
func CreateHeirarchy(c *fiber.Ctx) error {
	var createInput models.CreateHeirarchyRequest

	if parErr := c.BodyParser(&createInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "401",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	if createInput.Set_created_date == "" {
		createInput.Set_created_date = time.Now().Format("Jan 2, 2006")
	}

	createModel := models.CreateHeirarchyResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_heirarchy(?,?,?,?,?,?,?,?)", createInput.Set_branch_code, createInput.Set_branch_desc, createInput.Set_unit_code, createInput.Set_unit_desc, createInput.Set_center_code, createInput.Set_center_desc, createInput.Set_created_by, createInput.Set_created_date).Scan(&createModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
		RetCode: "200",
		Message: "Created Successfully",
	})
}
