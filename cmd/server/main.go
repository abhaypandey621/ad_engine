package main

import (
	"context"
	"fmt"
	"net/http"

	"os"

	"github.com/abhaypandey621/targeting-engine/internal/endpoint"
	"github.com/abhaypandey621/targeting-engine/internal/repository"
	"github.com/abhaypandey621/targeting-engine/internal/service"
	"github.com/abhaypandey621/targeting-engine/internal/transport"
	"github.com/abhaypandey621/targeting-engine/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	// Load config
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("failed to load config")
	}

	// Connect to MySQL
	database, err := repository.NewMySQLDBWithDSN(cfg.MySQL.DSN)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to DB")
	}
	defer database.Close()

	// Service
	svc := service.NewService(database, cfg)

	// Start campaign refresher
	logrus.WithField("interval", cfg.CampaignRefreshInterval()).Info("Starting campaign refresher")
	svc.StartCampaignRefresher(context.Background())

	// Endpoint
	endpt := endpoint.MakeServeAdEndpoint(svc)

	// Transport (HTTP handler)
	handler := transport.Handler(endpt)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logrus.WithField("addr", addr).Info("adservice listening")
	logrus.Info("Prometheus metrics available at /metrics")
	if err := http.ListenAndServe(addr, handler); err != nil {
		logrus.WithError(err).Fatal("server error")
	}
}
