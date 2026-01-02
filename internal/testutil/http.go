package testutil

import (
    "net/http"

    "yuno-faqman-reciever/internal/middleware"
)

func SetupTestServer(register func(*http.ServeMux),) http.Handler {
    mux := http.NewServeMux()
    register(mux)
    return middleware.Logging(mux)
}
