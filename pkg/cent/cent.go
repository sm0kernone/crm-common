package cent

import (
	"github.com/sm0kernone/crm-common/pkg/config"
	"github.com/centrifugal/gocent"
)

func New(cfg config.Centrifugo) *gocent.Client {
	return gocent.New(gocent.Config{
		Addr: cfg.URL,
		Key:  cfg.ApiKey,
	})
}
