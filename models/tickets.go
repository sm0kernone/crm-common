package models

type FWAdditionalInfo struct {
	FWId   int    `json:"fw_id"`
	FWType string `json:"fw_type"`
}

type Tickets struct {
	Id                    int               `sql:"id,pk" json:"id" label:"ID" type:"int" validate:"omitempty"`
	TicketID              string            `sql:"-" json:"ticket_id,omitempty" label:"Ticket ID" validate:"omitempty"`
	Name                  string            `sql:"name" json:"name" label:"Name" type:"string" validate:"required"`
	Description           string            `sql:"description" json:"description" label:"Description" type:"string" validate:"required"`
	Type                  TicketType        `sql:"type" json:"type" label:"Type" validate:"required,min=1" type:"string,autocomplete,readonly" related:"ticket_types"`
	Status                TicketStatus      `sql:"status" json:"status" label:"Status" type:"autocomplete,string,readonly" related:"ticket_statuses" bson:"status" validate:"required,min=1" `
	OpenedBy              int               `sql:"opened_by" json:"opened_by" label:"Opened By" bson:"opened_by" validate:"omitempty" type:"int,autocomplete,readonly" related:"users"`
	AssignedTo            int               `sql:"assigned_to" json:"assigned_to" label:"Assigned To" bson:"assigned_to" validate:"required" type:"int,autocomplete" related:"users"`
	AssignedAt            int               `sql:"assigned_at" json:"assigned_at" label:"Assigned At" bson:"assigned_at" validate:"omitempty" type:"datetime,readonly"`
	ApprovedBy            int               `sql:"approved_by" json:"approved_by, omitempty" label:"Approved By" type:"int,readonly,autocomplete" bson:"approved_by" related:"users" validate:"omitempty"`
	ApprovedAt            int               `sql:"approved_at" json:"approved_at, omitempty" label:"Approved At" type:"datetime,autocomplete" bson:"approved_at" validate:"omitempty"`
	StartedAt             int               `sql:"started_at, notnull" json:"started_at" label:"Started At" bson:"started_at" validate:"omitempty" type:"datetime" related:"users"`
	Priority              TicketPriority    `sql:"priority, notnull" json:"priority" label:"Priority" type:"string,autocomplete" related:"ticket_priorities" validate:"required,min=1"`
	ApproveByOwnerRole    bool              `sql:"approve_by_owner,notnull" pg:",use_zero" json:"approve_by_owner" label:"Approve By Owner" type:"bool,readonly,autocomplete" related:"yes_no" bson:"approve_by_owner" validate:"omitempty"`
	TicketRelations       []TicketRelations `sql:"-" pg:"fk:ticket_id" json:"relations, omitempty" label:"Ticket Relations" validate:"required" view:"hidden"`
	Comments              []Comments        `sql:"-" json:"comments,omitempty" validate:"omitempty" view:"hidden"`
	Contributors          []int64           `pg:",array" sql:"contributors" json:"contributors, omitempty" label:"Contributors" validate:"omitempty" view:"hidden"`
	CommentsCount         int               `sql:"-" json:"comments_count" label:"Comments count" bson:"comments_count" validate:"-" type:"int"`
	IsRead                bool              `sql:"-" json:"is_read" label:"Is Read" validate:"-" type:"bool"`
	FixTime               int               `sql:"fix_time" json:"fix_time" label:"Fix Time" type:"datetime,readonly" view:"watch"`
	FixDescription        string            `sql:"fix_description" json:"fix_description" label:"Fix Description" view:"hidden"`
	ApproveDescription    string            `sql:"approve_description" json:"approve_description" label:"Approve Description" view:"hidden"`
	DisapproveTime        int               `sql:"disapprove_time" json:"disapprove_time" label:"Disapprove Time" type:"datetime,readonly" view:"watch"`
	DisapproveDescription string            `sql:"disapprove_description" json:"disapprove_description" label:"Disapprove Description" view:"hidden"`
	DeletedAt             int               `sql:"deleted_at" json:"deleted_at,omitempty" label:"Deleted At" type:"datetime" view:"watch"`
	CreatedAt             int               `sql:"created_at" json:"created_at,omitempty" label:"Created At" type:"datetime" view:"watch"`
	UpdatedAt             int               `sql:"updated_at" json:"updated_at,omitempty" label:"Updated At" type:"datetime" view:"watch"`
	AdditionalInfo        FWAdditionalInfo  `sql:"additional_info" json:"additional_info" label:"AdditionalInfo"`

	// Related Fields to Applications
	ApplicationAttempts              int    `sql:"-" json:"application_attempts,omitempty" view:"hidden" label:"Application Attempts"`
	ApplicationLeadId                int    `sql:"-" json:"application_lead_id" view:"hidden" label:"Application Lead Id"`
	ApplicationOfficePhone           string `sql:"-" json:"application_office_phone" view:"hidden" label:"Application Office Phone"`
	ApplicationId                    int    `sql:"-" json:"application_id" view:"hidden" label:"Application Id"`
	ApplicationFullName              string `sql:"-" json:"application_full_name" view:"hidden" label:"Application Full Name"`
	ApplicationLeadNumber            string `sql:"-" json:"application_lead_number" view:"hidden" label:"Application Lead Number"`
	ApplicationFormStatus            string `sql:"-" json:"application_form_status" view:"hidden" label:"Application Form Status" type:"string,autocomplete" related:"form_status"`
	ApplicationCSRAssignTo           string `sql:"-" json:"application_csr_assign_to" view:"hidden" label:"Application CSR Assign To" type:"int,autocomplete,readonly" related:"users"`
	ApplicationMDCAssignTo           string `sql:"-" json:"application_mdc_assign_to" view:"hidden" label:"Application MDC Visa Assign To" type:"int,autocomplete,readonly" related:"users"`
	ApplicationMDCBackofficeAssignTo string `sql:"-" json:"application_mdc_backoffice_assign_to" view:"hidden" label:"Application MDC Backoffice Assign To" type:"int,autocomplete,readonly" related:"users"`
	ApplicationMDCEvaluationAssignTo string `sql:"-" json:"application_mdc_evaluation_assign_to" view:"hidden" label:"Application MDC Evaluation Assign To" type:"int,autocomplete,readonly" related:"users"`
	ApplicationEmail                 string `sql:"-" json:"application_email" view:"hidden" label:"Application Email" type:"string,readonly"`
	ApplicationProducts              string `sql:"-" json:"application_products" view:"hidden" label:"Application Products" type:"string,autocomplete,readonly" related:"products"`
	IsRisk                           bool   `sql:"-" json:"is_risk" view:"watch" pg:",use_zero" label:"Is Risk" type:"bool,autocomplete" related:"yes_no"`

	// Related Fields to Documents
	ResultUploadedBy int `sql:"-" json:"result_uploaded_by,omitempty" view:"hidden" label:"Result Uploaded By" type:"int,readonly,autocomplete" related:"users"`
	ResultUploadedAt int `sql:"-" json:"result_uploaded_at,omitempty" view:"hidden" label:"Result Uploaded At" type:"datetime,readonly"`
}
