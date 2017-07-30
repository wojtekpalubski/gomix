package main

import (
	"fmt"
	"os"
)

// pobiera ze środowiska zmienną PORT
func getPort() string {
	zmienna := "PORT"
	portDomyslny := "8080"
	port := os.Getenv(zmienna)
	if len(port) == 0 {
		fmt.Printf("Nie znaleziono zmienej %s, zwracam port domyslny %s\n", zmienna, portDomyslny)
		port = portDomyslny
	}
	return ":" + port
}
