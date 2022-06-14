package models

type Dashboards struct {
	Base
	Name          string `sql:"name" json:"name" label:"Name" validate:"required,min=1,max=255"`
	CreatedBy     int    `sql:"created_by" json:"created_by" label:"Created By" validate:"required" related:"users" structs:"-"`
	Collaborators []int  `sql:"collaborators" pg:",array" json:"collaborators,omitempty" label:"Collaborators" validate:"required" related:"users"`
}
