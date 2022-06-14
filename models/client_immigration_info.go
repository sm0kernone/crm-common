package models

type ClientsImmigrationInfo struct {
	Id             int                    `sql:"id" json:"id" label:"ID"`
	LeadCrmId      int                    `sql:"lead_crm_id" json:"lead_crm_id" label:"Lead Crm Id"`
	ClientUUID     string                 `sql:"client_uuid" json:"client_uuid" label:"Client UUID"`
	Status         int                    `sql:"status,notnull" pg:",use_zero" json:"status" label:"Status"`
	Timezone       string                 `sql:"timezone" json:"timezone" label:"Timezone"`
	AdditionalInfo map[string]interface{} `sql:"additional_info" json:"additional_info" label:"Additional Info"`
	CreatedAt      int                    `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt      int                    `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
	DeletedAt      int                    `sql:"deleted_at" json:"deleted_at" label:"Deleted At" type:"datetime"`
}
