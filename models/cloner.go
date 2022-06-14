package models

type ClonerRequest struct {
	ID           int    `json:"id" sql:"id"`
	URL          string `json:"url" sql:"url"`
	Method       string `json:"method" sql:"method"`
	Request      string `json:"request" sql:"req"`
	OriginatedAt int    `json:"originated_at" sql:"originated_at"`
}

type ClonerResponse struct {
	ID         int    `json:"id" sql:"id"`
	RequestID  int    `json:"request_id" sql:"request_id"`
	Status     int    `json:"status" sql:"status"`
	Response   string `json:"response" sql:"resp"`
	ReceivedAt int    `json:"received_at" sql:"received_at"`
}
