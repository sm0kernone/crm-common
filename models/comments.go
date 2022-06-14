package models

type Comments struct {
	Id          int    `sql:"id" json:"id"`
	CrmId       int    `sql:"crm_id" json:"-"`
	Author      string `sql:"-, author" json:"author,omitempty"`
	UserId      int    `sql:"user_id" json:"user_id"`
	RelatedId   int    `sql:"related_id" json:"related_id"`
	RelatedType string `sql:"related_type" json:"related_type"`
	Body        string `sql:"body" json:"body" validate:"required"`
	CreatedAt   int    `sql:"created_at" json:"created_at"`
	UpdatedAt   int    `sql:"updated_at" json:"updated_at"`
}
