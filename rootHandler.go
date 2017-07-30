package main

import "net/http"
import "fmt"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("W funkcji rootHandler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Gorilla!\n"))
}
