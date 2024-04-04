package models

import "time"

type Autorizacao struct {
	IdAutorizacao         string    `json:"id_autorizacao"`
	IdContaOperacao       string    `json:"id_conta_operacao"`
	DataHoraAutorizacao   time.Time `json:"data_hora_autorizacao"`
	QuantidadeContas      int       `json:"quantidade_contas"`
	QuantidadeTotal       int       `json:"quantidade_total"`
	ValorTotal            float64   `json:"valor_total"`
	QuantidadeNaoEfetuado int       `json:"quantidade_nao_efetuado"`
	ValorNaoEfetuado      float64   `json:"valor_nao_efetuado"`
	QuantidadeParcial     int       `json:"quantidade_parcial"`
	ValorParcial          float64   `json:"valor_parcial"`
	QuantidadeEfetuado    int       `json:"quantidade_efetuado"`
	ValorEfetuado         float64   `json:"valor_efetuado"`
	QuantidadeAgendado    int       `json:"quantidade_agendado"`
	ValorAgendado         float64   `json:"valor_agendado"`
	Situacao              string    `json:"situacao"`
	DataHoraCriacao       time.Time `json:"data_hora_criacao"`
	DataHoraAtualizacao   time.Time `json:"data_hora_atualizacao"`
	DataHoraDelecao       time.Time `json:"data_hora_delecao"`
}

func (Autorizacao) TableName() string {
	return "autorizacao"
}
