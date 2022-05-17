package routes

import (
	"github.com/Natanael-devops/api-numerais-romanos/controllers"
	"github.com/gin-gonic/gin"
)

func CarregaRotas() {
	r := gin.Default()
	r.GET("/", controllers.ApresentaNumeros)
	r.POST("/search", controllers.CriaPalavra)
	r.Run()
}
