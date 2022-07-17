package usecase

import (
	"context"
	"daily_report/internal/domain"
	"sync"
)

const (
	promoBalance   = 10
	promoLimitUser = 100

	maxPromoThread = 8
)

func (au AccountUsecase) CalculatePromo(ctx context.Context, accounts []*domain.Account) {
	wg := new(sync.WaitGroup)
	for i := 0; i < maxPromoThread; i++ {
		wg.Add(1)

		// will add promo asynchronously with maximum 8 go routine
		go func(threadIndex int, wg *sync.WaitGroup) {
			defer wg.Done()

			threadNum := au.threadNum.GetID()
			for i := threadIndex; i < promoLimitUser; i += maxPromoThread {
				if i >= 0 && i < len(accounts) {
					account := accounts[i]
					addPromo(account, threadNum)
				}
			}
		}(i, wg)
	}
	wg.Wait()
}

// addPromo will add promo to eligible account
func addPromo(account *domain.Account, threadID uint16) {
	if account.ID >= 1 && account.ID <= promoLimitUser {
		account.IncBalance(promoBalance)
		account.ThreadNum3 = threadID
	}
}
