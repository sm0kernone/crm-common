package models

const SettingCommissionType = `commission_type`
const SettingDayPassword = `day_password`
const SettingNightPassword = `night_password`
const SettingCascadePayment = `cascade_payment`
const SettingUseOldCrm = `use_old_crm`
const SettingAttemptTime = `attempt_time`
const SettingAnswerTime = `answer_time`
const SettingEffectiveTime = `effective_time`
const SettingAgentLimitAttempts = `agent_limit_attempts` // don't change value please
const SettingNewFeatureSipAttemptsBlocking = `enable_new_feature_sip_attempts_blocking`
const SettingNewFeatureAfterEffectiveCallBlocking = `enable_new_feature_effective_call_blocking`
const SettingWhiteListIP = `white_list_ip`
const SettingCallActivityNotification = `call_activity_notification`
const SettingEnableInstalments = `enable_instalments`
const SettingTicketPrefix = `ticket_prefix`
const OnlineProducts = `online_products`
const PositiveRiskStatuses = `positive_risk_statuses`
const SettingCallActivityWithNewCRM = `call_activity_new_crm`
const SettingShowSunCity = "show_sun_city"
const SettingEnableRestorePassword = "enable_restore_password"
const SettingRestorePasswordPeriod = "restore_password_period"
const SettingRestoreCallLimits = "restore_call_limits"
const SettingGMTCountries = "gmt_countries"
const SettingSipUrl = "sip_url"

type Settings struct {
	Id             int         `sql:"id" json:"id,omitempty"`
	Value          string      `sql:"value" json:"value"`
	AdditionalData interface{} `sql:"additional_data" json:"additional_data,omitempty"`
	Type           string      `sql:"type" json:"type,omitempty"`
	Name           string      `sql:"name" json:"name,omitempty"`
	Label          string      `sql:"label" json:"label"`
	CreatedAt      int         `sql:"created_at" json:"created_at,omitempty"`
	UpdatedAt      int         `sql:"updated_at" json:"updated_at,omitempty"`
}
