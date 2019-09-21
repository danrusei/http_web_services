package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Danr17/items-rest-api/pkg/model"
	"github.com/Danr17/items-rest-api/pkg/storage"
)

//Handlers host the dependencies
type Handlers struct {
	logger *log.Logger
	db     storage.Storage
}

//NewHandlers creates new Handlers struct
func NewHandlers(logger *log.Logger, db storage.Storage) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}

func (h *Handlers) handleLists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.db.ListsGoods()
		if err != nil {
			h.respond(w, r, data, http.StatusNotFound)
		}
		h.respond(w, r, data, http.StatusOK)
	}
}

func (h *Handlers) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var items []model.Item
		err := h.decode(w, r, &items)
		if err != nil {
			h.respond(w, r, items, http.StatusBadRequest)
		}

		data, err := h.db.AddGood(items...)
		if err != nil {
			h.respond(w, r, data, http.StatusBadRequest)
		}

		h.respond(w, r, data, http.StatusOK)
	}
}

func (h *Handlers) handleOpen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			h.respond(w, r, err, http.StatusBadRequest)
		}
		b, err := strconv.ParseBool(r.FormValue("open"))
		if err != nil {
			h.respond(w, r, err, http.StatusBadRequest)
		}
		data, err := h.db.OpenState(id, b)
		if err != nil {
			h.respond(w, r, err, http.StatusBadRequest)
		}
		h.respond(w, r, data, http.StatusOK)
	}
}

func (h *Handlers) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			h.respond(w, r, id, http.StatusBadRequest)
		}
		data, err := h.db.DelGood(id)
		if err != nil {
			h.respond(w, r, err, http.StatusBadRequest)
		}
		h.respond(w, r, data, http.StatusOK)
	}
}

func (h *Handlers) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Could not encode in json", status)
		}
	}
}

func (h *Handlers) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
