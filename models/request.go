package models

///clientprofile

type ClientProfileRequest struct {
	SearchClient string `json:"search_client"`
}

///hierarchy

type HierarchyRequest struct {
	SearchHierarchy string `json:"search_hierarchy"`
}

///listfailedenrollment

type FailedEnrollmentRequest struct {
	SearchFailedenrollment string `json:"search_failedenrollment"`
}

///listofagent

type ListofAgentRequest struct {
	SearchListofagent string `json:"search_listofagent"`
}

///listuseddevice

type UsedDeviceRequest struct {
	SearchUsedDevice string `json:"search_useddevice"`
}

///remittancelog

type RemittanceLogRequest struct {
	SearchLog string `json:"search_log"`
}

///rolemanagement

type RolesManagementRequest struct {
	SearchRole string `json:"search_role"`
}

//smslog

type SmsLogRequest struct {
	SearchSmslog string `json:"search_smslog"`
}

///transactionlog

type TransLogRequest struct {
	SearchTranslog string `json:"search_translog"`
}

///usermanagement

type UserManagementRequest struct {
	SearchName string `json:"search_name"`
}

///feestructure

type FeeStructureRequest struct {
	SearchFeeStructure string `json:"search_feestructure"`
}

///paramconfig

type ParamConfigRequest struct {
	SearchParamConfig string `json:"search_paramconfig"`
}

///atmloc

type AtmLocRequest struct {
	SearchAtmloc string `json:"search_atmloc"`
}

///productservices

type ProductServicesRequest struct {
	SearchProductServices string `json:"search_productservices"`
}

///servicedowntime

type ServiceDowntimeRequest struct {
	SearchServicedowntime string `json:"search_servicedowntime"`
}

///banknews

type BankNewsRequest struct {
	SearchBanknews string `json:"search_banknews"`
}

///institution

type InstitutionRequest struct {
	SearchInstitution string `json:"search_institution"`
}

///branch

type BranchRequest struct {
	SearchBranch string `json:"search_branch"`
}

///unit

type UnitRequest struct {
	SearchUnit string `json:"search_unit"`
}

///center

type CenterRequest struct {
	SearchCenter string `json:"search_center"`
}

///provider

type ProviderRequest struct {
	SearchProvider string `json:"search_provider"`
}

///producttype

type ProductTypeRequest struct {
	SearchProduct string `json:"search_product"`
}

///productcategory

type ProductCategoryRequest struct {
	SearchProductcategory string `json:"search_productcategory"`
}

///billerproduct

type BillerProductRequest struct {
	SearchBillerproduct string `json:"search_billerproduct"`
}

///billerproduct

type LoadProductRequest struct {
	SearchLoadproduct string `json:"search_loadproduct"`
}

///commissionsetup

type CommissionRequest struct {
	SearchCommission string `json:"search_commission"`
}

///banklist

type BankListRequest struct {
	SearchBanklist string `json:"search_banklist"`
}

///partnerlist

type PartnerListRequest struct {
	SearchPartnerlist string `json:"search_partnerlist"`
}

///splashscreen

type SplashScreenRequest struct {
	SearchSplashscreen string `json:"search_splashscreen"`
}
