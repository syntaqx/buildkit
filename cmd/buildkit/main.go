package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"

	"github.com/syntaqx/buildkit/pkg/cmd"
	"github.com/syntaqx/buildkit/pkg/config"
)

func main() {
	cfg := config.Load()

	app := &cli.App{
		Name:  "buildkit",
		Flags: globalFlags(cfg),
		Commands: []*cli.Command{
			cmd.Server(cfg),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func globalFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "set logging level",
			EnvVars:     []string{"BUILDKIT_LOG_LEVEL"},
			Destination: &cfg.Logging.Level,
		},
		&cli.StringFlag{
			Name:        "log-format",
			Value:       "json",
			Usage:       "set logging format",
			EnvVars:     []string{"BUILDKIT_LOG_FORMAT"},
			Destination: &cfg.Logging.Format,
		},
	}
}
