package mandrill

import (
	"errors"
	"github.com/keighl/mandrill"
)

func NewClient(apiKey string) (*mandrill.Client, error) {
	client := mandrill.ClientWithKey(apiKey)

	pong, err := client.Ping()
	if err != nil {
		return nil, err
	}

	if pong != "PONG!" {
		return nil, errors.New("failed creating mandrill client")
	}

	return client, nil
}
