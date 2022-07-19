package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Danr17/http_web_services/api_layered_arch/pkg/handlers"
	"github.com/Danr17/http_web_services/api_layered_arch/pkg/storage"
	"github.com/Danr17/http_web_services/api_layered_arch/pkg/storage/dbmemory"
)

var (
	listenAddr string
	dbType     string
)

// api holds dependencies
type api struct {
	mutex  sync.Mutex
	db     storage.Storage
	router *http.ServeMux
	logger *log.Logger
}

func newAPI() *api {
	a := &api{
		router: http.NewServeMux(),
		mutex:  sync.Mutex{},
	}
	return a
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.StringVar(&dbType, "db-type", "memory", "select database, memory or firestore")
	flag.Parse()

	API := newAPI()

	API.logger = log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	switch dbType {
	case "memory":
		API.db = dbmemory.NewMemory()
	default:
		API.db = dbmemory.NewMemory()
	}

	h := handlers.NewHandlers(API.logger, API.db)

	mux := API.router
	h.CreateRoutes(mux)

	storage.PopulateItems(API.db)

	server := http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	//channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("main : API listening on %s", listenAddr)
		serverErrors <- server.ListenAndServe()
	}()

	//channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	//blocking run and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting server: %s", err)

	case <-shutdown:
		log.Println("main : Start shutdown")

		//give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// asking listener to shutdown
		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = server.Close()
		}

		if err != nil {
			return fmt.Errorf("main : could not stop server gracefully : %v", err)
		}
	}

	return nil
}
