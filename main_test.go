package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ferrbruno/rest-api-gin/database"
	"github.com/ferrbruno/rest-api-gin/models"
	"github.com/ferrbruno/rest-api-gin/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	ID int
	aluno models.Aluno
)

func CreateAlunoMock() {
	aluno = models.Aluno{
		Nome: "Aluno Teste",
		RG: "12.345.678-9",
		CPF: "123.456.789-01",
	}

	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	database.DB.Delete(&models.Aluno{}, ID)
}

func SetupTestRoutes() *gin.Engine {
	database.Connect()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.LoadRoutes(r)

	return r
}

func TestGetAllAlunos(t *testing.T) {
	r := SetupTestRoutes()

	CreateAlunoMock()
	defer DeleteAlunoMock()

	req, _ := http.NewRequest(http.MethodGet, "/alunos", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetAlunoById(t *testing.T) {
	assert := assert.New(t)
	r := SetupTestRoutes()

	CreateAlunoMock()
	defer DeleteAlunoMock()

	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(http.StatusOK, response.Code)

	var alunoMock models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	if assert.NotNil(alunoMock) {
		assert.Equal(aluno.Nome, alunoMock.Nome)
		assert.Equal(aluno.CPF, alunoMock.CPF)
		assert.Equal(aluno.RG, alunoMock.RG)
	}
}

func TestUpdateAluno(t *testing.T) {
	assert := assert.New(t)
	r := SetupTestRoutes()

	CreateAlunoMock()
	defer DeleteAlunoMock()

	newAluno := models.Aluno{
		Nome: "Updated Aluno Teste",
		RG: "88.345.678-9",
		CPF: "883.456.789-01",
	}

	alunoJson, _ := json.Marshal(newAluno)

	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(alunoJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(http.StatusOK, response.Code)

	var UpdatedAlunoMock models.Aluno

	json.Unmarshal(response.Body.Bytes(), &UpdatedAlunoMock)

	if assert.NotNil(UpdatedAlunoMock) {
		assert.Equal(newAluno.Nome, UpdatedAlunoMock.Nome)
		assert.Equal(newAluno.CPF, UpdatedAlunoMock.CPF)
		assert.Equal(newAluno.RG, UpdatedAlunoMock.RG)
	}
}

func TestDeleteAluno(t *testing.T) {
	assert := assert.New(t)
	r := SetupTestRoutes()

	CreateAlunoMock()
	defer DeleteAlunoMock()

	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest(http.MethodDelete, path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(http.StatusNoContent, response.Code)
}
