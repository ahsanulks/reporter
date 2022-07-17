package usecase

import (
	"context"
)

// EndOfTheDay will calculate the balance, calculate the benefit and will be applying the promo for every account that eligible
func (au *AccountUsecase) EndOfTheDay(ctx context.Context) error {
	accounts, err := au.repo.GetReportBeforeEOD(ctx)
	if err != nil {
		return err
	}

	au.CalculateBalance(ctx, accounts)
	au.CalculateBenefit(ctx, accounts)
	au.CalculatePromo(ctx, accounts)

	err = au.repo.WriteReportAfterEOD(ctx, accounts)
	return err
}
