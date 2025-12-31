package http

import (
    "encoding/json"
    "net/http"
    "errors"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/mongo"

    "yuno-faqman-reciever/internal/domain"
    "yuno-faqman-reciever/internal/service"
)

func RegisterThemaRoutes(mux *http.ServeMux, client *mongo.Client) {
    mux.HandleFunc("/thema", themaHandler(client))
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

func resolveThemaSelector(r *http.Request) (uuid.UUID, string, error) {
    idStr := r.URL.Query().Get("id")
    title := r.URL.Query().Get("title")

    switch {
    case idStr != "" && title != "":
        return uuid.Nil, "", errors.New("only one of id or title allowed")

    case idStr != "":
        id, err := uuid.Parse(idStr)
        if err != nil {
            return uuid.Nil, "", errors.New("invalid uuid")
        }
        return id, "", nil

    case title != "":
        return uuid.Nil, title, nil

    default:
        return uuid.Nil, "", errors.New("missing id or title")
    }
}


func themaHandler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {

        case http.MethodGet:
            handleThemaGet(w, r, client)

        case http.MethodPost:
            handleThemaCreate(w, r, client)

        case http.MethodPut:
            handleThemaUpdate(w, r, client)

        case http.MethodDelete:
            handleThemaDelete(w, r, client)

        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

func handleThemaGet(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
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

func handleThemaCreate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    var payload struct {
        Title string `json:"title"`
    }

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    thema, err := service.CreateThema(ctx, client, payload.Title)
    if err == service.ErrDuplicateTitle {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(thema)
}

func handleThemaUpdate(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    idStr := r.URL.Query().Get("id")
    oldTitle := r.URL.Query().Get("title")

    var (
        id uuid.UUID
        err error
    )

    switch {
    case idStr != "":
        id, err = uuid.Parse(idStr)
        if err != nil {
            http.Error(w, "invalid uuid", http.StatusBadRequest)
            return
        }

    case oldTitle != "":
        // ok, handled below

    default:
        http.Error(w, "missing id or title", http.StatusBadRequest)
        return
    }

    var payload struct {
        Title string `json:"title"`
    }

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    if payload.Title == "" {
        http.Error(w, "title required", http.StatusBadRequest)
        return
    }

    if idStr != "" {
        err = service.UpdateThemaTitle(ctx, client, id, payload.Title)
    } else {
        err = service.UpdateThemaTitleByTitle(ctx, client, oldTitle, payload.Title)
    }

    if err == service.ErrDuplicateTitle {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    if err == mongo.ErrNoDocuments {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func handleThemaDelete(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
    ctx := r.Context()

    id, title, err := resolveThemaSelector(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = service.DeleteThema(ctx, client, id, title)
    if err == mongo.ErrNoDocuments {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
