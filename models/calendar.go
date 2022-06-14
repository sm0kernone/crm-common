package models

type Calendar struct {
	tableName   struct{}    `sql:"calendar" pg:"calendar"`
	Id          int         `sql:"id" json:"id" label:"ID"`
	Type        string      `sql:"type" json:"type" label:"Type" validate:"required,min=1,max=255,oneof=event schedule"`
	Name        string      `sql:"name" json:"name" label:"Name" validate:"required,min=1,max=255"`
	Description string      `sql:"description" json:"description" label:"Description" validate:"omitempty,min=1,max=255"`
	WorkingDays float64     `sql:"working_days,notnull" pg:",use_zero" json:"working_days" label:"Working Days" validate:"omitempty,max=1"`
	RelatedId   int         `sql:"related_id,notnull" pg:",use_zero" json:"related_id" label:"Related ID"`
	RelatedTo   string      `sql:"related_to,notnull" pg:",use_zero" json:"related_to" label:"Related To" validate:"required,oneof=user_groups all"`
	Data        interface{} `sql:"data" json:"data" label:"Additional Data"`
	Date        int         `sql:"date" json:"date" label:"Date" type:"datetime" validate:"required"`
	CreatedAt   int         `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt   int         `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}
