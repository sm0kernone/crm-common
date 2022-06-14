package models

type Installments struct {
	Id               int    `sql:"id" json:"id" label:"ID"`
	LeadNumber       string `sql:"lead_number" json:"lead_number" validate:"required,min=1,max=255" type:"string" label:"Lead Number"`
	Product          string `sql:"product" json:"product" validate:"required,min=1,max=255" type:"autocomplete,string" related:"products" label:"Product"`
	InitialPaymentId int    `sql:"initial_payment_id" json:"initial_payment_id" validate:"required,min=1" label:"Initial Payment ID"`
	OpportunityId    *int   `sql:"opportunity_id" json:"opportunity_id,omitempty" label:"Opportunity ID"`
	UserId           int    `sql:"user_id" json:"user_id" validate:"required,min=1,max=255" type:"int,disabled" related:"users" type:"autocomplete,int" label:"User ID"`
	LeadId           int    `sql:"lead_id" json:"lead_id" validate:"required,min=1,max=255" label:"Lead ID"`
	PaidAmount       Amount `sql:"paid_amount,notnull" pg:",use_zero" json:"paid_amount,omitempty" type:"float" label:"Paid Amount"`
	QuoteAmount      Amount `sql:"quote_amount,notnull" pg:",use_zero" json:"quote_amount,omitempty" type:"float" label:"Quote Amount"`
	Status           int    `sql:"status,notnull" pg:",use_zero" json:"status" label:"Status" type:"autocomplete,string" related:"instalments_statuses"`
	Method           int    `sql:"method,notnull" pg:",use_zero" json:"method" label:"Method" type:"autocomplete,string" related:"instalments_methods"`
	Period           int    `sql:"period,notnull" pg:",use_zero" json:"period" label:"Period" type:"autocomplete,string" related:"instalments_periods"`
	PeriodDays       int    `sql:"period_days,notnull" pg:",use_zero" json:"period_days" label:"Period Days" type:"int"`
	Deleted          bool   `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" label:"Deleted"`
	CreatedAt        int    `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt        int    `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`

	ChangedAttributes []string `sql:"-" json:"changed_attributes,omitempty"`
}

const (
	InstalmentsTypeLeadsValue = iota
	InstalmentsTypeApplicationsValue
	InstalmentsTypeOpportunitiesValue
	InstalmentsTypeSwipesValue
	InstalmentsTypePartialsValue

	// must be trim + lower
	InstalmentsTypeLeads         = "leads"
	InstalmentsTypeApplications  = "applications"
	InstalmentsTypeOpportunities = "opportunities"
	InstalmentsTypeSwipes        = "swipes"
	InstalmentsTypePartials      = "partials"
)

const (
	InstalmentsStatusCreatedValue = iota
	InstalmentsStatusSuccessValue
	InstalmentsStatusFailedValue
	InstalmentsStatusPendingValue
	InstalmentsStatusProcessingValue
	InstalmentsStatusExpiredValue
	InstalmentsStatusDeclinedValue

	// must be trim + lower
	InstalmentsStatusCreated = "created" // simple type
	InstalmentsStatusFailed  = "failed"  // simple type
	InstalmentsStatusSuccess = "success" // simple type

	InstalmentsStatusProcessing = "processing"
	InstalmentsStatusPending    = "pending"
	InstalmentsStatusExpired    = "expired"
	InstalmentsStatusDeclined   = "declined"
)

const (
	InstalmentsPeriodsEveryMonthValue = iota
	InstalmentsPeriodsEveryWeekValue
	InstalmentsPeriodsEveryDayValue
	InstalmentsPeriodsTwiceEveryMonthValue
	InstalmentsPeriodsTwiceEveryWeekValue
	InstalmentsPeriodsCustomValue

	// must be trim + lower
	InstalmentsPeriodsEveryMonth      = "every_month"
	InstalmentsPeriodsEveryWeek       = "every_week"
	InstalmentsPeriodsEveryDay        = "every_day"
	InstalmentsPeriodsTwiceEveryMonth = "twice_every_month"
	InstalmentsPeriodsTwiceEveryWeek  = "twice_every_week"
	InstalmentsPeriodsCustom          = "custom"
)

const (
	InstalmentsMethodEmailValue = iota
	InstalmentsMethodSmsValue
	InstalmentsMethodEmailSmsValue

	// must be trim + lower
	InstalmentsMethodEmail    = "email"
	InstalmentsMethodSms      = "sms"
	InstalmentsMethodEmailSms = "email_sms"
)
