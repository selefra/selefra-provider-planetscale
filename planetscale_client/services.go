package planetscale_client

import (
	"context"
	"github.com/planetscale/planetscale-go/planetscale"
	"os"
	"strings"
)

func Connect(_ context.Context, config *Config) (*planetscale.Client, error) {
	conn, err := planetscale.NewClient(planetscale.WithAccessToken(config.Token))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Organization(_ context.Context, config *Config) string {
	org := os.Getenv("PLANETSCALE_ORGANIZATION")

	if config.Organization != "" {
		org = config.Organization
	}

	return org
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Not Found")
}
