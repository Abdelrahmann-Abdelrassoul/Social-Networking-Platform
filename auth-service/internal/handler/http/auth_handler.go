package handlers

import "net/http"

type AuthHandler struct {}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "Login is not implemented yet",
        },
    })
}


func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "Callback is not implemented yet",
        },
    })
}


func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "Logout is not implemented yet",
        },
    })
}
