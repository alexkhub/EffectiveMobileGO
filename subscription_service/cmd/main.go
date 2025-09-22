package main

import (
	"effective_mobile/internal/app"
	"effective_mobile/internal/config"

	"effective_mobile/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// @title Subscription API
// @version 1.0
// @description API для управления подписками
// @host localhost:8080
// @BasePath /api
func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	err = logger.LoadLogger(cfg.LogLevel, cfg.LogFile)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Debug("config load")

	app, err := app.NewApp(cfg)
	if err != nil {
		logrus.Error(err)
	}

	go func() {
		if err := app.Handler.InitRouter().Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			logrus.Error(fmt.Sprintf("server didn't start: %v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.Close(); err != nil {
		logrus.Error(fmt.Sprintf("server error: %v", err))
	}

	logrus.Info("server stopped gracefully")

}
