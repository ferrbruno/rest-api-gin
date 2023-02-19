package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ferrbruno/rest-api-gin/database"
	"github.com/ferrbruno/rest-api-gin/models"
	"github.com/gin-gonic/gin"
)

func CreateAluno(ctx *gin.Context) {
	var aluno models.Aluno

	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.ValidateAluno(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&aluno)

	ctx.JSON(http.StatusCreated, aluno)
}

func GetAllAlunos(ctx *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	ctx.JSON(http.StatusOK, alunos)
}

func GetAlunoById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Panic(err.Error())
	}

	var aluno models.Aluno

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})

		return
	}

	ctx.JSON(http.StatusOK, aluno)
}

func GetAlunoByCPF(ctx *gin.Context) {
	cpf := ctx.Param("cpf")

	var aluno models.Aluno

	database.DB.Where(&models.Aluno{ CPF: cpf }).First(&aluno)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})

		return
	}

	ctx.JSON(http.StatusOK, aluno)
}

func UpdateAluno(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Panic(err.Error())
	}

	var aluno models.Aluno

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})

		return
	}

	if err := ctx.ShouldBindJSON(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := models.ValidateAluno(&aluno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Model(&aluno).Updates(aluno)

	ctx.JSON(http.StatusOK, aluno)
}

func DeleteAluno(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var aluno models.Aluno

	database.DB.Delete(&aluno, id)

	ctx.Status(http.StatusNoContent)
}
