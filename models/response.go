package models

import "net/http"

type (
	SearchResponse struct {
		Route      *http.Request `json:"-"`
		Page       int           `json:"page"`
		PageSize   int           `json:"page_size"`
		PageCount  int           `json:"page_count"`
		Rows       int           `json:"rows"`
		Data       interface{}   `json:"data"`
		Info       interface{}   `json:"info"`
		Additional interface{}   `json:"additional,omitempty"`
		User       interface{}   `json:"user,omitempty"`
	}
	InfoResponse struct {
		Route  *http.Request `json:"-"`
		Data   interface{}   `json:"data"`
		Info   interface{}   `json:"info"`
		Groups interface{}   `json:"groups"`
	}
)
