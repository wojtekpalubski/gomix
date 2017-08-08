package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("W funkcji healthHandler")
	w.WriteHeader(http.StatusOK)
}
