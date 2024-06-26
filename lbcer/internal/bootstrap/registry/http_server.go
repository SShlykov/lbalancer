package registry

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	loggerPkg "github.com/SShlykov/lbalancer/lbcer/internal/pkg/logger"
)

var target = "https://jsonplaceholder.typicode.com/todos/1"

func RunProxy(ctx context.Context, logger loggerPkg.Logger) error {
	handler := http.NewServeMux()
	uri, err := url.Parse(target)
	if err != nil {
		return err
	}
	proxy := httputil.NewSingleHostReverseProxy(uri)

	handler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host

		proxy.ServeHTTP(w, req)
	})

	server := http.Server{Addr: ":8080", Handler: handler}

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
