package apimiddleware

import (
	"bytes"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func RequestResponseLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var requestID string
		if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
			requestID = reqID.(string)
		}

		log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("_request_id", requestID)
		})

		var requestBody bytes.Buffer
		if r.Body != nil {
			body, _ := io.ReadAll(r.Body)
			requestBody.Write(body)
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		rw := &responseWriter{ResponseWriter: w}

		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("params", r.URL.Query().Encode()).
			//Str("request_body", requestBody.String()).
			Msg("Request received")

		next.ServeHTTP(rw, r)

		elapsed := time.Since(start)

		logEvent := log.Info()
		if rw.status >= 400 {
			logEvent = log.Error()
		}

		logEvent.
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", rw.status).
			Int("response_size_b", rw.size).
			Float64("response_time_µs", float64(elapsed.Microseconds())).
			Msg("Response sent")
	})
}
