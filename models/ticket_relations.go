package models

type TicketRelations struct {
	Id         int        `sql:"id, pk" json:"-" label:"ID" type:"int" validate:"omitempty"`
	TicketsID  int        `sql:"ticket_id" json:"ticket_id" label:"Ticket Id" type:"int" validate:"omitempty"`
	EntityID   int        `sql:"entity_id" json:"entity_id" label:"Entity Id" type:"int" validate:"required"`
	EntityType EntityType `sql:"entity_type" json:"entity_type" label:"Entity Type" type:"int" validate:"required"`
}
