package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

const (
	UpdateHO HistoryOperation = "UPDATE"
	InsertHO HistoryOperation = "INSERT"
	DeleteHO HistoryOperation = "DELETE"
)

type HistoryOperation string

type History struct {
	Id          int              `sql:"id" json:"id,omitempty"`
	Who         string           `sql:"who" json:"user_id"`
	Changes     interface{}      `sql:"changes" json:"changes"`
	Operation   HistoryOperation `sql:"operation" json:"operation"`
	OperationAt int              `sql:"operation_at" json:"operation_at"`
	EntityType  string           `sql:"-" json:"-"`
	EntityID    int              `sql:"entity_id" json:"-"`
}

type LeadsHistory struct {
	History
}

type ApplicationsHistory struct {
	History
}

type HistoryUpdate map[string]ChangedValues

// Make the Leads struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (h HistoryUpdate) Value() (driver.Value, error) {
	return json.Marshal(h)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (h *HistoryUpdate) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}

type ChangedValues struct {
	From interface{} `json:"from"`
	To   interface{} `json:"to"`
}

func GetChangedAttributes(changedValues HistoryUpdate) []string {
	var attributes []string

	for fieldName := range changedValues {
		attributes = append(attributes, fieldName)
	}

	return attributes
}

func GetUpdateChanges(changes interface{}) (map[string][]string, error) {
	changesInput := changes.(map[string]interface{})
	changeOutput := map[string][]string{}

	for key, val := range changesInput {
		valByte, err := json.Marshal(val)
		if err != nil {
			log.Println("failed marshalling changes value to bytes: ", err)
			return nil, fmt.Errorf("failed masrashalling changes value to bytes: %v", err)
		}

		changes := new(ChangedValues)

		err = json.Unmarshal(valByte, changes)
		if err != nil {
			log.Println("failed unmarsalling changes: ", err)
			return nil, fmt.Errorf("failed unmasrashalling changes: %v", err)
		}

		switch changes.From.(type) {
		case string:
			changeOutput[key] = []string{changes.To.(string), changes.From.(string)}
		case bool:
			changeOutput[key] = []string{strconv.FormatBool(changes.To.(bool)), strconv.FormatBool(changes.From.(bool))}
		case float64:
			changeOutput[key] = []string{fmt.Sprintf("%.0f", changes.To.(float64)), fmt.Sprintf("%.0f", changes.From.(float64))}
		case map[string]interface{}:
			for k, v := range changes.From.(map[string]interface{}) {
				cTo := changes.To.(map[string]interface{})
				changeOutput[key] = []string{cTo[k].(string), v.(string)}
			}
		case nil:
			switch changes.To.(type) {
			case float64:
				changeOutput[key] = []string{fmt.Sprintf("%.0f", changes.To.(float64)), "0"}
			case string:
				changeOutput[key] = []string{fmt.Sprintf("%.0f", changes.To.(string)), "0"}
			case bool:
				changeOutput[key] = []string{strconv.FormatBool(changes.To.(bool)), strconv.FormatBool(changes.From.(bool))}
			default:
				changeOutput[key] = []string{"0", "0"}
			}
		}
	}

	return changeOutput, nil
}
