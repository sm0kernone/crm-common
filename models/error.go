package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResp struct {
	Error interface{} `json:"error"`
}

type EmailError struct {
	EmailError string `json:"Email"`
}

var (
	ErrEmailAlreadyUsed = echo.NewHTTPError(
		http.StatusBadRequest, ErrorResp{EmailError{EmailError: "This email is already in use"}})
)
