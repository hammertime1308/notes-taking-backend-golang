package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"notes-taking-backend-golang/pkg/config"
	"notes-taking-backend-golang/pkg/controllers"
	"notes-taking-backend-golang/pkg/repository"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Init("./config/local.json")

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// we want to have filename for log like `file="sample.go:141`
			entirePath := strings.Split(f.File, "/")
			return "", fmt.Sprintf("%s:%d", entirePath[len(entirePath)-1], f.Line)
		},
	})
	// get log level from config
	level, err := logrus.ParseLevel(config.Get().LogLevel)
	// consider info level incase of parse failure
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}

func main() {

	// get repository
	db := repository.New(config.Get())

	// get controller
	controller := controllers.NewController(db)

	// connect to database
	err := db.Connect()
	if err != nil {
		logrus.Fatalf("error connecting database. error = %v", err.Error())
	}
	// register the database connection closure
	defer func() {
		err := db.Close()
		if err != nil {
			logrus.Errorf("error closing database. error = %v", err.Error())
		}
	}()

	r := mux.NewRouter()

	r.Use(panicRecovery)

	// register the routes and handlers
	controller.RegisterRoutes(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Get().ServeOn),
		Handler: r,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		logrus.Info("starting server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server: %v\n", err)
		}
	}()

	<-sigChan
	shutdown(server)
	logrus.Info("server shutdown complete")

}

// gracefull shutdown
func shutdown(server *http.Server) {
	// Create a context with timeout for shutdown
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	logrus.Info("shutting down server")
	err := server.Shutdown(timeoutCtx)
	if err != nil {
		log.Fatalf("error shutting down server: %v\n", err)
	}
}

// panic recovery middleware
func panicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				logrus.Panicf("recovered from panic. stack trace = %v", string(debug.Stack()))
			}
		}()
		h.ServeHTTP(w, r)
	})
}
