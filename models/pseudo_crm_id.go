package models

type PseudoCrmId struct {
	CrmId    int           `sql:"crm_id,pk" json:"crm_id,omitempty"`
	EntityId int           `sql:"entity_id" json:"entity_id"`
	Entity   CRMEntityName `sql:"entity" json:"entity"`
}

type CRMEntityName string

var (
	LeadsCRMName          CRMEntityName = "Leads"
	CallActivitiesCRMName CRMEntityName = "call_activities"
)
