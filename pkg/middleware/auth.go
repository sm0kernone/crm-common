package middleware

import (
	"bitbucket.org/ssinbeti/crm-common/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const authorizationHeader = "Authorization"
const secretKeyHeader = "secret_key"
const TimezoneHeader = "Timezone"

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// User session
	session   Session
	SecretKey string
}

func New(s Session, secret string) *Service {
	return &Service{session: s, SecretKey: secret}
}

// Session contains user info in Redis
type Session interface {
	Get(string) (*models.Users, error)
}

func (s *Service) MVFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var bearer = c.Request().Header.Get(authorizationHeader)

			if len(bearer) < 8 {
				return echo.ErrUnauthorized
			}

			tokenValue := bearer[7:]

			var name = bearer[:6]

			if name != "Bearer" {
				return echo.ErrUnauthorized
			}

			if tokenValue == "" {
				return echo.ErrUnauthorized
			}

			user, err := s.session.Get(tokenValue)
			if err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			c.Set("user", user)
			c.Set("token", tokenValue)
			c.Set(TimezoneHeader, s.getTimezone(c))

			return next(c)
		}
	}
}

func (s *Service) MVSecretKeyFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if c.Request().Header.Get(secretKeyHeader) != s.SecretKey {
				return echo.ErrUnauthorized
			}
			c.Set(TimezoneHeader, s.getTimezone(c))

			return next(c)
		}
	}
}

func (s *Service) getTimezone(c echo.Context) int {
	timezone := c.Request().Header.Get(TimezoneHeader)
	if len(timezone) == 0 {
		return 0
	}

	timezoneVal, _ := strconv.Atoi(timezone)

	return timezoneVal
}
