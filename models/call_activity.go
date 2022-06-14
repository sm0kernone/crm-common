package models

import "database/sql"

const EffectiveCallDuration = 300000

type CallActivityDB struct {
	Id               int             `sql:"id" json:"id" label:"ID"`
	CallID           sql.NullString  `sql:"call_id" json:"call_id" label:"Call ID"`
	CrmAgentId       int             `sql:"crm_agent_id" json:"crm_agent_id" label:"Agent Id"`
	CrmLeadId        int             `sql:"crm_lead_id" json:"crm_lead_id" label:"Lead Id"`
	CrmApplicationId int             `sql:"crm_application_id" json:"crm_application_id" label:"Application ID"`
	AgentCodeNew     string          `sql:"agent_code_new" json:"agent_code_new" label:"Agent Code New" group:"details"`
	Status           CallStatus      `sql:"status" json:"status" label:"Status" group:"details"`
	StartTime        sql.NullFloat64 `sql:"start_time" json:"start_time" label:"Start Time" type:"datetime"`
	EndTime          sql.NullFloat64 `sql:"end_time" json:"end_time" label:"End Time" type:"datetime"`
	WaitMS           sql.NullFloat64 `sql:"waitms" json:"waitms" label:"Wait(ms)" type:"datetime"`
	DurationMS       sql.NullFloat64 `sql:"durationms" json:"durationms" label:"Duration(ms)" type:"datetime"`
	ProgressMS       sql.NullFloat64 `sql:"progressms" json:"progressms" label:"Progress(ms)" type:"datetime"`
	BillMS           sql.NullFloat64 `sql:"billms" json:"billms" label:"Bill(ms)" type:"datetime"`
	CcWait           sql.NullFloat64 `sql:"cc_wait" json:"cc_wait" label:"CC Wait(ms)" type:"datetime"`
	HangupCause      sql.NullString  `sql:"hangup_cause" json:"hangup_cause" label:"Hangup Cause" group:"details"`
	CallSystemType   int             `sql:"call_system_type" json:"call_system_type"`
	CallReason       string          `sql:"call_reason" json:"call_reason"`
	CreatedAt        int             `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt        int             `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}

type CallActivity struct {
	Id               int        `sql:"id" json:"id" label:"ID"`
	CallID           string     `sql:"call_id" json:"call_id" label:"Call ID"`
	CrmAgentId       int        `sql:"crm_agent_id" json:"crm_agent_id" label:"Agent Id"`
	CrmLeadId        int        `sql:"crm_lead_id" json:"crm_lead_id" label:"Lead Id"`
	CrmApplicationId int        `sql:"crm_application_id" json:"crm_application_id" label:"Application ID"`
	AgentCodeNew     string     `sql:"agent_code_new" json:"agent_code_new" label:"Agent Code New" group:"details"`
	Status           CallStatus `sql:"status" json:"status" label:"Status" group:"details"`
	StartTime        float64    `sql:"start_time" json:"start_time" label:"Start Time" type:"datetime"`
	EndTime          float64    `sql:"end_time" json:"end_time" label:"End Time" type:"datetime"`
	WaitMS           float64    `sql:"waitms" json:"waitms" label:"Wait(ms)" type:"datetime"`
	DurationMS       float64    `sql:"durationms" json:"durationms" label:"Duration(ms)" type:"datetime"`
	ProgressMS       float64    `sql:"progressms" json:"progressms" label:"Progress(ms)" type:"datetime"`
	BillMS           float64    `sql:"billms" json:"billms" label:"Bill(ms)" type:"datetime"`
	CcWait           float64    `sql:"cc_wait" json:"cc_wait" label:"CC Wait(ms)" type:"datetime"`
	HangupCause      string     `sql:"hangup_cause" json:"hangup_cause" label:"Hangup Cause" group:"details"`
	CallSystemType   int        `sql:"call_system_type" json:"call_system_type"`
	CallReason       string     `sql:"call_reason" json:"call_reason"`
	CreatedAt        int        `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt        int        `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}

func PrepareCallActivityForCloner(callActivityDb *CallActivityDB) *CallActivity {
	return &CallActivity{
		Id:               callActivityDb.Id,
		CallID:           callActivityDb.CallID.String,
		CrmAgentId:       callActivityDb.CrmAgentId,
		CrmLeadId:        callActivityDb.CrmLeadId,
		CrmApplicationId: callActivityDb.CrmApplicationId,
		AgentCodeNew:     callActivityDb.AgentCodeNew,
		Status:           callActivityDb.Status,
		StartTime:        callActivityDb.StartTime.Float64,
		EndTime:          callActivityDb.EndTime.Float64,
		WaitMS:           callActivityDb.WaitMS.Float64,
		DurationMS:       callActivityDb.DurationMS.Float64,
		ProgressMS:       callActivityDb.ProgressMS.Float64,
		BillMS:           callActivityDb.BillMS.Float64,
		CcWait:           callActivityDb.CcWait.Float64,
		HangupCause:      callActivityDb.HangupCause.String,
		CallSystemType:   callActivityDb.CallSystemType,
		CallReason:       callActivityDb.CallReason,
		CreatedAt:        callActivityDb.CreatedAt,
		UpdatedAt:        callActivityDb.UpdatedAt,
	}
}

type CallStatus string

var (
	CallStartStatus  CallStatus = "call:start"
	CallRingStatus              = "call:ring"
	CallAnswerStatus            = "call:answer"
	CallEndStatus               = "call:end"
)
