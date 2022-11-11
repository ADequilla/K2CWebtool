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

	feestructureinput := c.Params("feetructureInput")
	fmt.Println("FeeStructureInput:", feestructureinput)

	feestructureModel := []models.FeeStructureResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.feestructure(?)", feestructureInput.SearchFeeStructure).Scan(&feestructureModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	paramconfiginput := c.Params("paramconfigInput")
	fmt.Println("ParamConfigInput:", paramconfiginput)

	paramconfigModel := []models.ParamConfigResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.paramconfig(?)", paramconfigInput.SearchParamConfig).Scan(&paramconfigModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	atmlocinput := c.Params("atmlocInput")
	fmt.Println("AtmLocInput:", atmlocinput)

	atmlocModel := []models.AtmLocResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.atmloc(?)", atmlocInput.SearchAtmloc).Scan(&atmlocModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	productservicesinput := c.Params("productservicesInput")
	fmt.Println("ProductServicesInput:", productservicesinput)

	productservicesModel := []models.ProductServicesResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.productservices(?)", productservicesInput.SearchProductServices).Scan(&productservicesModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	servicedowntimeinput := c.Params("servicedowntimeInput")
	fmt.Println("ServiceDowntimeInput:", servicedowntimeinput)

	servicedowntimeModel := []models.ServiceDowntimeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.servicedowntime(?)", servicedowntimeInput.SearchServicedowntime).Scan(&servicedowntimeModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	banknewsinput := c.Params("banknewsInput")
	fmt.Println("BankNewsInput:", banknewsinput)

	banknewsModel := []models.BankNewsResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.banknews(?)", banknewsInput.SearchBanknews).Scan(&banknewsModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	institutioninput := c.Params("institutionInput")
	fmt.Println("InstitutionInput:", institutioninput)

	institutionModel := []models.InstitutionResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.institution(?)", institutionInput.SearchInstitution).Scan(&institutionModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	branchbinput := c.Params("branchInput")
	fmt.Println("BranchInput:", branchbinput)

	branchModel := []models.BranchResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.branch(?)", branchInput.SearchBranch).Scan(&branchModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	unitinput := c.Params("unitInput")
	fmt.Println("UnitInput:", unitinput)

	unitModel := []models.UnitResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.unit(?)", unitInput.SearchUnit).Scan(&unitModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	centerinput := c.Params("centerInput")
	fmt.Println("CenterInput:", centerinput)

	centerModel := []models.CenterResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.center(?)", centerInput.SearchCenter).Scan(&centerModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	providerinput := c.Params("providerInput")
	fmt.Println("ProviderInput:", providerinput)

	providerModel := []models.ProviderResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.provider(?)", providerInput.SearchProvider).Scan(&providerModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	producttypeinput := c.Params("producttypeInput")
	fmt.Println("ProductTypeInput:", producttypeinput)

	producttypeModel := []models.ProductTypeResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.product(?)", producttypeInput.SearchProduct).Scan(&producttypeModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	productcategoryinput := c.Params("productcategoryInput")
	fmt.Println("ProductCategoryInput:", productcategoryinput)

	productcategoryModel := []models.ProductCategoryResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.productcategory(?)", productcategoryInput.SearchProductcategory).Scan(&productcategoryModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	billerproductinput := c.Params("billerproductInput")
	fmt.Println("BillerProductInput:", billerproductinput)

	billerproductModel := []models.BillerProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.billerproduct(?)", billerproductInput.SearchBillerproduct).Scan(&billerproductModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	loadproductinput := c.Params("loadproductInput")
	fmt.Println("LoadProductInput:", loadproductinput)

	loadproductModel := []models.LoadProductResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.loadproduct(?)", loadproductInput.SearchLoadproduct).Scan(&loadproductModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	commissioninput := c.Params("commissionInput")
	fmt.Println("CommissionInput:", commissioninput)

	commissionModel := []models.CommissionResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.commission(?)", commissionInput.SearchCommission).Scan(&commissionModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	banklistinput := c.Params("banklistInput")
	fmt.Println("BankListInput:", banklistinput)

	banklistModel := []models.BankListResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.banklist(?)", banklistInput.SearchBanklist).Scan(&banklistModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	partnerlistinput := c.Params("partnerlistInput")
	fmt.Println("PartnerListInput:", partnerlistinput)

	partnerlistModel := []models.PartnerListResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.partnerlist(?)", partnerlistInput.SearchPartnerlist).Scan(&partnerlistModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
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

	splashscreeninput := c.Params("splashscreenInput")
	fmt.Println("SplashScreenInput:", splashscreeninput)

	splashscreenModel := []models.SplashScreenResponse{}

	if dbErr := middleware.DBConn.Debug().Raw("select * from mfs.splashscreen(?)", splashscreenInput.SearchSplashscreen).Scan(&splashscreenModel).Error; dbErr != nil {
		return c.Status(http.StatusCreated).JSON(models.ResponseModel{
			RetCode: "400",
			Message: "Database Error",
			Data:    dbErr.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(models.ResponseModel{
		RetCode: "200",
		Message: "Succes",
		Data:    splashscreenModel,
	})
}
