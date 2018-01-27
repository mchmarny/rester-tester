package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mchmarny/rester-tester/ping"
)

const (
	defaultPort int = 8080
)

var (
	logger = log.New(os.Stdout, "[server] ", log.Lshortfile|log.Ldate|log.Ltime)
)

func main() {

	port := flag.Int("port", defaultPort, "HTTP port for server to listen on [8080]")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", getRoutes)
	router.HandleFunc("/_ah/health", healthCheckHandler)

	ping.LoadRouts(router)

	httpserver := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", *port),
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, router),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		logger.Printf("Starting server on port: %d", *port)
		log.Fatal(httpserver.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	logger.Printf("Shutdown down server on port: %d", *port)
	httpserver.Shutdown(context.Background())
	os.Exit(0)

}

func getRoutes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([...]string{
		"/ping/",
	})
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "ok")
}
