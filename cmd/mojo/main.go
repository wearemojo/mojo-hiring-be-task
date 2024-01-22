package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wearemojo/lead-be-task/session"
	"github.com/wearemojo/lead-be-task/session/app"
	"github.com/wearemojo/lead-be-task/session/db"
	"github.com/wearemojo/lead-be-task/session/service"
	"github.com/wearemojo/mojo-public-go/lib/clog"
	"github.com/wearemojo/mojo-public-go/lib/config"
	"github.com/wearemojo/mojo-public-go/lib/crpc"
	"github.com/wearemojo/mojo-public-go/lib/middleware/request"
)

type ServerConfig struct {
	Logging clog.Config `json:"logging"`

	Server config.Server `json:"server"`
}

func main() {
	cfg := &ServerConfig{
		Logging: clog.Config{
			Format: "text",
			Debug:  true,
		},

		Server: config.Server{
			Addr: "127.0.0.1:3000",
		},
	}

	log := cfg.Logging.Configure(context.Background())

	var svc session.Service = &service.Service{App: &app.App{
		DB: db.New(context.Background()),
	}}

	// create a new RPC server
	rpc := crpc.NewServer(unsafeNoAuthentication)

	// add logging middleware
	rpc.Use(crpc.Logger())

	rpc.Register("list_by_device_id", "2024-01-20", service.ListByDeviceIDSchema, svc.ListByDeviceID)
	rpc.Register("list_by_platform", "2024-01-20", service.ListByPlatformSchema, svc.ListByPlatform)

	mux := chi.NewRouter()
	mux.Use(request.Logger(log))
	mux.With(request.StripPrefix("/v1")).Handle("/v1/*", rpc)

	httpServer := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.WithField("addr", cfg.Server.Addr).Info("listening")
	if err := cfg.Server.ListenAndServe(httpServer); err != nil {
		log.WithError(err).Fatal("listen failed")
	}
}

func unsafeNoAuthentication(next crpc.HandlerFunc) crpc.HandlerFunc {
	return func(res http.ResponseWriter, req *crpc.Request) error {
		return next(res, req)
	}
}
