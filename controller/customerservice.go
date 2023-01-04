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

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_csr_hotline").Find(&csrhotlineModel, csrhotlineInput).Error; dbErr != nil {
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
// @Param       	allCsrHotlineInput body models.AllCsrHotlineRequest true "AllCsrHotline Input"
// @Success		  	200 {object} models.AllCsrHotlineResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_csrhotline/ [post]
func SelectCsrHotlinebyID(c *fiber.Ctx) error {
	csrhotlineInput := models.AllCsrHotlineRequest{}

	if parErr := c.BodyParser(&csrhotlineInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allcsrhotlineModel := []models.AllCsrHotlineResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_csrhotline(?)", csrhotlineInput.Get_id).Scan(&allcsrhotlineModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allcsrhotlineModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allcsrhotlineModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editCsrHotlineInput body models.EditCsrHotlineRequest true "EditCsrHotline Input"
// @Success		  	200 {object} models.EditCsrHotlineResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_csrhotline/ [post]
func EditCsrHotline(c *fiber.Ctx) error {
	editcsrhotlineInput := models.EditCsrHotlineRequest{}

	if parErr := c.BodyParser(&editcsrhotlineInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editcsrhotlineModel := models.EditCsrHotlineResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_csrhotline(?,?,?,?)", editcsrhotlineInput.Get_id, editcsrhotlineInput.Get_contact_number, editcsrhotlineInput.Get_network_provider, editcsrhotlineInput.Get_inst_code).Scan(&editcsrhotlineModel).Error; dbErr != nil {
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
// @Param       	dropcsrhotlineInput body models.DropCsrHotlineRequest true "DropCsrHotlineInput"
// @Success		  	200 {object} models.DropCsrHotlineResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_csrhotline/ [post]
func DropCsrHotline(c *fiber.Ctx) error {
	dropcsrhotlineInput := models.DropCsrHotlineRequest{}

	if parErr := c.BodyParser(&dropcsrhotlineInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropcsrhotlineModel := models.DropCsrHotlineResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_csrhotline(?,?)", dropcsrhotlineInput.Drop_id, dropcsrhotlineInput.Drop_csrhotline).Scan(&dropcsrhotlineModel).Error; dbErr != nil {
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
// @Param       	allConcernTypeInput body models.AllConcernTypeRequest true "AllConcernType Input"
// @Success		  	200 {object} models.AllConcernTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_concerntype/ [post]
func SelectConcernTypebyID(c *fiber.Ctx) error {
	concerntypeInput := models.AllConcernTypeRequest{}

	if parErr := c.BodyParser(&concerntypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allconcerntypeModel := []models.AllConcernTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_concerntype(?)", concerntypeInput.Get_id).Scan(&allconcerntypeModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allconcerntypeModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allconcerntypeModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editConcernTypeInput body models.EditConcernTypeRequest true "EditConcernType Input"
// @Success		  	200 {object} models.EditConcernTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_concerntype/ [post]
func EditConcernType(c *fiber.Ctx) error {
	editconcerntypeInput := models.EditConcernTypeRequest{}

	if parErr := c.BodyParser(&editconcerntypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editconcerntypeModel := models.EditConcernTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_concerntype(?,?,?,?,?)", editconcerntypeInput.Get_Concern_code, editconcerntypeInput.Get_Concern_name, editconcerntypeInput.Get_Concern_desc, editconcerntypeInput.Get_Concern_time, editconcerntypeInput.Get_Concern_level).Scan(&editconcerntypeModel).Error; dbErr != nil {
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
// @Param       	dropconcerntypeInput body models.DropConcernTypeRequest true "DropConcernTypeInput"
// @Success		  	200 {object} models.DropConcernTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_concerntype/ [post]
func DropConcernType(c *fiber.Ctx) error {
	dropconcerntypeInput := models.DropConcernTypeRequest{}

	if parErr := c.BodyParser(&dropconcerntypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropconcerntypeModel := models.DropConcernTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_concerntype(?,?)", dropconcerntypeInput.Drop_id, dropconcerntypeInput.Drop_concerntype).Scan(&dropconcerntypeModel).Error; dbErr != nil {
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

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allBroadcastSmsInput body models.AllBroadcastSmsRequest true "AllBroadcastSms Input"
// @Success		  	200 {object} models.AllBroadcastSmsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_broadcastsms/ [post]
func SelectBroadcastSmsbyID(c *fiber.Ctx) error {
	splabroadcastsmsInput := models.AllBroadcastSmsRequest{}

	if parErr := c.BodyParser(&splabroadcastsmsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allbroadcastsmsModel := []models.AllBroadcastSmsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_broadcastsms(?)", splabroadcastsmsInput.Get_id).Scan(&allbroadcastsmsModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allbroadcastsmsModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allbroadcastsmsModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editBroadcastSmsInput body models.EditBroadcastSmsRequest true "EditBroadcastSms Input"
// @Success		  	200 {object} models.EditBroadcastSmsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_broadcastsms/ [post]
func EditBroadcastSms(c *fiber.Ctx) error {
	editbroadcastsmsInput := models.EditBroadcastSmsRequest{}

	if parErr := c.BodyParser(&editbroadcastsmsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editbroadcastsmsModel := models.EditBroadcastSmsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_broadcastsms(?,?,?,?,?,?,?)", editbroadcastsmsInput.Get_inbox_id, editbroadcastsmsInput.Get_subject, editbroadcastsmsInput.Get_period_start, editbroadcastsmsInput.Get_period_end, editbroadcastsmsInput.Get_inbox_desc, editbroadcastsmsInput.Get_client_type, editbroadcastsmsInput.Get_branch_code).Scan(&editbroadcastsmsModel).Error; dbErr != nil {
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
// @Param       	dropbroadcastsmsInput body models.DropBroadcastSmsRequest true "DropBroadcastSmsInput"
// @Success		  	200 {object} models.DropBroadcastSmsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_broadcastsms/ [post]
func DropBroadcastSms(c *fiber.Ctx) error {
	dropbroadcastsmsInput := models.DropBroadcastSmsRequest{}

	if parErr := c.BodyParser(&dropbroadcastsmsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropbroadcastsmsModel := models.DropBroadcastSmsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_broadcastsms(?,?)", dropbroadcastsmsInput.Drop_id, dropbroadcastsmsInput.Drop_broadcastsms).Scan(&dropbroadcastsmsModel).Error; dbErr != nil {
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
