package models

type AutoAssignToUser struct {
	TableName    struct{} `sql:"auto_assign_to_user" json:"-"`
	UserId       int      `sql:"user_id,pk" json:"user_id"`
	AutoAssignID int      `sql:"auto_assign_id,pk" json:"auto_assign_id"`
}
