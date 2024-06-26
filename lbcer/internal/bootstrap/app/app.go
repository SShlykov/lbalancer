package app

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/SShlykov/lbalancer/lbcer/internal/bootstrap/registry"
	configPkg "github.com/SShlykov/lbalancer/lbcer/internal/config"
	loggerPkg "github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger"
)

type App struct {
	ctx    context.Context
	cancel context.CancelFunc
	config *configPkg.Config

	logger loggerPkg.Logger
}

func New(domainCtx context.Context, config *configPkg.Config) (*App, error) {
	ctx, cancel := context.WithCancel(domainCtx)
	app := &App{ctx: ctx, cancel: cancel, config: config}

	inits := []func() error{
		app.initLogger,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (app *App) Run() error {
	ctx, cancel := signal.NotifyContext(app.ctx, os.Interrupt)
	defer cancel()

	var wg sync.WaitGroup
	stoppedChan := make(chan struct{})

	app.logger.Info("starting application", loggerPkg.Any("config", app.config))
	app.logger.Debug("debug level enabled")

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer app.cancel()
		defer app.logger.Debug("http.Server stopped")
		if err := registry.RunProxy(ctx, app.logger, app.config.App); err != nil {
			app.logger.Error("http.Server stopped with error", loggerPkg.Error(err))
		}
	}()

	go func() {
		wg.Wait()
		stoppedChan <- struct{}{}
	}()

	return app.closer(ctx)
}
