package models

type AutoAssignConditions struct {
	TableName    struct{} `sql:"auto_assign_conditions" json:"-"`
	Id           int      `sql:"id" json:"id,omitempty" label:"ID"`
	Group        int      `sql:"group" json:"group" label:"Group" validate:"required,min=-1,max=255"`
	AutoAssignID int      `sql:"auto_assign_id" json:"auto_assign_id" label:"Auto Assign ID" validate:"required"`
	FieldName    string   `sql:"field_name" json:"field_name" label:"Field Name" validate:"required,min=1,max=255"`
	Condition    string   `sql:"condition" json:"condition" label:"Condition" validate:"required,min=1,max=255"`
	Operator     string   `sql:"operator" json:"operator" label:"operator" validate:"required,min=1,max=255"`
	Value        string   `sql:"value" json:"value" label:"value" validate:"required,min=1,max=255"`
	CreatedAt    int      `sql:"created_at" json:"created_at,omitempty" label:"Created At" type:"datetime"`
	UpdatedAt    int      `sql:"updated_at" json:"updated_at,omitempty" label:"Updated At" type:"datetime"`
}
