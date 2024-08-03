package main

import (
	"context"
	"flag"
	_ "flights/docs"
	"flights/internal/fetcher/radarbox"
	"flights/internal/handlers"
	"flights/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Flights API
// @version 1.0
// @description This is a sample server for managing flight information.

// @host localhost:8000
func main() {
	port := flag.String("port", "8000", "Port to server will be listening.")
	flag.Parse()

	fHandler := &handlers.FlightHandler{
		FlightsFetcher:    &radarbox.FlightsFetcherRadarbox{},
		FlightInfoFetcher: &radarbox.FlightInfoRadarbox{},
	}

	router := server.NewRouter(fHandler).(*mux.Router)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Print("Server Started")
	log.Printf("Listening on 0.0.0.0:%s", *port)
	<-done

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
