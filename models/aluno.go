package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=14,regexp=^\d{3}\.\d{3}\.\d{3}\-\d{2}$"`
	RG   string `json:"rg" validate:"len=12,regexp=^\d{2}\.\d{3}\.\d{3}\-\w{1}$"`
}

func ValidateAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}

	return nil
}
