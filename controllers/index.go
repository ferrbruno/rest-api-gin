package controllers

import (
	"net/http"

	"github.com/ferrbruno/rest-api-gin/database"
	"github.com/ferrbruno/rest-api-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}
