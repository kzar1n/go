package controllers

import (
	"encoding/json"
	"net/http"
	"simple-rest-api/database"
	"simple-rest-api/models"

	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	var pagamentos []models.Pagamento
	database.DB.Find(&pagamentos)
	json.NewEncoder(w).Encode(pagamentos)
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pagamento models.Pagamento
	database.DB.First(&pagamento, "id = ?", id)
	json.NewEncoder(w).Encode(pagamento)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var pagamento models.Pagamento
	json.NewDecoder(r.Body).Decode(&pagamento)
	database.DB.Create(&pagamento)
	json.NewEncoder(w).Encode(pagamento)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pagamento models.Pagamento
	database.DB.Delete(&pagamento, "id = ?", id)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var pagamento models.Pagamento
	database.DB.First(&pagamento, "id = ?", id)
	json.NewDecoder(r.Body).Decode(&pagamento)
	database.DB.Save(&pagamento)
	json.NewEncoder(w).Encode(pagamento)
}
