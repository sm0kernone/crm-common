package models

var (
	AgentType  = "agent"
	ClientType = "client"
)

type Documents struct {
	Name          string `sql:"name" json:"name" label:"name" type:"string"`
	Id            int    `sql:"id" json:"id,omitempty" label:"ID" type:"int"`
	EntityId      int    `sql:"entity_id" json:"entity_id,omitempty" label:"Entity ID" type:"int" bson:"entity_id"`
	EntityType    string `sql:"entity_type" json:"entity_type,omitempty" label:"Entity Type" validate:"required,oneof=lead opportunity application visas" type:"string" bson:"entity_type"`
	LeadNumber    string `sql:"lead_number" json:"lead_number,omitempty" label:"Lead Number" type:"string" validate:"required" bson:"lead_number"`
	Url           string `sql:"url" json:"url" label:"Url" type:"string" validate:"required"`
	Status        string `sql:"status" json:"status" label:"Status" validate:"required,oneof=ready signed expired" type:"string"`
	IsSign        bool   `sql:"is_sign,notnull" json:"is_sign" label:"Is Sign" validate:"required" type:"bool" bson:"is_sign"`
	Type          string `sql:"type" json:"type,omitempty" label:"Type" validate:"required,oneof=receipt retainer retainer_spouse result result_spouse visa_result custom cv pp ph rp ed tf cv_spouse pp_spouse ph_spouse rp_spouse ed_spouse tf_spouse ifi iur ofr potc adl rlff pof dip vielts loe vmc pcc veh vbio sppf cop loa wes vra st passport credit residence bank" type:"string"`
	System        string `sql:"system,notnull" json:"system" label:"System" validate:"required,oneof=cvb cvcd cvad" type:"string"`
	UpdatedByType string `sql:"updated_by_type,notnull" json:"updated_by_type" label:"Updated by Type" validate:"required,oneof=agent client" type:"string"`
	UpdatedByID   int    `sql:"updated_by_id,notnull" json:"updated_by_id" label:"Updated by ID" validate:"required" type:"int"`
	Hash          string `sql:"hash" json:"hash,omitempty" label:"Hash" type:"string"`
	Mime          string `sql:"mime" json:"mime,omitempty" label:"Mime" type:"string"`
	Size          int64  `sql:"size" json:"size,omitempty" label:"Size" type:"int"`
	ReviewStatus  string `sql:"review_status" json:"review_status,omitempty" label:"Review Status" type:"string,autocomplete" related:"review_statuses"`
	Ext           string `sql:"ext" json:"ext,omitempty" label:"Extension" type:"string"`
	TransactionID string `sql:"-" json:"transaction_id,omitempty"`
	SignTime      int    `sql:"sign_time" json:"sign_time,omitempty" label:"Sign Time" type:"datetime"`
	Class         int    `sql:"class" json:"class,omitempty" label:"Class" type:"int"`
	DeletedAt     int    `sql:"deleted_at" json:"deleted_at,omitempty"`
	CreatedAt     int    `sql:"created_at" json:"created_at,omitempty"`
	UpdatedAt     int    `sql:"updated_at" json:"updated_at,omitempty"`

	ChangedAttributes []string `sql:"-" json:"changed_attributes,omitempty"`
}
