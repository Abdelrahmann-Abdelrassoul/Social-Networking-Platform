package handlers

import "net/http"

type ProxyHandler struct {}

func NewProxyHandler() *ProxyHandler {
    return &ProxyHandler{}
}

func (h *ProxyHandler) ProxyAuth(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "ProxyAuth is not implemented yet",
        },
    })
}


func (h *ProxyHandler) ProxyUsers(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "ProxyUsers is not implemented yet",
        },
    })
}


func (h *ProxyHandler) ProxyPosts(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "ProxyPosts is not implemented yet",
        },
    })
}


func (h *ProxyHandler) ProxyFeed(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "ProxyFeed is not implemented yet",
        },
    })
}


func (h *ProxyHandler) ProxyNotifications(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "ProxyNotifications is not implemented yet",
        },
    })
}
