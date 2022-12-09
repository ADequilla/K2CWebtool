package models

import "time"

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

type AllHeirarchyRequest struct {
	Get_id string `json:"get_id"`
}

type EditHierarchyRequest struct {
	Get_hierarchy_id string `json:"get_hierarchy_id"`
	Get_branch_code  string `json:"get_branch_code"`
	Get_branch_desc  string `json:"get_branch_desc"`
	Get_unit_code    string `json:"get_unit_code"`
	Get_unit_desc    string `json:"get_unit_desc"`
	Get_center_code  string `json:"get_center_code"`
	Get_center_desc  string `json:"get_center_desc"`
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

type AllUseddeviceRequest struct {
	Get_id string `json:"get_id"`
}

type EditUseddeviceRequest struct {
	Get_id            string `json:"get_id"`
	Get_device_id     string `json:"get_device_id"`
	Get_device_model  string `json:"get_device_model"`
	Get_cid           string `json:"get_cid"`
	Get_branch_code   string `json:"get_branch_code"`
	Get_mobile_number string `json:"get_mobile_number"`
	Get_client_name   string `json:"get_client_name"`
	Get_client_type   string `json:"get_client_type"`
	Get_device_status int    `json:"get_device_status"`
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

type AllUserManagementRequest struct {
	Get_id string `json:"get_id"`
}

type EditUserManagementRequest struct {
	Get_user_id      string `json:"get_user_id"`
	Get_user_name    string `json:"get_user_name"`
	Get_given_name   string `json:"get_given_name"`
	Get_middle_name  string `json:"get_middle_name"`
	Get_last_name    string `json:"get_last_name"`
	Get_user_email   string `json:"get_user_email"`
	Get_user_phone   string `json:"get_user_phone"`
	Get_user_status  string `json:"get_user_status"`
	Get_check_status string `json:"get_check_status"`
	Get_role_name    string `json:"get_role_name"`
	Get_inst_desc    string `json:"get_inst_desc"`
	Get_branch_desc  string `json:"get_branch_desc"`
}

///feestructure

type FeeStructureRequest struct {
	Trans_type string `json:"trans_type"`
}

type AllFeestructureRequest struct {
	Get_id string `json:"get_id"`
}

type EditFeeStructureRequest struct {
	Get_fee_id              string `json:"get_fee_id"`
	Get_client_type         string `json:"get_client_type"`
	Get_trans_type          string `json:"get_trans_type"`
	Get_start_range         int    `json:"get_start_range"`
	Get_end_range           int    `json:"get_end_range"`
	Get_total_charge        int    `json:"get_total_charge"`
	Get_agent_income        int    `json:"get_agent_income"`
	Get_bank_income         int    `json:"get_bank_income"`
	Get_agent_target_income int    `json:"get_agent_target_income"`
	Get_bancnet_income      int    `json:"get_bancnet_income"`
}

///paramconfig

type ParamConfigRequest struct {
	App_type   string `json:"app_type" example:"Param Type"`
	Param_name string `json:"param_name"`
}

type AllParamConfigRequest struct {
	Get_id string `json:"get_id"`
}

type EditParamConfigRequest struct {
	Get_param_id    string `json:"get_param_id"`
	Get_app_type    string `json:"get_app_type"`
	Get_param_name  string `json:"get_param_name"`
	Get_param_value string `json:"get_param_value"`
	Get_param_desc  string `json:"get_param_desc"`
}

///atmloc

type AtmLocRequest struct {
	Inst_desc       string `json:"i.inst_desc"`
	Atm_description string `json:"c.atm_description"`
	Atm_address     string `json:"c.atm_address"`
	Atm_city        string `json:"c.atm_city"`
}

type AllAtmLocRequest struct {
	Get_id string `json:"get_id"`
}

type EditAtmLocRequest struct {
	Get_atm_id          string `json:"get_atm_id"`
	Get_atm_description string `json:"get_atm_description"`
	Get_atm_address     string `json:"get_atm_address"`
	Get_atm_city        string `json:"get_atm_city"`
	Get_atm_longitude   int    `json:"get_atm_longitude"`
	Get_atm_latitude    int    `json:"get_atm_latitude"`
}

///productservices

type ProductServicesRequest struct {
	Service_name string `json:"service_name"`
}

type AllProductServicesRequest struct {
	Get_id string `json:"get_id"`
}

type EditProductServicesRequest struct {
	Get_service_id          string `json:"get_service_id"`
	Get_service_name        string `json:"get_service_name"`
	Get_service_description string `json:"get_service_description"`
	Get_show                bool   `json:"get_show"`
	Get_service_banner      string `json:"get_service_banner"`
}

///servicedowntime

type ServiceDowntimeRequest struct {
	Downtime_desc  string `json:"downtime_desc"`
	Downtime_start string `json:"downtime_start"`
	Downtime_end   string `json:"downtime_end"`
}

type AllServiceDowntimeRequest struct {
	Get_id string `json:"get_id"`
}

type EditServiceDowntimeRequest struct {
	Get_downtime_id    string    `json:"get_downtime_id"`
	Get_downtime_start time.Time `json:"get_downtime_start"`
	Get_downtime_end   time.Time `json:"get_downtime_end"`
	Get_client_type    string    `json:"get_client_type"`
	Get_downtime_desc  string    `json:"get_downtime_desc"`
}

///banknews

type BankNewsRequest struct {
	Product_name string `json:"c.product_name"`
}

type AllBankNewsRequest struct {
	Get_id string `json:"get_id"`
}

type EditBankNewsRequest struct {
	Get_product_id            string    `json:"get_product_id"`
	Get_product_name          string    `json:"get_product_name"`
	Get_product_description   string    `json:"get_product_description"`
	Get_product_periode_start time.Time `json:"get_product_periode_start"`
	Get_product_periode_end   time.Time `json:"get_product_periode_end"`
	Get_product_img_name      string    `json:"get_product_img_name"`
}

///institution

type InstitutionRequest struct {
	Inst_code string `json:"inst_code"`
	Inst_desc string `json:"inst_desc"`
}

type AllInstiRequest struct {
	Get_id string `json:"get_id"`
}

type EditInstiRequest struct {
	Get_inst_id   string `json:"get_inst_id"`
	Get_inst_code string `json:"get_inst_code"`
	Get_inst_desc string `json:"get_inst_desc"`
}

///branch

type BranchRequest struct {
	Branch_code string `json:"branch_code"`
	Branch_desc string `json:"branch_desc"`
}

type AllBranchRequest struct {
	Get_id string `json:"get_id"`
}

type EditBranchRequest struct {
	Get_branch_id   string `json:"get_branch_id"`
	Get_branch_code string `json:"get_branch_code"`
	Get_branch_desc string `json:"get_branch_desc"`
}

///unit

type UnitRequest struct {
	Unit_code string `json:"unit_code"`
	Unit_desc string `json:"unit_desc"`
}

type AllUnitRequest struct {
	Get_id string `json:"get_id"`
}

type EditUnitRequest struct {
	Get_unit_id   string `json:"get_unit_id"`
	Get_unit_code string `json:"get_unit_code"`
	Get_unit_desc string `json:"get_unit_desc"`
}

///center

type CenterRequest struct {
	Center_code string `json:"center_code"`
	Center_desc string `json:"center_desc"`
}

type AllCenterRequest struct {
	Get_id string `json:"get_id"`
}

type EditCenterRequest struct {
	Get_center_id   string `json:"get_center_id"`
	Get_center_code string `json:"get_center_code"`
	Get_center_desc string `json:"get_center_desc"`
}

///provider

type ProviderRequest struct {
	Id            string `json:"id"`
	Provider_name string `json:"provider_name"`
}

type AllProviderRequest struct {
	Get_id string `json:"get_id"`
}

type EditProviderRequest struct {
	Get_provider_id    string `json:"get_provider_id"`
	Get_provider_name  string `json:"get_provider_name"`
	Get_description    string `json:"get_description"`
	Get_provider_alias string `json:"get_provider_alias"`
	Get_status         int    `json:"get_status"`
}

///producttype

type ProductTypeRequest struct {
	Product_type_id   string `json:"product_type_id"`
	Product_type_name string `json:"product_type_name"`
	Provider_name     string `json:"provider_name"`
}

type AllProductTypeRequest struct {
	Get_id string `json:"get_id"`
}

type EditProductTypeRequest struct {
	Get_producttype_id    string `json:"get_producttype_id"`
	Get_provider_name     string `json:"get_provider_name"`
	Get_product_type_id   int    `json:"get_product_type_id"`
	Get_product_type_name string `json:"get_product_type_name"`
	Get_description       string `json:"get_description"`
	Get_status            int    `json:"get_status"`
}

///productcategory

type ProductCategoryRequest struct {
	Product_category_id   string `json:"c.product_category_id"`
	Product_category_name string `json:"c.product_category_name"`
	Provider_name         string `json:"a.provider_name"`
	Product_type_name     string `json:"b.product_type_name"`
}

type AllProductCategoryRequest struct {
	Get_id string `json:"get_id"`
}

type EditProductCategoryRequest struct {
	Get_productcategory_id    string `json:"get_productcategory_id"`
	Get_product_type_name     string `json:"get_product_type_name"`
	Get_product_category_id   int    `json:"get_product_category_id"`
	Get_product_category_name string `json:"get_product_category_name"`
	Get_status                int    `json:"get_status"`
}

///billerproduct

type BillerProductRequest struct {
	Biller_product_id     string `json:"biller_product_id"`
	Biller_product_name   string `json:"biller_product_name"`
	Provider_name         string `json:"provider_name"`
	Product_category_name string `json:"product_category_name"`
}

type AllBillerProductRequest struct {
	Get_id string `json:"get_id"`
}

type EditBillerProductRequest struct {
	Get_Billerproduct_id      string `json:"get_billerproduct_id"`
	Get_Product_category_name string `json:"get_product_category_name"`
	Get_Biller_product_id     int    `json:"get_biller_product_id"`
	Get_Biller_product_name   string `json:"get_biller_product_name"`
	Get_Description           string `json:"get_description"`
	Get_Service_fee           string `json:"get_service_fee"`
	Get_Bank_commission       string `json:"get_bank_commission"`
	Get_Status                int    `json:"get_status"`
}

///loadproduct

type LoadProductRequest struct {
	Load_product_id       string `json:"load_product_id"`
	Load_product_name     string `json:"load_product_name"`
	Product_category_name string `json:"product_category_name"`
}

type AllLoadProductRequest struct {
	Get_id string `json:"get_id"`
}

type EditLoadProductRequest struct {
	Get_loadproduct_id        string `json:"get_loadproduct_id"`
	Get_product_category_name string `json:"get_product_category_name"`
	Get_load_product_id       int    `json:"get_load_product_id"`
	Get_load_product_name     string `json:"get_load_product_name"`
	Get_description           string `json:"get_description"`
	Get_status                int    `json:"get_status"`
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

///AgentDashboard

type AgentDashboardRequest struct {
	Branch_desc string `json:"branch_desc"`
}

///CsrDashboard

type CSRDashboardRequest struct {
	Ticket_id    string `json:"customer_ticket_id"`
	Client_cid   string `json:"client_cid"`
	Submitted_by string `json:"submitted_by"`
	Created_date string `json:"created_date"`
	Closed_date  string `json:"closed_date"`
	Concern_name string `json:"concern_name"`
	Status       string `json:"status"`
}

///TransConformation

type TransConfirmationRequest struct {
	Cid         string `json:"cid"`
	Trans_desc  string `json:"trans_desc"`
	Status      string `json:"status"`
	Branch_desc string `json:"branch_desc"`
	Trans_date  string `json:"trans_date"`
}

///Broadcastsms

type BroadcastSmsRequest struct {
	Subject string `json:"subject"`
}

///WebReport

type WebReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///TransactionReport

type TransactionReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///RemittanceSentReport

type RemittanceSentReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///RemittanceClaimedReport

type RemittanceClaimedReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///RemittanceCancelledReport

type RemittanceCancelledReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///ActivityHistoryReport

type ActivityHistoryReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///MpinChangeReport

type MpinChangeReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///SmsActivationReport

type SmsActivationReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///LoginLogoutReport

type LoginLogoutReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///UserActivityReport

type UserActivityReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///TransactionSuspiciousReport

type TransactionSuspiciousReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///TransactionValidReport

type TransactionValidReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///CsDashboardReport

type CsDashboardReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///ReconCcmReport

type ReconCcmReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///SmsLogReport

type SmsLogReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///AccountStatusReport

type AccountStatusReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///EnableAgentReport

type EnableAgentReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///IgateReconReport

type IgateReconReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///ActivatedMerchantReport

type ActivatedMerchantReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_id      string `json:"report_id"`
}

///DeactivatedMerchantReport

type DeactivatedMerchantReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_id      string `json:"report_id"`
}

///UsedDeviceReport

type UsedDeviceReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///FailedEnrollmentReport

type FailedEnrollmentReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///RegisteredClientReport

type RegisteredClientReportRequest struct {
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	Report_id      string `json:"report_id"`
	Branch_desc    string `json:"branch_desc"`
	User_name      string `json:"user_name"`
}

///Task

type TaskRequest struct {
	Task_id    string `json:"task_id"`
	Task_name  string `json:"task_name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
