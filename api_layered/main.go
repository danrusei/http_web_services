package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/Danr17/http_web_services/api_layered/routes"
)

var (
	listenAddr string
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	mux := routes.NewAPI()

	server := http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server couldn't start %v", err)
	}

}
