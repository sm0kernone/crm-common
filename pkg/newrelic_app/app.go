package newrelic_app

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewNewrelicApp(appName, license string) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
}
