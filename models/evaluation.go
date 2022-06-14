package models

const (
	AutomationPermissionRetainerAgreement = "automation-api-v1.retainer-agreement"
)

const (
	ApplicantTypeMain string = "main"
)

const AdminKey = "admin-lcgroups"

type ErrorsStruct struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Applicant struct {
	UUID      string `json:"uuid"`
	Weight    int    `json:"weight"`
	Status    int    `json:"status"`
	Type      string `json:"type"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

type CreateEvaluationReq struct {
	Email       string   `json:"email"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Key         string   `json:"key"`
	Password    string   `json:"password"`
	Timezone    string   `json:"timezone"`
	Username    string   `json:"username"`
	Evaluations []string `json:"evaluations"`
	ProjectUUID string `json:"project_uuid"`
}

type ClientStruct struct {
	Key string `json:"key"`
}

type EvaluationStruct struct {
	OriginEvaluationUUID string       `json:"origin_evaluation_uuid"`
	ClientEvaluationUUID string       `json:"client_evaluation_uuid"`
	MainApplicantUUID    string       `json:"main_applicant_uuid"`
	Name                 string       `json:"name"`
	Progress             int          `json:"progress"`
	Status               int          `json:"status"`
	AgentStatus          int          `json:"agent_status"`
	UpdatedByClient      string       `json:"updated_by_client"`
	UpdatedByUser        string       `json:"updated_by_user"`
	UpdatedAt            string       `json:"updated_at"`
	CreatedAt            string       `json:"created_at"`
	DeletedAt            string       `json:"deleted_at"`
	SubmittedAt          string       `json:"submitted_at"`
	Client               ClientStruct `json:"client"`
	Applicants           []Applicant  `json:"applicants"`
}

type ResultCreateResp struct {
	Username           string             `json:"username"`
	Email              string             `json:"email"`
	Key                string             `json:"key"`
	Status             int                `json:"status"`
	Timezone           string             `json:"timezone"`
	CreatedEvaluations []EvaluationStruct `json:"created_evaluations"`
	UpdatedEvaluations []EvaluationStruct `json:"updated_evaluations"`
}

type CreateResp struct {
	Success bool             `json:"success"`
	Code    float64          `json:"code"`
	Result  ResultCreateResp `json:"result"`
	Errors  interface{}
}

//Agents
type CreateUpdateAgentRequest struct {
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	AdminKey    string   `json:"admin_key"`
	UserKey     string   `json:"user_key"`
	Password    string   `json:"password"`
	Status      int      `json:"status"`
	Timezone    string   `json:"timezone"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	DeletedAt   int      `json:"deleted_at,omitempty"`
}

type CreateAgentResp struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	Result  struct {
		Key            string `json:"key"`
		UserParentID   int    `json:"user_parent_id"`
		PasswordHash   string `json:"password_hash"`
		Email          string `json:"email"`
		Name           string `json:"name"`
		Status         int    `json:"status"`
		Timezone       string `json:"timezone"`
		AdditionalInfo struct {
			CompanyID int `json:"company_id"`
			ProjectID int `json:"project_id"`
		} `json:"additional_info"`
		TokenAuthExpireAt interface{} `json:"token_auth_expire_at"`
		PasswordChangedAt interface{} `json:"password_changed_at"`
		CreatedAt         string      `json:"created_at"`
		UpdatedAt         string      `json:"updated_at"`
		ID                int         `json:"id"`
	} `json:"result"`
	Errors []ErrorsStruct `json:"errors"`
}
