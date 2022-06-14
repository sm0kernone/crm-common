package models

type InstalmentNotices struct {
	Id            int  `sql:"id" json:"id" label:"ID"`
	InstId        int  `sql:"inst_id" json:"inst_id" validate:"required,min=1" label:"Instalment ID"`
	InstPieceId   int  `sql:"inst_piece_id" json:"inst_piece_id" validate:"min=1" label:"Instalment Piece ID"`
	Type          int  `sql:"type,notnull" pg:",use_zero" json:"type" label:"Type" type:"autocomplete,string" related:"instalment_notices_types"`
	Status        int  `sql:"status,notnull" pg:",use_zero" json:"status" label:"Status" type:"autocomplete,string" related:"instalment_notices_statuses"`
	SentAt        *int `sql:"sent_at" json:"sent_at,omitempty" label:"Sent At" type:"datetime"`
	PlannedSendAt *int `sql:"planned_send_at" json:"planned_send_at,omitempty" label:"Planned Send At" type:"datetime"`
	CreatedAt     int  `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt     int  `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}

const (
	InstalmentNoticesStatusCreatedValue = iota
	InstalmentNoticesStatusSuccessValue
	InstalmentNoticesStatusRejectedValue
	InstalmentNoticesStatusPendingValue
	InstalmentNoticesStatusExpiredValue
	InstalmentNoticesStatusDeclinedValue
	InstalmentNoticesStatusFailedValue
	InstalmentNoticesStatusActivatedValue

	// must be trim + lower
	InstalmentNoticesStatusCreated  = "created"  // simple type
	InstalmentNoticesStatusRejected = "rejected" // simple type
	InstalmentNoticesStatusSuccess  = "success"  // simple type

	InstalmentNoticesStatusPending   = "pending"
	InstalmentNoticesStatusExpired   = "expired"
	InstalmentNoticesStatusDeclined  = "declined"
	InstalmentNoticesStatusFailed    = "failed"
	InstalmentNoticesStatusActivated = "activated"
)
