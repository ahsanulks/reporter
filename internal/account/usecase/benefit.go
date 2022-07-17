package usecase

import (
	"context"
	"daily_report/internal/domain"
	"sync"
)

const (
	benefitBalance = 25
)

func (au *AccountUsecase) CalculateBenefit(ctx context.Context, accounts []*domain.Account) {
	wg := new(sync.WaitGroup)
	for i := range accounts {
		wg.Add(1)

		// will calculate benefits asynchronously for each account
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			addBenefit(accounts[i], au.threadNum.GetID())
		}(i, wg)
	}
	wg.Wait()
}

// addBenefit to eligible account
// account who have balance 100 - 150 will be set free transfer to 5
// account who have balance > 150 will be incremented the balance by 25
func addBenefit(account *domain.Account, thradNum uint16) {
	if account.Balance() >= 100 && account.Balance() <= 150 {
		account.ThreadNum2a = thradNum
		account.FreeTransfer = 5
	} else if account.Balance() > 150 {
		account.ThreadNum2b = thradNum
		account.IncBalance(benefitBalance)
	}
}
