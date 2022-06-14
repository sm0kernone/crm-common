package assigner

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/curl"
)

type AssignType string
type ConditionType string

var (
	LeadsAssignType        AssignType = "leads"
	ApplicationsAssignType AssignType = "applications"
	CsrAssignType          AssignType = "csr"
)

const (
	AssignConditionType   ConditionType = "assign"
	ReassignConditionType ConditionType = "reassign"
)

type ReassignRequest struct {
	Type      AssignType    `json:"type" validate:"required,oneof=leads applications"`
	Ids       []int         `json:"ids" validate:"required"` //  crm_id of entity
	AssignTo  int           `json:"assign_to" validate:"required"`
	Condition ConditionType `json:"condition"`
}

// SendToAssignSystem sends request to assign system
func SendToAssignSystem(req *ReassignRequest, url string) ([]byte, int, error) {
	return curl.Post(url, req, map[string]string{}, curl.JsonContentType, true)
}
