package models

type UserSessions struct {
	Id         int    `sql:"id" json:"id" label:"ID"`
	UserId     int    `sql:"user_id" json:"user_id" label:"User ID"`
	Status     string `sql:"status" json:"status" label:"Status"`
	StatusTime int    `sql:"status_time" json:"status_time" label:"Status Time"`
	CreatedAt  int    `sql:"created_at" json:"created_at" label:"Created At"`
	UpdatedAt  int    `sql:"updated_at" json:"updated_at" label:"Updated At"`
}
