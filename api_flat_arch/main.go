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
)

var (
	listenAddr string
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	API := newAPI()

	API.PopulateItems()
	fmt.Printf("there are %d of items in database\n", len(API.db.Items))
	fmt.Printf("the format of a date is %v", API.db.Items[1].Created)

	server := http.Server{
		Addr:         listenAddr,
		Handler:      API,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
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

// api holds dependencies
type api struct {
	mutex  sync.Mutex
	db     Memory
	router *http.ServeMux
}

func newAPI() *api {
	a := &api{
		router: http.NewServeMux(),
		db:     Memory{},
		mutex:  sync.Mutex{},
	}
	a.routes()
	return a
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
