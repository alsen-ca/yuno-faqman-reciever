package http

import (
    "encoding/json"
    "net/http"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/domain"
    "yuno-faqman-reciever/internal/service"
)

func RegisterThemaRoutes(mux *http.ServeMux, client *mongo.Client) {
    mux.HandleFunc("/thema/new", newThemaHandler(client))
    mux.HandleFunc("/thema", getThemaHandler(client))
}

func respondSingle(w http.ResponseWriter, thema domain.Thema, err error) {
    if err == mongo.ErrNoDocuments {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(thema)
}



func newThemaHandler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if r.Method != http.MethodPost {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var req struct {
            Title string `json:"title"`
        }

        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }

        thema, err := service.CreateThema(r.Context(), client, req.Title)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(thema)
    }
}

func getThemaHandler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if r.Method != http.MethodGet {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

        ctx := r.Context()
        idStr := r.URL.Query().Get("id")
        title := r.URL.Query().Get("title")

        switch {
        case idStr != "":
            id, err := uuid.Parse(idStr)
            if err != nil {
                http.Error(w, "invalid uuid", http.StatusBadRequest)
                return
            }

            thema, err := service.GetThemaByID(ctx, client, id)
            respondSingle(w, thema, err)

        case title != "":
            thema, err := service.GetThemaByTitle(ctx, client, title)
            respondSingle(w, thema, err)

        default:
            themas, err := service.ListThemas(ctx, client)
            if err != nil {
                http.Error(w, "internal error", http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(themas)
        }
    }
}


