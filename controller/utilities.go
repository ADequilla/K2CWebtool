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
		Message: "Succes",
		Data:    feestructureModel,
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
		Message: "Succes",
		Data:    paramconfigModel,
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
		Message: "Succes",
		Data:    atmlocModel,
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
		Message: "Succes",
		Data:    productservicesModel,
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
		Message: "Succes",
		Data:    servicedowntimeModel,
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
		Message: "Succes",
		Data:    banknewsModel,
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
		Message: "Succes",
		Data:    institutionModel,
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
		Message: "Succes",
		Data:    branchModel,
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
		Message: "Succes",
		Data:    unitModel,
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
		Message: "Succes",
		Data:    centerModel,
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
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
		Message: "Succes",
		Data:    splashscreenModel,
	})
}
