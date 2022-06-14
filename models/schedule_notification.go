package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type NotificationType int

const (
	NotificationTypeBasic NotificationType = iota + 1
	NotificationTypeInstalment
)

type ScheduledNotificationData map[string]interface{}

type Notifications struct {
	Id               int                       `sql:"id" json:"id" label:"ID"`
	UserId           int                       `sql:"user_id" json:"user_id" label:"User Id"`
	Title            string                    `sql:"title" json:"title" label:"Title"`
	Text             string                    `sql:"text" json:"text" label:"Text"`
	Completed        bool                      `sql:"completed,notnull" pg:",use_zero" json:"completed" label:"Completed"`
	RelatedId        int                       `sql:"related_id" json:"related_id" label:"Related Id"`
	RelatedType      string                    `sql:"related_type" json:"related_type" label:"Related Type"`
	RelatedAt        int                       `sql:"related_at" json:"related_at" label:"Related At" type:"datetime"`
	CreatedAt        int                       `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt        int                       `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
	NotifyTime       int                       `sql:"notify_time" json:"notify_time" label:"Related At" type:"datetime"`
	NotificationType NotificationType          `sql:"notification_type" json:"notification_type,omitempty" label:"Notification Type"`
	Data             ScheduledNotificationData `sql:"data" json:"data" label:"Additional Data"`
	Delete           *bool                     `json:"delete,omitempty"`
}

// GetAllQueueRequest is used to send request to processor to retrieve all planned notifications
type GetAllQueueRequest struct {
	Get bool `json:"get"`
}

func GetListOfNotificationFields() []string {
	notification := new(Notifications)
	strValue := reflect.ValueOf(notification)
	var attributes []string

	for i := 0; i < strValue.Elem().NumField(); i++ {
		tagValue := strValue.Elem().Type().Field(i).Tag.Get("sql")
		if len(tagValue) > 0 {
			// if tag value contains , - some additional tags like notnull, do not add it to attributes
			if strings.Contains(tagValue, ",") {
				tagValue = strings.Split(tagValue, ",")[0]
			}
			attributes = append(attributes, tagValue)
		}
	}

	return attributes
}

// Make the Leads struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (h ScheduledNotificationData) Value() (driver.Value, error) {
	return json.Marshal(h)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (h *ScheduledNotificationData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}
