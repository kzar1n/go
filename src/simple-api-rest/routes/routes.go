package routes

import (
	"log"
	"net/http"
	"simple-rest-api/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/payments", controllers.List)
	r.HandleFunc("/payments/{id}", controllers.Get)
	log.Fatal(http.ListenAndServe(":8000", r))
}
