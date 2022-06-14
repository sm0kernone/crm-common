package session

import (
	"bitbucket.org/ssinbeti/crm-common/models"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/labstack/gommon/log"
)

// NewSession creates new redis session store
func NewSession(c *redis.Client) *Session {
	return &Session{client: c}
}

// Session represents in memory session store
type Session struct {
	client *redis.Client
}

// GetAuthUser tries to fetch existing session for given key
func (s *Session) Get(key string) (*models.Users, error) {
	val, err := s.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	user, err := s.decodeItem(val)
	if err != nil {
		log.Error("Failed decoding user from redis to golang struct, err: ", err)
		return nil, err
	}

	return user, nil
}

func (*Session) decodeItem(str string) (*models.Users, error) {
	u := new(models.Users)

	err := json.Unmarshal([]byte(str), u)
	return u, err
}
