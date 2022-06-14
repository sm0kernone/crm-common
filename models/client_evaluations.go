package models

type ClientEvaluations struct {
	Id                      int                    `sql:"id" json:"id" label:"ID" validate:"required"`
	ClientImmigrationInfoId int                    `sql:"client_immigration_info_id" json:"client_immigration_info_id" label:"Client Immigration Info Id"`
	EvaluationUUID          string                 `sql:"evaluation_uuid" json:"evaluation_uuid" label:"Evaluation UUID"`
	Status                  int                    `sql:"status,notnull" pg:",use_zero" json:"status" label:"Status"`
	AdditionalInfo          map[string]interface{} `sql:"additional_info" json:"additional_info" label:"Additional Info"`
	CreatedAt               int                    `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt               int                    `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
	DeletedAt               int                    `sql:"deleted_at" json:"deleted_at" label:"Deleted At" type:"datetime"`
}
