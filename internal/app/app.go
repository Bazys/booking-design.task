package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/router"
	"applicationDesignTest/internal/storage"
)

// TODO: вынести в конфиг
const port = ":8080"

func NewApp(logger logger.Logger) *App {
	return &App{
		Log: logger,
	}
}

type App struct {
	Log logger.Logger
}

func (a *App) Run(ctx context.Context) error {
	a.Log.Info("server: starting...")

	mux := router.NewRouter(storage.NewDB(), a.Log)

	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		a.Log.Info("server: listen monitor server on port", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			a.Log.Errorf("error listening for server: %s", err)
		}
	}()

	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	sig := <-termSignal
	a.Log.Info("server: shutting down... reason:", sig.String())

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	a.Log.Info("server closed")
	return nil
}
