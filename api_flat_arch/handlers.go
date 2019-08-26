package main

import "net/http"

func handlerLists(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here I'll list all goods"))

}

func handlerAdd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here you'll add new goods"))

}

func handlerModify(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here you'll modify selected good"))

}

func handlerDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here you'll delete selected good"))

}
