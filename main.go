package main

import (
	"log"
	"net/http"
	"os"

	"github.com/phaicom/maqe-template-styling/endpoints"
	"github.com/phaicom/maqe-template-styling/handlers"
	"github.com/phaicom/maqe-template-styling/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/api/forum", endpoints.ForumEndpoint).Methods(http.MethodGet)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	handler := ghandlers.CORS()(middleware.SecureHeaders(ghandlers.RecoveryHandler()(ghandlers.LoggingHandler(os.Stdout, r))))

	log.Fatal(http.ListenAndServe(":8080", handler))
}
