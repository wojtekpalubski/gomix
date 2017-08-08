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

func aIdHandlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("W funkcji aIdHandlerPost")
	vars := mux.Vars(r)

	fmt.Fprintf(w, "/a/id: %v\n", vars["id"])
	// w.Write([]byte("w aIdhandler"))
	w.WriteHeader(http.StatusOK)
}

func aNazwaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("W funkcji aNazwaHandler")
	vars := mux.Vars(r)

	fmt.Fprintf(w, "/a/nazwa: %v\n", vars["nazwa"])
	// w.Write([]byte("w aIdhandler"))
	w.WriteHeader(http.StatusOK)
}
