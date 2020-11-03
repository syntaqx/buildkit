package main

import (
	"github.com/urfave/cli/v2"

	"github.com/syntaqx/buildkit/pkg/config"
)

func databaseFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "database-dsn",
			Value:       "postgres://postgres@localhost/api",
			Usage:       "database dsn",
			EnvVars:     []string{"BUILDKIT_DATABASE_DSN", "BUILDKIT_DATABASE_URL"},
			Destination: &cfg.Database.DSN,
		},
	}
}

func tracingFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "open tracing endpoint",
			EnvVars:     []string{"BUILDKIT_TRACING_ENDPOINT"},
			Destination: &cfg.Tracing.Endpoint,
		},
	}
}

func metricsFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "metrics-addr",
			Value:       "0.0.0.0:8090",
			Usage:       "address to bind the metrics",
			EnvVars:     []string{"BUILDKIT_METRICS_ADDR"},
			Destination: &cfg.Metrics.Addr,
		},
		&cli.StringFlag{
			Name:        "metrics-token",
			Value:       "",
			Usage:       "token to make metrics secure",
			EnvVars:     []string{"BUILDKIT_METRICS_TOKEN"},
			Destination: &cfg.Metrics.Token,
		},
	}
}
