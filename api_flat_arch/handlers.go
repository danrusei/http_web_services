package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *api) handlerLists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := a.listsGoods()
		if err != nil {
			a.respond(w, r, data, http.StatusNotFound)
		}
		w.Write([]byte("Here I'll list all goods"))

	}
}

func (a *api) handlerAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Here you'll add new goods"))
	}
}

func (a *api) handlerModify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Here you'll modify selected good"))
	}
}

func (a *api) handlerDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("here you'll delete selected good"))
	}
}

func (a *api) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Errorf("could not decode to json %v", err)
		}
	}
}
