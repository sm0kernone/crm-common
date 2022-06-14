package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
	"strconv"
)

const (
	PaymentHistoryStatusCreatedValue = iota
	PaymentHistoryStatusSuccessValue
	PaymentHistoryStatusRejectedValue
	PaymentHistoryStatusPendingValue
	PaymentHistoryStatusExpiredValue
	PaymentHistoryStatusDeclinedValue
	PaymentHistoryStatusFailedValue
	PaymentHistoryStatusActivatedValue

	PaymentHistoryStatusCreated  = "created"  // simple type
	PaymentHistoryStatusRejected = "rejected" // simple type
	PaymentHistoryStatusSuccess  = "success"  // simple type

	PaymentHistoryStatusPending   = "pending"
	PaymentHistoryStatusExpired   = "expired"
	PaymentHistoryStatusDeclined  = "declined"
	PaymentHistoryStatusFailed    = "failed"
	PaymentHistoryStatusActivated = "activated"
)

const (
	PaymentHistoryServiceCascadeCheckoutSessionsValue = iota
	PaymentHistoryServiceUsersInfoValue
	PaymentHistoryServiceFunnelValue
	PaymentHistoryServicePaymentFunnelV3Value
	PaymentHistoryServicePaymentFunnelV2Value
	PaymentHistoryServiceBankTransferValue
	PaymentHistoryServiceNewImmiCVValue
	PaymentHistoryServiceOnlineBCValue
	PaymentHistoryServicePaymentFunnelV8Value
	PaymentHistoryServicePaymentFunnelV3NewValue
	PaymentHistoryServicePaymentFunnelV9Value

	PaymentHistoryServiceCascadeCheckoutSessions = "cascadecheckoutsessions"
	PaymentHistoryServiceUsersInfo               = "usersinfo"
	PaymentHistoryServiceFunnel                  = "funnel"
	PaymentHistoryServicePaymentFunnelV3         = "paymentfunnelv3"
	PaymentHistoryServicePaymentFunnelV2         = "paymentfunnelv2"
	PaymentHistoryServiceBankTransfer            = "banktransfer"
	PaymentHistoryServiceNewImmiCV               = "newimmicv"
	PaymentHistoryServiceOnlineBC                = "onlinebc"
	PaymentHistoryServicePaymentFunnelV8         = "funnelv8"
	PaymentHistoryServicePaymentFunnelV3New      = "funnelv3"
	PaymentHistoryServicePaymentFunnelV9         = "funnelv9"
)

var FunnelServices = []int{PaymentHistoryServiceFunnelValue, PaymentHistoryServicePaymentFunnelV3Value,
	PaymentHistoryServicePaymentFunnelV2Value, PaymentHistoryServicePaymentFunnelV8Value, PaymentHistoryServicePaymentFunnelV3NewValue,
	PaymentHistoryServicePaymentFunnelV9Value,
}

const (
	PaymentHistoryMethodVoguePayCurValue = iota + 1
	PaymentHistoryMethodTrustPay3dValue
	PaymentHistoryMethodDirectPayValue
	PaymentHistoryMethodSolid3dValue
	PaymentHistoryMethodSolidValue
	PaymentHistoryMethodTrustPayValue
	PaymentHistoryMethodVoguePayValue
	PaymentHistoryMethodFlutterWaveValue
	PaymentHistoryMethodCascadingValue
	PaymentHistoryMethodTrust3dValue
	PaymentHistoryMethodTrustValue
	PaymentHistoryMethodPayPalValue
	PaymentHistoryMethodDectaValue
	PaymentHistoryMethodDectaMotoValue
	PaymentHistoryMethodCreditCardValue
	PaymentHistoryMethodCardPayValue

	// must be trim + lower
	PaymentHistoryMethodVoguePayCur = "voguepaycur"
	PaymentHistoryMethodTrustPay3d  = "trustpay3d"
	PaymentHistoryMethodDirectPay   = "directpay"
	PaymentHistoryMethodSolid3d     = "solid3d"
	PaymentHistoryMethodSolid       = "solid"
	PaymentHistoryMethodTrustPay    = "trustpay"
	PaymentHistoryMethodVoguePay    = "voguepay"
	PaymentHistoryMethodFlutterWave = "flutterwave"
	PaymentHistoryMethodCascading   = "cascading"
	PaymentHistoryMethodTrust3d     = "trust3d"
	PaymentHistoryMethodTrust       = "trust"
	PaymentHistoryMethodPayPal      = "paypal"
	PaymentHistoryMethodDecta       = "decta"
	PaymentHistoryMethodDectaMoto   = "dectamoto"
	PaymentHistoryMethodCreditCard  = "creditcard"
	PaymentHistoryMethodCardPay     = "cardpay"
)

var PaymentHistoryServiceInfo = map[int]struct {
	Label  string
	Public string
}{
	PaymentHistoryServiceCascadeCheckoutSessionsValue: {Label: PaymentHistoryServiceCascadeCheckoutSessions, Public: "Cascade"},
	PaymentHistoryServiceUsersInfoValue:               {Label: PaymentHistoryServiceUsersInfo, Public: "Payment Link"},
	PaymentHistoryServiceFunnelValue:                  {Label: PaymentHistoryServiceFunnel, Public: "Funnel"},
	PaymentHistoryServicePaymentFunnelV3Value:         {Label: PaymentHistoryServiceFunnel, Public: "Funnel V3"},
	PaymentHistoryServicePaymentFunnelV2Value:         {Label: PaymentHistoryServiceFunnel, Public: "Funnel V2"},
	PaymentHistoryServiceBankTransferValue:            {Label: PaymentHistoryServiceBankTransfer, Public: "Bank Transfer"},
	PaymentHistoryServiceNewImmiCVValue:               {Label: PaymentHistoryServiceNewImmiCV, Public: "New Immi CV"},
	PaymentHistoryServiceOnlineBCValue:                {Label: PaymentHistoryServiceOnlineBC, Public: "Online BC"},
	PaymentHistoryServicePaymentFunnelV8Value:         {Label: PaymentHistoryServicePaymentFunnelV8, Public: "Funnel V8"},
	PaymentHistoryServicePaymentFunnelV3NewValue:      {Label: PaymentHistoryServicePaymentFunnelV3New, Public: "Immi funnel"},
	PaymentHistoryServicePaymentFunnelV9Value:         {Label: PaymentHistoryServicePaymentFunnelV9, Public: "Funnel V9"},
}

var PaymentHistoryStatusInfo = map[int]struct {
	Label  string
	Public string
}{
	PaymentHistoryStatusCreatedValue:   {Label: PaymentHistoryStatusCreated, Public: "Created"},
	PaymentHistoryStatusRejectedValue:  {Label: PaymentHistoryStatusRejected, Public: "Rejected"},
	PaymentHistoryStatusSuccessValue:   {Label: PaymentHistoryStatusSuccess, Public: "Success"},
	PaymentHistoryStatusPendingValue:   {Label: PaymentHistoryStatusPending, Public: "Pending"},
	PaymentHistoryStatusExpiredValue:   {Label: PaymentHistoryStatusExpired, Public: "Expired"},
	PaymentHistoryStatusDeclinedValue:  {Label: PaymentHistoryStatusDeclined, Public: "Declined"},
	PaymentHistoryStatusFailedValue:    {Label: PaymentHistoryStatusFailed, Public: "Failed"},
	PaymentHistoryStatusActivatedValue: {Label: PaymentHistoryStatusActivated, Public: "Activated"},
}

var PaymentHistoryMethodInfo = map[int]struct {
	Label  string
	Public string
}{
	PaymentHistoryMethodVoguePayCurValue: {Label: PaymentHistoryMethodVoguePayCur, Public: "VoguePayCur"},
	PaymentHistoryMethodTrustPay3dValue:  {Label: PaymentHistoryMethodTrustPay3d, Public: "TrustPay3D"},
	PaymentHistoryMethodDirectPayValue:   {Label: PaymentHistoryMethodDirectPay, Public: "DirectPay"},
	PaymentHistoryMethodSolid3dValue:     {Label: PaymentHistoryMethodSolid3d, Public: "Solid3d"},
	PaymentHistoryMethodSolidValue:       {Label: PaymentHistoryMethodSolid, Public: "Solid"},
	PaymentHistoryMethodTrustPayValue:    {Label: PaymentHistoryMethodTrustPay, Public: "TrustPay"},
	PaymentHistoryMethodVoguePayValue:    {Label: PaymentHistoryMethodVoguePay, Public: "VoguePay"},
	PaymentHistoryMethodFlutterWaveValue: {Label: PaymentHistoryMethodFlutterWave, Public: "FlutterWave"},
	PaymentHistoryMethodCascadingValue:   {Label: PaymentHistoryMethodCascading, Public: "Cascading"},
	PaymentHistoryMethodTrust3dValue:     {Label: PaymentHistoryMethodTrust3d, Public: "Trust3d"},
	PaymentHistoryMethodTrustValue:       {Label: PaymentHistoryMethodTrust, Public: "Trust"},
	PaymentHistoryMethodPayPalValue:      {Label: PaymentHistoryMethodPayPal, Public: "PayPal"},
	PaymentHistoryMethodDectaValue:       {Label: PaymentHistoryMethodDecta, Public: "Decta"},
	PaymentHistoryMethodDectaMotoValue:   {Label: PaymentHistoryMethodDectaMoto, Public: "DectaMoto"},
	PaymentHistoryMethodCreditCardValue:  {Label: PaymentHistoryMethodCreditCard, Public: "Credit Card"},
	PaymentHistoryMethodCardPayValue:     {Label: PaymentHistoryMethodCardPay, Public: "Card Pay"},
}

var RetentionProducts = []string{
	"RCIC Fees Payment - UPGRADE",
	"Registration Fee",
	"File Opening",
	"IELTS Ashton - Retention - General Course",
	"IELTS Ashton - Retention - Academic Course",
	"File Opening Referral",
	"IELTS Online Course",
	"Documents Verification",
}

type PaymentDetails map[string]interface{}

func (a PaymentDetails) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *PaymentDetails) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type SpouseInfo struct {
	Id                 int    `json:"id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	CountryOfResidence string `json:"country_of_residence"`
}

type PaymentHistory struct {
	tableName       struct{}       `sql:"payment_history" pg:"payment_history" json:"-"`
	Id              int            `sql:"id" json:"id" label:"ID"`
	PaymentId       string         `sql:"payment_id" json:"payment_id" validate:"required,min=1,max=255" label:"Transaction ID"`
	RequestId       *int           `sql:"request_id" json:"request_id" label:"Transaction Request ID"`
	UserCrmId       int            `sql:"user_crm_id,notnull" json:"user_crm_id" validate:"required" label:"User Crm ID"`
	LeadCrmId       int            `sql:"lead_crm_id,notnull" json:"lead_crm_id" validate:"required" label:"Lead Crm ID"`
	Product         string         `sql:"product" json:"product" validate:"required,min=1,max=255" type:"autocomplete,string" related:"products" label:"Product"`
	FromRelCrmId    int            `sql:"from_rel_crm_id,notnull" json:"from_rel_crm_id" validate:"required" label:"Related Crm Id"`
	FromRelType     int            `sql:"from_rel_type,notnull" pg:",use_zero" json:"from_rel_type" label:"Related Type" type:"autocomplete,string" related:"transactions_types"`
	Status          int            `sql:"status,notnull" pg:",use_zero" json:"status" label:"Status" type:"autocomplete,string" related:"transactions_statuses"`
	Service         int            `sql:"service,notnull" pg:",use_zero" json:"service" label:"Service" type:"autocomplete,string" related:"transactions_services"`
	Method          *int           `sql:"method" json:"method,omitempty" type:"autocomplete,string" label:"Method" related:"transactions_methods"`
	AgentType       int            `sql:"agent_type" json:"agent_type" label:"Agent Type" type:"autocomplete,string" related:"transactions_agent_types"`
	Action          *int           `sql:"action" json:"action,omitempty" type:"autocomplete,string" label:"Action" related:"transactions_actions"`
	PaidAt          *int           `sql:"paid_at" json:"paid_at,omitempty" label:"Paid At" type:"datetime"`
	PaidAmount      Amount         `sql:"paid_amount,notnull" pg:",use_zero" json:"paid_amount,omitempty" type:"float" label:"Paid Amount"`
	QuoteAmount     Amount         `sql:"quote_amount,notnull" pg:",use_zero" json:"quote_amount,omitempty" type:"float" label:"Quote Amount"`
	Details         interface{}    `sql:"details" pg:",json_use_number" json:"details,omitempty" label:"Details"`
	PaymentDetails  PaymentDetails `sql:"payment_details" pg:",json_use_number" json:"payment_details,omitempty" label:"Payment Details"`
	PaymentResponse string         `sql:"-" json:"payment_response,omitempty" label:"Payment Response" type:"string,readonly"`

	Deleted     bool        `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" label:"Deleted"`
	CreatedAt   int         `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt   int         `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
	CreatedTime int         `sql:"-" json:"created_time" label:"Created Time" type:"datetime"`
	AssignTime  int         `sql:"-" json:"assign_time" label:"Assign Time" type:"datetime"`
	Ip          string      `sql:"ip" json:"ip" label:"IP"`
	IpDetails   interface{} `sql:"ip_details" pg:",json_use_number" json:"ip_details,omitempty" label:"IP Details"`

	IpCountry string `sql:"-" json:"ip_country,omitempty" type:"string,readonly" view:"hidden" type:"autocomplete,string" related:"country" label:"Ip Country"`
	IpCity    string `sql:"-" json:"ip_city,omitempty" type:"string,readonly" view:"hidden" type:"string" label:"Ip City"`
	IpRegion  string `sql:"-" json:"ip_region,omitempty" type:"string,readonly" view:"hidden" type:"string" label:"Ip Region"`

	IsDuplicate string `sql:"-" json:"is_duplicate,omitempty" type:"string,readonly" view:"hidden" type:"autocomplete,string" related:"yes_no" label:"Is Duplicate"`
	IsIPRisk    string `sql:"-" json:"is_ip_risk,omitempty" type:"string,readonly" view:"hidden" type:"autocomplete,string" related:"yes_no" label:"Is IP Risk"`
	IsBinRisk   string `sql:"-" json:"is_bin_risk,omitempty" type:"string,readonly" view:"hidden" type:"autocomplete,string" related:"yes_no" label:"Is BIN Risk"`

	UserId int `sql:"-" json:"user_id,omitempty" type:"int,disabled" related:"users" type:"autocomplete,int" label:"User ID"`

	LeadId                 int    `sql:"-" json:"lead_id,omitempty" type:"int,disabled" view:"hidden" label:"Lead ID"`
	LeadAccountNumber      string `sql:"-" json:"lead_account_number,omitempty" type:"string,readonly" view:"hidden" label:"Account"`
	LeadFullName           string `sql:"-" json:"lead_full_name,omitempty" type:"string,readonly" view:"hidden" label:"Full Name"`
	LeadCountryOfBirth     string `sql:"-" json:"lead_country_of_birth,omitempty" view:"hidden" type:"autocomplete,string,readonly" related:"country" label:"Country Of Birth"`
	LeadCountryOfResidence string `sql:"-" json:"lead_country_of_residence,omitempty" view:"hidden" type:"autocomplete,string,readonly" related:"country" label:"Country Of Residence"`
	LeadCampaignSource     string `sql:"-" json:"lead_campaign_source,omitempty" type:"string,readonly" view:"hidden" label:"Campaign Source"`

	IsInst     bool   `sql:"-" json:"is_inst,omitempty" type:"bool,autocomplete" view:"hidden" label:"Is Installment" related:"yes_no"`
	InstStatus string `sql:"-" json:"inst_status,omitempty" type:"autocomplete,string,readonly" view:"hidden" label:"Installments Status" related:"instalments_status" `
	InstPos    string `sql:"-" json:"inst_pos,omitempty" type:"string,readonly" view:"hidden" label:"Installment Position"`

	OpportunityId int `sql:"-" json:"opportunity_id,omitempty" type:"int,disabled" view:"hidden" label:"Opportunity Id"`

	Bin         string `sql:"-" json:"bin,omitempty" type:"string,disabled" view:"hidden" label:"Bin"`
	Last4Digits string `sql:"-" json:"last_4_digits,omitempty" string:"int,disabled" view:"hidden" label:"Last 4 Digits"`
	BinCountry  string `sql:"-" json:"bin_country,omitempty" type:"string,disabled" view:"hidden" label:"Bin Country"`

	PaymentPaidAmount   Amount `sql:"-" json:"payment_paid_amount" view:"hidden" label:"Payment Paid Amount" type:"float"`
	PaymentQuoteAmount  Amount `sql:"-" json:"payment_quote_amount" view:"hidden" label:"Payment Quote Amount" type:"float"`
	OutstandingAmount   Amount `sql:"-" json:"outstanding_amount" view:"hidden" label:"Outstanding Amount" type:"float"`
	AgentName           string `sql:"-" json:"agent_name" label:"Agent Name" view:"hidden" type:"string"`
	Installment         string `sql:"-" json:"installment" label:"Installment" view:"hidden" related:"yes_no" `
	NextInstallmentDate int    `sql:"-" json:"next_installment_date" label:"Next Installment Date" view:"hidden"`

	ChangedAttributes []string `sql:"-" json:"changed_attributes,omitempty"`
}

type TransactionRefundStats struct {
	tableName struct{} `sql:"transaction_refund_stats" pg:"select:transaction_refund_stats" json:"-"`
	Uuid      string   `sql:"uuid" json:"uuid" label:"Uuid"`
	UserId    int      `sql:"user_id" json:"user_id" label:"User Id"`
	Amount    float64  `sql:"amount" json:"amount" label:"Amount"`
	Product   string   `sql:"product" json:"product" label:"Product"`
	At        int      `sql:"at" json:"at" label:"At" type:"datetime"`
	IsSale    bool     `sql:"is_sale" json:"is_sale" label:"Is Sale"`
}

type Amount float64

func (n Amount) Value() (driver.Value, error) {
	return math.Round(float64(n)*100) / 100, nil
}

func (n *Amount) Scan(value interface{}) error {
	var wp sql.NullFloat64
	var err error
	var aVal interface{}

	err = wp.Scan(value)

	if err != nil {
		return err
	}

	aVal, err = wp.Value()

	if err != nil {
		return err
	}

	if aVal != nil {
		num, _ := strconv.ParseFloat(
			strconv.FormatFloat(
				math.Round(aVal.(float64)*100)/100,
				'f',
				2,
				64,
			),
			64,
		)
		*n = Amount(num)
	}

	return nil
}
