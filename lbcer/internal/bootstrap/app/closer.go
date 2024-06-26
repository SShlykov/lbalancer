package app

import (
	"context"
	"fmt"
	"time"
)

func (app *App) closer(ctx context.Context) error {
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(app.ctx, 15*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return fmt.Errorf("closing app timed out")
	}
}
