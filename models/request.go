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
	Created_date      string `json:"created_date" example:"Activated Date Start"`
	Last_updated_date string `json:"last_updated_date" example:"Activated Date End"`
	Date_of_birth     string `json:"date_of_birth"`
	Client_type       string `json:"client_type" example:"Client Type"`
	Account_number    string `json:"account_number"`
	Mobile_number     string `json:"mobile_number"`
	Error_message     string `json:"error_message"`
}

///listofagent

type ListofAgentRequest struct {
	Date_And_Time string `json:"c.date_and_time" example:"Activated Date Start"`
	Date_and_time string `json:"q.date_and_time" example:"Activated Date End"`
	Cid           string `json:"c.cid"`
	Inst_desc     string `json:"a.inst_desc"`
	Branch_desc   string `json:"c.branch_desc"`
	Mobile_no     string `json:"c.mobile_no"`
	Unit_desc     string `json:"c.unit_desc"`
	Center_desc   string `json:"c.center_desc"`
}

///listuseddevice

type UsedDeviceRequest struct {
	Created_date      string `json:"created_date" example:"Activated Date Start"`
	Last_updated_date string `json:"last_updated_date" example:"Activated Date End"`
	Cid               string `json:"cid"`
	Client_type       string `json:"client_type" example:"Client Type"`
	Device_id         string `json:"device_id"`
	Mobile_number     string `json:"mobile_number"`
	Device_status     string `json:"device_status" example:"Status"`
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
	Core_refno      string `json:"core_refno" example:"Core_ID"`
	Trans_type_code string `json:"trans_type_code" example:"Transaction_Type"`
	Source_cid      string `json:"source_cid"`
	Source_account  string `json:"source_account"`
	Status          string `json:"status" `
	Mobile_refno    string `json:"mobile_refno" example:"Mobile_ID"`
	Trans_date      string `json:"trans_date" example:"Date_Time"`
	Post_date       string `json:"post_date" `
	Target_cid      string `json:"target_cid" `
	Target_account  string `json:"target_account" `
}

///usermanagement

type UserManagementRequest struct {
	User_login   string `json:"user_login"`
	Given_name   string `json:"given_name"`
	Middle_name  string `json:"middle_name"`
	Last_name    string `json:"last_name"`
	Branch_names string `json:"branch_names"`
	Check_status string `json:"check_status"`
}

///feestructure

type FeeStructureRequest struct {
	Trans_type string `json:"trans_type"`
}

///paramconfig

type ParamConfigRequest struct {
	App_type   string `json:"app_type" example:"Param Type"`
	Param_name string `json:"param_name"`
}

///atmloc

type AtmLocRequest struct {
	Inst_desc       string `json:"i.inst_desc"`
	Atm_description string `json:"c.atm_description"`
	Atm_address     string `json:"c.atm_address"`
	Atm_city        string `json:"c.atm_city"`
}

///productservices

type ProductServicesRequest struct {
	Service_name string `json:"service_name"`
}

///servicedowntime

type ServiceDowntimeRequest struct {
	Downtime_desc  string `json:"downtime_desc"`
	Downtime_start string `json:"downtime_start"`
	Downtime_end   string `json:"downtime_end"`
}

///banknews

type BankNewsRequest struct {
	Product_name string `json:"c.product_name"`
}

///institution

type InstitutionRequest struct {
	Inst_code string `json:"inst_code"`
	Inst_desc string `json:"inst_desc"`
}

///branch

type BranchRequest struct {
	Branch_code string `json:"branch_code"`
	Branch_desc string `json:"branch_desc"`
}

///unit

type UnitRequest struct {
	Unit_code string `json:"unit_code"`
	Unit_desc string `json:"unit_desc"`
}

///center

type CenterRequest struct {
	Center_code string `json:"center_code"`
	Center_desc string `json:"center_desc"`
}

///provider

type ProviderRequest struct {
	Id            string `json:"id"`
	Provider_name string `json:"provider_name"`
}

///producttype

type ProductTypeRequest struct {
	Product_type_id   string `json:"c.product_type_id"`
	Product_type_name string `json:"c.product_type_name"`
	Provider_name     string `json:"a.provider_name"`
}

///productcategory

type ProductCategoryRequest struct {
	Product_category_id   string `json:"c.product_category_id"`
	Product_category_name string `json:"c.product_category_name"`
	Provider_name         string `json:"a.provider_name"`
	Product_type_name     string `json:"b.product_type_name"`
}

///billerproduct

type BillerProductRequest struct {
	Biller_product_id     string `json:"c.biller_product_id"`
	Biller_product_name   string `json:"c.biller_product_name"`
	Provider_name         string `json:"a.provider_name"`
	Product_category_name string `json:"b.product_category_name"`
}

///billerproduct

type LoadProductRequest struct {
	Load_product_id       string `json:"c.load_product_id"`
	Load_product_name     string `json:"c.load_product_name"`
	Product_category_name string `json:"a.product_category_name"`
}

///commissionsetup

type CommissionRequest struct {
	Trans_type      string `json:"trans_type"`
	Commission_type string `json:"commission_type"`
}

///banklist

type BankListRequest struct {
	Bank_code  string `json:"bank_code"`
	Bank_name  string `json:"bank_name"`
	Short_name string `json:"short_name"`
}

///partnerlist

type PartnerListRequest struct {
	Partner_id   string `json:"partner_id"`
	Partner_desc string `json:"partner_desc"`
	Partner_name string `json:"partner_name"`
	Status       string `json:"status"`
}

///splashscreen

type SplashScreenRequest struct {
	Action string `json:"action"`
	Title  string `json:"title"`
}

///csrhotline

type CsrHotlineRequest struct {
	Contact_number   string `json:"c.contact_number"`
	Network_provider string `json:"c.network_provider"`
}

///concerntype

type ConcernTypeRequest struct {
	Concern_name  string `json:"concern_name"`
	Concern_level string `json:"concern_level"`
}

///login

type LoginRequest struct {
	User_login  string `json:"c.user_login" validate:"required"`
	User_passwd string `json:"c.user_passwd" validate:"required"`
}

///Clientlistforregistration

type ListforRegistrationRequest struct {
	Cid             string `json:"cid"`
	Branch_desc     string `json:"branch_desc"`
	Unit_desc       string `json:"unit_desc"`
	Center_desc     string `json:"center_desc"`
	Mobile_no       string `json:"mobile_no"`
	Enrolled        string `json:"enrolled"`
	Re_enrolled     string `json:"re_enrolled"`
	Approved        string `json:"approved"`
	Re_approved     string `json:"re_approved"`
	Acc_status_desc string `json:"acc_status_desc"`
	Ams_status_desc string `json:"sms_status_desc"`
	Account_number  string `json:"account_number"`
}

///SlfRequest

type SlfRequestRequest struct {
	Cid         string `json:"cid"`
	Branch_desc string `json:"branch_desc"`
	Trans_date  string `json:"trans_date"`
}

///OperationDashboard

type OperationDashboardRequest struct {
	Trans_post_date   string `json:"trans_post_date"`
	Last_updated_date string `json:"last_updated_date"`
	Ref_branch        string `json:"ref_branch"`
}

///AuthorResetPasswword

type ARPRequest struct {
	Cid    string `json:"cid"`
	Status string `json:"status"`
}
