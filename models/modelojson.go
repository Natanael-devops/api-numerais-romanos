package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Numero struct {
	gorm.Model
	Text   string `json:"text" validate:"nonzero"`
	Number string `json:"number"`
	Value  int    `json: "value"`
}

func NovoNumero() *Numero {
	return &Numero{}
}

func ValidaNumero(numero *Numero) error {
	if err := validator.Validate(numero); err != nil {
		return err
	}
	return nil
}
