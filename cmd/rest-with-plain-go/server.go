package main

import (
	"adzo261/backend-with-go/cmd/rest-with-plain-go/handlers"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var port string
	flag.StringVar(&port, "port", ":8080", "Port address for the server")
	flag.Parse()

	logger := log.New(os.Stdout, "API", log.LstdFlags)

	//Create user handler with os.Stdout variant logger
	uh := handlers.NewUsers(logger)

	//Create new serve mux and register this handler
	serverMux := http.NewServeMux()
	serverMux.Handle("/", uh)

	// create a new server
	server := http.Server{
		Addr:         port,              // configure the bind address
		Handler:      uh,                // set the default handler
		ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Printf("Starting server on port %s\n", port)

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Catch sigterm or interupt into channel and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server shutdown gracefully")
}
