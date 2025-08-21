package main

import (
	"log"
	"net/http"

	"cmp-firebird/api"
	"cmp-firebird/config"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", api.ApiHealthEndpoint)
	mux.HandleFunc("POST /json/query", api.ApiJsonQueryEndpoint)

	servidor := config.Servidor()
	log.Println("Iniciando servidor en ", servidor)
	log.Fatal(http.ListenAndServe(servidor, mux))
}
