package models

type CsrPoints struct {
	Id               int     `sql:"id" json:"id,omitempty" label:"Id" type:"int"`
	DocumentPointsID int     `json:"document_points_id" sql:"document_points_id"`
	CsrAgentCrmID    int     `json:"csr_agent_crm_id" sql:"csr_agent_id"`
	Points           float32 `json:"points" sql:"points"`
	LeadNumber       string  `json:"lead_number" sql:"lead_number"`
	CreatedAt        int     `sql:"created_at" json:"created_at,omitempty" label:"Created At" type:"datetime"`
	UpdatedAt        int     `sql:"updated_at" json:"updated_at,omitempty" label:"Updated At" type:"datetime"`
}
