package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/adding"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"
)

//Handlers holds the dependencies
type Handlers struct {
	lister listing.Service
	adder  adding.Service
}

//NewHandlers is the constructor for Handlers struct
func NewHandlers(l listing.Service, a adding.Service) *Handlers {
	return &Handlers{
		lister: l,
		adder:  a,
	}
}

//SetupRoutes define the mux routes
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.logger(h.handleLists(h.lister)))
	mux.HandleFunc("/", h.logger(h.handleAdd(h.adder)))
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

func (h *Handlers) handleLists(s listing.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, err := s.ListGoods()
		if err != nil {
			log.Printf("could not list the items: %v", err)
		}
		json.NewEncoder(w).Encode(list)
	}
}

func (h *Handlers) handleAdd(s adding.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newItem adding.Item

		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddItem(newItem)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New item added.")
	}
}
