package main

import (
	"loan/internal/app"
	"loan/internal/pkg/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()

	r := mux.NewRouter()
	app.Router(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
