package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (a *api) handleLists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := a.listsGoods()
		if err != nil {
			a.respond(w, r, data, http.StatusNotFound)
		}
		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) handleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var items []Item
		err := a.decode(w, r, &items)
		if err != nil {
			a.respond(w, r, items, http.StatusBadRequest)
		}

		data, err := a.addGood(items...)
		if err != nil {
			a.respond(w, r, data, http.StatusBadRequest)
		}

		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) handleOpen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			a.respond(w, r, id, http.StatusBadRequest)
		}
		b, err := strconv.ParseBool(r.FormValue("open"))
		data, err := a.openState(id, b)
		if err != nil {
			a.respond(w, r, data, http.StatusBadRequest)
		}
		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			a.respond(w, r, id, http.StatusBadRequest)
		}
		data, err := a.delGood(id)
		if err != nil {
			a.respond(w, r, err, http.StatusBadRequest)
		}
		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Could not encode in json", status)
		}
	}
}

func (a *api) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
