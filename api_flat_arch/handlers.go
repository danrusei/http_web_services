package main

import (
	"encoding/json"
	"net/http"
)

func (a *api) handlerLists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := a.listsGoods()
		if err != nil {
			a.respond(w, r, data, http.StatusNotFound)
		}
		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) handlerAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := Item{}
		a.decode(w, r, item)
		data, err := a.addGood(item)
		if err != nil {
			a.respond(w, r, data, http.StatusBadRequest)
		}
		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) handlerModify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := Item{}
		a.decode(w, r, item)
		data, err := a.modifyGood(item)
		if err != nil {
			a.respond(w, r, data, http.StatusBadRequest)
		}
		a.respond(w, r, data, http.StatusOK)
	}
}

func (a *api) handlerDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := Item{}
		a.decode(w, r, item)
		data, err := a.delGood(item)
		if err != nil {
			a.respond(w, r, data, http.StatusBadRequest)
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
