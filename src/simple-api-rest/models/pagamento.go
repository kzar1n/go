package models

type Pagamento struct {
	Id               string  `json:"id"`
	Data             string  `json:"data"`
	Tipo             string  `json:"tipo"`
	Valor            float64 `json:"valor"`
	IdPagador        string  `json:"id_pagador"`
	IdContaPagador   string  `json:"id_conta_pagador"`
	IdRecebedor      string  `json:"id_recebedor"`
	BancoRecebedor   string  `json:"banco_recebedor"`
	AgenciaRecebedor string  `json:"agencia_recebedor"`
	ContaRecebedor   string  `json:"conta_recebedor"`
}
