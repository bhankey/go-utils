package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/bhankey/go-utils/pkg/logger"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
)

const RequestIDHeader = "x-request-id"

type ContextKey int

const (
	RequestIDKey ContextKey = iota + 1
)

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func LoggingMiddleware(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		requestID := uuid.NewUUID().String()
		f := func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get(RequestIDHeader) != "" {
				requestID = RequestIDHeader
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, RequestIDKey, requestID)
			r = r.WithContext(ctx)

			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.WithFields(logrus.Fields{
						"err":        err,
						"request_id": requestID,
					},
					).Error(
						"panic",
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			log := log.WithFields(logrus.Fields{
				"status":     wrapped.status,
				"method":     r.Method,
				"path":       r.URL.EscapedPath(),
				"duration":   time.Since(start),
				"request_id": requestID,
			})

			log.Info("request")
		}

		return http.HandlerFunc(f)
	}
}
