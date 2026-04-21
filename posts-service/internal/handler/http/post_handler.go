package handlers

import "net/http"

type PostHandler struct {}

func NewPostHandler() *PostHandler {
    return &PostHandler{}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "CreatePost is not implemented yet",
        },
    })
}


func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "GetPost is not implemented yet",
        },
    })
}


func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "UpdatePost is not implemented yet",
        },
    })
}


func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusNotImplemented, map[string]any{
        "success": false,
        "error": map[string]any{
            "code": "NOT_IMPLEMENTED",
            "message": "DeletePost is not implemented yet",
        },
    })
}
