package context

import (
	"github.com/sm0kernone/crm-common/models"
	"github.com/sm0kernone/crm-common/pkg/middleware"
	"github.com/labstack/echo/v4"
)

// New instantiates new service session
func New() *Context {
	return &Context{}
}

// Context represents service session
type Context struct{}

// GetUser returns user data stored in jwt token
func (s *Context) GetUser(c echo.Context) *models.Users {
	userI := c.Get("user")
	if user, ok := userI.(*models.Users); ok {
		return user
	}

	return nil
}

func (s *Context) GetTimezone(c echo.Context) int {
	return c.Get(middleware.TimezoneHeader).(int)
}
