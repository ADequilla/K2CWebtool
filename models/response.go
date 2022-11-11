package models

import (
	"time"
)

///clientprofile

type ClientProfileResponse struct {
	CostumerID     string `json:"CostumerID"`
	Username       string `json:"Username"`
	MobileNo       string `json:"MobileNo"`
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
	BranchCode string `json:"branch_code"`
	BranchDesc string `json:"branch_desc"`
	UnitCode   string `json:"unit_code"`
	UnitDesc   string `json:"unit_desc"`
	CenterCode string `json:"center_code"`
	CenterDesc string `json:"center_desc"`
}

///listfailedenrollment

type FailedEnrollmentResponse struct {
	DateTime      time.Time `json:"Date_Time"`
	AccountNumber string    `json:"Account_Number"`
	BirthDate     string    `json:"Birth_Date"`
	MobileNumber  string    `json:"Mobile_Number"`
	ClientType    string    `json:"Client_Type"`
	DeviceId      string    `json:"Device_Id"`
	DeviceModel   string    `json:"Device_Model"`
	ErrorMessage  string    `json:"Error_Message"`
}

///listofagent

type ListofAgentResponse struct {
	DateTimeEnabled time.Time `json:"Date_Time_Enabled"`
	EnabledBy       string    `json:"Enabled_By"`
	Cid             string    `json:"Cid"`
	FullName        string    `json:"Full_Name"`
	MobileNumber    string    `json:"Mobile_Number"`
	Institution     string    `json:"Institution"`
	Branch          string    `json:"Branch"`
	Unit            string    `json:"Unit"`
	Center          string    `json:"Center"`
}

///listuseddevice

type UsedDeviceResponse struct {
	DateTimeActivated time.Time `json:"DateTime_Activated" example:"2022-02-08T19:33:06Z"`
	DeviceId          string    `json:"Device_Id"`
	DeviceModel       string    `json:"Device_Model"`
	AndroidVersion    string    `json:"Android_Version"`
	Cid               string    `json:"Cid"`
	Branch            string    `json:"Branch"`
	MobileNumber      string    `json:"Mobile_Number"`
	CostumerName      string    `json:"Costumer_Name"`
	CostumerType      string    `json:"Costumer_Type"`
	Status            int       `json:"Status"`
}

///remittancelog

type RemittanceLogResponse struct {
	Mobile_Ref_ID         string    `json:"Mobile_Ref_ID"`
	Core_Ref_ID           string    `json:"Core_Ref_ID"`
	Remittance_Ref_ID     string    `json:"Remittance _Ref_ID"`
	SenderName            string    `json:"Sender_Name"`
	Receiver_Name         string    `json:"Receiver_Name"`
	Amount                string    `json:"Amount"`
	Sender_Mobile_Number  string    `json:"Sender_Mobile_Number"`
	DateTime_Send         time.Time `json:"DateTime_Send"`
	SourceBranch          string    `json:"Source_Branch"`
	Processed_By_Fullname string    `json:"Processed_By_Fullname"`
	DateTime_Receive      time.Time `json:"DateTime_Receive"`
	Target_Branch         string    `json:"Target_Branch"`
	Disbursed_By_Fullname string    `json:"Disbursed_By_Fullname"`
	DateTime_Cancelled    time.Time `json:"DateTime_Cancelled"`
	Cancelled_By_Fullname string    `json:"Cancelled_By_Fullname"`
	Status                string    `json:"Status"`
	C_core_Ref_ID         string    `json:"C_core_Ref_ID"`
	M_mobile_Ref_ID       string    `json:"M_mobile_Ref_ID"`
}

///remittancestatus

type RemittanceStatusResponse struct {
	Pending   int `json:"pending"`
	Sent      int `json:"sent"`
	Claimed   int `json:"claimed"`
	Cancelled int `json:"cancelled"`
}

///rolemanagement

type RolesManagementResponse struct {
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
}

///smslog

type SmsLogResponse struct {
	MsgId       string `json:"msg_id"`
	MsgSentDate string `json:"msg_sent_date"`
	Msisdn      string `json:"msisdn"`
	Cid         string `json:"cid"`
	Name        string `json:"name"`
	MsgCommand  string `json:"msg_command"`
	Activity    string `json:"activity"`
	Msg_Status  string `json:"msg_status"`
}

//transactionlog

type TransLogResponse struct {
	MobileId           string    `json:"Mobile_Id"`
	CoreId             string    `json:"Core_Id"`
	SourceBranch       string    `json:"Source_Branch"`
	SourceCid          string    `json:"Source_Cid"`
	SourceIdClientType string    `json:"Source_Id_Client_Type"`
	SourceAccountType  string    `json:"Source_Account_Type"`
	SourceAccount      string    `json:"Source_Account"`
	SourceName         string    `json:"Source_Name"`
	TargetBranch       string    `json:"Target_Branch"`
	TargetCid          string    `json:"Target_Cid"`
	BankName           string    `json:"Bank_Name"`
	TargetIdClientType string    `json:"Target_Id_Client_Type"`
	TargetAccountType  string    `json:"Target_Account_Type"`
	TargetAccount      string    `json:"Target_Account"`
	TargetName         string    `json:"Target_Name"`
	TransactionType    string    `json:"Transaction_Type"`
	TransactionAmount  string    `json:"Transaction_Amount"`
	TransactionCharge  string    `json:"Transaction_Charge"`
	AgentIncome        string    `json:"Agent_Income"`
	BankIncome         string    `json:"Bank_Income"`
	APBancnetInstapay  string    `json:"Ap_Bancnet_Instapay"`
	DateTime           time.Time `json:"Date_Time"`
	PostDateTime       time.Time `json:"Post_Date_Time"`
	Status             string    `json:"Status"`
	Message            string    `json:"Mes_sage"`
	WithAgentFeature   int       `json:"With_Agent_Feature"`
	ArOrNumber         string    `json:"Ar_Or_Number"`
	ChargesCorerefID   string    `json:"Charges_Coreref_Id"`
}

///usermanagement

type UserManagementResponse struct {
	UserName    string `json:"user_name"`
	GivenName   string `json:"given_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	BranchNames string `json:"branch_names"`
	Roles       string `json:"roles"`
	CheckStatus string `json:"check_status"`
}

///feestructure

type FeeStructureResponse struct {
	Transaction       string `json:"Transaction"`
	Range             string `json:"Range"`
	TotalCharge       string `json:"Total_Charge"`
	AgentIncome       string `json:"Agent_Income"`
	BankIncome        string `json:"Bank_Income"`
	AgentTargetIncome string `json:"Agent_Target_Income"`
	ApBancnetIncome   string `json:"Ap_Bancnet_Income"`
}

///paramconfig

type ParamConfigResponse struct {
	ParameterType  string `json:"Parameter_Type"`
	ParameterName  string `json:"Parameter_Name"`
	ParameterValue string `json:"Parameter_Value"`
	Description    string `json:"Description"`
}

///atmloc

type AtmLocResponse struct {
	InstiDescription string `json:"Insti_Description"`
	Description      string `json:"Description"`
	StreetBrgy       string `json:"Street_Brgy"`
	CityProvince     string `json:"City_Province"`
}

///productservices

type ProductServicesResponse struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Show        bool   `json:"Show"`
}

///servicedowntime

type ServiceDowntimeResponse struct {
	Description string    `json:"Description"`
	StartDate   time.Time `json:"Start_Date"`
	EndDate     time.Time `json:"End_Date"`
	ClientType  string    `json:"Client_Type"`
}

///banknews

type BankNewsResponse struct {
	Date   time.Time `json:"Date_"`
	Sender string    `json:"Sender"`
	Topic  string    `json:"Topic"`
}

///institution

type InstitutionResponse struct {
	Code        string    `json:"Code"`
	Description string    `json:"Description"`
	Topic       time.Time `json:"Created_Date"`
}

///branch

type BranchResponse struct {
	Code        string    `json:"Code"`
	Description string    `json:"Description"`
	Topic       time.Time `json:"Created_Date"`
}

///unit

type UnitResponse struct {
	Code        string    `json:"Code"`
	Description string    `json:"Description"`
	Topic       time.Time `json:"Created_Date"`
}

///center

type CenterResponse struct {
	Code        string    `json:"Code"`
	Description string    `json:"Description"`
	Topic       time.Time `json:"Created_Date"`
}

///provider

type ProviderResponse struct {
	ProviderId    int    `json:"Provider_Id"`
	ProviderName  string `json:"Provider_Name"`
	Description   string `json:"Description"`
	ProviderAlias string `json:"Provider_Alias"`
	Status        int    `json:"Status"`
}

///producttype

type ProductTypeResponse struct {
	Provider_Name     string `json:"Provider_Name"`
	Product_Type_Id   int    `json:"Product_Type_Id"`
	Product_Type_Name string `json:"Product_Type_Name"`
	Description       string `json:"Description"`
	Status            int    `json:"Status"`
}

///productcategory

type ProductCategoryResponse struct {
	ProviderName        string `json:"Provider_Name"`
	ProductTypeName     string `json:"Product_Type_Name"`
	ProductCategoryId   int    `json:"Product_Category_Id"`
	ProductCategoryName string `json:"Product_Category_Name"`
	Status              int    `json:"Status"`
}

///billerproduct

type BillerProductResponse struct {
	ProviderName        string  `json:"Provider_Name"`
	ProductCategoryName string  `json:"Product_Category_Name"`
	BillerProductId     int     `json:"Biller_Product_Id"`
	BillerProductName   string  `json:"Biller_Product_Name"`
	Description         string  `json:"Description"`
	BankCommission      float32 `json:"Bank_Commission"`
	ServiceFee          float32 `json:"Service_Fee"`
	Status              int     `json:"Status"`
}

///loadproduct

type LoadProductResponse struct {
	ProductCategoryName string `json:"Product_Category_Name"`
	LoadProductId       int    `json:"Load_Product_Id"`
	LoadProductName     string `json:"Load_Product_Name"`
	Description         string `json:"Description"`
	Status              int    `json:"Status"`
}

///loadproduct

type CommissionResponse struct {
	TransactionType         string  `json:"Transaction_Type"`
	CommissionType          string  `json:"Commission_Type"`
	CustomerIncome          float32 `json:"Customer_Income"`
	AgentIncome             float32 `json:"Agent_Income"`
	BankIncome              float32 `json:"Bank_Income"`
	StaBankPartnerIncometus float32 `json:"Bank_Partner_Income"`
}

//banklist

type BankListResponse struct {
	BankCode  string `json:"Bank_Code"`
	BankName  string `json:"Bank_Name"`
	ShortName string `json:"Short_Name"`
	BankBic   string `json:"Bank_Bic"`
}

//partnerlist

type PartnerListResponse struct {
	PartnerId                  string `json:"Partner_Id"`
	PartnerName                string `json:"Partner_Name"`
	Description                string `json:"Description"`
	PartnerApiUrl              string `json:"Partner_Api_Url"`
	MerchantPaymentCallbackUrl string `json:"Merchant_Payment_Callback_Url"`
	MerchantIdPrefix           string `json:"Merchant_Id_Prefix"`
	MriGroup                   int    `json:"Mri_Group"`
	Status                     int    `json:"Status"`
}

//splashscreen

type SplashScreenResponse struct {
	Action     string `json:"Action"`
	Title      string `json:"Title"`
	Message    string `json:"Message"`
	SubMessage string `json:"Sub_Message"`
	ImageUrl   string `json:"Image_Url"`
	Show       string `json:"Show"`
}
