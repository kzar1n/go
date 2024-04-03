package controllers

import (
	"encoding/json"
	"net/http"
	"simple-rest-api/database"
	"simple-rest-api/models"

	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	var pagamento []models.Pagamento
	database.DB.InnerJoins("Pagador").InnerJoins("Recebedor").InnerJoins("JOIN `contas` `contaRecebedor` ON `Recebedor`.`id_conta` = `contaRecebedor`.`id_conta`").InnerJoins("JOIN `contas` `contaPagador` ON `Pagador`.`id_conta` = `contaPagador`.`id_conta`").Find(&pagamento)
	// database.DB.InnerJoins("Pagador").InnerJoins("Recebedor").Find(&pagamento)
	json.NewEncoder(w).Encode(pagamento)
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, v := range models.Pagamentos {
		if v.Id == id {
			json.NewEncoder(w).Encode(v)
		}
	}
}
