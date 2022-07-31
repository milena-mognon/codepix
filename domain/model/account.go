package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct { // GO não tem classes
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"` // serialização
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"name" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

// Método
func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

// função
func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) { // *Bank é um ponteiro
	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil // retorno &back pois estou retornando um endereço na memória e não o objeto em si
}
