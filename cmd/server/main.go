package main

import (
	"context"
	"flag"
	"flights/internal/fetcher/radarbox"
	"flights/internal/handlers"
	"flights/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := flag.String("port", "8000", "Port to server will be listening.")
	flag.Parse()

	fHandler := &handlers.FlightHandler{
		FlightFetcher: &radarbox.FlightFetcherRadarbox{},
	}
	router := server.NewRouter(fHandler)

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
