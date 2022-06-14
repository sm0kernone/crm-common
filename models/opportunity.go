package models

type Opportunities struct {
	Id                     int    `sql:"id" json:"id,omitempty" view:"hidden" type:"int,disabled" label:"ID"`
	CrmId                  int    `sql:"crm_id" json:"crm_id,omitempty" view:"hidden" type:"int,disabled" label:"CRM ID" type:"readonly"`
	PotentialName          string `sql:"potential_name" json:"potential_name,omitempty" type:"string" view:"watch" label:"Opportunity Name" group:"details,short"`
	QuoteAmount            Amount `sql:"quote_amount" json:"quote_amount,omitempty" view:"watch" type:"float" label:"Quote Amount" group:"details,short"`
	PaidAmount             Amount `sql:"paid_amount" json:"paid_amount,omitempty" view:"watch" label:"Paid Amount" type:"float" group:"details,short"`
	RefundAmount           Amount `sql:"refund_amount" json:"refund_amount,omitempty" view:"watch" label:"Refund Amount" type:"float" group:"details,short"`
	AssignedTo             int    `sql:"assigned_to" json:"assigned_to,omitempty" view:"watch" label:"Assigned to" group:"details,short" related:"users" type:"autocomplete,int"`
	ActualCloseDate        int    `sql:"actual_close_date" json:"actual_close_date,omitempty" view:"hidden" label:"Actual Close date" group:"details"`
	CreatedTime            int    `sql:"created_time" json:"created_time,omitempty" label:"Created Time" view:"watch" type:"datetime,readonly" group:"details"`
	CampaignSource         string `sql:"campaign_source" json:"campaign_source,omitempty" label:"Campaign Source" view:"watch" type:"readonly,string" group:"details"`
	ModifiedTime           int    `sql:"modified_time" json:"modified_time,omitempty" label:"Modified Time" view:"watch" type:"datetime,readonly"`
	RefundType             string `sql:"refund_type" json:"refund_type,omitempty" label:"Refund Type" view:"watch" group:"details,short" type:"autocomplete,string" related:"refund_type"`
	Product                string `sql:"product" json:"product,omitempty" label:"Product" group:"details,short" view:"watch" type:"autocomplete,string" related:"products"`
	StatusReason           string `sql:"status_reason" json:"status_reason,omitempty" view:"watch" related:"status_reason_opportunity" type:"autocomplete,string" label:"Status Reason" group:"details,short"`
	LeadCreated            int    `sql:"lead_created" json:"lead_created,omitempty" view:"watch" type:"datetime" label:"Lead Created"`
	LeadSource             string `sql:"lead_source" json:"lead_source,omitempty" view:"watch" label:"Lead Source" group:"details" type:"autocomplete,string" related:"lead_source"`
	CountryOfResidence     string `sql:"country_of_residence" json:"country_of_residence,omitempty" view:"watch" label:"Country Of Residence" type:"autocomplete,string" group:"details,short" related:"country"`
	ConvertedFromLead      bool   `sql:"converted_from_lead,notnull" pg:",use_zero" json:"converted_from_lead,omitempty" view:"watch" label:"Converted From Lead?" type:"bool,autocomplete" related:"yes_no" group:"details,short"`
	CountryOfBirth         string `sql:"country_of_birth" json:"country_of_birth,omitempty" label:"Country Of Birth" view:"watch" type:"autocomplete,string" group:"details,short" related:"country"`
	PaymentMethod          string `sql:"payment_method" json:"payment_method,omitempty" label:"Payment Method" view:"watch" group:"details,short" type:"autocomplete,string" related:"payment_method"`
	ProductEmailStatus     string `sql:"product_email_status" json:"product_email_status,omitempty" view:"watch" label:"Product Email Status" type:"string" group:"details"`
	AccountNumber          string `sql:"account_number" json:"account_number,omitempty" label:"Account Number" view:"hidden" group:"details,short" type:"readonly,string"`
	OpportunityNumber      string `sql:"opportunity_number" json:"opportunity_number,omitempty" view:"hidden" label:"Opportunity Number" type:"readonly,string" group:"details"`
	FunnelStatusReason     string `sql:"funnel_status_reason" json:"funnel_status_reason,omitempty" view:"watch" label:"Funnel Status Reason" group:"details" type:"autocomplete,string" related:"funnel_status_reason"`
	PreviousCampaignSource string `sql:"previous_campaign_source" json:"previous_campaign_source,omitempty" view:"watch" label:"Previous Campaign Source" group:"details,short" type:"string"`
	Email                  string `sql:"email" json:"email,omitempty" label:"Email" view:"hidden" group:"details,short"`
	Password               string `sql:"password" json:"password,omitempty" label:"Password" view:"hidden" group:"details,short" type:"readonly,string"`
	ContactMe              bool   `sql:"contact_me,notnull" pg:",use_zero" json:"contact_me,omitempty" view:"watch" label:"Contact Me" type:"bool,autocomplete" related:"yes_no" group:"details,short"`
	RefundTime             int    `sql:"refund_time" json:"refund_time,omitempty" label:"Refund Time" view:"watch" type:"datetime" group:"details,short"`
	PaymentTime            int    `sql:"payment_time" json:"payment_time,omitempty" label:"Payment Time" view:"watch" type:"datetime"`
	SingleOrMarried        string `sql:"single_or_married" json:"single_or_married,omitempty" label:"Single Or Married" view:"watch" group:"details,short" type:"autocomplete,string" related:"marital_status"`
	MonthlySalary          string `sql:"monthly_salary" json:"monthly_salary,omitempty" label:"Monthly Salary" view:"watch" group:"details,short" type:"autocomplete,string" related:"salary_range"`
	Gender                 string `sql:"gender" json:"gender,omitempty" label:"Gender" group:"details,short" view:"hidden" type:"autocomplete,string" related:"gender"`
	Birthday               string `sql:"birthday" json:"birthday,omitempty" label:"Birthday" view:"watch" type:"string" group:"details"`
	VisaType               string `sql:"visa_type" json:"visa_type,omitempty" label:"Visa Type" view:"watch" group:"questionnaire" type:"autocomplete,string" related:"visa_type"`
	PreferredDestination   string `sql:"preferred_destination" json:"preferred_destination,omitempty" view:"hidden" label:"Preferred destination" group:"questionnaire" type:"autocomplete,string" related:"preferred_destination"`
	Occupation             string `sql:"occupation" json:"occupation,omitempty" label:"Occupation" view:"watch" group:"questionnaire"`
	CompletedHighSchool    string `sql:"completed_high_school" json:"completed_high_school,omitempty" view:"hidden" label:"Completed High School" type:"autocomplete,string" related:"yes_no" group:"questionnaire"`
	ChooseCurrency         string `sql:"choose_currency" json:"choose_currency,omitempty" label:"Choose Currency" view:"hidden" group:"questionnaire" type:"autocomplete,string" related:"currency"`
	Nationality            string `sql:"nationality" json:"nationality,omitempty" label:"Nationality"  group:"questionnaire" view:"hidden" type:"autocomplete,string" related:"nationality"`
	CurrentlyEmployed      string `sql:"currently_employed" json:"currently_employed,omitempty" label:"Currently Employed" view:"watch" type:"autocomplete,string" related:"yes_no" group:"questionnaire"`
	Employed               string `sql:"employed" json:"employed,omitempty" label:"Employed" view:"hidden" group:"questionnaire"`
	YearsOfEmployment      string `sql:"years_of_employment" json:"years_of_employment,omitempty" view:"watch" label:"Years Of Employment" group:"questionnaire" type:"autocomplete,string" related:"years_of_employment"`
	EducationOrTraining    string `sql:"education_or_training" json:"education_or_training,omitempty" view:"watch" label:"Education Or Training" type:"autocomplete,string" related:"yes_no" group:"questionnaire"`
	ExactBirthday          string `sql:"exact_birthday" json:"exact_birthday,omitempty" label:"Exact Birthday" view:"hidden" group:"questionnaire" type:"datetime,readonly"`
	NetWorth               string `sql:"net_worth" json:"net_worth,omitempty" label:"Net Worth" group:"questionnaire" view:"watch" type:"autocomplete,string" related:"net_worth"`
	TrackerId              string `sql:"tracker_id" json:"tracker_id,omitempty" label:"Medium" type:"readonly,string" view:"hidden" group:"custom"` //This is medium field
	AdGroup                string `sql:"adgroup" json:"adgroup,omitempty" label:"Ad Group" type:"readonly,string" view:"hidden" group:"custom"`
	Keyword                string `sql:"keyword" json:"keyword,omitempty" label:"Keyword" type:"string" view:"hidden" group:"custom,short"`
	MatchType              string `sql:"matchtype" json:"matchtype,omitempty" label:"Match Type" view:"hidden" type:"readonly,string" group:"custom"`
	AdId                   string `sql:"adid" json:"adid,omitempty" label:"Ad ID" type:"readonly,string" view:"hidden" group:"custom"`
	Source                 string `sql:"source" json:"source,omitempty" label:"Source" type:"readonly,string" view:"hidden" group:"custom"`
	AdPosition             string `sql:"adposition" json:"adposition,omitempty" label:"Ad Position" view:"hidden" type:"readonly,string" group:"custom"`
	Placement              string `sql:"placement" json:"placement,omitempty" label:"Placement" view:"hidden" type:"readonly,string" group:"custom"`
	GclId                  string `sql:"gclid" json:"gclid,omitempty" label:"GCL ID" type:"readonly,string" view:"hidden" group:"custom"`
	Device                 string `sql:"device" json:"device,omitempty" label:"Device" type:"readonly,string" view:"hidden" group:"custom"`
	GeoId                  string `sql:"geoid" json:"geoid,omitempty" label:"GEO ID" type:"readonly,string" view:"hidden" group:"custom"`
	Campaign               string `sql:"campaign" json:"campaign,omitempty" label:"Campaign" type:"readonly,string" view:"hidden" group:"custom"`
	Age                    int    `sql:"age" json:"age,omitempty" view:"watch" label:"Age" type:"int"`
	AssistedBySenior       string `sql:"assisted_by_senior" json:"assisted_by_senior,omitempty" view:"watch" label:"Assisted By Senior" type:"autocomplete,string" related:"yes_no" group:"details"`
	Last4                  string `sql:"last_4" json:"last_4,omitempty" type:"string,readonly" view:"watch" label:"Last 4" type:"string" group:"details"`
	PaymentBin             string `sql:"payment_bin" json:"payment_bin,omitempty" type:"string,readonly" view:"watch" label:"Payment Bin" type:"string" group:"details"`
	Attempt                int    `sql:"-" json:"attempt" view:"watch" label:"Attempt"`
	OfficePhone            string `sql:"-" json:"office_phone,omitempty" type:"phone,string" view:"hidden" label:"Lead's Phone"`
	CreatedAt              int    `sql:"created_at" json:"created_at,omitempty" view:"watch" label:"Created At" type:"datetime" type:"readonly.int"`
	UpdatedAt              int    `sql:"updated_at" json:"updated_at,omitempty" view:"watch" label:"Updated At" type:"datetime" type:"readonly,int"`
	PPStats                string `sql:"pp_stats" json:"pp_stats,omitempty" view:"watch" label:"PP Stats" type:"autocomplete,string" related:"pp_stats"`
	WhoClosed              int    `sql:"who_closed" json:"who_closed,omitempty" view:"watch" label:"Who Closed" type:"autocomplete,int" related:"users"`
	ContactName            string `sql:"contact_name" json:"contact_name,omitempty" label:"Contact Name" type:"string" group:"details,short" view:"hidden"`
	EmailCampaign          string `sql:"email_campaign" json:"email_campaign,omitempty" label:"Email Campaign" type:"string" group:"details,short" view:"hidden"`
	EmailTemplate          string `sql:"email_template" json:"email_template,omitempty" label:"Email Template" type:"string" group:"details,short" view:"hidden"`
	EmailSegment           string `sql:"email_segment" json:"email_segment,omitempty" label:"Email Segment" group:"details,short" type:"string" view:"hidden"`
	Deleted                bool   `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" view:"watch" label:"Deleted" type:"bool,autocomplete" related:"yes_no"`
	CardType               string `sql:"-" json:"card_type,omitempty" view:"watch" label:"Card Type" type:"string,autocomplete" related:"card_type"`
	Card                   string `sql:"-" json:"card,omitempty" view:"watch" label:"Card" type:"string,autocomplete" related:"card"`
	CardClass              string `sql:"-" json:"card_class,omitempty" view:"watch" label:"Card Class" type:"string,autocomplete" related:"card_class"`
	BinCountry             string `sql:"-" json:"bin_country,omitempty" view:"watch" label:"Bin Countries" type:"readonly,string"`

	Online bool `sql:"-" json:"online" view:"watch" label:"Online" view:"watch" type:"bool,autocomplete" related:"yes_no"`

	FBCLid           string `sql:"fbclid" json:"fbclid,omitempty" label:"FBCLid" group:"details" type:"string" view:"hidden"`
	UtmTerm          string `sql:"utm_term" json:"utm_term,omitempty" label:"Utm Term" group:"details" type:"string" view:"hidden"`
	DeviceModel      string `sql:"devicemodel" json:"devicemodel,omitempty" label:"Devicemodel" group:"details" type:"string" view:"hidden"`
	OrganizationName string `sql:"organization_name" json:"organization_name,omitempty" label:"Organization Name" group:"details" type:"string" view:"hidden"`
	Phone            string `sql:"phone" json:"phone,omitempty" label:"Phone" group:"details" type:"string" view:"watch"`
	LastEventTime    int    `sql:"last_event_time" json:"last_event_time,omitempty" view:"watch" label:"Last Event Time" group:"details,short" type:"datetime,readonly"`
	PaymentService   string `sql:"payment_service" json:"payment_service,omitempty" view:"watch" label:"Payment Service" type:"string"`
	IsRiskBin        bool   `sql:"is_risk_bin,notnull" pg:",use_zero" json:"is_risk_bin,omitempty" view:"watch" label:"Is Risk Bin" type:"bool,autocomplete" related:"yes_no" group:"-"`
	IsRiskIp         bool   `sql:"is_risk_ip,notnull" pg:",use_zero" json:"is_risk_ip,omitempty" view:"watch" label:"Is Risk IP" type:"bool,autocomplete" related:"yes_no" group:"-"`
	IsDuplicate      bool   `sql:"is_duplicate,notnull" pg:",use_zero" json:"is_duplicate,omitempty" view:"watch" label:"Is Duplicate" type:"bool,autocomplete" related:"yes_no" group:"-"`

	// Related Fields From Leads
	LeadId                 int    `sql:"-" json:"lead_id,omitempty" type:"int,disabled" view:"hidden" label:"Lead ID" `
	LeadQuoteAmount        string `sql:"-" json:"lead_quote_amount,omitempty" view:"hidden" type:"float,disabled" label:"Lead Quote amount"`
	LeadPassengerOne       string `sql:"-" json:"lead_passenger_one,omitempty" view:"hidden" type:"string,readonly" label:"Lead Passenger One"`
	LeadNumberOfPassengers int    `sql:"-" json:"lead_number_of_passengers,omitempty" view:"hidden" type:"int,disabled" label:"Lead Number Of Passengers"`
	LeadProduct            string `sql:"-" json:"lead_product,omitempty" view:"hidden" type:"string,readonly" label:"Lead Product"`

	// Related Fields From Applications
	AppMailingStreet string `sql:"-" json:"app_mailing_street,omitempty" view:"watch" label:"Receipt Signature" type:"string,readonly"`
	AppMailingPoBox  string `sql:"-" json:"app_mailing_po_box,omitempty" view:"watch" label:"Receipt Status" type:"autocomplete,string,readonly" related:"receipt_status"`
	AppFormStatus    string `sql:"-" json:"app_form_status,omitempty" view:"watch" label:"Form Status" type:"autocomplete,string,readonly" related:"form_status"`
	IsRisk           bool   `sql:"-" json:"is_risk" view:"watch" pg:",use_zero" label:"Is Risk" type:"bool,autocomplete" related:"yes_no"`
	ContactStatus    string `sql:"-" json:"contact_status" view:"watch" label:"Contact Status" type:"readonly,autocomplete,string" related:"contact_me_status"`

	// Related Fields From Instalments
	InstalmentStatus int `sql:"-" json:"instalment_status,omitempty" label:"Installments Status" type:"autocomplete,string,readonly" view:"watch" related:"instalments_status"`
	InstalmentCount  int `sql:"-" json:"instalment_count,omitempty" label:"Installments Count" type:"int,readonly" view:"watch"`

	// Related Fields to Forms
	HasForm                bool   `sql:"-" json:"has_form,omitempty" label:"Has Feedback Form" type:"readonly,bool" related:"yes_no"`
	FeedbackFormTypes      string `sql:"-" json:"feedback_form_types,omitempty" label:"Form Feedback Types" type:"readonly,string,autocomplete" view:"watch" related:"feedback_form_types"`
	FeedbackFormStatuses   string `sql:"-" json:"feedback_form_statuses,omitempty" label:"Form Feedback Statuses" type:"readonly,string,autocomplete" view:"watch" related:"feedback_form_statuses"`
	FeedbackFormAnswer     string `sql:"-" json:"feedback_form_answer,omitempty" label:"Form Feedback Answer" type:"readonly,string,autocomplete" view:"watch" related:"feedback_form_answer"`
	FeedbackFormFieldTypes string `sql:"-" json:"feedback_form_field_types,omitempty" label:"Form Feedback Field Types" type:"readonly,string,autocomplete" view:"watch" related:"feedback_form_field_types"`
	FeedbackFieldNames     string `sql:"-" json:"feedback_field_names,omitempty" label:"Form Feedback Question" type:"readonly,string,autocomplete" view:"watch" related:"feedback_field_names"`
}

var OnlineProductsArray = []string{"Canada Online Evaluation"}
