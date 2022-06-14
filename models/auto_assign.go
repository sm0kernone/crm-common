package models

type AutoAssignType string

var (
	LeadsAutoAssignType        = "leads"
	ApplicationsAutoAssignType = "applications"
)

type AutoAssign struct {
	TableName   struct{}       `sql:"auto_assign" json:"-"`
	Id          int            `sql:"id" json:"id" label:"ID"`
	Name        string         `sql:"name" json:"name" label:"Name" validate:"required,min=1,max=255"`
	Type        AutoAssignType `sql:"type" json:"type" label:"Type" validate:"required,min=1,max=255"`
	Active      bool           `sql:"active,notnull" pg:",use_zero" json:"active" label:"Active"`
	AssignLimit int            `sql:"assign_limit" json:"assign_limit" label:"Assign Limit"`
	DeletedAt   int            `sql:"deleted_at" json:"deleted_at,omitempty" label:"Deleted At" type:"datetime" view:"watch"`
	CreatedAt   int            `sql:"created_at" json:"created_at,omitempty" label:"Created At" type:"datetime" view:"watch"`
	UpdatedAt   int            `sql:"updated_at" json:"updated_at,omitempty" label:"Updated At" type:"datetime" view:"watch"`
}
