package models

type Applications struct {
	Id                      int               `sql:"id" json:"id,omitempty" view:"hidden" type:"int,disabled" label:"ID"`
	CrmId                   int               `sql:"crm_id" json:"crm_id,omitempty" view:"hidden" type:"int,disabled" label:"CRM ID" bson:"crm_id"`
	CreatedUserCrmId        int               `sql:"created_user_crm_id" view:"watch" json:"created_user_crm_id,omitempty" label:"Created User CRM ID" related:"users" type:"autocomplete,int" bson:"created_user_crm_id"`
	AssignedUserCrmId       int               `sql:"assigned_user_crm_id" view:"watch" json:"assigned_user_crm_id,omitempty" label:"Assigned To" group:"short,details" type:"autocomplete,int" related:"users" bson:"assigned_user_crm_id"`
	ModifiedUserCrmId       int               `sql:"modified_user_crm_id" view:"watch" json:"modified_user_crm_id,omitempty" label:"Modified User CRM ID" related:"users" bson:"modified_user_crm_id"`
	OpportunityNumber       string            `sql:"opportunity_number" view:"watch" json:"opportunity_number,omitempty" label:"Opportunity Number" type:"readonly,string" group:"portal" bson:"opportunity_number"`
	FirstName               string            `sql:"first_name" json:"first_name,omitempty" view:"watch" label:"First Name" group:"details,short" bson:"first_name"`
	LastName                string            `sql:"last_name" json:"last_name,omitempty" view:"hidden" label:"Last Name" group:"details,short" bson:"last_name"`
	SubmittedCV             bool              `sql:"submitted_cv,notnull" pg:",use_zero" json:"submitted_cv,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Submitted CV?" group:"details,short" bson:"submitted_cv"`
	SubmittedCertificate    bool              `sql:"submitted_certificate,notnull" pg:",use_zero" json:"submitted_certificate,omitempty" type:"bool,autocomplete" related:"yes_no" view:"watch" label:"Submitted IELTS/TOEFL?" group:"details,short" bson:"submitted_certificate"`
	SubmittedPhoto          bool              `sql:"submitted_photo,notnull" pg:",use_zero" json:"submitted_photo,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Submitted Photo?" group:"details,short" bson:"submitted_photo"`
	SubmittedEducation      bool              `sql:"submitted_education,notnull" pg:",use_zero" json:"submitted_education,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Submitted Education?" group:"details,short" bson:"submitted_education"`
	SubmittedPassportCopy   bool              `sql:"submitted_passport_copy,notnull" pg:",use_zero" json:"submitted_passport_copy,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Submitted Passport Copy?" group:"details,short" bson:"submitted_passport_copy"`
	SubmittedUtilityBill    bool              `sql:"submitted_utility_bill,notnull" pg:",use_zero" view:"watch" json:"submitted_utility_bill,omitempty" type:"bool,autocomplete" related:"yes_no" label:"Submitted Utility Bill?" group:"details,short" bson:"submitted_utility_bill"`
	ApplicationCreated      bool              `sql:"application_created,notnull" pg:",use_zero" view:"hidden" json:"application_created,omitempty" label:"Application Created?" group:"details,short" bson:"application_created"`
	PortalUser              bool              `sql:"portal_user,notnull" pg:",use_zero" json:"portal_user,omitempty" view:"hidden" label:"Portal User?" group:"portal" bson:"portal_user"`
	NotifyOwner             bool              `sql:"notify_owner,notnull" pg:",use_zero" json:"notify_owner,omitempty" view:"hidden" label:"Notify Owner?" group:"portal" bson:"notify_owner"`
	EmailOptOut             bool              `sql:"email_opt_out,notnull" pg:",use_zero" json:"email_opt_out,omitempty" view:"hidden" label:"Email Opt Out?" group:"portal" bson:"email_opt_out"`
	DoNotCall               bool              `sql:"do_not_call,notnull" pg:",use_zero" json:"do_not_call,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Do Not Call" group:"portal"`
	ConvertedFromLead       bool              `sql:"converted_from_lead,notnull" pg:",use_zero" json:"converted_from_lead,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Converted from lead?" group:"portal"`
	Product                 string            `sql:"product" json:"product,omitempty" label:"Product" group:"details,short" view:"watch" type:"autocomplete,string" related:"products"`
	CountryOfBirth          string            `sql:"country_of_birth" json:"country_of_birth,omitempty" label:"Country Of Birth" view:"watch" group:"details" type:"autocomplete,string" related:"country"`
	LeadNumber              string            `sql:"lead_number,notnull" json:"lead_number,omitempty" label:"Lead Number" type:"readonly,string" view:"watch" group:"details,short" bson:"lead_number"`
	CountryOfResidence      string            `sql:"country_of_residence" json:"country_of_residence,omitempty" label:"Country Of Residence" view:"watch" group:"details,short" type:"autocomplete,string" related:"country" bson:"country_of_residence"`
	Disqualified            string            `sql:"disqualified" json:"disqualified,omitempty" label:"Disqualified" view:"watch" group:"details,short" type:"autocomplete,string" related:"disqualified"`
	VisaType                string            `sql:"visa_type" json:"visa_type,omitempty" label:"Visa Type" view:"watch" group:"details" type:"autocomplete,string" related:"visa_type" bson:"visa_type"`
	ContactStatus           string            `sql:"contact_status" json:"contact_status,omitempty" view:"watch" label:"Contact Status" group:"details" type:"autocomplete,string" related:"contact_me_status" bson:"contact_status"`
	MobilePhone             string            `sql:"mobile_phone" json:"mobile_phone,omitempty" view:"hidden" label:"Mobile Phone" group:"details" bson:"mobile_phone"`
	SecondaryEmail          string            `sql:"secondary_email" json:"secondary_email,omitempty" view:"watch" type:"string" label:"Secondary Email" group:"details" bson:"secondary_email"`
	Email                   string            `sql:"email" json:"email,omitempty" label:"Email" view:"watch" type:"string" group:"details,short"`
	CampaignSource          string            `sql:"campaign_source" json:"campaign_source,omitempty" view:"watch" label:"Campaign Source" type:"string" group:"details" bson:"campaign_source"`
	Password                string            `sql:"password" json:"password,omitempty" label:"Password" view:"hidden" type:"readonly,string" group:"details,short"`
	FormStatus              string            `sql:"form_status" json:"form_status,omitempty" label:"Form Status" view:"watch" group:"details,short" type:"autocomplete,string" related:"form_status" bson:"form_status"`
	StatusApplication       string            `sql:"status_application" json:"status_application,omitempty" view:"watch" label:"Status Application" group:"details,short" type:"autocomplete,string" related:"status_application" bson:"status_application"`
	OrganizationName        string            `sql:"organization_name" json:"organization_name,omitempty" view:"hidden" label:"Organization Name" group:"portal" bson:"organization_name"`
	Assistant               string            `sql:"assistant" json:"assistant,omitempty" label:"Assistant" view:"hidden" group:"portal"`
	Reference               bool              `sql:"reference,notnull" pg:",use_zero" json:"reference,omitempty" view:"watch" type:"bool,autocomplete" related:"yes_no" label:"Reference" group:"portal"`
	Deleted                 bool              `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" view:"watch" label:"Deleted" type:"bool,autocomplete" related:"yes_no" group:"-"`
	Title                   string            `sql:"title" json:"title,omitempty" label:"Title" view:"watch" group:"portal,short"`
	LeadSource              string            `sql:"lead_source" json:"lead_source,omitempty" view:"watch" label:"LeadSource" group:"portal" bson:"lead_source"`
	Fax                     string            `sql:"fax" json:"fax,omitempty" label:"Fax" view:"hidden" group:"portal"`
	AssistantPhone          string            `sql:"assistant_phone" json:"assistant_phone,omitempty" view:"watch" label:"Assistant Phone" group:"portal" bson:"assistant_phone"`
	HomePhone               string            `sql:"home_phone" json:"home_phone,omitempty" label:"Home Phone" view:"hidden" group:"portal" bson:"home_phone"`
	OfficePhone             string            `sql:"office_phone" json:"office_phone,omitempty" view:"hidden" label:"Office Phone" group:"portal,short" bson:"office_phone"`
	SecondaryPhone          string            `sql:"secondary_phone" json:"secondary_phone,omitempty" view:"hidden" label:"Secondary Phone" group:"portal" bson:"secondary_phone"`
	ReportsTo               string            `sql:"reports_to" json:"reports_to,omitempty" view:"hidden" label:"Reports To" bson:"reports_to"`
	ContactNo               string            `sql:"contact_no" json:"contact_no,omitempty" view:"hidden" label:"Contact No" bson:"contact_no"`
	Department              string            `sql:"department" json:"department,omitempty" view:"watch" label:"Waived Debt" group:"portal"`
	MailingStreet           string            `sql:"mailing_street" json:"mailing_street,omitempty" view:"watch" label:"Receipt Signature" type:"string,readonly" group:"receipt" bson:"mailing_street"`
	MailingPoBox            string            `sql:"mailing_po_box" json:"mailing_po_box,omitempty" view:"watch" label:"Receipt Status" type:"autocomplete,string" related:"receipt_status" group:"address" bson:"mailing_po_box"`
	MailingCity             string            `sql:"mailing_city" json:"mailing_city,omitempty"  view:"hidden" label:"Mailing City" group:"address" bson:"mailing_city"`
	MailingState            string            `sql:"mailing_state" json:"mailing_state,omitempty" view:"hidden" label:"Mailing State" group:"address" bson:"mailing_state"`
	MailingZip              string            `sql:"mailing_zip" json:"mailing_zip,omitempty" view:"hidden" label:"Mailing ZIP" group:"address" bson:"mailing_zip"`
	MailingCountry          string            `sql:"mailing_country" json:"mailing_country,omitempty" view:"hidden" label:"Mailing Country" group:"address" type:"autocomplete,string" related:"country" bson:"mailing_country"`
	MailingStreetOne        string            `sql:"mailing_street_1" json:"mailing_street_1,omitempty" view:"watch" label:"Receipt Signed(First)" type:"string,readonly" group:"receipt" bson:"mailing_street_1"`
	MailingPoBoxOne         string            `sql:"mailing_po_box_1" json:"mailing_po_box_1,omitempty" view:"watch" label:"Receipt Status(First)" type:"autocomplete,readonly" related:"receipt_status"  group:"address" bson:"mailing_po_box_1"`
	MailingCityOne          string            `sql:"mailing_city_1" json:"mailing_city_1,omitempty" view:"hidden" label:"Other City" group:"address" bson:"mailing_city_1"`
	MailingStateOne         string            `sql:"mailing_state_1" json:"mailing_state_1,omitempty" view:"hidden" label:"Other State" group:"address" bson:"mailing_state_1"`
	MailingZipOne           string            `sql:"mailing_zip_1" json:"mailing_zip_1,omitempty" view:"hidden" label:"Other ZIP" group:"address" bson:"mailing_zip_1"`
	MailingCountryOne       string            `sql:"mailing_country_1" json:"mailing_country_1,omitempty" view:"hidden" label:"Other Country" group:"address" bson:"mailing_country_1"`
	Program                 string            `sql:"program" json:"program,omitempty" label:"Program" view:"watch" type:"autocomplete,string" related:"program" group:"details,short"`
	RangeOfSalary           string            `sql:"range_of_salary" json:"range_of_salary,omitempty" view:"hidden" label:"Range Of Salary" type:"autocomplete,string" related:"salary_range" group:"details,short" bson:"range_of_salary"`
	Attempts                int               `sql:"attempts,notnull" pg:",use_zero" json:"attempts,omitempty" view:"watch" label:"Attempts" type:"int" group:"details,short"`
	CreatedTime             int               `sql:"created_time" json:"created_time,omitempty" view:"watch" label:"Created Time" type:"datetime" bson:"created_time"`
	ModifiedTime            int               `sql:"modified_time" json:"modified_time,omitempty" view:"watch" label:"Modified Time" type:"datetime,readonly" group:"portal"`
	LeadCreated             int               `sql:"lead_created" json:"lead_created,omitempty" view:"watch" label:"Lead Created" type:"datetime,readonly" group:"details,short" bson:"lead_created"`
	SupportEndDate          int               `sql:"support_end_date" json:"support_end_date,omitempty" view:"hidden" label:"Support End Date" type:"datetime" bson:"support_end_date"`
	DateOfBirth             int               `sql:"date_of_birth" json:"date_of_birth,omitempty" view:"watch" label:"Date Of Birth" type:"datetime" group:"portal" bson:"date_of_birth"`
	SupportStartDate        int               `sql:"support_start_date" json:"support_start_date,omitempty" view:"hidden" label:"Support Start Date" type:"datetime" bson:"support_start_date"`
	LastEventTime           int               `sql:"last_event_time" json:"last_event_time,omitempty" view:"watch" label:"Last Event Time" group:"details,short" type:"datetime,readonly" bson:"last_event_time"`
	AttemptTime             int               `sql:"attempt_time" json:"attempt_time,omitempty" view:"watch" label:"Attempt Time" group:"details,short" type:"datetime" bson:"attempt_time"`
	CreatedAt               int               `sql:"created_at" json:"created_at,omitempty" view:"watch" label:"Created At" type:"datetime"`
	UpdatedAt               int               `sql:"updated_at" json:"updated_at,omitempty" view:"watch" label:"Updated At" type:"datetime"`
	Products                string            `sql:"-" json:"products,omitempty" type:"autocomplete,string" view:"watch" related:"products" label:"Products"`
	Last4                   string            `sql:"last_4" json:"last_4,omitempty" type:"string,readonly" view:"watch" label:"Last 4" group:"details" bson:"last_4"`
	PaymentBin              string            `sql:"payment_bin" json:"payment_bin,omitempty" view:"watch" type:"string,readonly" label:"Payment Bin" group:"details" bson:"payment_bin"`
	PaidAmount              float64           `sql:"-" json:"paid_amount,omitempty" view:"watch" type:"float" label:"Paid amount" bson:"paid_amount"`
	QuoteAmount             float64           `sql:"-" json:"quote_amount,omitempty" view:"watch" type:"float" label:"Quote amount" bson:"quote_amount"`
	OppQuoteAmount          float64           `sql:"-" pg:"-" json:"opp_quote_amount,omitempty" view:"watch" type:"float" label:"Opportunity Quote amount" bson:"opp_quote_amount"`
	LeadId                  int               `sql:"-" json:"lead_id,omitempty" type:"int,disabled" view:"hidden" label:"Lead ID" bson:"lead_id"`
	LeadQuoteAmount         string            `sql:"-" json:"lead_quote_amount,omitempty" view:"hidden" type:"float,disabled" label:"Lead Quote amount" bson:"lead_quote_amount"`
	LeadPassengerOne        string            `sql:"-" json:"lead_passenger_one,omitempty" view:"hidden" type:"string,readonly" label:"Lead Passenger One" bson:"lead_passenger_one"`
	LeadNumberOfPassengers  int               `sql:"-" json:"lead_number_of_passengers,omitempty" view:"hidden" type:"int,disabled" label:"Lead Number Of Passengers" bson:"lead_number_of_passengers"`
	LeadProduct             string            `sql:"-" json:"lead_product,omitempty" view:"hidden" type:"string,readonly" label:"Lead Product"bson:"lead_product"`
	ReAssessmentDocs        string            `sql:"re_assessment_docs" json:"re_assessment_docs,omitempty" view:"watch" label:"Re Assessment Docs" type:"autocomplete,string" related:"re_assessment_docs" bson:"re_assessment_docs"`
	NoAnswerCounts          int               `sql:"no_answer_counts" json:"no_answer_counts,omitempty" view:"watch" label:"No Answer Counts" type:"int,readonly" bson:"no_answer_counts"`
	Gender                  string            `sql:"gender" json:"gender,omitempty" view:"hidden" label:"Gender"`
	NoSentEmails            int               `sql:"no_sent_emails" json:"no_sent_emails,omitempty" view:"watch" label:"No Sent Email" bson:"no_sent_emails"`
	ScheduleTime            int               `sql:"schedule_time" json:"schedule_time,omitempty" label:"Schedule Time" type:"datetime,readonly" view:"watch" group:"details" bson:"schedule_time"`
	ContactMeTime           int               `sql:"contact_me_time" json:"contact_me_time,omitempty" label:"Contact Me Time" type:"datetime,readonly" view:"watch" group:"details" bson:"contact_me_time"`
	DisqualifiedTime        int               `sql:"disqualified_time" json:"disqualified_time,omitempty" label:"Disqualified Time" type:"datetime,readonly" view:"watch" group:"details" bson:"disqualified_time"`
	AssignTime              int               `sql:"assign_time" json:"assign_time,omitempty" label:"Assign Time" type:"datetime,readonly" view:"watch" group:"details" bson:"assign_time"`
	EffectiveTime           int               `sql:"effective_time" json:"effective_time,omitempty" label:"Effective Time" type:"datetime,readonly" view:"watch" group:"details" bson:"effective_time"`
	AnswerTime              int               `sql:"answer_time" json:"answer_time,omitempty" label:"Answer Time" type:"datetime,readonly" view:"watch" group:"details" bson:"answer_time"`
	RiskStatus              string            `sql:"risk_status" json:"risk_status,omitempty" label:"Risk Status"  type:"autocomplete,string" related:"risk_statuses" bson:"risk_status"`
	RiskSource              string            `sql:"risk_source" json:"risk_source,omitempty" label:"Risk Source"  type:"autocomplete,string" view:"watch" related:"risk_sources" bson:"risk_source"`
	DateRefund              int               `sql:"date_refund" json:"date_refund,omitempty" label:"Date Refund" type:"datetime" view:"watch" bson:"date_refund"`
	IsCardHolder            *bool             `sql:"is_card_holder" json:"is_card_holder,omitempty" label:"Is Card Holder" view:"watch" type:"bool,autocomplete" related:"yes_no" bson:"is_card_holder"`
	CardHolderRelation      string            `sql:"card_holder_relation" json:"card_holder_relation,omitempty" label:"Card Holder Relation"  type:"autocomplete,string" view:"watch" related:"card_holder_relations" bson:"card_holder_relation"`
	StudentProcess          string            `sql:"student_process" json:"student_process,omitempty" label:"Student Process"  type:"autocomplete,string" view:"watch" related:"student_processes" bson:"student_process"`
	FileOpening             FileOpeningStatus `sql:"file_opening" json:"file_opening,omitempty" label:"File Opening"  type:"autocomplete,string" view:"watch" related:"file_openings" bson:"file_opening"`
	Active                  bool              `sql:"active,notnull" pg:",use_zero" json:"active,omitempty" view:"watch" label:"Active" type:"bool,autocomplete" related:"yes_no" group:"-"`
	ResultUrl               string            `sql:"result_url" json:"result_url,omitempty" label:"Result URL"  type:"string" view:"watch"`
	ResultTime              int               `sql:"result_time" json:"result_time,omitempty" label:"Result Time" type:"datetime" view:"watch"`
	RcicTime                int               `sql:"rcic_time" json:"rcic_time,omitempty" label:"RCIC Time" type:"datetime" view:"watch"`
	CSRAssignTo             int               `sql:"csr_assign_to" json:"csr_assign_to,omitempty" label:"CSR Assign To" type:"autocomplete,int" related:"users"`
	CSRPreviousAssignTo     int               `sql:"csr_previous_assign_to" json:"csr_previous_assign_to,omitempty" label:"CSR Previous Assign To" type:"readonly,autocomplete,int" related:"users"`
	CSRAssignTime           int               `sql:"csr_assign_time" json:"csr_assign_time,omitempty" label:"CSR Assign Time" type:"readonly,datetime" view:"watch"`
	DocumentStatus          string            `sql:"document_status" bson:"document_status,omitempty" json:"document_status,omitempty" label:"Document Status" type:"autocomplete,string" view:"watch" related:"document_status_app"`
	DocumentReviewTime      int               `sql:"document_review_time" json:"document_review_time" label:"Document Review Time" type:"datetime" view:"watch"`
	MDCAssignTo             int               `sql:"mdc_assign_to" json:"mdc_assign_to,omitempty" label:"MDC Assign To" type:"autocomplete,int" related:"users"`
	MDCAssignTime           int               `sql:"mdc_assign_time" json:"mdc_assign_time,omitempty" label:"MDC Assign Time" type:"readonly,datetime" view:"watch"`
	FormStatusChangeTime    int               `sql:"form_status_change_time" json:"form_status_change_time,omitempty" label:"Form Status Change Time" type:"readonly,datetime" view:"watch" bson:"form_status_change_time"`
	FormStatusCompleteTime  int               `sql:"form_status_complete_time" json:"form_status_complete_time,omitempty" label:"Form Status Complete Time" type:"readonly,datetime" view:"watch" bson:"form_status_complete_time"`
	RiskAssignTo            int               `sql:"risk_assign_to" json:"risk_assign_to,omitempty" label:"Risk Assign To" type:"autocomplete,int" related:"users" view:"watch" bson:"risk_assign_to"`
	RiskCreatedTime         int               `sql:"risk_created_time" json:"risk_created_time,omitempty" label:"Risk Created Time" type:"datetime" view:"watch" bson:"risk_created_time"`
	RiskResolvedTime        int               `sql:"risk_resolved_time" json:"risk_resolved_time,omitempty" label:"Risk Resolved Time" type:"datetime" view:"watch" bson:"risk_resolved_time"`
	RiskAssignTime          int               `sql:"risk_assign_time" json:"risk_assign_time,omitempty" label:"Risk Assign Time" type:"datetime" view:"watch" bson:"risk_assign_time"`
	MDCBackofficeAssignTo   int               `sql:"mdc_backoffice_assign_to" json:"mdc_backoffice_assign_to,omitempty" label:"MDC Backoffice Assign To" type:"autocomplete,int" related:"users" bson:"mdc_backoffice_assign_to"`
	MDCBackofficeAssignTime int               `sql:"mdc_backoffice_assign_time" json:"mdc_backoffice_assign_time,omitempty" label:"MDC Backoffice Assign Time" type:"datetime" view:"watch" bson:"mdc_backoffice_assign_time"`
	CollectedDocumentsTime  int               `sql:"collected_documents_time" json:"collected_documents_time" label:"Collected Documents Time" type:"datetime" view:"watch" bson:"collected_documents_time"`
	MDCEvaluationAssignTo   int               `sql:"mdc_evaluation_assign_to" json:"mdc_evaluation_assign_to,omitempty" label:"MDC Evaluation Assign To" type:"autocomplete,int" related:"users" bson:"mdc_evaluation_assign_to"`
	MDCEvaluationAssignTime int               `sql:"mdc_evaluation_assign_time" json:"mdc_evaluation_assign_time,omitempty" label:"MDC Evaluation Assign Time" type:"datetime" view:"watch" bson:"mdc_evaluation_assign_time"`
	Diploma                 string            `sql:"diploma" json:"diploma,omitempty" label:"Diploma" type:"readonly,string"`

	EjobboStatus EjobboStatus `sql:"ejobbo_status" json:"ejobbo_status,omitempty" label:"Ejobbo Status" type:"readonly,string,autocomplete" related:"ejobbo_statuses" view:"watch"`

	Online        bool   `sql:"-" json:"online" view:"watch" pg:",use_zero" label:"Online" view:"watch" type:"bool,autocomplete" related:"yes_no"`
	PaymentStatus string `sql:"-" json:"payment_status,omitempty" view:"watch" label:"Payment Status" type:"string,autocomplete" related:"status_reason_opportunity"`
	CardType      string `sql:"-" json:"card_type,omitempty" view:"watch" label:"Card Type" type:"string,autocomplete" related:"card_type"`
	Card          string `sql:"-" json:"card,omitempty" view:"watch" label:"Card" type:"string,autocomplete" related:"card"`
	CardClass     string `sql:"-" json:"card_class,omitempty" view:"watch" label:"Card Class" type:"string,autocomplete" related:"card_class"`
	IsRisk        bool   `sql:"-" json:"is_risk" view:"watch" pg:",use_zero" label:"Is Risk" type:"bool,autocomplete" related:"yes_no"`

	// Related Fields to Documents
	DocumentsType                 string       `sql:"-" json:"documents_type,omitempty" view:"watch" label:"Documents Type" type:"readonly,string,autocomplete" related:"document_types"`
	DocumentsStatus               string       `sql:"-" json:"documents_status,omitempty" view:"watch" label:"Documents Status" type:"readonly,string,autocomplete" related:"document_statuses"`
	DocumentsUrl                  string       `sql:"-" json:"documents_url,omitempty" view:"watch" label:"Documents Url" type:"readonly,string"`
	RetainerSignTime              int          `sql:"-" json:"retainer_sign_time" view:"watch" label:"Retainer Sign Time" type:"datetime,readonly"`
	DocumentsUploadVisaResultTime string       `sql:"-" json:"documents_upload_visa_result_time" view:"watch" label:"Documents Upload Visa Result Time" type:"datetime,readonly"`
	ResultScores                  ResultScores `sql:"result_scores" json:"result_scores" label:"Result Scores" type:"readonly"`
	DocumentRCICResultUploadedBy  string       `sql:"-" json:"document_rcic_result_uploaded_by,omitempty" label:"Document RCIC Result Uploaded By" type:"readonly,string,autocomplete" related:"users"`
	DocumentRCICUploadedBySystem  string       `sql:"-" json:"document_rcic_uploaded_by_system" label:"Document RCIC Uploaded By System" type:"readonly,autocomplete,string" related:"document_systems"`

	// Related Fields to VISA
	HasVisa                 bool   `sql:"-" json:"has_visa,omitempty" label:"Has Visa" type:"readonly,bool,autocomplete" related:"yes_no"`
	VisaTypes               string `sql:"-" json:"visa_types,omitempty" label:"Visa Types" type:"readonly,string,autocomplete" related:"visa_types" view:"watch"`
	VisaStatuses            string `sql:"-" json:"visa_statuses,omitempty" label:"Visa Statuses" type:"readonly,string,autocomplete" related:"visa_statuses" view:"watch"`
	VisaProgresses          int    `sql:"-" json:"visa_progresses,omitempty" label:"Visa Progresses" type:"readonly,int" view:"watch"`
	VisaOpenTimes           string `sql:"-" json:"visa_open_times,omitempty" label:"Visa Open Times" type:"datetime,readonly" view:"watch"`
	VisaResultStatuses      string `sql:"-" json:"visa_result_statuses,omitempty" label:"Visa Result Statuses" type:"string,autocomplete,readonly" view:"watch" related:"visa_result_statuses"`
	VisaSubmittedToGovTimes string `sql:"-" json:"visa_submitted_to_gov_times,omitempty" label:"Visa Submitted To Gov Times" type:"datetime,readonly" view:"watch"`
	RequirementTypes        string `sql:"-" json:"requirement_types,omitempty" label:"Requirement Types" type:"readonly,string,autocomplete" view:"watch" related:"requirement_types"`
	RequirementStatuses     string `sql:"-" json:"requirement_statuses,omitempty" label:"Requirement Statuses" type:"readonly,string,autocomplete" view:"watch" related:"requirement_statuses"`

	// Related Fields to Tickets
	TicketTypes    string `sql:"-" json:"ticket_types,omitempty" label:"Ticket Types" type:"readonly,string,autocomplete" view:"watch" related:"ticket_types"`
	TicketStatuses string `sql:"-" json:"ticket_statuses,omitempty" label:"Ticket Statuses" type:"readonly,string,autocomplete" view:"watch" related:"ticket_statuses"`
	HasTicket      bool   `sql:"-" json:"has_ticket,omitempty" label:"Has Ticket" type:"readonly,bool,autocomplete" related:"yes_no"`

	// Related Fields to Forms
	HasForm              bool   `sql:"-" json:"has_form,omitempty" label:"Has Feedback Form" type:"readonly,bool,autocomplete" related:"yes_no"`
	FeedbackFormTypes    string `sql:"-" json:"feedback_form_types,omitempty" label:"Form Feedback Types" type:"readonly,string,autocomplete" view:"watch" related:"feedback_form_types"`
	FeedbackFormStatuses string `sql:"-" json:"feedback_form_statuses,omitempty" label:"Form Feedback Statuses" type:"readonly,string,autocomplete" view:"watch" related:"feedback_form_statuses"`

	ChangedAttributes []string `sql:"-" json:"changed_attributes,omitempty"`

	NeedCreateNewImmi bool `sql:"-" json:"need_create_new_immi"`
}

type EjobboStatus int

var (
	ActiveEjobboStatus   = 1
	InActiveEjobboStatus = 2
)

type FileOpeningStatus string

var (
	ClosedFOStatus     FileOpeningStatus = "Closed"
	ProcessingFOStatus FileOpeningStatus = "Processing"
	PendingFOStatus    FileOpeningStatus = "Pending"
)

const (
	HighAssessmentPriority   = "High"
	MediumAssessmentPriority = "Medium"
	LowAssessmentPriority    = "Low"
)

const (
	SentToRcicFormStatus    = "Sent to RCIC"
	FormCompletedFormStatus = "Form Completed"
)

type ResultScores struct {
	Work         int `json:"work"`
	English      int `json:"english"`
	Education    int `json:"education"`
	Age          int `json:"age"`
	Offer        int `json:"offer"`
	Adaptability int `json:"adaptability"`
	Total        int `json:"total"`
}

var PositiveRiskStatusesArray = []string{"Saved & Continuing", "Partial Refund & Continuing", "Clear To Go", "Online Risk Good To Go"}
