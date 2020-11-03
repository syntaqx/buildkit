package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	doc "github.com/utahta/swagger-doc"
	"go.bobheadxi.dev/zapx/util/contextx"
	"go.bobheadxi.dev/zapx/zhttp"
	"go.uber.org/zap"

	"github.com/syntaqx/buildkit/pkg/config"

	apiv1 "github.com/syntaqx/buildkit/pkg/api/v1"
	restapiv1 "github.com/syntaqx/buildkit/pkg/api/v1/restapi"
)

func NewRouter(cfg *config.Config, log *zap.Logger) http.Handler {
	r := chi.NewRouter()

	chilog := zhttp.NewMiddleware(log, zhttp.LogFields{
		func(ctx context.Context) zap.Field {
			return zap.String("req.id", contextx.String(ctx, middleware.RequestIDKey))
		},
	})

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(chilog.Logger)
	r.Use(chilog.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.AllowAll().Handler)
	r.Use(secure.New(secure.Options{
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "script-src $NONCE",
	}).Handler)

	r.Route("/v1", func(v1 chi.Router) {
		api, err := apiv1.New()
		if err != nil {
			log.Fatal("unable to initialize apiv1 routes", zap.Error(err))
		}

		if api != nil {
			v1.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(restapiv1.SwaggerJSON))
			})

			v1.Handle("/docs", doc.NewRedocHandler("/v1/swagger", "/v1/docs"))
			v1.Mount("/", middleware.NoCache(api.Handler))
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"v1": "%s/v1"}`, cfg.Server.BaseURL)
	})

	if cfg.Server.Debug {
		r.Mount("/debug", middleware.Profiler())
	}

	return r
}
