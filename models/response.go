package models

///*Search*///

///clientprofile

type ClientProfileResponse struct {
	Id             string `json:"id"`
	Cid            string `json:"cid"`
	Username       string `json:"username"`
	Mobile         string `json:"mobile"`
	FullName       string `json:"FullName"`
	Birthday       string `json:"Birthday"`
	Insti_name     string `json:"insti_name"`
	Agent          string `json:"Agent"`
	Enable         string `json:"Enable"`
	Merchant       string `json:"Merchant"`
	Account_name   string `json:"account_name"`
	Account_number string `json:"account_number"`
	Branch         string `json:"Branch"`
	Unit           string `json:"Unit"`
	Center         string `json:"Center"`
	Type           string `json:"Type"`
	Classification string `json:"Classification"`
	Is_enabled     string `json:"is_enabled"`
}

///hierarchy

type HierarchyResponse struct {
	Hierarchy_id string `json:"hierarchy_id"`
	Branch_code  string `json:"branch_code"`
	Branch_desc  string `json:"branch_desc"`
	Unit_code    string `json:"unit_code"`
	Unit_desc    string `json:"unit_desc"`
	Center_code  string `json:"center_code"`
	Center_desc  string `json:"center_desc"`
	Inst_desc    string `json:"i.inst_desc"`
}

type AllHeirarchyResponse struct {
	Hierarchy_id string `json:"hierarchy_id"`
	Branch_code  string `json:"branch_code"`
	Branch_desc  string `json:"branch_desc"`
	Unit_code    string `json:"unit_code"`
	Unit_desc    string `json:"unit_desc"`
	Center_code  string `json:"center_code"`
	Center_desc  string `json:"center_desc"`
}

type EditHeirarchyResponse struct {
	Update_hierarchy string `json:"update_hierarchy"`
}

///listfailedenrollment

type FailedEnrollmentResponse struct {
	Id             string `json:"id"`
	Created_date   string `json:"created_date"`
	Account_number string `json:"account_number"`
	Date_of_birth  string `json:"date_of_birth"`
	Mobile_number  string `json:"mobile_number"`
	Client_type    string `json:"client_type"`
	Device_id      string `json:"device_id"`
	Device_model   string `json:"device_model"`
	Error_message  string `json:"error_message"`
}

///listofagent

type ListofAgentResponse struct {
	Date_and_time string `json:"c.date_and_time"`
	User_name     string `json:"i.user_name"`
	Cid           string `json:"c.cid"`
	Fullname      string `json:"c.fullname"`
	Mobile_no     string `json:"c.mobile_no"`
	Inst_desc     string `json:"a.inst_desc"`
	Branch_desc   string `json:"c.branch_desc"`
	Unit_desc     string `json:"c.unit_desc"`
	Center_desc   string `json:"c.center_desc"`
}

///listuseddevice

type UsedDeviceResponse struct {
	Created_date    string `json:"created_date"`
	Device_id       string `json:"device_id"`
	Device_model    string `json:"device_model"`
	Android_version string `json:"android_version"`
	Cid             string `json:"cid"`
	Branch_code     string `json:"branch_code"`
	Mobile_number   string `json:"mobile_number"`
	Client_name     string `json:"client_name"`
	Client_type     string `json:"client_type"`
	Device_status   string `json:"device_status"`
}

type AllUseddeviceResponse struct {
	Id            string `json:"id"`
	Device_id     string `json:"device_id"`
	Device_model  string `json:"device_model"`
	Cid           string `json:"cid"`
	Branch_code   string `json:"branch_code"`
	Mobile_number string `json:"mobile_number"`
	Client_name   string `json:"client_name"`
	Client_type   string `json:"client_type"`
	Device_status string `json:"device_status"`
}

type EditUseddeviceResponse struct {
	Update_listuseddevice string `json:"update_listuseddevice"`
}

///remittancelog

type RemittanceLogResponse struct {
	Id                    string `json:"id"`
	Sent_mobile_ref_id    string `json:"sent_mobile_ref_id"`
	Sent_core_ref_id      string `json:"sent_core_ref_id"`
	Reference_number      string `json:"reference_number _Ref_ID"`
	Sender_name           string `json:"sender_name"`
	Receiver_Name         string `json:"Receiver_Name"`
	Amount                string `json:"amount"`
	Sender_Mobile_Number  string `json:"sender_mobile_number"`
	Created_date          string `json:"created_date"`
	Source_branch         string `json:"source_branch"`
	Processed_By_Fullname string `json:"Processed_By_Fullname"`
	Last_updated_date     string `json:"last_updated_date"`
	Target_branch         string `json:"target_branch"`
	Disbursed_By_Fullname string `json:"Disbursed_By_Fullname"`
	Cancelled_date        string `json:"cancelled_date"`
	Cancelled_By_Fullname string `json:"Cancelled_By_Fullname"`
	Status                string `json:"status"`
	Claimed_core_ref_id   string `json:"claimed_core_ref_id"`
	Claimed_mobile_ref_id string `json:"claimed_mobile_ref_id"`
}

///remittancestatus

type RemittanceStatusResponse struct {
	PENDING   string `json:"PENDING"`
	SENT      string `json:"SENT"`
	CLAIMED   string `json:"CLAIMED"`
	CANCELLED string `json:"CANCELLED"`
}

///rolemanagement

type RolesManagementResponse struct {
	Role_id  string `json:"role_id"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
}

///smslog

type SmsLogResponse struct {
	Msg_id        string `json:"msg_id"`
	Msg_sent_date string `json:"msg_sent_date"`
	Msisdn        string `json:"msisdn"`
	Cid           string `json:"cid"`
	Name          string `json:"name"`
	Msg_command   string `json:"msg_command"`
	Activity      string `json:"activity"`
	Msg_status    string `json:"msg_status"`
}

//transactionlog

type TransLogResponse struct {
	Trans_log           string `json:"trans_log"`
	Mobile_refno        string `json:"mobile_refno"`
	Core_refno          string `json:"core_refno"`
	Source_branch       string `json:"source_branch"`
	Source_cid          string `json:"source_cid"`
	Source_client_type  string `json:"source_client_type"`
	Source_account_type string `json:"source_account_type"`
	Source_account      string `json:"source_account"`
	Source_name         string `json:"source_name"`
	Target_branch       string `json:"target_branch"`
	Target_cid          string `json:"target_cid"`
	Bank_name           string `json:"bank_name"`
	Target_client_type  string `json:"target_client_type"`
	Target_account_type string `json:"target_account_type"`
	Target_account      string `json:"target_account"`
	Target_name         string `json:"target_name"`
	Trans_type_code     string `json:"trans_type_code"`
	Amount              string `json:"amount"`
	Trans_amount_fee    string `json:"trans_amount_fee"`
	Agent_income        string `json:"agent_income"`
	Bank_income         string `json:"bank_income"`
	Bancnet_income      string `json:"bancnet_income"`
	Trans_date          string `json:"trans_date"`
	Post_date           string `json:"post_date"`
	Status              string `json:"status"`
	Message             string `json:"message"`
	Agent_feature       string `json:"agent_feature"`
	Ar_or_number        string `json:"ar_or_number"`
	Cust_cid            string `json:"cust_cid"`
}

///usermanagement

type UserManagementResponse struct {
	User_id      string `json:"user_id"`
	User_login   string `json:"user_login"`
	Given_name   string `json:"given_name"`
	Middle_name  string `json:"middle_name"`
	Last_name    string `json:"last_name"`
	Branch_names string `json:"branch_names"`
	Roles        string `json:"roles"`
	Check_status string `json:"check_status"`
}

type AllUserManagementResponse struct {
	Got_user_name    string `json:"got_user_name"`
	Got_given_name   string `json:"got_given_name"`
	Got_middle_name  string `json:"got_middle_name"`
	Got_last_name    string `json:"got_last_name"`
	Got_user_email   string `json:"got_user_email"`
	Got_user_phone   string `json:"got_user_phone"`
	Got_user_status  string `json:"got_user_status"`
	Got_check_status string `json:"got_check_status"`
	Got_role_name    string `json:"got_role_name"`
	Got_inst_desc    string `json:"got_inst_desc"`
	Got_branch_desc  string `json:"got_branch_desc"`
}

type EditUserManagementResponse struct {
	Update_usermanagement string `json:"update_usermanagement"`
}

///feestructure

type FeeStructureResponse struct {
	Fee_id             string `json:"fee_id"`
	Trans_type         string `json:"trans_type"`
	Range              string `json:"Range"`
	Total_charge       string `json:"total_charge"`
	Agent_income       string `json:"agent_income"`
	Bank_income        string `json:"bank_income"`
	Agen_target_income string `json:"agent_target_income"`
	Bancnet_income     string `json:"bancnet_income"`
}

type AllFeestructureResponse struct {
	Fee_id              string `json:"fee_id"`
	Client_type         string `json:"client_type"`
	Trans_type          string `json:"trans_type"`
	Start_range         string `json:"start_range"`
	End_range           string `json:"end_range"`
	Total_charge        string `json:"total_charge"`
	Agent_income        string `json:"agent_income"`
	Bank_income         string `json:"bank_income"`
	Agent_target_income string `json:"agent_target_income"`
	Bancnet_income      string `json:"bancnet_income"`
}

type EditFeeStructureResponse struct {
	Update_feestructure string `json:"update_feestructure"`
}

///paramconfig

type ParamConfigResponse struct {
	Param_id    string `json:"param_id"`
	App_type    string `json:"app_type"`
	Param_name  string `json:"param_name"`
	Param_value string `json:"param_value"`
	Param_desc  string `json:"param_desc"`
}

type AllParamConfigResponse struct {
	Param_id    string `json:"param_id"`
	App_type    string `json:"app_type"`
	Param_name  string `json:"param_name"`
	Param_value string `json:"param_value"`
	Param_desc  string `json:"param_desc"`
}

type EditParamConfigResponse struct {
	Update_paramconfig string `json:"update_paramconfig"`
}

///atmloc

type AtmLocResponse struct {
	Atm_id          string `json:"c.atm_id"`
	Inst_desc       string `json:"i.inst_desc"`
	Atm_description string `json:"c.atm_description"`
	Atm_address     string `json:"c.atm_address"`
	Atm_city        string `json:"c.atm_city"`
}

type AllAtmLocResponse struct {
	Atm_id          string `json:"atm_id"`
	Atm_description string `json:"atm_description"`
	Atm_address     string `json:"atm_address"`
	Atm_city        string `json:"atm_city"`
	Atm_longitude   string `json:"atm_longitude"`
	Atm_latitude    string `json:"atm_latitude"`
}

type EditAtmLocResponse struct {
	Update_atmloc string `json:"update_atmloc"`
}

///productservices

type ProductServicesResponse struct {
	Service_id          string `json:"service_id"`
	Service_name        string `json:"service_name"`
	Service_description string `json:"service_description"`
	Show                bool   `json:"show"`
}

///servicedowntime

type ServiceDowntimeResponse struct {
	Downtime_id    string `json:"downtime_id"`
	Downtime_desc  string `json:"downtime_desc"`
	Downtime_start string `json:"downtime_start"`
	Downtime_end   string `json:"downtime_end"`
	Client_type    string `json:"client_type"`
}

///banknews

type BankNewsResponse struct {
	Product_id   string `json:"c.product_id"`
	Product_date string `json:"c.product_date"`
	Given_name   string `json:"t.given_name"`
	Product_name string `json:"c.product_name"`
}

///institution

type InstitutionResponse struct {
	Inst_code    string `json:"inst_code"`
	Inst_desc    string `json:"inst_desc"`
	Created_date string `json:"created_date"`
}

///branch

type BranchResponse struct {
	Branch_code  string `json:"branch_code"`
	Branch_desc  string `json:"branch_desc"`
	Created_date string `json:"created_date"`
}

///unit

type UnitResponse struct {
	Unit_code    string `json:"unit_code"`
	Unit_desc    string `json:"unit_desc"`
	Created_date string `json:"created_date"`
}

///center

type CenterResponse struct {
	Center_code  string `json:"center_code"`
	Center_desc  string `json:"center_desc"`
	Created_date string `json:"created_date"`
}

///provider

type ProviderResponse struct {
	Id             string `json:"id"`
	Provider_name  string `json:"provider_name"`
	Description    string `json:"description"`
	Provider_alias string `json:"provider_alias"`
	Status         string `json:"status"`
}

///producttype

type ProductTypeResponse struct {
	Id                string `json:"c.id"`
	Provider_name     string `json:"a.provider_name"`
	Product_type_Id   string `json:"c.product_type_id"`
	Product_type_name string `json:"c.product_type_name"`
	Description       string `json:"c.description"`
	Status            string `json:"c.status"`
}

///productcategory

type ProductCategoryResponse struct {
	Id                    string `json:"c.id"`
	Provider_name         string `json:"a.provider_name"`
	Product_type_name     string `json:"b.product_type_name"`
	Product_category_id   string `json:"c.product_category_id"`
	Product_category_name string `json:"c.product_category_name"`
	Status                string `json:"c.status"`
}

///billerproduct

type BillerProductResponse struct {
	Id                    string  `json:"c.id"`
	Provider_name         string  `json:"a.provider_name"`
	Product_category_name string  `json:"b.product_category_name"`
	Biller_product_id     int     `json:"c.biller_product_id"`
	Biller_product_name   string  `json:"c.biller_product_name"`
	Description           string  `json:"c.description"`
	Bank_commission       float32 `json:"c.bank_commission"`
	Service_fee           float32 `json:"c.service_fee"`
	Status                int     `json:"c.status"`
}

///loadproduct

type LoadProductResponse struct {
	Id                    string `json:"c.id"`
	Product_category_name string `json:"a.product_category_name"`
	Load_product_id       string `json:"c.load_product_id"`
	Load_product_name     string `json:"c.load_product_name"`
	Description           string `json:"c.description"`
	Status                string `json:"c.status"`
}

///commission

type CommissionResponse struct {
	Id                  string  `json:"id"`
	Trans_type          string  `json:"trans_type"`
	Commission_type     string  `json:"commission_type"`
	Customer_income     float32 `json:"customer_income"`
	Agent_income        float32 `json:"agent_income"`
	Bank_income         float32 `json:"bank_income"`
	Bank_partner_income float32 `json:"bank_partner_income"`
}

//banklist

type BankListResponse struct {
	Id         string `json:"id"`
	Bank_code  string `json:"bank_code"`
	Bank_name  string `json:"bank_name"`
	Short_name string `json:"short_name"`
	Bank_bic   string `json:"bank_bic"`
}

//partnerlist

type PartnerListResponse struct {
	Partner_id                    string `json:"partner_id"`
	Partner_name                  string `json:"partner_name"`
	Partner_desc                  string `json:"partner_desc"`
	Partner_api_url               string `json:"partner_api_url"`
	Merchant_payment_callback_url string `json:"merchant_payment_callback_url"`
	Merchant_id_prefix            string `json:"merchant_id_prefix"`
	Mri_group                     string `json:"mri_group"`
	Status                        string `json:"status"`
}

//splashscreen

type SplashScreenResponse struct {
	Action      string `json:"action"`
	Title       string `json:"title"`
	Message     string `json:"message"`
	Sub_message string `json:"sub_message"`
	Image_url   string `json:"image_url"`
	Show        string `json:"show"`
}

//csrhotline

type CsrHotlineResponse struct {
	Id               string `json:"c.id"`
	Contact_number   string `json:"c.contact_number"`
	Network_provider string `json:"c.network_provider"`
	Inst_desc        string `json:"a.inst_desc"`
}

//corncerntype

type ConcernTypeResponse struct {
	Concern_code  string `json:"concern_code"`
	Concern_name  string `json:"concern_name"`
	Concern_time  string `json:"concern_time"`
	Concern_level string `json:"concern_level"`
}

//login

type LoginResponse struct {
	User_id                   uint   `json:"c.user_id"`
	User_login                string `json:"c.user_login"`
	User_passwd               string `json:"c.user_passwd"`
	User_name                 string `json:"c.user_name"`
	User_email                string `json:"c.user_email"`
	User_phone                string `json:"c.user_phone"`
	User_expired_passwd       string `json:"c.user_expired_passwd"`
	User_type                 string `json:"c.user_type"`
	User_status               string `json:"c.user_status"`
	User_position             string `json:"c.user_position"`
	Login_attempts            string `json:"c.login_attempts"`
	User_enabled              string `json:"c.user_enabled"`
	Inst_desc                 string `json:"i.inst_desc"`
	Branch_desc               string `json:"b.branch_desc"`
	Unit_desc                 string `json:"u.unit_desc"`
	Center_desc               string `json:"a.center_desc"`
	Last_name                 string `json:"c.last_name"`
	Given_name                string `json:"c.given_name"`
	Middle_name               string `json:"c.middle_name"`
	Passwd_default            string `json:"c.passwd_default"`
	Activation_limit          string `json:"c.activation_limit"`
	Check_status              string `json:"c.check_status"`
	Last_login_date           string `json:"c.last_login_date"`
	Last_password_chaged_date string `json:"c.last_password_chaged_date"`
}

///Clientlistforregistration

type ListforRegistrationResponse struct {
	Id               string `json:"id"`
	Enrolled         string `json:"enrolled"`
	Approved         string `json:"approved"`
	Act_code_expired string `json:"act_code_expired"`
	Cid              string `json:"cid"`
	Account_number   string `json:"account_number"`
	Type_account     string `json:"type_account"`
	Mobile_no        string `json:"mobile_no"`
	Fullname         string `json:"fullname"`
	Client_type      string `json:"client_type"`
	Branch_desc      string `json:"branch_desc"`
	Unit_desc        string `json:"unit_desc"`
	Center_desc      string `json:"center_desc"`
	Acc_status_desc  string `json:"acc_status_desc"`
	Sms_status_desc  string `json:"sms_status_desc"`
}

//slfrequest

type SlfRequestResponse struct {
	Trans_id        string `json:"trans_id"`
	Trans_date      string `json:"trans_date"`
	Branch_desc     string `json:"branch_desc"`
	Cid             string `json:"cid"`
	Client_name     string `json:"client_name"`
	Amount          string `json:"amount"`
	Amount_approved string `json:"amount_approved"`
	Status          string `json:"status"`
}

//OperationDashboard

type OperationDashboardResponse struct {
	Trans_id     string `json:"trans_id"`
	Member       string `json:"Member"`
	NonMember    string `json:"NonMember"`
	Trans_type   string `json:"trans_type"`
	Number       string `json:"number"`
	Arans_amount string `json:"trans_amount"`
	Agent_income string `json:"agent_income"`
	Bank_income  string `json:"bank_income"`
}

//AuthorResetPassword

type ARPResponse struct {
	Id                string `json:"id"`
	Created_date      string `json:"created_date"`
	Cid               string `json:"cid"`
	Client_name       string `json:"client_name"`
	Client_mobile_no  string `json:"client_mobile_no"`
	Branch_code       string `json:"branch_code"`
	Unit_code         string `json:"unit_code"`
	Center_code       string `json:"center_code"`
	Type              string `json:"type"`
	Last_updated_date string `json:"last_updated_date"`
}

///AgentDashboard

type AgentDashboardResponse struct {
	Id                     string `json:"id"`
	Branch_desc            string `json:"branch_desc"`
	Cash_in                string `json:"cash_in"`
	Cash_out               string `json:"cash_out"`
	Agent_assisted_payment string `json:"agent_assisted_payment"`
	Bill_payment           string `json:"bill_payment"`
	Total                  string `json:"total"`
	Total_income           string `json:"total_income"`
	Client_name            string `json:"client_name"`
	Ci                     string `json:"ci"`
	Co                     string `json:"co"`
	Aap                    string `json:"aap"`
	Bp                     string `json:"bp"`
	Sum                    string `json:"sum"`
	Sum_income             string `json:"sum_income"`
}

///CSRDashboard

type CSRDashboardResponse struct {
	Ticket_id          string `json:"ticket_id"`
	Customer_ticket_id string `json:"customer_ticket_id"`
	Created_date       string `json:"created_date"`
	Client_cid         string `json:"client_cid"`
	Branch_desc        string `json:"branch_desc"`
	Unit_desc          string `json:"unit_desc"`
	Center_desc        string `json:"center_desc"`
	Client_name        string `json:"client_name"`
	Client_mobile_no   string `json:"client_mobile_no"`
	Submitted_by       string `json:"submitted_by"`
	Action_detail      string `json:"action_detail"`
	Concern_desc       string `json:"concern_desc"`
	Concern_level      string `json:"concern_level"`
	Concern_type       string `json:"concern_type"`
	Trans_type         string `json:"trans_type"`
	Assigned_to        string `json:"assigned_to"`
	Action_taken       string `json:"action_taken"`
}

///TransConformation

type TransConfirmationResponse struct {
	Suspicious_id    string `json:"suspicious_id"`
	Trans_date       string `json:"trans_date"`
	Trans_desc       string `json:"trans_desc"`
	Client_mobile_no string `json:"client_mobile_no"`
	Cid              string `json:"cid"`
	Client_name      string `json:"client_name"`
	Branch_desc      string `json:"branch_desc"`
	Unit_desc        string `json:"unit_desc"`
	Center_desc      string `json:"center_desc"`
	Note             string `json:"note"`
	Status           string `json:"status"`
}

///BroadcastSms

type BroadcastSmsResponse struct {
	Inbox_id     string `json:"inbox_id"`
	Inbox_date   string `json:"inbox_date"`
	Subject      string `json:"subject"`
	Inbox_desc   string `json:"inbox_desc"`
	Period_start string `json:"period_start"`
	Period_end   string `json:"period_end"`
}

///WebReport

type WebReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///TransactionReport

type TransactionReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///RemittanceSentReport

type RemittanceSentReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///RemittanceClaimedReport

type RemittanceClaimedReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///RemittanceCancelledReport

type RemittanceCancelledReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///ActivityHistoryReport

type ActivityHistoryReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///MpinChangeReport

type MpinChangeReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///SmsActivationReport

type SmsActivationReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///LoginLogoutReport

type LoginLogoutReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///UserActivityReport

type UserActivityReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///TransactionSuspiciousReport

type TransactionSuspiciousReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///TransactionValidReport

type TransactionValidReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///CsDashboardReport

type CsDashboardReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///ReconCcmReport

type ReconCcmReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///SmsLogReport

type SmsLogReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///AccountStatusReport

type AccountStatusReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///EnableAgentReport

type EnableAgentReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///IgateReconReport

type IgateReconReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///ActivatedMerchantReport

type ActivatedMerchantReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
}

///DeactivatedMerchantReport

type DeactivatedMerchantReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
}

///UsedDeviceReport

type UsedDeviceReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///FailedEnrollmentReport

type FailedEnrollmentReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///RegisteredClientReport

type RegisteredClientReportResponse struct {
	Report_id      string `json:"report_id"`
	Report_param   string `json:"report_param"`
	User_name      string `json:"user_name"`
	Branch_desc    string `json:"branch_desc"`
	Submited_date  string `json:"submited_date"`
	Completed_date string `json:"completed_date"`
	Report_status  string `json:"report_status"`
	File_type      string `json:"file_type"`
	Remark         string `json:"remark"`
}

///Task

type TaskResponse struct {
	Task_id    string `json:"task_id"`
	Task_name  string `json:"task_name"`
	Chain_id   string `json:"chain_id"`
	Chain_name string `json:"chain_name"`
	Task_order string `json:"task_order"`
	Created_at string `json:"created_at"`
}
