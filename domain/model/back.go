package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct { // GO não tem classes
	Base     `valid:"required"`
	Code     string     `json:"code" valid:"notnull"` // serialização
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

// Método
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

// função
func NewBank(code string, name string) (*Bank, error) { // *Bank é um ponteiro
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil // retorno &back pois estou retornando um endereço na memória e não o objeto em si
}
