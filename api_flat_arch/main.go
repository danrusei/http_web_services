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
	mutex  *sync.RWMutex
	db     *Memory
	router *http.ServeMux
}

func newAPI() *api {
	a := &api{
		router: http.NewServeMux(),
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

	db := new(Memory)
	db.PopulateItems()
	fmt.Printf("there are %d of items in database", len(db.Items))

	API = &api{
		db: db,
	}

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
