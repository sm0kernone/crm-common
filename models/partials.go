package models

type Partials struct {
	TableName          struct{}       `sql:"partials" json:"-"`
	Id                 int            `sql:"id" json:"id,omitempty" view:"hidden" type:"int,disabled"`
	CrmId              int            `sql:"crm_id" json:"crm_id,omitempty" view:"hidden" type:"int,disabled" label:"CRM ID"`
	Label              string         `sql:"label" json:"label,omitempty" view:"hidden" label:"GetLabel"`
	Subject            string         `sql:"subject" json:"subject,omitempty" view:"hidden" label:"Subject" type:"readonly" group:"details"`
	OpportunityCrmId   int            `sql:"opportunity_crm_id" json:"opportunity_crm_id,omitempty" view:"hidden" type:"int,disabled" label:"Opportunity CRM ID"`
	ApplicationCrmId   int            `sql:"application_crm_id" json:"application_crm_id,omitempty" view:"hidden" type:"int,disabled" label:"Application CRM ID"`
	PartialNumber      string         `sql:"partial_number" json:"partial_number,omitempty" view:"hidden" label:"Partial Number"`
	Subtotal           float64        `sql:"subtotal" json:"subtotal,omitempty" view:"hidden" label:"Subtotal"`
	Adjustment         float64        `sql:"adjustment" json:"adjustment,omitempty" view:"hidden" label:"Adjustment"`
	Total              float64        `sql:"total" json:"total,omitempty" label:"Total" view:"watch" group:"item"`
	TaxType            string         `sql:"tax_type" json:"tax_type,omitempty" view:"hidden" label:"Tax Type"`
	PreTaxTotal        float64        `sql:"pre_tax_total" json:"pre_tax_total,omitempty" view:"hidden" label:"Pre Tax Total" group:"item"`
	CreatedUserCrmId   int            `sql:"created_user_crm_id" json:"created_user_crm_id,omitempty" view:"watch" label:"Created By" type:"autocomplete,readonly,int" group:"details" related:"users"`
	AssignedUserCrmId  int            `sql:"assigned_user_crm_id" json:"assigned_user_crm_id,omitempty" view:"watch" type:"autocomplete,readonly,int" label:"Assigned To" related:"users"`
	ModifiedUserCrmId  int            `sql:"modified_user_crm_id" json:"modified_user_crm_id,omitempty" view:"watch" label:"Modified CRM ID" type:"autocomplete,int" related:"users"`
	CreatedTime        int            `sql:"created_time" json:"created_time,omitempty" label:"Created Time" view:"watch" type:"datetime,readonly" group:"details"`
	ModifiedTime       int            `sql:"modified_time" json:"modified_time,omitempty" label:"Modified Time" view:"watch" type:"datetime,readonly" group:"details"`
	Deleted            bool           `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" view:"watch" label:"Deleted" type:"bool,autocomplete" related:"yes_no"`
	AccountNumber      string         `sql:"-" json:"account_number,omitempty" view:"hidden" label:"Account number"`
	PotentialName      string         `sql:"-" json:"potential_name,omitempty" view:"watch" label:"Opportunity name"`
	Product            string         `sql:"-" json:"product,omitempty" label:"Product" view:"watch" type:"autocomplete,string" related:"products"`
	CountryOfResidence string         `sql:"-" json:"country_of_residence,omitempty" view:"watch" label:"Country Of Residence" type:"autocomplete,string" related:"country"`
	CountryOfBirth     string         `sql:"-" json:"country_of_birth,omitempty" view:"watch" label:"Country Of Birth" type:"autocomplete,string" related:"country"`
	NeedToPay          float64        `sql:"-" json:"need_to_pay,omitempty" type:"float" view:"watch" label:"Need To Pay Amount"`
	QuoteAmount        float64        `sql:"-" json:"quote_amount,omitempty" view:"watch" type:"readonly,float" label:"Quote"`
	LastEventTime      int            `sql:"-" json:"last_event_time,omitempty" view:"watch" label:"Last Event Time" type:"datetime"`
	AttemptTime        int            `sql:"-" json:"attempt_time,omitempty" view:"watch" label:"Last Attempt Time" type:"datetime"`
	CreatedAt          int            `sql:"created_at" json:"created_at,omitempty" view:"watch" label:"Created At" type:"datetime"`
	UpdatedAt          int            `sql:"updated_at" json:"updated_at,omitempty" view:"watch" label:"Updated At" type:"datetime"`
	Opportunities      *Opportunities `sql:"-" json:"opportunity" view:"hidden"`

	//Fields Related To Applications
	IsRisk bool `sql:"-" json:"is_risk" view:"watch" pg:",use_zero" label:"Is Risk" type:"bool,autocomplete" related:"yes_no"`
}
