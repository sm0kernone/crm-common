package models

type Refunds struct {
	Id            int          `sql:"id" json:"id" label:"ID"`
	AccountNumber string       `sql:"account_number" json:"account_number,omitempty" label:"Account Number" type:"readonly,string"`
	AssignedTo    int          `sql:"assigned_to" json:"assigned_to,omitempty" view:"watch" label:"Assigned to" related:"users" type:"autocomplete,int"`
	ClosedAt      int          `sql:"closed_at" json:"closed_at,omitempty" view:"watch" label:"Closed At" type:"datetime" related:"users" type:"autocomplete,int"`
	ClosedBy      int          `sql:"closed_by" json:"closed_by,omitempty" view:"watch" label:"Closed By" type:"int,autocomplete" related:"users"`
	PaymentSystem string       `sql:"payment_system" json:"payment_system,omitempty" label:"Payment System" view:"watch" type:"string"`
	RelatedTo     int          `sql:"related_to" json:"related_to,omitempty" view:"watch" label:"Related To" type:"readonly,int"`
	RelatedType   RefundType   `sql:"related_type" json:"related_type,omitempty" label:"Related Type" type:"string"`
	RefundAmount  Amount       `sql:"refund_amount" json:"refund_amount,omitempty" label:"Refund Amount"`
	RefundStatus  RefundStatus `sql:"refund_status" json:"refund_status,omitempty" label:"Refund Status" type:"autocomplete,string" view:"watch" related:"refund_statuses"`
	RefundTime    int          `sql:"refund_time" json:"refund_time,omitempty" label:"Refund Time" view:"watch" type:"datetime"`
	RefundType    string       `sql:"refund_type" json:"refund_type,omitempty" label:"Refund Type" type:"autocomplete,string" view:"watch" related:"refund_type"`
	RefundReason  string       `sql:"refund_reason" json:"refund_reason,omitempty" label:"Refund Reason" type:"autocomplete,string" view:"watch" related:"refund_reasons"`
	SendEmail     bool         `sql:"send_email,notnull" pg:",use_zero" json:"send_email,omitempty" view:"watch" label:"Send Email" type:"bool,autocomplete" related:"yes_no"`
	Deleted       bool         `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" view:"watch" label:"Deleted" type:"bool,autocomplete" related:"yes_no"`
	CreatedAt     int          `sql:"created_at" json:"created_at,omitempty" view:"watch" label:"Created At" type:"datetime" type:"readonly,int"`
	UpdatedAt     int          `sql:"updated_at" json:"updated_at,omitempty" view:"watch" label:"Updated At" type:"datetime" type:"readonly,int"`
	TransactionId int          `sql:"transaction_id" json:"transaction_id" view:"hidden" label:"Transaction Id" type:"int,readonly"`
	AgentType     int          `sql:"agent_type" json:"agent_type" view:"hidden" label:"Agent Type" type:"int,readonly"`

	LeadId                 int    `sql:"-" json:"lead_id,omitempty" type:"int,disabled" view:"hidden" label:"Lead ID" `
	LeadQuoteAmount        string `sql:"-" json:"lead_quote_amount,omitempty" view:"hidden" type:"float,disabled" label:"Lead Quote amount"`
	LeadPassengerOne       string `sql:"-" json:"lead_passenger_one,omitempty" view:"hidden" type:"string,readonly" label:"Lead Passenger One"`
	LeadNumberOfPassengers int    `sql:"-" json:"lead_number_of_passengers,omitempty" view:"hidden" type:"int,disabled" label:"Lead Number Of Passengers"`
	LeadProduct            string `sql:"-" json:"lead_product,omitempty" view:"hidden" type:"string,readonly" label:"Lead Product"`
	LeadEmail              string `sql:"-" json:"lead_email,omitempty" view:"hidden" type:"string,readonly" label:"Lead Email"`

	OppPotentialName string `sql:"-" json:"opp_potential_name" view:"hidden" type:"string,readonly" label:"Opportunity Name"`

	PhPaymentId string `sql:"-" json:"ph_payment_id,omitempty" view:"hidden" type:"string,readonly" label:"PaymentId"`
	PhMethod    *int   `sql:"-" json:"ph_method,omitempty" view:"hidden" type:"autocomplete,string" label:"Method" related:"transactions_methods"`
	PhPaidAt    *int   `sql:"-" json:"ph_paid_at,omitempty" view:"hidden" label:"Paid At" type:"datetime"`

	// Fields Related To Application
	IsRisk       bool `sql:"-" json:"is_risk" view:"watch" pg:",use_zero" label:"Is Risk" type:"bool,autocomplete" related:"yes_no"`
	RiskAssignTo int  `sql:"-" json:"risk_assign_to,omitempty" label:"Risk Assign To" type:"autocomplete,int,readonly" related:"users" view:"watch"`
}

type RefundType string

type RefundStatus string

var (
	RefundTypeChargeback    RefundType = "Chargeback"
	RefundTypeFullRefund    RefundType = "Full Refund"
	RefundTypePartialRefund RefundType = "Partial Refund"

	RefundStatusOpen     RefundStatus = "Open"
	RefundStatusApproved RefundStatus = "Approved"
	RefundStatusRejected RefundStatus = "Rejected"
)
