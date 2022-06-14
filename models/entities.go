package models

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/utils"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

type EntityType int

// All possible entities
const (
	LeadEntityType EntityType = iota + 1
	OpportunityEntityType
	RefundEntityType
	ApplicationEntityType
	SwipesEntityType
	DocumentEntityType
	VisaEntityType
	RequirementEntityType
	TicketsEntityType
	FeedbackFormEntityType
	PaymentHistoryEntityType
	BinsEntityType
	InstallmentsEntityType
	UsersEntityType
	PartialsEntityType
	CallActivityEntityType

	LeadEntity                   = "leads"
	OpportunityEntity            = "opportunities"
	RefundEntity                 = "refunds"
	ApplicationEntity            = "applications"
	SwipesEntity                 = "swipes"
	DocumentEntity               = "documents"
	VisaEntity                   = "visas"
	TicketEntity                 = "tickets"
	RequirementEntity            = "requirements"
	UsersEntity                  = "users"
	ScheduledNotificationsEntity = "scheduled_notifications"
	PartialsEntity               = "partials"
	CallActivityEntity           = "call_activity"
)

var EntityList = []string{LeadEntity, OpportunityEntity, RefundEntity, ApplicationEntity, SwipesEntity, VisaEntity}

var EntityTypeToString = map[EntityType]string{
	LeadEntityType:         LeadEntity,
	OpportunityEntityType:  OpportunityEntity,
	RefundEntityType:       RefundEntity,
	ApplicationEntityType:  ApplicationEntity,
	SwipesEntityType:       SwipesEntity,
	DocumentEntityType:     DocumentEntity,
	VisaEntityType:         VisaEntity,
	RequirementEntityType:  RequirementEntity,
	TicketsEntityType:      TicketEntity,
	UsersEntityType:        UsersEntity,
	CallActivityEntityType: CallActivityEntity,
}

var EventTimeFromModel = map[EntityType]func(model []byte) int{
	OpportunityEntityType: func(model []byte) int {
		opp := new(Opportunities)
		err := json.Unmarshal(model, opp)
		if err != nil {
			log.Error("failed unmarshalling bytes to opp: ", err)
			return 0
		}
		return opp.PaymentTime
	},
	PaymentHistoryEntityType: func(model []byte) int {
		ph := new(PaymentHistory)
		err := json.Unmarshal(model, ph)
		if err != nil {
			log.Error("failed unmarshalling bytes to ph: ", err)
			return 0
		}

		if !utils.InListInt(ph.Status, []int{PaymentHistoryStatusRejectedValue, PaymentHistoryStatusSuccessValue, PaymentHistoryStatusFailedValue}) {
			return 0
		}

		return ph.UpdatedAt
	},
	RefundEntityType: func(model []byte) int {
		r := new(Refunds)
		json.Unmarshal(model, r)
		return r.CreatedAt
	},
	PartialsEntityType: func(model []byte) int {
		p := new(Partials)
		json.Unmarshal(model, p)
		return p.CreatedAt
	},
}
