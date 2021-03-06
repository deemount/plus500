package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/deemount/plus500/api/app"
	"github.com/deemount/plus500/api/middlewares"
)

// Server is a struct
type Server struct {
	App app.App
}

type key int

const (
	requestIDKey key = 0
)

// Initialize is a method
func (server *Server) Initialize() error {

	var err error

	// set new router instance
	server.App.Router = mux.NewRouter()

	// register new routes with matcher for path
	server.App.V1 = server.App.Router.PathPrefix(server.App.API.Path).Subrouter()
	server.App.V1.Use(middlewares.JSON)

	err = server.initializeRoutes()
	if err != nil {
		return err
	}

	return nil

}

// Run calls listen-and-serve and implements logging handler
func (server *Server) Run() error {

	var err error

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	go log.Printf("Plus500 API v%d is ready to listen and serve on port %s", server.App.API.Version, server.App.API.Port)

	srv := &http.Server{
		Addr:         ":" + server.App.API.Port,
		Handler:      tracing(nextRequestID)(logging(logger)(server.App.Router)), //handlers.LoggingHandler(os.Stdout, server.App.Router)
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil

}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
