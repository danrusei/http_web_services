package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/adding"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/opening"
)

//Handlers holds the dependencies
type Handlers struct {
	lister listing.Service
	adder  adding.Service
	opener opening.Service
}

//NewHandlers is the constructor for Handlers struct
func NewHandlers(l listing.Service, a adding.Service, o opening.Service) *Handlers {
	return &Handlers{
		lister: l,
		adder:  a,
		opener: o,
	}
}

//SetupRoutes define the mux routes
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.logger(h.handleList(h.lister)))
	mux.HandleFunc("/add", h.logger(h.handleAdd(h.adder)))
	mux.HandleFunc("/state", h.logger(h.handleOpen(h.opener)))
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

func (h *Handlers) handleList(s listing.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, err := s.ListItems()
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
		json.NewEncoder(w).Encode("New item has been added.")
	}
}

func (h *Handlers) handleOpen(s opening.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "the id is not a number value", http.StatusBadRequest)
		}
		b, err := strconv.ParseBool(r.FormValue("open"))
		if err != nil {
			http.Error(w, "the open is not a boolean value", http.StatusBadRequest)
		}

		data := opening.OpenRequest{
			ID:     id,
			IsOpen: b,
		}

		err = s.OpenItem(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
