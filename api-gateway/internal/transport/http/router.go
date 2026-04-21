package httptransport

import (
    "net/http"

    handlers "social-networking-platform/api-gateway/internal/handler/http"
    "social-networking-platform/api-gateway/internal/middleware"
)

func NewRouter(serviceName string) http.Handler {
    mux := http.NewServeMux()

    healthHandler := handlers.NewHealthHandler(serviceName)
    featureHandler := handlers.NewProxyHandler()

    mux.HandleFunc("/health", healthHandler.Health)

    mux.HandleFunc("/api/v1/auth/", featureHandler.ProxyAuth)
    mux.HandleFunc("/api/v1/users/", featureHandler.ProxyUsers)
    mux.HandleFunc("/api/v1/posts/", featureHandler.ProxyPosts)
    mux.HandleFunc("/api/v1/feed", featureHandler.ProxyFeed)
    mux.HandleFunc("/api/v1/notifications", featureHandler.ProxyNotifications)


    return middleware.RequestID(
        middleware.Logging(serviceName)(
            middleware.Recovery(mux),
        ),
    )
}
