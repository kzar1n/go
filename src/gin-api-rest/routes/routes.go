package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

const path = "/minhas-autorizacoes"

func HandleRequests() {
	r := gin.Default()
	r.GET(path, controllers.GetAllByIdContaAndDate)
	r.GET(path+"/:id_autorizacao/totais", controllers.GetSummaryByID)
	r.Run()
}
