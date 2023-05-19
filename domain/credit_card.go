package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreditCard struct {
	ID 	 string
	Name 	 string
	Number 	 string
	ExpMonth int32
	ExpYear  int32
	CVV 	 int32
	Balance float64
	Limit 	 float64
	CreatedAt time.Time
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.New().String()
	c.CreatedAt = time.Now()
	return c
}