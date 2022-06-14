package models

type InstalmentPieces struct {
	Id          int     `sql:"id" json:"id" label:"ID"`
	InstId      int     `sql:"inst_id" json:"inst_id" validate:"required,min=1" label:"Instalment ID"`
	PaymentId   *int    `sql:"payment_id" json:"payment_id,omitempty" validate:"required,min=1" label:"Payment ID"`
	PaidAmount  Amount  `sql:"paid_amount,notnull" pg:",use_zero" json:"paid_amount,omitempty" type:"float" label:"Paid Amount"`
	QuoteAmount Amount  `sql:"quote_amount,notnull" pg:",use_zero" json:"quote_amount,omitempty" type:"float" label:"Quote Amount"`
	PaidAt      *int    `sql:"paid_at" json:"paid_at,omitempty" label:"Paid At" type:"datetime"`
	PayLink     *string `sql:"pay_link" json:"pay_link,omitempty" label:"Pay Link" type:"string,readonly"`
	Status      int     `sql:"status,notnull" pg:",use_zero" json:"status" label:"Status" type:"autocomplete,string" related:"instalment_pieces_statuses"`
	CreatedAt   int     `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt   int     `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}

const (
	InstalmentPiecesStatusCreatedValue = iota
	InstalmentPiecesStatusSuccessValue
	InstalmentPiecesStatusRejectedValue
	InstalmentPiecesStatusProcessingValue
	InstalmentPiecesStatusPendingValue
	InstalmentPiecesStatusExpiredValue
	InstalmentPiecesStatusDeclinedValue
	InstalmentPiecesStatusFailedValue
	InstalmentPiecesStatusActivatedValue

	// must be trim + lower
	InstalmentPiecesStatusCreated  = "created"  // simple type
	InstalmentPiecesStatusRejected = "rejected" // simple type
	InstalmentPiecesStatusSuccess  = "success"  // simple type

	InstalmentPiecesStatusProcessing = "processing"
	InstalmentPiecesStatusPending    = "pending"
	InstalmentPiecesStatusExpired    = "expired"
	InstalmentPiecesStatusDeclined   = "declined"
	InstalmentPiecesStatusFailed     = "failed"
	InstalmentPiecesStatusActivated  = "activated"
)
