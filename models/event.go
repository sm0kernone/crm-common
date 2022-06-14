package models

type (
	ObjectType  string
	EventAction string
)

const (
	EventDocumentType    ObjectType = "document"
	EventTicketType      ObjectType = "ticket"
	EventApplicationType ObjectType = "application"
	EventRequirementType ObjectType = "requirement"

	ActionTicketCreate      EventAction = "ticket_create"
	ActionTicketReply       EventAction = "ticket_reply"
	ActionTicketClose       EventAction = "ticket_close"
	ActionTicketFixed       EventAction = "ticket_fixed"
	ActionTicketDisapproved EventAction = "ticket_disapproved"

	ActionTicketFixedByClient EventAction = "ticket_fixed_by_client"
	ActionTicketVisaRespond   EventAction = "ticket_visa_respond"

	ActionDocumentUploadedByClient    EventAction = "document_uploaded_by_client"
	ActionDocumentLoaUploadedByClient EventAction = "loa_document_uploaded_by_client"
	ActionDocumentPreApproved         EventAction = "document_pre_approved"
	ActionDocumentApproved            EventAction = "document_approved"

	ActionVisaProcessingStatusChange EventAction = "application_visa_processing_status"

	UsersTopicPattern = "events-%v"
)

type Event struct {
	UserCrmID      int                    `json:"user_id"`
	ObjectType     ObjectType             `json:"object_type"`
	ObjectModel    interface{}            `json:"object_model"`
	Action         EventAction            `json:"action"`
	AdditionalInfo map[string]interface{} `json:"additional_info"`
	IsRead         bool                   `json:"is_read"`
	IsClicked      bool                   `json:"is_clicked"`
}
