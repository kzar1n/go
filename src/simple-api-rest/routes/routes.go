package routes

import (
	"log"
	"net/http"
	"simple-rest-api/controllers"
	"simple-rest-api/middleware"

	"github.com/gorilla/mux"
)

const API = "/pagamentos"

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.SetContentType)
	r.HandleFunc(API, controllers.List).Methods("Get")
	r.HandleFunc(API, controllers.Create).Methods("Post")
	r.HandleFunc(API+"/{id}", controllers.Get).Methods("Get")
	r.HandleFunc(API+"/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc(API+"/{id}", controllers.Update).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
