package models

import (
	"time"
)

type ShiftType int

func (s *ShiftType) IsValid() bool {
	for _, shiftType := range AllShiftTypes {
		if *s == shiftType {
			return true
		}
	}

	return false
}

const (
	ShiftTypeDefault ShiftType = iota
	ShiftTypeRetentionAssign
	ShiftTypeConversionAssign
	ShiftTypeRetentionFOAssign
)

var AllShiftTypes = []ShiftType{ShiftTypeDefault, ShiftTypeRetentionAssign, ShiftTypeConversionAssign, ShiftTypeRetentionFOAssign}

type ShiftHours int

const (
	ShiftHoursDefault ShiftHours = iota
	ShiftHoursDay
	ShiftHoursAfternoon
	ShiftHoursNight
	ShiftHoursAll
)

var AllShiftHours = []ShiftHours{ShiftHoursDefault, ShiftHoursDay, ShiftHoursAfternoon, ShiftHoursNight, ShiftHoursAll}

func (s *ShiftHours) IsValid() bool {
	for _, shiftHour := range AllShiftHours {
		if *s == shiftHour {
			return true
		}
	}

	return false
}

type UserGroups struct {
	tableName         struct{}   `sql:"user_groups" pg:"user_groups" json:"-"`
	Id                int        `sql:"id" json:"id" label:"ID"`
	Goal              int        `sql:"goal,notnull" pg:",use_zero" json:"goal" label:"Goal"`
	Name              string     `sql:"name" json:"name" label:"Name" validate:"required,min=1,max=255"`
	Icon              string     `sql:"icon" json:"icon" label:"Icon" validate:"required,min=1,max=255"`
	Description       string     `sql:"description" json:"description" label:"Description" validate:"required,min=1,max=255"`
	Active            bool       `sql:"active,notnull" pg:",use_zero" json:"active" label:"Active"`
	ShiftType         ShiftType  `sql:"shift_type,notnull" pg:",use_zero" json:"shift_type" label:"Shift Type"`
	ShiftHours        ShiftHours `sql:"shift_hours,notnull" pg:",use_zero" json:"shift_hours" label:"Shift Work Hours"`
	Hour              int        `sql:"hour" json:"hour" label:"Hour"`
	Minute            int        `sql:"minute" json:"minute" label:"Minute"`
	Timezone          string     `sql:"timezone" json:"timezone"`
	UseCallLimitation bool       `sql:"use_call_limitation, notnull" pg:",use_zero" json:"use_call_limitation" label:"Use Call Limitation"`
	CreatedAt         int        `sql:"created_at" json:"created_at" label:"Created At" type:"datetime"`
	UpdatedAt         int        `sql:"updated_at" json:"updated_at" label:"Updated At" type:"datetime"`
}

func (ug *UserGroups) setTimeStamp() {
	ug.CreatedAt = int(time.Now().UTC().Unix())
	ug.UpdatedAt = int(time.Now().UTC().Unix())
	return
}

func (ug *UserGroups) updateTime() {
	ug.UpdatedAt = int(time.Now().UTC().Unix())
}

func (ug *UserGroups) isUpdate() bool {
	return ug.Id > 0
}
