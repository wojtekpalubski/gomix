package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func aIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("W funkcji aIdHandler")
	vars := mux.Vars(r)

	fmt.Fprintf(w, "/a/id: %v\n", vars["id"])
	// w.Write([]byte("w aIdhandler"))
	w.WriteHeader(http.StatusOK)
}
