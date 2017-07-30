package main

import (
	"encoding/json"
	"net/http"
)

type Hostinfo struct {
	Name string
	Body string
	Time int64
}

func hostinfoHandler(w http.ResponseWriter, r *http.Request) {
	m := Hostinfo{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
