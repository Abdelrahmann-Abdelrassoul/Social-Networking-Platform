package bootstrap

import (
    "fmt"
    "net/http"

    "social-networking-platform/api-gateway/internal/config"
    httptransport "social-networking-platform/api-gateway/internal/transport/http"
)

type App struct {
    Router http.Handler
}

func NewApp(cfg config.Config) (*App, error) {
    router := httptransport.NewRouter(cfg.ServiceName)
    if router == nil {
        return nil, fmt.Errorf("failed to initialize router")
    }
    return &App{Router: router}, nil
}
