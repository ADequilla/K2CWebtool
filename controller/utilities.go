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
// @Param       	feestructureInput body models.FeeStructureRequest true "FeeStructure Input"
// @Success		  	200 {object} models.FeeStructureResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_feestructure/ [post]
func GetFeeStructure(c *fiber.Ctx) error {
	feestructureInput := models.FeeStructureRequest{}

	if parErr := c.BodyParser(&feestructureInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	feestructureModel := []models.FeeStructureResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_feestracture").Find(&feestructureModel, feestructureInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(feestructureModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    feestructureModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allFeestructureInput body models.AllFeestructureRequest true "AllFeestructure Input"
// @Success		  	200 {object} models.AllFeestructureResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_feestructure/ [post]
func SelectFeeStructurebyID(c *fiber.Ctx) error {
	feestructureInput := models.AllFeestructureRequest{}

	if parErr := c.BodyParser(&feestructureInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allfeestructureModel := []models.AllFeestructureResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_feestructure(?)", feestructureInput.Get_id).Scan(&allfeestructureModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allfeestructureModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allfeestructureModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editfeestructureInput body models.EditFeeStructureRequest true "EditFeeStructure Input"
// @Success		  	200 {object} models.EditFeeStructureResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_feestructure/ [post]
func EditFeeStructure(c *fiber.Ctx) error {
	editfeestructureInput := models.EditFeeStructureRequest{}

	if parErr := c.BodyParser(&editfeestructureInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editfeestructureModel := models.EditFeeStructureResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_feestructure(?,?,?,?,?,?,?,?,?,?)", editfeestructureInput.Get_fee_id, editfeestructureInput.Get_client_type, editfeestructureInput.Get_trans_type, editfeestructureInput.Get_start_range, editfeestructureInput.Get_end_range, editfeestructureInput.Get_total_charge, editfeestructureInput.Get_agent_income, editfeestructureInput.Get_bank_income, editfeestructureInput.Get_agent_target_income, editfeestructureInput.Get_bancnet_income).Scan(&editfeestructureModel).Error; dbErr != nil {
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
// @Param       	dropfeestructureInput body models.DropFeeStructureRequest true "DropFeeStructure Input"
// @Success		  	200 {object} models.DropFeeStructureResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_feestructure/ [post]
func DropFeeStructure(c *fiber.Ctx) error {
	dropfeestructureInput := models.DropFeeStructureRequest{}

	if parErr := c.BodyParser(&dropfeestructureInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropfeestructureModel := models.DropFeeStructureResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_feestructure(?,?)", dropfeestructureInput.Drop_id, dropfeestructureInput.Drop_feestructure).Scan(&dropfeestructureModel).Error; dbErr != nil {
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
// @Param       	createFeeStructureInput body models.CreateFeeStructureRequest true "FeeStructure Input"
// @Success		  	200 {object} models.CreateFeeStructureResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_feestructure/ [post]
func CreateFeeStructure(c *fiber.Ctx) error {
	createInput := models.CreateFeeStructureRequest{}

	if parErr := c.BodyParser(&createInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "401",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	createInput.Set_total_charge = createInput.Set_agent_income + createInput.Set_bank_income + createInput.Set_agent_target_income

	if createInput.Set_created_date == "" {
		createInput.Set_created_date = time.Now().Format("Jan 2, 2006")
	}

	createModel := models.CreateFeeStructureResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_feestructure(?,?,?,?,?,?,?,?,?,?,?)", createInput.Set_start_range, createInput.Set_end_range, createInput.Set_total_charge, createInput.Set_agent_income, createInput.Set_bank_income, createInput.Set_created_date, createInput.Set_created_by, createInput.Set_trans_type, createInput.Set_client_type, createInput.Set_agent_target_income, createInput.Set_bancnet_income).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	paramconfigInput body models.ParamConfigRequest true "ParamConfig Input"
// @Success		  	200 {object} models.ParamConfigResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_paramconfig/ [post]
func GetParamConfig(c *fiber.Ctx) error {
	paramconfigInput := models.ParamConfigRequest{}

	if parErr := c.BodyParser(&paramconfigInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	paramconfigModel := []models.ParamConfigResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_paramconfig").Find(&paramconfigModel, paramconfigInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(paramconfigModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    paramconfigModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allParamConfigInput body models.AllParamConfigRequest true "AllParamConfig Input"
// @Success		  	200 {object} models.AllFeestructureResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_paramconfig/ [post]
func SelectParamConfigbyID(c *fiber.Ctx) error {
	paramconfigInput := models.AllParamConfigRequest{}

	if parErr := c.BodyParser(&paramconfigInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allparamconfigModel := []models.AllParamConfigResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_paramsconfig(?)", paramconfigInput.Get_id).Scan(&allparamconfigModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allparamconfigModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allparamconfigModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editparamconfigInput body models.EditParamConfigRequest true "EditParamConfig Input"
// @Success		  	200 {object} models.EditParamConfigResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_paramconfig/ [post]
func EditParamConfig(c *fiber.Ctx) error {
	editparamconfigInput := models.EditParamConfigRequest{}

	if parErr := c.BodyParser(&editparamconfigInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editparamconfigModel := models.EditParamConfigResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_paramconfig(?,?,?,?,?)", editparamconfigInput.Get_param_id, editparamconfigInput.Get_app_type, editparamconfigInput.Get_param_name, editparamconfigInput.Get_param_value, editparamconfigInput.Get_param_desc).Scan(&editparamconfigModel).Error; dbErr != nil {
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
// @Param       	createParamConfigInput body models.CreateParamConfigRequest true "ParamConfig Input"
// @Success		  	200 {object} models.CreateParamConfigResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_paramconfig/ [post]
func CreateParamConfig(c *fiber.Ctx) error {
	createInput := models.CreateParamConfigRequest{}

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

	createModel := models.CreateParamConfigResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_paramconfig(?,?,?,?,?,?)", createInput.Set_param_type, createInput.Set_param_name, createInput.Set_param_value, createInput.Set_param_desc, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	atmlocInput body models.AtmLocRequest true "AtmLoc Input"
// @Success		  	200 {object} models.AtmLocResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_atmloc/ [post]
func GetAtmLoc(c *fiber.Ctx) error {
	atmlocInput := models.AtmLocRequest{}

	if parErr := c.BodyParser(&atmlocInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	atmlocModel := []models.AtmLocResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_atmloc").Find(&atmlocModel, atmlocInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(atmlocModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    atmlocModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allAtmLocInput body models.AllAtmLocRequest true "AllAtmLoc Input"
// @Success		  	200 {object} models.AllAtmLocResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_atmloc/ [post]
func SelectAtmLocbyID(c *fiber.Ctx) error {
	atmlocInput := models.AllAtmLocRequest{}

	if parErr := c.BodyParser(&atmlocInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allatmlocModel := []models.AllAtmLocResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_atmloc(?)", atmlocInput.Get_id).Scan(&allatmlocModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allatmlocModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allatmlocModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editatmlocInput body models.EditAtmLocRequest true "EditAtmLoc Input"
// @Success		  	200 {object} models.EditAtmLocResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_atmloc/ [post]
func EditAtmLoc(c *fiber.Ctx) error {
	editatmlocInput := models.EditAtmLocRequest{}

	if parErr := c.BodyParser(&editatmlocInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editatmlocModel := models.EditAtmLocResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_atmloc(?,?,?,?,?,?)", editatmlocInput.Get_atm_id, editatmlocInput.Get_atm_description, editatmlocInput.Get_atm_address, editatmlocInput.Get_atm_city, editatmlocInput.Get_atm_longitude, editatmlocInput.Get_atm_latitude).Scan(&editatmlocModel).Error; dbErr != nil {
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
// @Param       	dropatmlocInput body models.DropAtmLocRequest true "DropAtmLoc Input"
// @Success		  	200 {object} models.DropAtmLocResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_atmloc/ [post]
func DropAtmLoc(c *fiber.Ctx) error {
	dropatmlocInput := models.DropAtmLocRequest{}

	if parErr := c.BodyParser(&dropatmlocInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropatmlocModel := models.DropAtmLocResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_atmloc(?,?)", dropatmlocInput.Drop_id, dropatmlocInput.Drop_atmloc).Scan(&dropatmlocModel).Error; dbErr != nil {
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
// @Param       	createAtmLocInput body models.CreateAtmLocRequest true "AtmLoc Input"
// @Success		  	200 {object} models.CreateAtmLocResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_atmloc/ [post]
func CreateAtmLoc(c *fiber.Ctx) error {
	createInput := models.CreateAtmLocRequest{}

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

	createModel := models.CreateAtmLocResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_atmloc(?,?,?,?,?,?,?,?)", createInput.Set_atm_description, createInput.Set_atm_address, createInput.Set_atm_city, createInput.Set_atm_longitude, createInput.Set_atm_latitude, createInput.Set_insti_desc, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	productservicesInput body models.ProductServicesRequest true "ProductServices Input"
// @Success		  	200 {object} models.ProductServicesResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_productservices/ [post]
func GetProductServices(c *fiber.Ctx) error {
	productservicesInput := models.ProductServicesRequest{}

	if parErr := c.BodyParser(&productservicesInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	productservicesModel := []models.ProductServicesResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_productservice").Find(&productservicesModel, productservicesInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(productservicesModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    productservicesModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allProductServicesInput body models.AllProductServicesRequest true "AllProductServices Input"
// @Success		  	200 {object} models.AllProductServicesResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_productservices/ [post]
func SelectProductServicesbyID(c *fiber.Ctx) error {
	productservicesInput := models.AllProductServicesRequest{}

	if parErr := c.BodyParser(&productservicesInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allproductservicesModel := []models.AllProductServicesResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_productservices(?)", productservicesInput.Get_id).Scan(&allproductservicesModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allproductservicesModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allproductservicesModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editproductservicesInput body models.EditProductServicesRequest true "EditProductServices Input"
// @Success		  	200 {object} models.EditProductServicesResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_productservices/ [post]
func EditProductServices(c *fiber.Ctx) error {
	editproductservicesInput := models.EditProductServicesRequest{}

	if parErr := c.BodyParser(&editproductservicesInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editproductservicesModel := models.EditProductServicesResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_productservices(?,?,?,?,?)", editproductservicesInput.Get_service_id, editproductservicesInput.Get_service_name, editproductservicesInput.Get_service_description, editproductservicesInput.Get_show, editproductservicesInput.Get_service_banner).Scan(&editproductservicesModel).Error; dbErr != nil {
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
// @Param       	dropproductservicesInput body models.DropProductServicesRequest true "DropProductServices Input"
// @Success		  	200 {object} models.DropProductServicesResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_productservices/ [post]
func DropProductServices(c *fiber.Ctx) error {
	dropproductservicesInput := models.DropProductServicesRequest{}

	if parErr := c.BodyParser(&dropproductservicesInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropproductservicesModel := models.DropProductServicesResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_productservices(?,?)", dropproductservicesInput.Drop_id, dropproductservicesInput.Drop_productservices).Scan(&dropproductservicesModel).Error; dbErr != nil {
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
// @Param       	createProductServicesInput body models.CreateProductServicesRequest true "ProductServices Input"
// @Success		  	200 {object} models.CreateProductServicesResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_productservices/ [post]
func CreateProductServices(c *fiber.Ctx) error {
	createInput := models.CreateProductServicesRequest{}

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

	createModel := models.CreateProductServicesResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_productandservices(?,?,?,?,?,?)", createInput.Set_service_name, createInput.Set_service_description, createInput.Set_service_banner, createInput.Set_created_date, createInput.Set_created_by, createInput.Set_show).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	servicedowntimeInput body models.ServiceDowntimeRequest true "ServiceDowntime Input"
// @Success		  	200 {object} models.ServiceDowntimeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_servicedowntime/ [post]
func GetServiceDowntime(c *fiber.Ctx) error {
	servicedowntimeInput := models.ServiceDowntimeRequest{}

	if parErr := c.BodyParser(&servicedowntimeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	servicedowntimeModel := []models.ServiceDowntimeResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_servicedowntime").Find(&servicedowntimeModel, servicedowntimeInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(servicedowntimeModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    servicedowntimeModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allServiceDowntimeInput body models.AllServiceDowntimeRequest true "AllServiceDowntime Input"
// @Success		  	200 {object} models.AllServiceDowntimeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_servicedowntime/ [post]
func SelectServiceDowntimebyID(c *fiber.Ctx) error {
	servicedowntimeInput := models.AllServiceDowntimeRequest{}

	if parErr := c.BodyParser(&servicedowntimeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allservicedowntimeModel := []models.AllServiceDowntimeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_servicedowntime(?)", servicedowntimeInput.Get_id).Scan(&allservicedowntimeModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allservicedowntimeModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allservicedowntimeModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editservicedowntimeInput body models.EditServiceDowntimeRequest true "EditDowntime Input"
// @Success		  	200 {object} models.EditServiceDowntimeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_servicedowntime/ [post]
func EditServiceDowntime(c *fiber.Ctx) error {
	editservicedowntimeInput := models.EditServiceDowntimeRequest{}

	if parErr := c.BodyParser(&editservicedowntimeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editservicedowntimeModel := models.EditServiceDowntimeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_servicedowntime(?,?,?,?,?)", editservicedowntimeInput.Get_downtime_id, editservicedowntimeInput.Get_downtime_start, editservicedowntimeInput.Get_downtime_end, editservicedowntimeInput.Get_client_type, editservicedowntimeInput.Get_downtime_desc).Scan(&editservicedowntimeModel).Error; dbErr != nil {
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
// @Param       	dropservicedowntimeInput body models.DropServiceDowntimeRequest true "DropServiceDowntime Input"
// @Success		  	200 {object} models.DropServiceDowntimeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_servicedowntime/ [post]
func DropServiceDowntime(c *fiber.Ctx) error {
	dropservicedowntimeInput := models.DropServiceDowntimeRequest{}

	if parErr := c.BodyParser(&dropservicedowntimeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropservicedowntimeModel := models.DropServiceDowntimeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_servicedowntime(?,?)", dropservicedowntimeInput.Drop_id, dropservicedowntimeInput.Drop_servicedowntime).Scan(&dropservicedowntimeModel).Error; dbErr != nil {
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
// @Param       	createServiceDowntimeInput body models.CreateServiceDowntimeRequest true "ServiceDowntime Input"
// @Success		  	200 {object} models.CreateServiceDowntimeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_servicedowntime/ [post]
func CreateServiceDowntime(c *fiber.Ctx) error {
	createInput := models.CreateServiceDowntimeRequest{}

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

	createModel := models.CreateServiceDowntimeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_servicedowntime(?,?,?,?,?,?)", createInput.Set_downtime_start, createInput.Set_downtime_end, createInput.Set_downtime_desc, createInput.Set_created_date, createInput.Set_created_by, createInput.Set_client_type).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	backnewsInput body models.BankNewsRequest true "BankNews Input"
// @Success		  	200 {object} models.BankNewsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_banknews/ [post]
func GetBankNews(c *fiber.Ctx) error {
	banknewsInput := models.BankNewsRequest{}

	if parErr := c.BodyParser(&banknewsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	banknewsModel := []models.BankNewsResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_banknew").Find(&banknewsModel, banknewsInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(banknewsModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    banknewsModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allBankNewsInput body models.AllBankNewsRequest true "AllBankNews Input"
// @Success		  	200 {object} models.AllBankNewsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_banknews/ [post]
func SelectBankNewsbyID(c *fiber.Ctx) error {
	banknewsInput := models.AllBankNewsRequest{}

	if parErr := c.BodyParser(&banknewsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allbanknewsModel := []models.AllBankNewsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_banknews(?)", banknewsInput.Get_id).Scan(&allbanknewsModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allbanknewsModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allbanknewsModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editbanknewsnput body models.EditBankNewsRequest true "EditBankNews Input"
// @Success		  	200 {object} models.EditBankNewsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_banknews/ [post]
func EditBankNews(c *fiber.Ctx) error {
	editbanknewsInput := models.EditBankNewsRequest{}

	if parErr := c.BodyParser(&editbanknewsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editbanknewsModel := models.EditBankNewsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_banknews(?,?,?,?,?,?)", editbanknewsInput.Get_product_id, editbanknewsInput.Get_product_name, editbanknewsInput.Get_product_description, editbanknewsInput.Get_product_periode_start, editbanknewsInput.Get_product_periode_end, editbanknewsInput.Get_product_img_name).Scan(&editbanknewsModel).Error; dbErr != nil {
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
// @Param       	dropbanknewsInput body models.DropBankNewsRequest true "DropBankNews Input"
// @Success		  	200 {object} models.DropBankNewsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_banknews/ [post]
func DropBankNews(c *fiber.Ctx) error {
	dropbanknewsInput := models.DropBankNewsRequest{}

	if parErr := c.BodyParser(&dropbanknewsInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropbanknewsModel := models.DropBankNewsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_banknews(?,?)", dropbanknewsInput.Drop_id, dropbanknewsInput.Drop_banknews).Scan(&dropbanknewsModel).Error; dbErr != nil {
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
// @Param       	createBankNewsInput body models.CreateBankNewsRequest true "BankNews Input"
// @Success		  	200 {object} models.CreateBankNewsResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_banknews/ [post]
func CreateBankNews(c *fiber.Ctx) error {
	createInput := models.CreateBankNewsRequest{}

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

	createModel := models.CreateBankNewsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_banknews(?,?,?,?,?,?,?)", createInput.Set_product_name, createInput.Set_product_description, createInput.Set_product_periode_start, createInput.Set_product_periode_end, createInput.Set_product_img_name, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	institutionInput body models.InstitutionRequest true "Institution Input"
// @Success		  	200 {object} models.InstitutionResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_insti/ [post]
func GetInstitution(c *fiber.Ctx) error {
	institutionInput := models.InstitutionRequest{}

	if parErr := c.BodyParser(&institutionInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	institutionModel := []models.InstitutionResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_insti").Find(&institutionModel, institutionInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(institutionModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    institutionModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allInstiInput body models.AllInstiRequest true "AllInsti Input"
// @Success		  	200 {object} models.AllInstiResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_insti/ [post]
func SelectInstibyID(c *fiber.Ctx) error {
	instiInput := models.AllInstiRequest{}

	if parErr := c.BodyParser(&instiInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allinstiModel := []models.AllInstiResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_insti(?)", instiInput.Get_id).Scan(&allinstiModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allinstiModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allinstiModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editinstiInput body models.EditInstiRequest true "EditInsti Input"
// @Success		  	200 {object} models.EditInstiResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_insti/ [post]
func EditInsti(c *fiber.Ctx) error {
	editinstiInput := models.EditInstiRequest{}

	if parErr := c.BodyParser(&editinstiInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editinstiModel := models.EditInstiResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_insti(?,?,?)", editinstiInput.Get_inst_id, editinstiInput.Get_inst_code, editinstiInput.Get_inst_desc).Scan(&editinstiModel).Error; dbErr != nil {
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
// @Param       	dropinstiInput body models.DropInstiRequest true "DropInsti Input"
// @Success		  	200 {object} models.DropInstiResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_insti/ [post]
func DropInstitution(c *fiber.Ctx) error {
	dropinstiInput := models.DropInstiRequest{}

	if parErr := c.BodyParser(&dropinstiInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropInstiModel := models.DropInstiResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_insti(?,?)", dropinstiInput.Drop_id, dropinstiInput.Drop_insti).Scan(&dropInstiModel).Error; dbErr != nil {
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
// @Param       	createInstiInput body models.CreateInstiRequest true "Insti Input"
// @Success		  	200 {object} models.CreateInstiResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_insti/ [post]
func CreateInsti(c *fiber.Ctx) error {
	createInput := models.CreateInstiRequest{}

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

	createModel := models.CreateInstiResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_insti(?,?,?,?)", createInput.Set_inst_code, createInput.Set_inst_desc, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	branchInput body models.BranchRequest true "Branch Input"
// @Success		  	200 {object} models.BranchResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_branch/ [post]
func GetBranch(c *fiber.Ctx) error {
	branchInput := models.BranchRequest{}

	if parErr := c.BodyParser(&branchInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	branchModel := []models.BranchResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_branch").Find(&branchModel, branchInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(branchModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    branchModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allBranchInput body models.AllBranchRequest true "AllBranch Input"
// @Success		  	200 {object} models.AllBranchResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_branch/ [post]
func SelectBranchbyID(c *fiber.Ctx) error {
	branchInput := models.AllBranchRequest{}

	if parErr := c.BodyParser(&branchInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allbranchModel := []models.AllBranchResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_branch(?)", branchInput.Get_id).Scan(&allbranchModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allbranchModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allbranchModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editbranchInput body models.EditBranchRequest true "EditBranch Input"
// @Success		  	200 {object} models.EditBranchResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_branch/ [post]
func EditBranch(c *fiber.Ctx) error {
	editbranchInput := models.EditBranchRequest{}

	if parErr := c.BodyParser(&editbranchInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editbranchModel := models.EditBranchResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_branch(?,?,?)", editbranchInput.Get_branch_id, editbranchInput.Get_branch_code, editbranchInput.Get_branch_desc).Scan(&editbranchModel).Error; dbErr != nil {
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
// @Param       	dropbranchInput body models.DropBranchRequest true "DropBranch Input"
// @Success		  	200 {object} models.DropBranchResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_branch/ [post]
func DropBranch(c *fiber.Ctx) error {
	dropbranchInput := models.DropBranchRequest{}

	if parErr := c.BodyParser(&dropbranchInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropbranchModel := models.DropBranchResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_branch(?,?)", dropbranchInput.Drop_id, dropbranchInput.Drop_branch).Scan(&dropbranchModel).Error; dbErr != nil {
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
// @Param       	createBranchInput body models.CreateBranchRequest true "Branch Input"
// @Success		  	200 {object} models.CreateBranchResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_branch/ [post]
func CreateBranch(c *fiber.Ctx) error {
	createInput := models.CreateBranchRequest{}

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

	createModel := models.CreateBranchResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_branch(?,?,?,?)", createInput.Set_branch_code, createInput.Set_branch_desc, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	unitInput body models.UnitRequest true "Unit Input"
// @Success		  	200 {object} models.UnitResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_unit/ [post]
func GetUnit(c *fiber.Ctx) error {
	unitInput := models.UnitRequest{}

	if parErr := c.BodyParser(&unitInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	unitModel := []models.UnitResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_unit").Find(&unitModel, unitInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(unitModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    unitModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allUnitInput body models.AllUnitRequest true "AllUnit Input"
// @Success		  	200 {object} models.AllUnitResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_unit/ [post]
func SelectUnitbyID(c *fiber.Ctx) error {
	unitInput := models.AllUnitRequest{}

	if parErr := c.BodyParser(&unitInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allunitModel := []models.AllUnitResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_unit(?)", unitInput.Get_id).Scan(&allunitModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allunitModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allunitModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editunitInput body models.EditUnitRequest true "EditUnit Input"
// @Success		  	200 {object} models.EditUnitResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_unit/ [post]
func EditUnit(c *fiber.Ctx) error {
	editunitInput := models.EditUnitRequest{}

	if parErr := c.BodyParser(&editunitInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editunitModel := models.EditUnitResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_unit(?,?,?)", editunitInput.Get_unit_id, editunitInput.Get_unit_code, editunitInput.Get_unit_desc).Scan(&editunitModel).Error; dbErr != nil {
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
// @Param       	dropunitInput body models.DropUnitRequest true "DropUnit Input"
// @Success		  	200 {object} models.DropUnitResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_unit/ [post]
func DropUnit(c *fiber.Ctx) error {
	dropunitInput := models.DropUnitRequest{}

	if parErr := c.BodyParser(&dropunitInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropunitModel := models.DropUnitResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_unit(?,?)", dropunitInput.Drop_id, dropunitInput.Drop_unit).Scan(&dropunitModel).Error; dbErr != nil {
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
// @Param       	createUnitInput body models.CreateUnitRequest true "Unit Input"
// @Success		  	200 {object} models.CreateUnitResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_unit/ [post]
func CreateUnit(c *fiber.Ctx) error {
	createInput := models.CreateUnitRequest{}

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

	createModel := models.CreateUnitResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_unit(?,?,?,?)", createInput.Set_unit_code, createInput.Set_unit_desc, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	centerInput body models.CenterRequest true "Center Input"
// @Success		  	200 {object} models.CenterResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_center/ [post]
func GetCenter(c *fiber.Ctx) error {
	centerInput := models.CenterRequest{}

	if parErr := c.BodyParser(&centerInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	centerModel := []models.CenterResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_center").Find(&centerModel, centerInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(centerModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    centerModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allCenterInput body models.AllCenterRequest true "AllCenter Input"
// @Success		  	200 {object} models.AllCenterResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_center/ [post]
func SelectCenterbyID(c *fiber.Ctx) error {
	centerInput := models.AllCenterRequest{}

	if parErr := c.BodyParser(&centerInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allcenterModel := []models.AllCenterResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_center(?)", centerInput.Get_id).Scan(&allcenterModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allcenterModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allcenterModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editCenterInput body models.EditCenterRequest true "EditCenter Input"
// @Success		  	200 {object} models.EditCenterResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_center/ [post]
func EditCenter(c *fiber.Ctx) error {
	editcenterInput := models.EditCenterRequest{}

	if parErr := c.BodyParser(&editcenterInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editcenterModel := models.EditCenterResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_center(?,?,?)", editcenterInput.Get_center_id, editcenterInput.Get_center_code, editcenterInput.Get_center_desc).Scan(&editcenterModel).Error; dbErr != nil {
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
// @Param       	dropcenterInput body models.DropCenterRequest true "DropCenter Input"
// @Success		  	200 {object} models.DropCenterResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_center/ [post]
func DropCenter(c *fiber.Ctx) error {
	dropcenterInput := models.DropCenterRequest{}

	if parErr := c.BodyParser(&dropcenterInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropcenterModel := models.DropCenterResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_center(?,?)", dropcenterInput.Drop_id, dropcenterInput.Drop_center).Scan(&dropcenterModel).Error; dbErr != nil {
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
// @Param       	createCenterInput body models.CreateCenterRequest true "Center Input"
// @Success		  	200 {object} models.CreateCenterResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_center/ [post]
func CreateCenter(c *fiber.Ctx) error {
	createInput := models.CreateCenterRequest{}

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

	createModel := models.CreateCenterResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_center(?,?,?,?)", createInput.Set_center_code, createInput.Set_center_desc, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	providerInput body models.ProviderRequest true "Provider Input"
// @Success		  	200 {object} models.ProviderResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_provider/ [post]
func GetProvider(c *fiber.Ctx) error {
	providerInput := models.ProviderRequest{}

	if parErr := c.BodyParser(&providerInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	providerModel := []models.ProviderResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_provider").Find(&providerModel, providerInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(providerModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    providerModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allProviderInput body models.AllProviderRequest true "AllProvider Input"
// @Success		  	200 {object} models.AllProviderResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_provider/ [post]
func SelectProviderbyID(c *fiber.Ctx) error {
	providerInput := models.AllProviderRequest{}

	if parErr := c.BodyParser(&providerInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allproviderModel := []models.AllProviderResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_provider(?)", providerInput.Get_id).Scan(&allproviderModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allproviderModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allproviderModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editProviderInput body models.EditProviderRequest true "EditProvider Input"
// @Success		  	200 {object} models.EditProviderResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_provider/ [post]
func EditProvider(c *fiber.Ctx) error {
	editproviderInput := models.EditProviderRequest{}

	if parErr := c.BodyParser(&editproviderInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editproviderModel := models.EditProviderResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_provider(?,?,?,?,?)", editproviderInput.Get_provider_id, editproviderInput.Get_provider_name, editproviderInput.Get_description, editproviderInput.Get_provider_alias, editproviderInput.Get_status).Scan(&editproviderModel).Error; dbErr != nil {
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
// @Param       	dropproviderInput body models.DropProviderRequest true "DropProviderInput"
// @Success		  	200 {object} models.DropProviderResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_provider/ [post]
func DropProvider(c *fiber.Ctx) error {
	dropproviderInput := models.DropProviderRequest{}

	if parErr := c.BodyParser(&dropproviderInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropproviderModel := models.DropProviderResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_provider(?,?)", dropproviderInput.Drop_id, dropproviderInput.Drop_provider).Scan(&dropproviderModel).Error; dbErr != nil {
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
// @Param       	createProviderInput body models.CreateProviderRequest true "Provider Input"
// @Success		  	200 {object} models.CreateProviderResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_provider/ [post]
func CreateProvider(c *fiber.Ctx) error {
	createInput := models.CreateProviderRequest{}

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

	createModel := models.CreateProviderResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_provider(?,?,?,?,?,?)", createInput.Set_provider_name, createInput.Set_description, createInput.Set_provider_alias, createInput.Set_status, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	producttypeInput body models.ProductTypeRequest true "ProductType Input"
// @Success		  	200 {object} models.ProductTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_producttype/ [post]
func GetProductType(c *fiber.Ctx) error {
	producttypeInput := models.ProductTypeRequest{}

	if parErr := c.BodyParser(&producttypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	producttypeModel := []models.ProductTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_product_type").Find(&producttypeModel, producttypeInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(producttypeModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    producttypeModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allProductTypeInput body models.AllProductTypeRequest true "AllProductType Input"
// @Success		  	200 {object} models.AllProductTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_producttype/ [post]
func SelectProductTypebyID(c *fiber.Ctx) error {
	producttypeInput := models.AllProductTypeRequest{}

	if parErr := c.BodyParser(&producttypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allproducttypeModel := []models.AllProductTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_producttype(?)", producttypeInput.Get_id).Scan(&allproducttypeModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allproducttypeModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allproducttypeModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editProductTypeInput body models.EditProductTypeRequest true "EditProductType Input"
// @Success		  	200 {object} models.EditProductTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_producttype/ [post]
func EditProductType(c *fiber.Ctx) error {
	editproducttypeInput := models.EditProductTypeRequest{}

	if parErr := c.BodyParser(&editproducttypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editproducttypeModel := models.EditProductTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_producttype(?,?,?,?,?,?)", editproducttypeInput.Get_producttype_id, editproducttypeInput.Get_provider_name, editproducttypeInput.Get_product_type_id, editproducttypeInput.Get_product_type_name, editproducttypeInput.Get_description, editproducttypeInput.Get_status).Scan(&editproducttypeModel).Error; dbErr != nil {
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
// @Param       	dropproducttypeInput body models.DropProductTypeRequest true "DropProductTypeInput"
// @Success		  	200 {object} models.DropProductTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_producttype/ [post]
func DropProductType(c *fiber.Ctx) error {
	dropproducttypeInput := models.DropProductTypeRequest{}

	if parErr := c.BodyParser(&dropproducttypeInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropproducttypeModel := models.DropProductTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_producttype(?,?)", dropproducttypeInput.Drop_id, dropproducttypeInput.Drop_producttype).Scan(&dropproducttypeModel).Error; dbErr != nil {
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
// @Param       	createProductTypeInput body models.CreateProductTypeRequest true "ProductType Input"
// @Success		  	200 {object} models.CreateProductTypeResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_product_type/ [post]
func CreateProductType(c *fiber.Ctx) error {
	createInput := models.CreateProductTypeRequest{}

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

	createModel := models.CreateProductTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_product_type(?,?,?,?,?,?,?)", createInput.Set_provider_name, createInput.Set_product_type_id, createInput.Set_product_type_name, createInput.Set_description, createInput.Set_status, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	productcategoryInput body models.ProductCategoryRequest true "ProductCategory Input"
// @Success		  	200 {object} models.ProductCategoryResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_productcategory/ [post]
func GetProductCategory(c *fiber.Ctx) error {
	productcategoryInput := models.ProductCategoryRequest{}

	if parErr := c.BodyParser(&productcategoryInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	productcategoryModel := []models.ProductCategoryResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_productcategory").Find(&productcategoryModel, productcategoryInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(productcategoryModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    productcategoryModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allProductCategoryInput body models.AllProductCategoryRequest true "AllProductCategory Input"
// @Success		  	200 {object} models.AllProductCategoryResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_productcategory/ [post]
func SelectProductCategorybyID(c *fiber.Ctx) error {
	productcategoryInput := models.AllProductCategoryRequest{}

	if parErr := c.BodyParser(&productcategoryInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allproductcategoryModel := []models.AllProductCategoryResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_productcategory(?)", productcategoryInput.Get_id).Scan(&allproductcategoryModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allproductcategoryModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allproductcategoryModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editProductCategoryInput body models.EditProductCategoryRequest true "EditProductCategory Input"
// @Success		  	200 {object} models.EditProductCategoryResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_productcategory/ [post]
func EditProductCategory(c *fiber.Ctx) error {
	editproductcategoryInput := models.EditProductCategoryRequest{}

	if parErr := c.BodyParser(&editproductcategoryInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editproducttypeModel := models.EditProductCategoryResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_productcategory(?,?,?,?,?)", editproductcategoryInput.Get_productcategory_id, editproductcategoryInput.Get_product_type_name, editproductcategoryInput.Get_product_category_id, editproductcategoryInput.Get_product_category_name, editproductcategoryInput.Get_status).Scan(&editproducttypeModel).Error; dbErr != nil {
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
// @Param       	dropproductcategoryInput body models.DropProductCategoryRequest true "DropProductCategoryInput"
// @Success		  	200 {object} models.DropProductCategoryResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_productcategory/ [post]
func DropProductCategory(c *fiber.Ctx) error {
	dropproductcategoryInput := models.DropProductCategoryRequest{}

	if parErr := c.BodyParser(&dropproductcategoryInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropproductcategoryModel := models.DropProductCategoryResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_productcategory(?,?)", dropproductcategoryInput.Drop_id, dropproductcategoryInput.Drop_productcategory).Scan(&dropproductcategoryModel).Error; dbErr != nil {
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
// @Param       	createProductCategoryInput body models.CreateProductCategoryRequest true "ProductCategory Input"
// @Success		  	200 {object} models.CreateProductCategoryResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_productcategory/ [post]
func CreateProductCategory(c *fiber.Ctx) error {
	createInput := models.CreateProductCategoryRequest{}

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

	createModel := models.CreateProductCategoryResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_productcategory(?,?,?,?,?,?,?)", createInput.Set_product_type_name, createInput.Set_product_category_id, createInput.Set_product_category_name, createInput.Set_description, createInput.Set_status, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	billerproductInput body models.BillerProductRequest true "BillerProduct Input"
// @Success		  	200 {object} models.BillerProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_billerproduct/ [post]
func GetBillerProduct(c *fiber.Ctx) error {
	billerproductInput := models.BillerProductRequest{}

	if parErr := c.BodyParser(&billerproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	billerproductModel := []models.BillerProductResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_biller_product").Find(&billerproductModel, billerproductInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(billerproductModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    billerproductModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allBillerProductInput body models.AllBillerProductRequest true "AllBillerProduct Input"
// @Success		  	200 {object} models.AllBillerProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_billerproduct/ [post]
func SelectBillerProductbyID(c *fiber.Ctx) error {
	billerproductInput := models.AllBillerProductRequest{}

	if parErr := c.BodyParser(&billerproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allbillerproductModel := []models.AllBillerProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_billerproduct(?)", billerproductInput.Get_id).Scan(&allbillerproductModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allbillerproductModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allbillerproductModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editBillerProductInput body models.EditBillerProductRequest true "EditBillerProduct Input"
// @Success		  	200 {object} models.EditBillerProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_billerproduct/ [post]
func EditBillerProduct(c *fiber.Ctx) error {
	editbillerproductInput := models.EditBillerProductRequest{}

	if parErr := c.BodyParser(&editbillerproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editbillerproductModel := models.EditBillerProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_billerproduct(?,?,?,?,?,?,?,?)", editbillerproductInput.Get_Billerproduct_id, editbillerproductInput.Get_Product_category_name, editbillerproductInput.Get_Biller_product_id, editbillerproductInput.Get_Biller_product_name, editbillerproductInput.Get_Description, editbillerproductInput.Get_Service_fee, editbillerproductInput.Get_Bank_commission, editbillerproductInput.Get_Status).Scan(&editbillerproductModel).Error; dbErr != nil {
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
// @Param       	dropbillerproductInput body models.DropBillerProductRequest true "DropBillerProductInput"
// @Success		  	200 {object} models.DropBillerProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_billerproduct/ [post]
func DropBillerProduct(c *fiber.Ctx) error {
	dropbillerproductInput := models.DropBillerProductRequest{}

	if parErr := c.BodyParser(&dropbillerproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropbillerproductModel := models.DropBillerProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_billerproduct(?,?)", dropbillerproductInput.Drop_id, dropbillerproductInput.Drop_billerproduct).Scan(&dropbillerproductModel).Error; dbErr != nil {
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
// @Param       	createBillerProductInput body models.CreateBillerProductRequest true "BillerProduct Input"
// @Success		  	200 {object} models.CreateBillerProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_billerproduct/ [post]
func CreateBillerProduct(c *fiber.Ctx) error {
	createInput := models.CreateBillerProductRequest{}

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

	createModel := models.CreateBillerProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_billerproduct(?,?,?,?,?,?,?,?,?)", createInput.Set_product_category_name, createInput.Set_biller_product_id, createInput.Set_biller_product_name, createInput.Set_description, createInput.Set_bank_commission, createInput.Set_service_fee, createInput.Set_status, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	loadproductInput body models.LoadProductRequest true "LoadProduct Input"
// @Success		  	200 {object} models.LoadProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_loadproduct/ [post]
func GetLoadProduct(c *fiber.Ctx) error {
	loadproductInput := models.LoadProductRequest{}

	if parErr := c.BodyParser(&loadproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	loadproductModel := []models.LoadProductResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_load_product").Find(&loadproductModel, loadproductInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(loadproductModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    loadproductModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allLoadProductInput body models.AllLoadProductRequest true "AllLoadProduct Input"
// @Success		  	200 {object} models.AllLoadProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_loadproduct/ [post]
func SelectLoadProductbyID(c *fiber.Ctx) error {
	loadproductInput := models.AllLoadProductRequest{}

	if parErr := c.BodyParser(&loadproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allloadproductModel := []models.AllLoadProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_loadproduct(?)", loadproductInput.Get_id).Scan(&allloadproductModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allloadproductModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allloadproductModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editLoadProductInput body models.EditLoadProductRequest true "EditLoadProduct Input"
// @Success		  	200 {object} models.EditLoadProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_loadproduct/ [post]
func EditLoadProduct(c *fiber.Ctx) error {
	editloadproductInput := models.EditLoadProductRequest{}

	if parErr := c.BodyParser(&editloadproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editloadproductModel := models.EditLoadProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_loadproduct(?,?,?,?,?,?)", editloadproductInput.Get_loadproduct_id, editloadproductInput.Get_product_category_name, editloadproductInput.Get_load_product_id, editloadproductInput.Get_load_product_name, editloadproductInput.Get_description, editloadproductInput.Get_status).Scan(&editloadproductModel).Error; dbErr != nil {
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
// @Param       	dropbillerproductInput body models.DropLoadProductRequest true "DropLoadProductInput"
// @Success		  	200 {object} models.DropLoadProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_loadproduct/ [post]
func DropLoadProduct(c *fiber.Ctx) error {
	droploadproductInput := models.DropLoadProductRequest{}

	if parErr := c.BodyParser(&droploadproductInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	droploadproductModel := models.DropLoadProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_loadproduct(?,?)", droploadproductInput.Drop_id, droploadproductInput.Drop_loadproduct).Scan(&droploadproductModel).Error; dbErr != nil {
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
// @Param       	createLoadProductInput body models.CreateLoadProductRequest true "LoadProduct Input"
// @Success		  	200 {object} models.CreateLoadProductResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_loadproduct/ [post]
func CreateLoadProduct(c *fiber.Ctx) error {
	createInput := models.CreateLoadProductRequest{}

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

	createModel := models.CreateLoadProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_loadproduct(?,?,?,?,?,?,?)", createInput.Set_product_category_name, createInput.Set_load_product_id, createInput.Set_load_product_name, createInput.Set_description, createInput.Set_status, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	commssionInput body models.CommissionRequest true "Commission Input"
// @Success		  	200 {object} models.CommissionResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_commission/ [post]
func GetCommission(c *fiber.Ctx) error {
	commissionInput := models.CommissionRequest{}

	if parErr := c.BodyParser(&commissionInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	commissionModel := []models.CommissionResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_commission").Find(&commissionModel, commissionInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(commissionModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    commissionModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allCommisionInput body models.AllCommissionRequest true "AllCommission Input"
// @Success		  	200 {object} models.AllCommissionResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_commission/ [post]
func SelectCommissionbyID(c *fiber.Ctx) error {
	commissionInput := models.AllCommissionRequest{}

	if parErr := c.BodyParser(&commissionInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allcommissionModel := []models.AllCommissionResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_commission(?)", commissionInput.Get_id).Scan(&allcommissionModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allcommissionModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allcommissionModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editCommissionInput body models.EditCommissionRequest true "EditCommission Input"
// @Success		  	200 {object} models.EditCommissionResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_commission/ [post]
func EditCommission(c *fiber.Ctx) error {
	editcommissionInput := models.EditCommissionRequest{}

	if parErr := c.BodyParser(&editcommissionInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editcommissionModel := models.EditCommissionResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_commission(?,?,?,?,?,?,?)", editcommissionInput.Get_Id, editcommissionInput.Get_Trans_type, editcommissionInput.Get_Commission_type, editcommissionInput.Get_Customer_income, editcommissionInput.Get_Agent_income, editcommissionInput.Get_Bank_income, editcommissionInput.Get_Bank_partner_income).Scan(&editcommissionModel).Error; dbErr != nil {
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
// @Param       	dropcommissionInput body models.DropCommissionRequest true "DropCommissionInput"
// @Success		  	200 {object} models.DropCommissionResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_commission/ [post]
func DropCommission(c *fiber.Ctx) error {
	dropcommissionInput := models.DropCommissionRequest{}

	if parErr := c.BodyParser(&dropcommissionInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropcommissionModel := models.DropCommissionResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_commission(?,?)", dropcommissionInput.Drop_id, dropcommissionInput.Drop_commission).Scan(&dropcommissionModel).Error; dbErr != nil {
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
// @Param       	createCommissionInput body models.CreateCommissionRequest true "Commission Input"
// @Success		  	200 {object} models.CreateCommissionResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_commission/ [post]
func CreateCommission(c *fiber.Ctx) error {
	createInput := models.CreateCommissionRequest{}

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

	createModel := models.CreateCommissionResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_commission(?,?,?,?,?,?,?,?)", createInput.Set_trans_type, createInput.Set_commission_type, createInput.Set_customer_income, createInput.Set_agent_income, createInput.Set_bank_income, createInput.Set_bank_partner_income, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	banklistInput body models.BankListRequest true "BankList Input"
// @Success		  	200 {object} models.BankListResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_banklist/ [post]
func GetBankList(c *fiber.Ctx) error {
	banklistInput := models.BankListRequest{}

	if parErr := c.BodyParser(&banklistInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	banklistModel := []models.BankListResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_banklist").Find(&banklistModel, banklistInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(banklistModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    banklistModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allBanklistInput body models.AllBanklistRequest true "AllBanklist Input"
// @Success		  	200 {object} models.AllBanklistResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_banklist/ [post]
func SelectBanklistbyID(c *fiber.Ctx) error {
	banklistInput := models.AllBanklistRequest{}

	if parErr := c.BodyParser(&banklistInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allbanklistModel := []models.AllBanklistResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_banklist(?)", banklistInput.Get_id).Scan(&allbanklistModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allbanklistModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allbanklistModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editBanklistInput body models.EditBanklistRequest true "EditBanklist Input"
// @Success		  	200 {object} models.EditBanklistResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_banklist/ [post]
func EditBanklist(c *fiber.Ctx) error {
	editbanklistnInput := models.EditBanklistRequest{}

	if parErr := c.BodyParser(&editbanklistnInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editbanklistModel := models.EditBanklistResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_banklist(?,?,?,?,?)", editbanklistnInput.Get_Banklist_id, editbanklistnInput.Get_Bank_code, editbanklistnInput.Get_Bank_name, editbanklistnInput.Get_Short_name, editbanklistnInput.Get_Bank_bic).Scan(&editbanklistModel).Error; dbErr != nil {
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
// @Param       	dropbanklistInput body models.DropBanklistRequest true "DropBanklistInput"
// @Success		  	200 {object} models.DropBanklistResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_banklist/ [post]
func DropBanklist(c *fiber.Ctx) error {
	dropbanklistInput := models.DropBanklistRequest{}

	if parErr := c.BodyParser(&dropbanklistInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropbanklistModel := models.DropBanklistResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_banklist(?,?)", dropbanklistInput.Drop_id, dropbanklistInput.Drop_banklist).Scan(&dropbanklistModel).Error; dbErr != nil {
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
// @Param       	createBankListInput body models.CreateBankListRequest true "BankList Input"
// @Success		  	200 {object} models.CreateBankListResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_bank/ [post]
func CreateBankList(c *fiber.Ctx) error {
	createInput := models.CreateBankListRequest{}

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

	createModel := models.CreateBankListResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_bank(?,?,?,?,?,?)", createInput.Set_bank_code, createInput.Set_bank_name, createInput.Set_short_name, createInput.Set_bank_bic, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	partnerlistInput body models.PartnerListRequest true "PartnerList Input"
// @Success		  	200 {object} models.PartnerListResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_partnerlist/ [post]
func GetPartnerList(c *fiber.Ctx) error {
	partnerlistInput := models.PartnerListRequest{}

	if parErr := c.BodyParser(&partnerlistInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	partnerlistModel := []models.PartnerListResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_partnerlist").Find(&partnerlistModel, partnerlistInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(partnerlistModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    partnerlistModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allPartnerlistInput body models.AllPartnerlistRequest true "AllPartnerlist Input"
// @Success		  	200 {object} models.AllPartnerlistResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_partnerlist/ [post]
func SelectPartnerlistbyID(c *fiber.Ctx) error {
	partnerlistInput := models.AllPartnerlistRequest{}

	if parErr := c.BodyParser(&partnerlistInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allpartnerlistModel := []models.AllPartnerlistResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_partnerlist(?)", partnerlistInput.Get_id).Scan(&allpartnerlistModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allpartnerlistModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allpartnerlistModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editPartnerlistInput body models.EditPartnerlistRequest true "EditPartnerlist Input"
// @Success		  	200 {object} models.EditPartnerlistResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_partnerlist/ [post]
func EditPartnerlist(c *fiber.Ctx) error {
	editpartnerlistnInput := models.EditPartnerlistRequest{}

	if parErr := c.BodyParser(&editpartnerlistnInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editpartnerlistModel := models.EditPartnerlistResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_partnerlist(?,?,?,?,?,?,?,?,?,?)", editpartnerlistnInput.Get_Partnerlist_id, editpartnerlistnInput.Get_Partner_id, editpartnerlistnInput.Get_Partner_name, editpartnerlistnInput.Get_Partner_desc, editpartnerlistnInput.Get_Partner_account, editpartnerlistnInput.Get_Partner_api_url, editpartnerlistnInput.Get_Merchant_payment_callback_url, editpartnerlistnInput.Get_Merchant_id_prefix, editpartnerlistnInput.Get_Mri_group, editpartnerlistnInput.Get_Status).Scan(&editpartnerlistModel).Error; dbErr != nil {
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
// @Param       	droppartnerlistInput body models.DropPartnerListRequest true "DropPartnerListInput"
// @Success		  	200 {object} models.DropPartnerListResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_partnerlist/ [post]
func DropPartnerList(c *fiber.Ctx) error {
	droppartnerlistInput := models.DropPartnerListRequest{}

	if parErr := c.BodyParser(&droppartnerlistInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	droppartnerlistModel := models.DropPartnerListResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_partnerlist(?,?)", droppartnerlistInput.Drop_id, droppartnerlistInput.Drop_partnerlist).Scan(&droppartnerlistModel).Error; dbErr != nil {
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
// @Param       	createPartnerListInput body models.CreatePartnerListRequest true "PartnerList Input"
// @Success		  	200 {object} models.CreatePartnerListResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_partner/ [post]
func CreatePartnerList(c *fiber.Ctx) error {
	createInput := models.CreatePartnerListRequest{}

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

	createModel := models.CreatePartnerListResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_partner(?,?,?,?,?,?,?,?,?,?,?)", createInput.Set_partner_id, createInput.Set_partner_name, createInput.Set_partner_desc, createInput.Set_partner_account, createInput.Set_partner_api_url, createInput.Set_merchant_payment_callback_url, createInput.Set_merchant_id_prefix, createInput.Set_mri_group, createInput.Set_status, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
// @Param       	splashscreenInput body models.SplashScreenRequest true "SplashScreen Input"
// @Success		  	200 {object} models.SplashScreenResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/get_splashscreen/ [post]
func GetSplashScreen(c *fiber.Ctx) error {
	splashscreenInput := models.SplashScreenRequest{}

	if parErr := c.BodyParser(&splashscreenInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	splashscreenModel := []models.SplashScreenResponse{}

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_splashscreen").Find(&splashscreenModel, splashscreenInput).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(splashscreenModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Success",
		Data:    splashscreenModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	allSplashScreenInput body models.AllSplashScreenRequest true "AllSplashScreen Input"
// @Success		  	200 {object} models.AllSplashScreenResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/select_splashscreen/ [post]
func SelectSplashScreenbyID(c *fiber.Ctx) error {
	splashscreenInput := models.AllSplashScreenRequest{}

	if parErr := c.BodyParser(&splashscreenInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	allsplashscreenModel := []models.AllSplashScreenResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.get_splashscreen(?)", splashscreenInput.Get_id).Scan(&allsplashscreenModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}

	if len(allsplashscreenModel) == 0 {
		return c.Status(http.StatusCreated).JSON(models.ResponseWoModel{
			RetCode: "400",
			Message: "No Data Available in Table",
		})
	}

	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    allsplashscreenModel,
	})
}

// @summary 	  	Fetch User Data
// @Description	  	Fetch User Data
// @Tags		  	Webtool
// @Accept		  	json
// @Produce		  	json
// @Param       	editSplashScreenInput body models.EditSplashScreenRequest true "EditSplashScreen Input"
// @Success		  	200 {object} models.EditSplashScreenResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/edit_splashscreen/ [post]
func EditSplashScreen(c *fiber.Ctx) error {
	editsplashscreenInput := models.EditSplashScreenRequest{}

	if parErr := c.BodyParser(&editsplashscreenInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	editsplashscreenModel := models.EditSplashScreenResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.update_splashscreen(?,?,?,?,?,?,?)", editsplashscreenInput.Get_id, editsplashscreenInput.Get_Action, editsplashscreenInput.Get_Title, editsplashscreenInput.Get_Message, editsplashscreenInput.Get_Sub_message, editsplashscreenInput.Get_Image_url, editsplashscreenInput.Get_Show).Scan(&editsplashscreenModel).Error; dbErr != nil {
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
// @Param       	dropsplashscreenInput body models.DropSplashScreenRequest true "DropSplashScreenInput"
// @Success		  	200 {object} models.DropSplashScreenResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/drop_splashscreen/ [post]
func DropSplashScreen(c *fiber.Ctx) error {
	dropsplashscreenInput := models.DropSplashScreenRequest{}

	if parErr := c.BodyParser(&dropsplashscreenInput); parErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Error",
			Data:    parErr.Error(),
		})
	}

	dropsplashscreenModel := models.DropSplashScreenResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.delete_splashscreen(?,?)", dropsplashscreenInput.Drop_id, dropsplashscreenInput.Drop_splashscreen).Scan(&dropsplashscreenModel).Error; dbErr != nil {
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
// @Param       	createSplashScreenInput body models.CreateSplashScreenRequest true "SplashScreen Input"
// @Success		  	200 {object} models.CreateSplashScreenResponse
// @Failure 		400 {object} models.ResponseModel
// @Router			/create_splashscreen/ [post]
func CreateSplashScreen(c *fiber.Ctx) error {
	createInput := models.CreateSplashScreenRequest{}

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

	createModel := models.CreateSplashScreenResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.create_splashscreen(?,?,?,?,?,?,?,?)", createInput.Set_action, createInput.Set_title, createInput.Set_message, createInput.Set_sub_message, createInput.Set_show, createInput.Set_set_img_url, createInput.Set_created_date, createInput.Set_created_by).Scan(&createModel).Error; dbErr != nil {
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
