package registry

import (
	"context"
	"net/http"

	"github.com/SShlykov/lbalancer/lbcer/internal/app/urls"
	configPkg "github.com/SShlykov/lbalancer/lbcer/internal/config"
	loggerPkg "github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger"
)

func RunProxy(ctx context.Context, logger loggerPkg.Logger, config configPkg.App) error {
	handler := http.NewServeMux()
	urlGetter := urls.New(config.Hosts)

	handler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host

		urlGetter.Get().ServeHTTP(w, req)
	})

	server := http.Server{Addr: config.Port, Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("server stopped with error", loggerPkg.Error(err))
		}
	}()

	return closer(ctx, &server)
}

func closer(ctx context.Context, srv *http.Server) error {
	select {
	case <-ctx.Done():
		return srv.Shutdown(ctx)
	}
}
