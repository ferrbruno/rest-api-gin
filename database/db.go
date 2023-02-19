package database

import (
	"log"

	"github.com/ferrbruno/rest-api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func Connect() {
	connectionStr := "postgres://root:root@localhost/root"

	DB, err = gorm.Open(postgres.Open(connectionStr))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}