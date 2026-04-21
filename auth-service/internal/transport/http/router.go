package httptransport

import (
    "net/http"

    handlers "social-networking-platform/auth-service/internal/handler/http"
    "social-networking-platform/auth-service/internal/middleware"
)

func NewRouter(serviceName string) http.Handler {
    mux := http.NewServeMux()

    healthHandler := handlers.NewHealthHandler(serviceName)
    featureHandler := handlers.NewAuthHandler()

    mux.HandleFunc("/health", healthHandler.Health)

    mux.HandleFunc("/api/v1/auth/login", featureHandler.Login)
    mux.HandleFunc("/api/v1/auth/callback", featureHandler.Callback)
    mux.HandleFunc("/api/v1/auth/logout", featureHandler.Logout)


    return middleware.RequestID(
        middleware.Logging(serviceName)(
            middleware.Recovery(mux),
        ),
    )
}
