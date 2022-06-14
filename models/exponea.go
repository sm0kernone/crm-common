package models

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/utils"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type ExponiaReq struct {
	CustomerIDs map[string]interface{} `json:"customer_ids"`
	EventType   string                 `json:"event_type,omitempty"`
	Timestamp   float64                `json:"timestamp,omitempty"`
	Properties  interface{}            `json:"properties"`
}

type CustomerExponiaReq struct {
	FirstName                           string            `json:"first_name" trigger:"first_name" sql:"leads.first_name"`
	LastName                            string            `json:"last_name,omitempty" trigger:"last_name" sql:"leads.last_name"`
	Company                             string            `json:"company,omitempty" trigger:"company" sql:"leads.company"`
	Gender                              string            `json:"gender,omitempty" trigger:"gender" sql:"leads.gender"`
	Email                               string            `json:"email" trigger:"email" sql:"leads.email"`
	Phone                               string            `json:"phone" trigger:"phone" sql:"leads.phone"`
	BirthDate                           string            `json:"birth_date,omitempty" trigger:"exact_birthday" sql:"leads.exact_birthday"`
	LeadNumber                          string            `json:"lead_number" trigger:"lead_number" sql:"leads.lead_number"`
	CountryOfResidence                  string            `json:"country_of_residence" trigger:"country_of_residence" sql:"leads.country_of_residence"`
	CountryOfBirth                      string            `json:"country_of_birth" trigger:"country_of_birth" sql:"leads.country_of_birth"`
	StatusReason                        string            `json:"status_reason" trigger:"status_reason" sql:"leads.status_reason"`
	CampaignSource                      string            `json:"campaign_source,omitempty" trigger:"campaign_source" sql:"leads.campaign_source"`
	PreviousCampaignSource              string            `json:"previous_campaign_source,omitempty" trigger:"previous_campaign_source" sql:"leads.previous_campaign_source"`
	SpeakFrench                         string            `json:"speak_french,omitempty" trigger:"speak_french" sql:"leads.speak_french"`
	SpeakEnglish                        string            `json:"speak_english,omitempty" trigger:"speak_english" sql:"leads.speak_english"`
	MaritalStatus                       string            `json:"marital_status" trigger:"marital_status" sql:"leads.marital_status"`
	Disqualified                        string            `json:"disqualified" trigger:"disqualified" sql:"leads.disqualified"`
	DisqualifiedTime                    int               `json:"disqualified_time" trigger:"disqualified_time" sql:"leads.disqualified_time"`
	VisaType                            string            `json:"visa_type,omitempty" trigger:"visa_type" sql:"leads.visa_type"`
	Nationality                         string            `json:"nationality,omitempty" trigger:"nationality" sql:"leads.nationality"`
	FunnelStatusReason                  string            `json:"funnel_status_reason,omitempty" trigger:"funnel_status_reason" sql:"leads.funnel_status_reason"`
	ApplicationType                     string            `json:"application_type" trigger:"application_type" sql:"leads.application_type"`
	Employed                            string            `json:"employed,omitempty" trigger:"employed" sql:"leads.employed"`
	Assignable                          string            `json:"assignable" trigger:"assignable" sql:"leads.assignable"`
	ScheduleTime                        int               `json:"schedule_time" trigger:"schedule_time" sql:"leads.schedule_time"`
	ContactMeTime                       int               `json:"contact_me_time" trigger:"contact_me_time" sql:"leads.contact_me_time"`
	ContactStatus                       string            `json:"contact_status" trigger:"contact_status" sql:"leads.contact_status"`
	AssignedTo                          int               `json:"assigned_to" trigger:"assigned_to" sql:"leads.assigned_to"`
	AssignTime                          int               `json:"assign_time" trigger:"assign_time" sql:"leads.assign_time"`
	Age                                 int               `json:"age" trigger:"age" sql:"leads.age"`
	AdGroup                             string            `json:"adgroup" trigger:"adgroup" sql:"leads.adgroup"`
	AdId                                string            `json:"adid" trigger:"adid" sql:"leads.adid"`
	AdPosition                          string            `json:"adposition" trigger:"adposition" sql:"leads.adposition"`
	Campaign                            string            `json:"campaign" trigger:"campaign" sql:"leads.campaign"`
	Device                              string            `json:"device" trigger:"device" sql:"leads.device"`
	GclId                               string            `json:"gclid" trigger:"gclid" sql:"leads.gclid"`
	GeoId                               string            `json:"geoid" trigger:"geoid" sql:"leads.geoid"`
	Keyword                             string            `json:"keyword" trigger:"keyword" sql:"leads.keyword"`
	MatchType                           string            `json:"matchtype" trigger:"matchtype" sql:"leads.matchtype"`
	Placement                           string            `json:"placement" trigger:"placement" sql:"leads.placement"`
	Source                              string            `json:"source" trigger:"source" sql:"leads.source"`
	Medium                              string            `json:"medium" trigger:"medium" sql:"leads.medium"`
	Occupation                          string            `json:"occupation" trigger:"occupation" sql:"leads.occupation"`
	ApplicationEmail                    string            `json:"application_email,omitempty" trigger:"email" sql:"applications.email"`
	ApplicationPhone                    string            `json:"application_phone,omitempty" trigger:"office_phone" sql:"applications.office_phone"`
	ApplicationDisqualified             string            `json:"application_disqualified" trigger:"disqualified" sql:"applications.disqualified"`
	ApplicationDisqualifiedTime         int               `json:"application_disqualified_time" trigger:"disqualified_time" sql:"applications.disqualified_time"`
	ApplicationVisaType                 string            `json:"application_visa_type,omitempty" trigger:"visa_type" sql:"applications.visa_type"`
	ApplicationContactStatus            string            `json:"application_contact_status" trigger:"contact_status" sql:"applications.contact_status"`
	ApplicationContactMeTime            int               `json:"application_contact_me_time" trigger:"contact_me_time" sql:"applications.contact_me_time"`
	ApplicationCampaignSource           string            `json:"application_campaign_source,omitempty" trigger:"campaign_source" sql:"applications.campaign_source"`
	ApplicationPassword                 string            `json:"application_password" trigger:"password" sql:"applications.password"`
	ApplicationFormStatus               string            `json:"application_form_status" trigger:"form_status" sql:"applications.form_status"`
	ApplicationStatusApplication        string            `json:"application_status_application,omitempty" trigger:"status_application" sql:"applications.status_application"`
	ApplicationRetainerEvaluation       string            `json:"application_retainer_evaluation,omitempty" trigger:"mailing_street" sql:"applications.mailing_street"`
	ApplicationRetainerEvaluationStatus string            `json:"application_retainer_evaluation_status,omitempty" trigger:"mailing_po_box" sql:"applications.mailing_po_box"`
	ApplicationProgram                  string            `json:"application_program,omitempty" trigger:"program" sql:"applications.program"`
	ApplicationAssignTime               int               `json:"application_assign_time,omitempty" trigger:"assign_time" sql:"applications.assign_time"`
	ApplicationAssignedUser             int               `json:"application_assigned_user,omitempty" trigger:"assigned_user_crm_id" sql:"applications.assigned_user_crm_id"`
	ApplicationRiskStatus               string            `json:"application_risk_status" trigger:"risk_status" sql:"applications.risk_status"`
	ApplicationRiskSource               string            `json:"application_risk_source" trigger:"risk_source" sql:"applications.risk_source"`
	ApplicationFileOpening              string            `json:"application_file_opening" trigger:"file_opening" sql:"applications.file_opening"`
	ApplicationResultUrl                string            `json:"application_result_url" trigger:"result_url" sql:"applications.result_url"`
	ApplicationResultTime               int               `json:"application_result_time" trigger:"result_time" sql:"applications.result_time"`
	ApplicationRcicTime                 int               `json:"application_rcic_time" trigger:"rcic_time" sql:"applications.rcic_time"`
	ApplicationHoldStatus               string            `json:"application_hold_status" trigger:"hold_status" sql:"applications.hold_status"`
	ApplicationRetainerUpgrade          string            `json:"application_retainer_upgrade,omitempty" sql:"string_agg(documents.url::text, ', ')"`
	ApplicationRetainerUpgradeStatus    string            `json:"application_retainer_upgrade_status,omitempty" sql:"string_agg(documents.status::text, ', ')"`
	ApplicationVisaStatus               string            `json:"application_visa_status,omitempty" sql:"string_agg(coalesce(visas.visa_status::text, '0'), ', ')"`
	ApplicationVisaProgress             string            `json:"application_visa_progress,omitempty" sql:"string_agg(coalesce(visas.progress::text, '0'), ', ')"`
	ApplicationVisaResultStatus         string            `json:"application_visa_result_status,omitempty" sql:"string_agg(visas.additional_info ->>'result_status'::text, ', ')"`
	ApplicationVisaSubmittedToGov       string            `json:"application_visa_submitted_to_gov,omitempty" sql:"string_agg(visas.submitted_to_go_time::text, ', ')"`
	ApplicationVisaTypeHash             string            `json:"application_visa_type_hash,omitempty" sql:"string_agg(visas.visa_type::text, ', ')"`
	ApplicationCreatedTime              int               `json:"application_created_time,omitempty" trigger:"created_at" sql:"applications.created_time"`
	LeadsCreatedTime                    int               `json:"leads_created_time" trigger:"created_at" sql:"leads.created_time"`
	ApplicationUpgradeStatus            string            `json:"application_upgrade_status"`
	ApplicationIeltsStatus              string            `json:"application_ielts_status"`
	ApplicationIsOnline                 bool              `json:"application_is_online" sql:"o.online as application_is_online"`
	ApplicationEvaluationCV             string            `json:"application_evaluation_cv,omitempty"`
	ApplicationEvaluationPP             string            `json:"application_evaluation_pp,omitempty"`
	ApplicationEvaluationRP             string            `json:"application_evaluation_rp,omitempty"`
	ApplicationEvaluationPH             string            `json:"application_evaluation_ph,omitempty"`
	ApplicationEvaluationED             string            `json:"application_evaluation_ed,omitempty"`
	ApplicationScheduleTime             int               `json:"application_schedule_time" trigger:"schedule_time" sql:"applications.schedule_time"`
	ApplicationMDCAssignTo              int               `json:"application_mdc_assign_to" trigger:"mdc_assign_to" sql:"applications.mdc_assign_to"`
	VisaMissingFields                   map[string]string `json:"visa_missing_fields,omitempty"`
	EmailHash                           string            `json:"email_hash"`
	AutoLoginHash                       string            `json:"auto_login_hash"`
	GMT                                 string            `json:"gmt"`
	ClientUUID                          string            `json:"client_uuid"`
	LeadID                              int               `json:"lead_id" sql:"leads.crm_id"`
}

type VisasRequirementsExponea []map[string]interface{}

// Make the VisasRequirementsExponea struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (v VisasRequirementsExponea) Value() (driver.Value, error) {
	return json.Marshal(v)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (v *VisasRequirementsExponea) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &v)
}

type reqStatusProgress struct {
	Progress int
	Status   int
	VisaType string
}

type visaRespStruct struct {
	VisaType string `json:"visa_type"`
	Type     string `json:"type"`
	Status   int    `json:"status"`
	Progress int    `json:"progress"`
}

func (v VisasRequirementsExponea) GetAllRequirements() map[string]string {
	visasResp := new([]visaRespStruct)

	respBytes, _ := json.Marshal(v)

	json.Unmarshal(respBytes, visasResp)

	reqNamePattern := "applications_%v_%v"
	topRequirementsMap := make(map[string]reqStatusProgress)
	for _, visaRequirement := range *visasResp {
		reqType := strings.ToLower(visaRequirement.Type)
		progress := visaRequirement.Progress
		status := visaRequirement.Status
		visaType := visaRequirement.VisaType
		// if reqType is already in map - check whether progress in this requirement is bigger than in the one in map
		// if it's bigger - set this requirement
		if val, ok := topRequirementsMap[reqType]; ok {
			if progress > val.Progress {
				topRequirementsMap[reqType] = reqStatusProgress{
					Progress: progress,
					Status:   status,
					VisaType: visaType,
				}
			}
		} else {
			topRequirementsMap[reqType] = reqStatusProgress{
				Progress: progress,
				Status:   status,
				VisaType: visaType,
			}
		}
	}

	requirementsStatusesMap := make(map[string]string)
	for requirementType, statusProgress := range topRequirementsMap {
		requirementsStatusesMap[fmt.Sprintf(reqNamePattern, statusProgress.VisaType, requirementType)] = RequirementStatuses[statusProgress.Status]
	}

	return requirementsStatusesMap
}

var RequirementStatuses = map[int]string{
	0:  "Required",
	-1: "Optional",
	1:  "Review",
	2:  "Feedback",
	3:  "Approve",
	4:  "Awaiting ITA",
	5:  "Awaiting LOA",
	6:  "Pre Approved",
}

func (c *CustomerExponiaReq) NormalizeFields() {
	reqValuePtr := reflect.ValueOf(c)
	reqValue := reqValuePtr.Elem()
	for i := 0; i < reqValue.NumField(); i++ {
		field := reqValue.Field(i)

		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		str = strings.TrimSpace(str)
		field.SetString(str)
	}

	c.FirstName = utils.NormalizeField(c.FirstName)
	c.LastName = utils.NormalizeField(c.LastName)
	c.Email = strings.ToLower(c.Email)
	c.ApplicationEmail = strings.ToLower(c.ApplicationEmail)

	if c.ApplicationScheduleTime < int(time.Now().Unix()) {
		c.ApplicationScheduleTime = 0
	}

	if c.ScheduleTime < int(time.Now().Unix()) {
		c.ScheduleTime = 0
	}
}

func GetListOfCustomerExponiaAttributes(tagType string) []string {
	customerStr := new(CustomerExponiaReq)
	strValue := reflect.ValueOf(customerStr)
	var attributes []string

	for i := 0; i < strValue.Elem().NumField(); i++ {
		tagValue := strValue.Elem().Type().Field(i).Tag.Get(tagType)
		if len(tagValue) > 0 {
			attributes = append(attributes, tagValue)
		}
	}

	return attributes
}

func GetListOfCustomerExponiaGroupByFields() []string {
	customerStr := new(CustomerExponiaReq)
	strValue := reflect.ValueOf(customerStr)
	var attributes []string

	for i := 0; i < strValue.Elem().NumField(); i++ {
		tagValue := strValue.Elem().Type().Field(i).Tag.Get("sql")

		tagArray := strings.Split(tagValue, " ")

		tagValue = tagArray[0]

		if !strings.Contains(strings.ToLower(tagValue), "string_agg") && len(tagValue) > 0 {
			attributes = append(attributes, tagValue)
		}
	}
	return attributes
}

type ProductList struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type PurchaseExponiaReq struct {
	PurchaseID         string        `json:"purchase_id"`
	PurchaseStatus     string        `json:"purchase_status"`
	PurchaseSourceType string        `json:"purchase_source_type"`
	PurchaseSourceName string        `json:"purchase_source_name"`
	VoucherCode        string        `json:"voucher_code,omitempty"`
	VoucherPercentage  string        `json:"voucher_percentage,omitempty"`
	VoucherValue       string        `json:"voucher_value,omitempty"`
	PaymentType        string        `json:"payment_type"`
	ProductList        []ProductList `json:"product_list"`
	ProductIds         []string      `json:"product_ids"`
	PaidAmount         float64       `json:"paid_amount"`
	TotalPrice         float64       `json:"total_price"`
	Language           string        `json:"language"`
	FlowType           string        `json:"flow_type"`
}

type PurchaseInstallmentExponiaReq struct {
	PaymentQuoteAmount  float64 `json:"payment_quote_amount"`
	PaymentPaidAmount   float64 `json:"payment_paid_amount"`
	OutstandingAmount   float64 `json:"outstanding_amount"`
	AgentName           string  `json:"agent_name"`
	Installment         string  `json:"installment"`
	NextInstallmentDate int     `json:"next_installment_date"`
	Product             string  `json:"product"`
}

type PurchaseItemExponiaReq struct {
	PurchaseID         string  `json:"purchase_id"`
	PurchaseStatus     string  `json:"purchase_status"`
	PurchaseSourceType string  `json:"purchase_source_type"`
	ProductID          string  `json:"product_id"`
	ProductName        string  `json:"product_name"`
	PurchaseSourceName string  `json:"purchase_source_name,omitempty"`
	DiscountPercentage string  `json:"discount_percentage,omitempty"`
	DiscountValue      string  `json:"discount_value,omitempty"`
	OriginalPrice      float64 `json:"original_price"`
	PaidAmount         float64 `json:"paid_amount"`
	TotalPrice         float64 `json:"total_price"`
	LocalCurrency      string  `json:"local_currency"`
	Language           string  `json:"language"`
	SpouseFirstName    string  `json:"spouse_first_name,omitempty"`
	SpouseLastName     string  `json:"spouse_last_name,omitempty"`
	SpouseEmail        string  `json:"spouse_email,omitempty"`
	SpousePhone        string  `json:"spouse_phone,omitempty"`
	SpouseCountry      string  `json:"spouse_country,omitempty"`
}

type ConsentExponeaReq struct {
	Action     ConsentAction `json:"action"`
	Category   string        `json:"category"`
	ValidUntil string        `json:"valid_until"`
	Source     string        `json:"source"`
}

type ConsentAction string

var (
	AcceptConsentAction ConsentAction = "accept"
	RejectConsentAction ConsentAction = "reject"

	DefaultValidUntil = "unlimited"
)

type CallActivityExponeaReq struct {
	AgentName     string                        `json:"agent_name"`
	AgentRole     string                        `json:"agent_role"`
	AgentLanguage string                        `json:"language"`
	Duration      int                           `json:"duration"`
	StartTime     int                           `json:"start_time"`
	EndTime       int                           `json:"end_time"`
	Wait          int                           `json:"wait"`
	Type          CallActivityExponeaEntityType `json:"type"`
}

type CallActivityExponeaEntityType string

var (
	LeadCallActivityEntityType        CallActivityExponeaEntityType = "Lead"
	ApplicationCallActivityEntityType CallActivityExponeaEntityType = "Application"
)

type SendClientExponeaReq struct {
	AgentName string `json:"agent_name"`
	URL       string `json:"url"`
}

type SendDocumentExponea struct {
	AgentName string `json:"agent_name"`
	URL       string `json:"url"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	AgentType string `json:"agent_type"`
}

type SendVsaRCICAssignDataReq struct {
	AgentName string `json:"agent_name"`
}

type SentToRCICTimeReq struct {
	SentToRCICTime int `json:"sent_to_rcic_time"`
}

type InstallmentPieceExponeaRequest struct {
	AgentName          string `json:"agent_name"`
	InstallmentPrice   Amount `json:"installment_price"`
	ProductName        string `json:"product_name"`
	PlannedPayment     int    `json:"planned_payment"`
	PaidAmount         Amount `json:"paid_amount"`
	QuoteAmount        Amount `json:"quote_amount"`
	InstallmentPieceId int    `json:"installment_piece_id"`
}

type CallActivityLogStruct struct {
	Date            string `json:"date"`
	Requests        int    `json:"requests"`
	SuccessRequests int    `json:"success_requests"`
	Role            string `json:"role"`
}
