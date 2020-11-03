package cmd

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/syntaqx/buildkit/pkg/config"
	"github.com/syntaqx/buildkit/pkg/router"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "start integrated server",
		Flags:  serverFlags(cfg),
		Before: serverBefore(cfg),
		Action: serverAction(cfg),
	}
}

func serverFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "server-addr",
			Value:       "0.0.0.0:8080",
			Usage:       "address to bind the server",
			EnvVars:     []string{"BUILDKIT_SERVER_ADDR"},
			Destination: &cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "server-port",
			Value:       "",
			Usage:       "port to bind to, overrides server-addr (heroku, etc.)",
			EnvVars:     []string{"PORT"},
			Hidden:      true,
			Destination: &cfg.Server.Port,
		},
		&cli.BoolFlag{
			Name:        "server-debug",
			Usage:       "enable server debug endpoint",
			EnvVars:     []string{"BUILDKIT_SERVER_DEBUG"},
			Destination: &cfg.Server.Debug,
		},
		&cli.StringFlag{
			Name:        "server-base-url",
			Value:       "http://localhost:8080",
			Usage:       "public url to the server",
			EnvVars:     []string{"BUILDKIT_SERVER_BASE_URL"},
			Destination: &cfg.Server.BaseURL,
		},
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
		&cli.StringFlag{
			Name:        "database-dsn",
			Value:       "postgres://postgres@localhost/api",
			Usage:       "database dsn",
			EnvVars:     []string{"BUILDKIT_DATABASE_DSN", "BUILDKIT_DATABASE_URL"},
			Destination: &cfg.Database.DSN,
		},
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Value:       false,
			Usage:       "enable open tracing",
			EnvVars:     []string{"BUILDKIT_TRACING_ENABLED"},
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "open tracing endpoint",
			EnvVars:     []string{"BUILDKIT_TRACING_ENDPOINT"},
			Destination: &cfg.Tracing.Endpoint,
		},
	}
}

func serverBefore(cfg *config.Config) cli.BeforeFunc {
	return func(c *cli.Context) error {
		return nil
	}
}

func serverAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		logger, err := setupLogger(cfg)
		if err != nil {
			return err
		}
		if logger != nil {
			defer logger.Sync() // nolint
		}

		tracing, err := setupTracing(cfg)
		if err != nil {
			return err
		}
		if tracing != nil {
			defer tracing.Close()
		}

		db, err := setupStorage(cfg)
		if err != nil {
			return err
		}
		if db != nil {
			defer db.Close()
		}

		listenAddr := cfg.Server.Addr
		if cfg.Server.Port != "" {
			host, _, _ := net.SplitHostPort(listenAddr)
			listenAddr = net.JoinHostPort(host, cfg.Server.Port)
		}

		srv := &http.Server{
			Addr:         listenAddr,
			Handler:      router.NewRouter(cfg, logger),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go func() {
			logger.Info("starting http server", zap.String("addr", srv.Addr))
			if err := srv.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					logger.Error("http server closed unexpectedly", zap.Error(err))
				}

				logger.Info("http server closed", zap.Error(err))
			}
		}()

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("failed to shutdown gracefully", zap.Error(err))
			return nil
		}

		logger.Info("shutdown gracefully")
		return nil
	}
}
