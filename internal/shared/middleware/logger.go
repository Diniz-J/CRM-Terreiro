package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// MIDDLEWARE LOGGER

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("response time: %v", time.Since(start))
	})
}

// AUTHENTICATION MIDDLEWARE

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// CORS MIDDLEWARE

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RATE LIMITING MIDDLEWARE
var requestCounts = make(map[string]int)

// *** ADD SYNC MUTEX
func rateLimitMiddleware(maxRequests int, windowSeconds int) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			requestCounts[ip]++

			if requestCounts[ip] > maxRequests {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			go func() {
				time.Sleep(time.Duration(windowSeconds) * time.Second)
				requestCounts[ip] = 0
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// CUSTOM RESPONSE WRAPPER

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r)
		log.Printf("[METRICS] %s %s - Status: %d", r.Method, r.RequestURI, rw.statusCode)
	})
}

// CHAIN DE MIDDLEWARES

func chain(handlers ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		for i := len(handlers) - 1; i >= 0; i-- {
			final = handlers[i](final)
		}
		return final
	}
}
