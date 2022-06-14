package models

type Requirements struct {
	Id             int                    `sql:"id" json:"id" label:"ID"`
	VisaId         int                    `sql:"visa_id" json:"visa_id" label:"Visa Id"`
	ApproveTime    int                    `sql:"approve_time" json:"approve_time" label:"Approve Time" type:"datetime"`
	Type           string                 `sql:"type" json:"type" label:"Type" type:"string,autocomplete" related:"requirement_types"`
	Status         *int                   `sql:"status" json:"status" label:"Status" type:"string,autocomplete" related:"requirement_statuses"`
	AdditionalInfo map[string]interface{} `sql:"additional_info,notnull" json:"additional_info" label:"Additional Info"`
	Settings       map[string]interface{} `sql:"settings,notnull" json:"settings" label:"Settings"`
	Position       int                    `sql:"position" json:"position" label:"Position"`
	Documents      []int                  `sql:"-" json:"documents,omitempty" label:"Documents"`
	CreatedAt      int                    `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt      int                    `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}
