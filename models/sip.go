package models

type SipRequest struct {
	Event            string `json:"event"`
	UserID           int    `json:"user_id"`
	ID               string `json:"id"`
	CallID           string `json:"call_id"`
	TS               int    `json:"ts"`
	Type             string `json:"type"`
	QueueID          int    `json:"queue_id"`
	CidNum           string `json:"cid_num"`
	DstNum           string `json:"dst_num"`
	A2               string `json:"a2"`
	DstUserId        int    `json:"dst_user_id"`
	AC               int    `json:"ac"`
	HangupCause      string `json:"hangup_cause"`
	DeviceID         int    `json:"device_id"`
	DeviceName       string `json:"device_name"`
	DstDeviceID      int    `json:"dst_device_id"`
	Progress         int    `json:"progress"`
	Progressms       int    `json:"progressms"`
	Waitms           int    `json:"waitms"`
	Wait             int    `json:"wait"`
	Durationms       int    `json:"durationms"`
	Duration         int    `json:"duration"`
	CcWait           int    `json:"cc_wait"`
	SrcDeviceID      int    `json:"src_device_id"`
	SrcUserID        int    `json:"src_user_id"`
	LeadID           string `json:"lead_id"`
	Billms           int    `json:"billms"`
	Bill             int    `json:"bill"`
	Effective        bool   `json:"effective"`
	Efficient        bool   `json:"efficient"`
	ApplicationCrmID string `json:"crm_application_id"`
	LeadNumber       string `json:"lead_number"`
	DialerID         string `json:"dialer_id"`
	CrmApplicationID string `json:"$crm_application_id"`
	Country          string `json:"$country"`
	CallReason       string `json:"$call_reason"`
}

type SipOriginateCallResponse struct {
	Success bool   `json:"Success"`
	UUID    string `json:"uuid"`
	Msg     string `json:"msg"`
}
