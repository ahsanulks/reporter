package usecase

import (
	"context"
	"daily_report/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountUsecase_CalculateBalance(t *testing.T) {
	type args struct {
		ctx      context.Context
		accounts []*domain.Account
	}
	tests := []struct {
		name            string
		args            args
		initialBalance  uint64
		expectedBalance float64
	}{
		{
			name: "success",
			args: args{context.Background(), []*domain.Account{
				{
					PrevBalance: 100,
				},
			}},
			initialBalance:  123,
			expectedBalance: 111.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := AccountUsecase{
				threadNum: new(thread),
			}
			tt.args.accounts[0].SetBalance(tt.initialBalance)
			au.CalculateBalance(tt.args.ctx, tt.args.accounts)
			assert.Equal(t, tt.expectedBalance, tt.args.accounts[0].AvgBalance)
		})
	}
}
