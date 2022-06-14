package segmentation

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/config"
	"bitbucket.org/ssinbeti/crm-common/pkg/curl"
	"bitbucket.org/ssinbeti/crm-common/pkg/logger"
	"encoding/json"
	"fmt"
)

var (
	sendEmailAPI    = "/api/email"
	updateEntityAPI = "/api/update-entity"
)

type Client struct {
	baseURL                   string
	apiKey                    string
	apiSecret                 string
	callScheduleURlPattern    string
	approveScheduleURLPattern string
	addToCalendarUrlPattern   string
	log                       logger.Logger
}

func New(cfg config.Segmentation, log logger.Logger) *Client {
	return &Client{
		baseURL:                   cfg.BaseURL,
		apiKey:                    cfg.ApiKey,
		apiSecret:                 cfg.ApiSecret,
		callScheduleURlPattern:    cfg.CallScheduleURL,
		approveScheduleURLPattern: cfg.ApproveScheduleURL,
		addToCalendarUrlPattern:   cfg.AddToCalendarURL,
		log:                       log,
	}
}

func (c *Client) SendScheduleNotificationEvent(email, firstName, scheduleTime, scheduleDate, time, hash, eventHash, campaignSource string) (map[string]interface{}, int, error) {
	data := map[string]string{
		"template":                   "CV - Confirm Call From Agent",
		"global[email]":              email,
		"global[subject]":            "Please Confirm Schedule Time",
		"vars[first_name]":           firstName,
		"vars[schedule_time]":        scheduleTime,
		"vars[Date]":                 scheduleDate,
		"vars[Time]":                 time,
		"vars[hash]":                 hash,
		"vars[call_schedule_url]":    c.getCallScheduleURL(hash, eventHash),
		"vars[approve_schedule_url]": c.getApproveScheduleURL(hash, eventHash),
		"vars[event_hash]":           eventHash,
		"vars[campaign_source]":      campaignSource,
	}

	url := c.baseURL + sendEmailAPI

	resp, statusCode, err := curl.PostForm(url, data, c.getAuthHeaders())
	if c.log.CheckError(err, c.SendScheduleNotificationEvent) != nil {
		return nil, 0, err
	}

	respMap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respMap)
	if c.log.CheckError(err, c.SendScheduleNotificationEvent) != nil {
		return nil, 0, err
	}

	return respMap, statusCode, nil
}

func (c *Client) SendCallConfirmedEvent(email, firstName, scheduleTime, scheduleDate, hash, eventHash string) (map[string]interface{}, int, error) {
	data := map[string]string{
		"template":                "CV - Confirm Call",
		"global[email]":           email,
		"global[subject]":         "Thank for set schedule time",
		"vars[FIRST_NAME]":        firstName,
		"vars[EVENT_HASH]":        eventHash,
		"vars[ENTITY_HASH]":       hash,
		"vars[Date]":              scheduleDate,
		"vars[Time]":              scheduleTime,
		"vars[CALL_SCHEDULE_URL]": c.getCallScheduleURL(hash, eventHash),
		"vars[ADD_TO_CALENDAR]":   c.getAddToCalendarURL(scheduleTime, scheduleDate),
	}

	url := c.baseURL + sendEmailAPI

	resp, statusCode, err := curl.PostForm(url, data, c.getAuthHeaders())
	if c.log.CheckError(err, c.SendCallConfirmedEvent) != nil {
		return nil, 0, err
	}

	respMap := make(map[string]interface{})
	err = json.Unmarshal(resp, &respMap)
	if c.log.CheckError(err, c.SendCallConfirmedEvent) != nil {
		return nil, 0, err
	}

	return respMap, statusCode, nil
}

func (c *Client) UpdateEntity(entityType, leadNumber, fields string) ([]byte, int, error) {
	data := map[string]string{
		"type":           entityType,
		"fields":         fields,
		"account_number": leadNumber,
	}

	url := c.baseURL + updateEntityAPI

	return curl.PostForm(url, data, c.getAuthHeaders())
}

func (c *Client) getCallScheduleURL(hash, eventHash string) string {
	return fmt.Sprintf(c.callScheduleURlPattern, hash, eventHash)
}

func (c *Client) getApproveScheduleURL(hash, eventHash string) string {
	return fmt.Sprintf(c.approveScheduleURLPattern, hash, eventHash)
}

func (c *Client) getAddToCalendarURL(scheduleTime, scheduleDate string) string {
	start := scheduleTime + "%20" + scheduleDate

	return fmt.Sprintf(c.addToCalendarUrlPattern, start)
}

func (c *Client) getAuthHeaders() map[string]string {
	return map[string]string{
		"Api-Key":    c.apiKey,
		"Api-Secret": c.apiSecret,
	}
}
