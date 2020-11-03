package main

import (
	"net"
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/syntaqx/buildkit/pkg/config"
	"github.com/syntaqx/buildkit/pkg/router"
)

type Server struct {
	cfg *config.Config
	log *zap.Logger
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		log: zap.NewNop(),
	}
}

func (s *Server) Command() *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "start integrated server",
		Flags:  s.Flags(),
		Before: s.Before,
		Action: s.Action,
	}
}

func (s *Server) Before(c *cli.Context) error {
	logger, err := setupLogger(s.cfg)
	if err != nil {
		return err
	}

	s.log = logger
	return nil
}

func (s *Server) Action(c *cli.Context) error {
	defer s.log.Sync() // nolint

	tracing, err := setupTracing(s.cfg)
	if err != nil {
		return err
	}
	if tracing != nil {
		defer tracing.Close()
	}

	db, err := setupStorage(s.cfg)
	if err != nil {
		return err
	}
	if db != nil {
		_ = db
		// 	defer db.Close()
	}

	listenAddr := s.cfg.Server.Addr
	if s.cfg.Server.Port != "" {
		host, _, _ := net.SplitHostPort(listenAddr)
		listenAddr = net.JoinHostPort(host, s.cfg.Server.Port)
	}

	srv := &http.Server{
		Addr:         listenAddr,
		Handler:      router.NewRouter(s.cfg, s.log),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.log.Info("starting http server", zap.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			s.log.Error("http server closed unexpectedly", zap.Error(err))
			return err
		}
	}

	return nil
}

func (s *Server) Flags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "server-addr",
			Value:       "0.0.0.0:8080",
			Usage:       "address to bind the server",
			EnvVars:     []string{"BUILDKIT_SERVER_ADDR"},
			Destination: &s.cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "server-port",
			Value:       "",
			Usage:       "port to bind to, overrides server-addr (heroku, etc.)",
			EnvVars:     []string{"PORT"},
			Hidden:      true,
			Destination: &s.cfg.Server.Port,
		},
		&cli.BoolFlag{
			Name:        "server-debug",
			Usage:       "enable server debug endpoint",
			EnvVars:     []string{"BUILDKIT_SERVER_DEBUG"},
			Destination: &s.cfg.Server.Debug,
		},
		&cli.StringFlag{
			Name:        "server-base-url",
			Value:       "http://localhost:8080",
			Usage:       "public url to the server",
			EnvVars:     []string{"BUILDKIT_SERVER_BASE_URL"},
			Destination: &s.cfg.Server.BaseURL,
		},
	}

	flags = append(flags, databaseFlags(s.cfg)...)
	flags = append(flags, metricsFlags(s.cfg)...)
	flags = append(flags, tracingFlags(s.cfg)...)

	return flags
}
