package httptransport

import (
    "net/http"

    handlers "social-networking-platform/posts-service/internal/handler/http"
    "social-networking-platform/posts-service/internal/middleware"
)

func NewRouter(serviceName string) http.Handler {
    mux := http.NewServeMux()

    healthHandler := handlers.NewHealthHandler(serviceName)
    featureHandler := handlers.NewPostHandler()

    mux.HandleFunc("/health", healthHandler.Health)

    mux.HandleFunc("/api/v1/posts", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            featureHandler.CreatePost(w, r)
            return
        }
        http.NotFound(w, r)
    })
    mux.HandleFunc("/api/v1/posts/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            featureHandler.GetPost(w, r)
        case http.MethodPatch:
            featureHandler.UpdatePost(w, r)
        case http.MethodDelete:
            featureHandler.DeletePost(w, r)
        default:
            http.NotFound(w, r)
        }
    })


    return middleware.RequestID(
        middleware.Logging(serviceName)(
            middleware.Recovery(mux),
        ),
    )
}
