package app

import loggerPkg "github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger"

func (app *App) initLogger() error {
	level, err := loggerPkg.LevelFromString(app.config.Logger.Level)
	if err != nil {
		return err
	}
	mode, err := loggerPkg.ModeFromString(app.config.Logger.Mode)
	if err != nil {
		return err
	}

	logger, err := loggerPkg.Setup(level, mode)
	if err != nil {
		return err
	}

	app.logger = logger
	return nil
}
