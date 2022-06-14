package models

import (
	"github.com/labstack/echo/v4"
	"strings"
)

var (
	RetentionAgentRole  = "retention-agent"
	ConversionAgentRole = "agent"
)

type Users struct {
	Id            int    `sql:"id" json:"id,omitempty" label:"ID"`
	CrmId         int    `sql:"crm_id" json:"crm_id" label:"CRM ID"`
	CrmApiKey     string `sql:"crm_api_key" json:"crm_api_key" label:"CRM API Key"`
	Username      string `sql:"username" json:"username,omitempty" group:"short" label:"Username"`
	Email         string `sql:"email" json:"email,omitempty" group:"short" label:"Email"`
	FullName      string `sql:"full_name" json:"full_name,omitempty" label:"Full Name" group:"short"`
	Avatar        string `sql:"avatar" json:"avatar,omitempty" label:"Avatar"`
	Password      string `sql:"password" json:"-" label:"Password"`
	Role          string `sql:"role" json:"role,omitempty" values:"admin,agent" group:"short" label:"Role"`
	Points        int    `sql:"points,notnull" pg:",use_zero" json:"points" label:"Total Points" group:"short"`
	LevelPoints   int    `sql:"level_points,notnull" pg:",use_zero" json:"level_points" label:"Points" group:"short"`
	LevelId       int    `sql:"level_id" json:"level_id" label:"Level ID"`
	CreatedAt     int    `sql:"created_at" json:"created_at,omitempty" group:"short" label:"Created At" type:"datetime"`
	UpdatedAt     int    `sql:"updated_at" json:"updated_at,omitempty" group:"short" label:"Updated At" type:"datetime"`
	DialCode      string `sql:"dial_code" json:"dial_code,omitempty" group:"short" label:"Dial Code" `
	ExtensionCode string `sql:"extension_code" json:"extension_code,omitempty" group:"short" label:"Extension Code" `
	AgentCode     string `sql:"agent_code" json:"agent_code,omitempty" group:"short" label:"Agent Code"`
	Deleted       bool   `sql:"deleted,notnull" pg:",use_zero" json:"deleted,omitempty" label:"Deleted"`
	IsBlocked     bool   `sql:"is_blocked,notnull" json:"is_blocked,omitempty" label:"Is Blocked" pg:",use_zero"`
	Token         string `sql:"-" json:"token,omitempty" label:"Token"`
	CentToken     string `sql:"-" json:"cent_token,omitempty" label:"Cent Token"`
	WorkType      string `sql:"work_type" json:"work_type,omitempty" label:"Work Type"`
	ImmiKey       string `sql:"immi_key" json:"immi_key" label:"Immi Key"`

	ChangedAttributes []string `sql:"-" json:"changed_attributes,omitempty"`
}

func (u *Users) IsRetention() bool {
	return strings.Contains(u.Role, "retention")
}

func (u *Users) IsConversion() bool {
	return u.Role == ConversionAgentRole
}

// Session represents keyvalue session interface
type Session interface {
	Get(string) (*Users, error)
}

// Context represents context interface
type Context interface {
	GetUser(echo.Context) *Users
	GetTimezone(c echo.Context) int
}
