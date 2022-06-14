package models

type TicketType int
type TicketStatus int
type TicketPriority int

const (
	AssessmentTicketType TicketType = iota + 1
	AssessmentResultTicketType
	VisaTypeTicketType
	QueriesFromClientDashTicketType
	BackofficeVisaTicketType
	RCICTicketType
	RetentionToBO
)

const (
	ToDoTicketStatus TicketStatus = iota + 1
	InProgressTicketStatus
	ApprovedTicketStatus
	FixedTicketStatus
	DisapprovedTicketStatus
)

const (
	LowTicketPriority TicketPriority = iota + 1
	MediumTicketPriority
	HighTicketPriority
	UrgentTicketPriority
)

const (
	ActionCreate  = "1"
	ActionUpdate  = "2"
	TicketTypeCD  = "1"
	CommentTypeCD = "2"
)
