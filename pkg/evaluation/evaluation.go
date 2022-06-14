package evaluation

const (
	ImmiAgentCSRRole        = "agent_csr"
	ImmiAgentRetentionRole  = "agent_retention"
	ImmiAgentConversionRole = "agent_conversion"
	ImmiAgentReviewRole     = "agent_review"
	ImmiAgentRCICRole       = "agent_rcic"
)

const (
	CRMAgentCSRRole                  = "csr-agent"
	CRMManagerCSRRole                = "csr_manager"
	CRMAgentRetentionRole            = "retention-agent"
	CrmAgentRetentionFileOpeningRole = "retention_file_opening"
	CRMManagerRetentionRole          = "retention_agent_manager"
	CRMAgentReviewRole               = "review"
	CRMManagerReviewRole             = "review_manager"
	CRMAgentMDCRole                  = "mdc_dep"
	CrmManagerMDCDavidRole           = "mdc_david"
)

func GetEvaluationRole(userRole string) string {
	switch userRole {
	case CRMAgentCSRRole, CRMManagerCSRRole:
		return ImmiAgentCSRRole
	case CRMAgentRetentionRole, CRMManagerRetentionRole, CrmAgentRetentionFileOpeningRole:
		return ImmiAgentRetentionRole
	case CRMAgentReviewRole, CRMManagerReviewRole:
		return ImmiAgentReviewRole
	case CRMAgentMDCRole, CrmManagerMDCDavidRole:
		return ImmiAgentRCICRole
	default:
		return ImmiAgentConversionRole
	}
}
