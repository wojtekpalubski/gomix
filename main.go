package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
