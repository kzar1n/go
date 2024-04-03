package main

import (
	"fmt"
	"simple-rest-api/database"
	"simple-rest-api/models"
	"simple-rest-api/routes"
)

func main() {

	database.OpenDB()

	models.Pagamentos = []models.Pagamento{}
	contaPagadora := criaConta("bb878f12-9e0c-4658-a835-59ee21f3eb88", "341", "1500", "005201", "2", "conta_corrente")
	contaRecebedora := criaConta("794c260c-33ae-4c4d-bd7e-40263e90f646", "001", "00001", "0001234567", "X", "conta_digital")
	pagador := criaPessoa("Richard E. Gibbon", "886.001.270-84", "CPF", contaPagadora)
	recebedor := criaPessoa("Susan E. Brey", "635.030.330-09", "CPF", contaRecebedora)

	models.Pagamentos = []models.Pagamento{
		criaPagamento("c3cbef93-b18e-47f3-9580-92a9de16a82e", "2024-12-31", "PIX", 200.0, pagador, recebedor),
		criaPagamento("91f562bd-d1fb-4b5f-ba93-38f3e6fbbd21", "2024-12-31", "TED", 200.0, pagador, recebedor),
		criaPagamento("91f562bd-d1fb-4b5f-ba93-38f3e6fbbd21", "2024-12-31", "TED", 200.0, pagador, recebedor),
	}

	fmt.Println("Iniciando...")
	routes.HandleRequest()
}

func criaConta(id string, banco string, agencia string, conta string, dac string, tipoConta string) models.Conta {
	c := models.Conta{}
	return *c.Create(id, banco, agencia, conta, dac, tipoConta)
}

func criaPessoa(nome string, documento string, tipoDocumento string, conta models.Conta) models.Pessoa {
	c := models.Pessoa{}
	return *c.Create(nome, documento, tipoDocumento, conta)
}

func criaPagamento(id string, data string, tipo string, valor float64, pagador models.Pessoa, recebedor models.Pessoa) models.Pagamento {
	c := models.Pagamento{}
	return *c.Create(id, data, tipo, valor, pagador, recebedor)

}
