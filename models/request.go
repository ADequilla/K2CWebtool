package models

///clientprofile

type ClientProfileRequest struct {
	Username string `json:"username"`
	MobileNo string `json:"mobile"`
	Costumer string `json:"cid"`
}

///hierarchy

type HierarchyRequest struct {
	Inst_desc   string `json:"i.inst_desc"`
	Unit_desc   string `json:"u.unit_desc"`
	Branch_desc string `json:"b.branch_desc"`
	Center_desc string `json:"c.center_desc"`
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
	Created_date         string `json:"created_date" example:"StartSendDate"`
	Last_updated_date    string `json:"last_updated_date" example:"EndSendDate"`
	Sender_mobile_number string `json:"sender_mobile_number"`
	Sent_mobile_ref_id   string `json:"sent_mobile_ref_id" example:"MobileReference"`
	Status               string `json:"status"`
	Source_branch        string `json:"source_branch"`
	Target_branch        string `json:"target_branch"`
}

///rolemanagement

type RolesManagementRequest struct {
	Role_name string `json:"role_name"`
	Role_desc string `json:"role_desc"`
}

//smslog

type SmsLogRequest struct {
	Cid           string `json:"cid" example:"Cid"`
	Msg_sent_date string `json:"msg_sent_date" example:"EnrolledDateStart"`
	Msg_rsp_date  string `json:"msg_rsp_date" example:"EnrolledDateEnd"`
	Msisdn        string `json:"msisdn" example:"MobileNumber"`
	Activity      string `json:"activity" example:"MessageType"`
	Msg_status    string `json:"msg_status" example:"Message"`
}

///transactionlog

type TransLogRequest struct {
	SearchTranslog string `json:"search_translog"`
}

///usermanagement

type UserManagementRequest struct {
	User_name    string `json:"user_name"`
	Given_name   string `json:"given_name"`
	MiddleName   string `json:"middle_name"`
	Last_name    string `json:"last_name"`
	Branch_names string `json:"branch_names"`
	Check_status string `json:"check_status"`
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

///csrhotline

type CsrHotlineRequest struct {
	SearchCsrhotline string `json:"search_csrhotline"`
}

///concerntype

type ConcernTypeRequest struct {
	SearchConcerntype string `json:"search_concerntype"`
}
