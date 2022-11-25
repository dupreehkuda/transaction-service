package api

import (
	"context"
	"errors"
	"github.com/dupreehkuda/transaction-service/internal/config"
	"github.com/dupreehkuda/transaction-service/internal/fileKeeper"
	"github.com/dupreehkuda/transaction-service/internal/handlers"
	i "github.com/dupreehkuda/transaction-service/internal/interfaces"
	"github.com/dupreehkuda/transaction-service/internal/logger"
	"github.com/dupreehkuda/transaction-service/internal/processors"
	"github.com/dupreehkuda/transaction-service/internal/storage"
	"github.com/dupreehkuda/transaction-service/internal/worker"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type api struct {
	handlers i.Handlers
	logger   *zap.Logger
	config   *config.Config
}

func NewByConfig() *api {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	store := storage.New(cfg.DatabasePath, log)
	store.CreateSchema()

	fkeep := fileKeeper.New(cfg.IndexFile, log)

	proc := processors.New(fkeep, log)

	handle := handlers.New(proc, log)

	wrkr := worker.New(fkeep, proc, store, log)
	wrkr.Run()

	return &api{
		handlers: handle,
		logger:   log,
		config:   cfg,
	}
}

// Run runs the service
func (a api) Run() {
	serv := &http.Server{Addr: a.config.Address, Handler: a.router()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				a.logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := serv.Shutdown(shutdownCtx)
		if err != nil {
			a.logger.Fatal("Error shutting down", zap.Error(err))
		}
		a.logger.Info("Server shut down", zap.String("port", a.config.Address))
		serverStopCtx()
	}()

	a.logger.Info("Server started", zap.String("port", a.config.Address))
	err := serv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		a.logger.Fatal("Cant start server", zap.Error(err))
	}

	<-serverCtx.Done()
}