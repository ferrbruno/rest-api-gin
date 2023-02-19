package router

import (
	"net/http"

	"github.com/ferrbruno/rest-api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {
	router.GET("", controllers.Index)

	alunos := router.Group("/alunos")
	{
		alunos.GET("", controllers.GetAllAlunos)
		alunos.POST("", controllers.CreateAluno)
		alunos.GET("/:id", controllers.GetAlunoById)
		alunos.PUT("/:id", controllers.UpdateAluno)
		alunos.DELETE("/:id", controllers.DeleteAluno)
		alunos.GET("/cpf/:cpf", controllers.GetAlunoByCPF)
	}

	router.NoRoute(NotFound)
}

func NotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", nil)
}