package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"
)

//Handlers holds the dependencies
type Handlers struct {
	lister listing.Service
}

//NewHandlers is the constructor for Handlers struct
func NewHandlers(l listing.Service) *Handlers {
	return &Handlers{
		lister: l,
	}
}

//SetupRoutes define the mux routes
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.logger(h.handleLists(h.lister)))
}

//GetServer returns an http.Server
func (h *Handlers) GetServer(listenAddr string) *http.Server {

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	server := http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server
}

func (h *Handlers) handleLists(l listing.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, err := l.ListGoods()
		if err != nil {
			log.Printf("could not list the items: %v", err)
		}
		json.NewEncoder(w).Encode(list)
	}
}
