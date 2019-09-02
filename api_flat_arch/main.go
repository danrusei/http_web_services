package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	listenAddr string
)

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

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	API := newAPI()

	API.PopulateItems()
	fmt.Printf("there are %d of items in database\n", len(API.db.Items))

	server := http.Server{
		Addr:         listenAddr,
		Handler:      API,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server couldn't start %v", err)
	}

}
