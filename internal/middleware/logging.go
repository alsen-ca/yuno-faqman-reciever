package middleware

import (
    "net"
    "net/http"
    "strings"
    "encoding/json"
    "os"
    "time"
)

type statusWriter struct {
    http.ResponseWriter
    status int
}

func LogJSON(v any) {
    b, _ := json.Marshal(v)
    os.Stdout.Write(b)
    os.Stdout.Write([]byte("\n"))
}

func (w *statusWriter) WriteHeader(code int) {
    w.status = code
    w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        sw := &statusWriter{
            ResponseWriter: w,
            status: http.StatusOK, // default
        }

        start := time.Now()
        next.ServeHTTP(sw, r)

        LogJSON(map[string]any{
            "level":       "info",
            "type":        "http_request",
            "client_ip":   ClientIP(r),
            "method":      r.Method,
            "path":        r.URL.Path,
            "status":      sw.status,
            "duration_ms": time.Since(start).String(),
        })
    })
}

func ClientIP(r *http.Request) string {
    // X-Forwarded-For: client, proxy1, proxy2
    if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
        parts := strings.Split(xff, ",")
        if len(parts) > 0 {
            return strings.TrimSpace(parts[0])
        }
    }

    // X-Real-IP (e.g., from Nginx)
    if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
        return xrip
    }

    // RemoteAddr fallback
    host, _, err := net.SplitHostPort(r.RemoteAddr)
    if err == nil {
        return host
    }

    return r.RemoteAddr
}
