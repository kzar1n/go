package routes

import (
	"loja/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.ListenAndServe(":8000", nil)
}
