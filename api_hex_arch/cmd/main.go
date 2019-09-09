package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/adding"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/opening"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/storage/memory"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/storage/seed"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/transport/rest"
)

var (
	listenAddr  string
	storageType string
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

	store := new(memory.Storage)

	var lister listing.Service
	var adder adding.Service
	var opener opening.Service
	//var remover removing.Service

	lister = listing.NewService(store)
	adder = adding.NewService(store)
	opener = opening.NewService(store)

	//seed the database
	adder.AddSampleItem(seed.PopulateItems())

	// set up the HTTP server
	h := rest.NewHandlers(lister, adder, opener)
	server := h.GetServer(listenAddr)

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
