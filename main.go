package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/a", rootHandler)
	r.HandleFunc("/a/{id:[0-9]+}", aIdHandler)
	r.HandleFunc("/a/{nazwa:[a-zA-z][a-zA-z0-9]+}", aNazwaHandler)

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		// p will contain a regular expression that is compatible with regular expressions in Perl, Python, and other languages.
		// For example, the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'.
		// p, err := route.GetPathRegexp()
		// if err != nil {
		// 	return err
		// }
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println("Mapowanie:", strings.Join(m, ","), t)
		return nil
	})

	serwerport := getPort()
	fmt.Println("Startuje serwer", serwerport)
	log.Fatal(http.ListenAndServe(serwerport, r))
	fmt.Println("Serwer wystartowal")
}
