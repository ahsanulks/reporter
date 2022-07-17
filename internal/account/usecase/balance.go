package usecase

import (
	"context"
	"daily_report/internal/domain"
	"sync"
)

func (au AccountUsecase) CalculateBalance(ctx context.Context, accounts []*domain.Account) {
	wg := new(sync.WaitGroup)
	for i := range accounts {
		wg.Add(1)

		// will calculate average balance asynchronously for each account
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			calAvgBalance(accounts[i], au.threadNum.GetID())
		}(i, wg)
	}
	wg.Wait()
}

// calAvgBalance will calculate avarage balance of account
// average balance = average(balance + previous balance)
func calAvgBalance(account *domain.Account, threadNum uint16) {
	account.ThreadNum1 = threadNum
	account.AvgBalance = float64(account.Balance()+account.PrevBalance) / 2
}
