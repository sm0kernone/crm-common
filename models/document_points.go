package models

type DocumentPoints struct {
	Id             int     `sql:"id" json:"id,omitempty" label:"Id" type:"int" validate:"omitempty"`
	DocumentType   string  `json:"document_type" sql:"document_type" validate:"required" type:"string"`
	Points         float32 `json:"points" sql:"points" validate:"required"`
	NeedToBeSigned bool    `json:"need_to_be_signed" sql:"need_to_be_signed" validate:"required"`
	Active         bool    `json:"active" sql:"active" validate:"required"`
	CreatedAt      int     `sql:"created_at" json:"created_at,omitempty" label:"Created At" type:"datetime"`
	UpdatedAt      int     `sql:"updated_at" json:"updated_at,omitempty" label:"Updated At" type:"datetime"`
}
