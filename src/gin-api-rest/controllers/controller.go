package controllers

import (
	"errors"
	"gin-api-rest/models"
	"gin-api-rest/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSummaryByID(c *gin.Context) {
	id_autorizacao := c.Param("id_autorizacao")

	if id_autorizacao == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "É obrigatório informar o campo id_autorizacao"})
		return
	}

	var autorizacao models.Autorizacao

	// query := repository.DB.Where(&models.Autorizacao{IdAutorizacao: id_autorizacao})
	// query.First(&autorizacao)

	// if autorizacao.IdAutorizacao == "" {
	// 	c.JSON(204, nil)
	// 	return
	// }

	c.JSON(200, autorizacao)
}

func GetAllByIdContaAndDate(context *gin.Context) {

	id_conta := context.Query("id_conta")

	if id_conta == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "É obrigatório informar o campo id_conta"})
		return
	}

	data_inicial, err := convertData(context.Query("data_inicial"), false)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data inicial inválida"})
		return
	}

	data_final, err := convertData(context.Query("data_final"), true)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data final inválida"})
		return
	}

	autorizacoes, err := repository.FindAllByIdContaAndDate(id_conta, *data_inicial, *data_final)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if len(autorizacoes) == 0 {
		context.JSON(http.StatusNoContent, nil)
		return
	}

	context.JSON(200, autorizacoes)
}

func convertData(data string, fim bool) (*time.Time, error) {

	if data == "" {
		return nil, errors.New("empty data")
	} else {
		dt, err := time.Parse("2006-01-02", data)

		if err != nil {
			return nil, err
		}

		if fim {
			dt = dt.Add((24 * time.Hour) - 1*time.Second)
		}

		return &dt, nil
	}
}
