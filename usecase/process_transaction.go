package usecase

import (
	"time"

	"github.com/fabiofa87/codebank/domain"
	"github.com/fabiofa87/codebank/dto"
)

type UseCaseTransaction struct {
	TransactionRepoository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepoository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(transactionDto)
	ccBalanceAndLimit, err := u.TransactionRepoository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.newTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAdnValidate(creditCard)
	
	err = u.TransactionRepoository.SaveTransaction(*t, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	return *t, nil
}

func (UseCaseTransaction) hydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard{
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpMonth = transactionDto.ExpMonth
	creditCard.ExpYear = transactionDto.ExpYear
	creditCard.CVV = transactionDto.CVV
	return creditCard
}
func (UseCaseTransaction) newTransaction(transactionDto dto.Transaction, cc domain.CreditCard) *domain.Transaction{
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transactionDto.Amount
	t.Store = transactionDto.Store
	t.Description = transactionDto.Description
	t.CreatedAt = time.Now()
	return t
	
}