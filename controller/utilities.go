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
// @Param       	editcenterInput body models.EditCenterRequest true "EditCenter Input"
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

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_producttype").Find(&producttypeModel, producttypeInput).Error; dbErr != nil {
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

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_billerproduct").Find(&billerproductModel, billerproductInput).Error; dbErr != nil {
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

	if dbErr := middleware.DBConn.Debug().Table("mfs.view_loadproduct").Find(&loadproductModel, loadproductInput).Error; dbErr != nil {
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
