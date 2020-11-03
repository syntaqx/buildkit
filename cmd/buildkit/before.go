package main

import (
	"io"

	"github.com/uber/jaeger-client-go"
	tracecfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/syntaqx/buildkit/pkg/config"
	"github.com/syntaqx/buildkit/pkg/store"
)

func setupLogger(cfg *config.Config) (*zap.Logger, error) {
	var lvl zapcore.Level
	if err := lvl.Set(cfg.Logging.Level); err != nil {
		return nil, err
	}

	zconfig := zap.NewProductionConfig()

	zconfig.Level.SetLevel(lvl)

	zconfig.Encoding = cfg.Logging.Format
	zconfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := zconfig.Build(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, err
	}
	// defer logger.Sync()

	zap.ReplaceGlobals(logger)
	return logger, nil
}

func setupTracing(cfg *config.Config) (io.Closer, error) {
	if cfg.Tracing.Endpoint != "" {
		closer, err := tracecfg.Configuration{
			Sampler: &tracecfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &tracecfg.ReporterConfig{
				LocalAgentHostPort: cfg.Tracing.Endpoint,
			},
		}.InitGlobalTracer("buildkit")

		if err != nil {
			return nil, err
		}

		return closer, nil
	}

	return nil, nil
}

func setupStorage(cfg *config.Config) (store.Store, error) {
	return nil, nil
}
