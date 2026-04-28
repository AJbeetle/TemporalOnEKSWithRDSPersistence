package temporal

import (
	"learningTemp/internal/config"

	"go.temporal.io/sdk/client"
)

func NewClient(cfg *config.Config) (client.Client, error) {
	return client.Dial(client.Options{
		HostPort:  cfg.TemporalHost,
		Namespace: "default",
	})
}
