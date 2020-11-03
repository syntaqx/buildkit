package main

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/syntaqx/buildkit/pkg/config"
)

type Migrate struct {
	cfg *config.Config
	log *zap.Logger
}

func NewMigrate(cfg *config.Config) *Migrate {
	return &Migrate{
		cfg: cfg,
		log: zap.NewNop(),
	}
}

func (s *Migrate) Command() *cli.Command {
	return &cli.Command{
		Name:   "migrate",
		Usage:  "auto migrate database changes",
		Flags:  s.Flags(),
		Before: s.Before,
		Action: s.Action,
	}
}

func (s *Migrate) Before(c *cli.Context) error {
	logger, err := setupLogger(s.cfg)
	if err != nil {
		return err
	}

	s.log = logger
	return nil
}

func (s *Migrate) Action(c *cli.Context) error {
	defer s.log.Sync() // nolint

	return nil
}

func (s *Migrate) Flags() []cli.Flag {
	flags := []cli.Flag{}

	flags = append(flags, databaseFlags(s.cfg)...)
	flags = append(flags, metricsFlags(s.cfg)...)
	flags = append(flags, tracingFlags(s.cfg)...)

	return flags
}
