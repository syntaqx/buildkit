package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/syntaqx/buildkit"
	"github.com/syntaqx/buildkit/middleware"
)

func main() {
	logger := buildkit.NewDefaultLogger()

	r := buildkit.New()

	r.Use(middleware.Logger(logger))

	r.AddRoute(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	r.Get("/users/:userId", func(w http.ResponseWriter, r *http.Request) {
		userId := buildkit.GetParam(r, "userId")
		fmt.Fprintf(w, "Hello, %s!", userId)
	})

	srv := &http.Server{
		Addr:    net.JoinHostPort("", "8080"),
		Handler: r,
	}

	logger.Info("http server listening at %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("http server error %v", err)
	}
}
