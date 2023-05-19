package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID 	 string
	Amount 	 float64
	Status 	 string
	Description string
	Store 	 string
	CreditCardId string
	CreatedAt time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.New().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAdnValidate(creditCard *CreditCard) {
	if t.Amount + creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}