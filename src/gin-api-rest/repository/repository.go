package repository

import (
	"fmt"
	"gin-api-rest/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func FindAllByIdContaAndDate(id_conta string, dt_ini time.Time, dt_fim time.Time) ([]models.Autorizacao, error) {
	query := "SELECT * FROM autorizacao WHERE id_conta_operacao = ? AND data_hora_autorizacao BETWEEN ? AND ?"
	// query := "SELECT * FROM autorizacao"
	// query := `
	// SELECT id_autorizacao, id_conta_operacao, data_hora_autorizacao, quantidade_contas,
	// quantidade_total, valor_total, quantidade_nao_efetuado, valor_nao_efetuado,
	// quantidade_parcial, valor_parcial, quantidade_efetuado, valor_efetuado,
	// quantidade_agendado, valor_agendado, situacao,
	// data_hora_criacao, data_hora_atualizacao, data_hora_delecao
	// FROM autorizacao
	// `
	_, err := DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	// rows, err := DB.Query(query)
	rows, err := DB.Query(query, id_conta, dt_ini, dt_fim)

	if err != nil {
		return nil, err
	}

	var autorizacoes []models.Autorizacao

	for rows.Next() {
		var autorizacao models.Autorizacao

		fmt.Println(rows.Columns())

		err := rows.Scan(&autorizacao.IdAutorizacao, &autorizacao.IdContaOperacao, &autorizacao.DataHoraAutorizacao, &autorizacao.QuantidadeContas,
			&autorizacao.QuantidadeTotal, &autorizacao.ValorTotal, &autorizacao.QuantidadeNaoEfetuado, &autorizacao.ValorNaoEfetuado,
			&autorizacao.QuantidadeParcial, &autorizacao.ValorParcial, &autorizacao.QuantidadeEfetuado, &autorizacao.ValorEfetuado,
			&autorizacao.QuantidadeAgendado, &autorizacao.ValorAgendado, &autorizacao.Situacao,
			&autorizacao.DataHoraCriacao, &autorizacao.DataHoraAtualizacao, &autorizacao.DataHoraDelecao)

		// err := rows.Scan(
		// 	&autorizacao.IdAutorizacao, &autorizacao.IdContaOperacao, &autorizacao.DataHoraAutorizacao, &autorizacao.QuantidadeContas,
		// 	&autorizacao.QuantidadeTotal, &autorizacao.ValorTotal, &autorizacao.QuantidadeNaoEfetuado, &autorizacao.ValorNaoEfetuado,
		// 	&autorizacao.QuantidadeParcial, &autorizacao.ValorParcial, &autorizacao.QuantidadeEfetuado, &autorizacao.ValorEfetuado,
		// 	&autorizacao.QuantidadeAgendado, &autorizacao.ValorAgendado, &autorizacao.Situacao,
		// 	&autorizacao.DataHoraCriacao, &autorizacao.DataHoraAtualizacao, &autorizacao.DataHoraDelecao)

		if err != nil {
			return nil, err
		}

		autorizacoes = append(autorizacoes, autorizacao)
	}

	rows.Close()

	return autorizacoes, nil
}

// [id_autorizacao id_conta_operacao data_hora_autorizacao quantidade_contas
// quantidade_total valor_total quantidade_nao_efetuado valor_nao_efetuado
// quantidade_parcial valor_parcial quantidade_efetuado valor_efetuado
// quantidade_agendado valor_agendado situacao
// data_hora_criacao data_hora_atualizacao data_hora_delecao]
