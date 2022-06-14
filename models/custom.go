package models

type IeltsEmailCustomRequest struct {
	LeadName                   string `json:"lead_name"`
	Email                      string `json:"email"`
	LeadSourceSubsection       string `json:"lead_source_subsection"`
	PurchasedProgramEnrollment string `json:"purchased_program_enrollment"`
}
