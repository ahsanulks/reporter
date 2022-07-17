package usecase

import (
	"context"
	"daily_report/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountUsecase_CalculateBenefit(t *testing.T) {
	type args struct {
		ctx            context.Context
		accounts       []*domain.Account
		initialBalance []uint64
	}
	type expected struct {
		balance      uint64
		freeTransfer uint8
		thread2a     bool
		thread2b     bool
	}
	tests := []struct {
		name  string
		args  args
		expec []expected
	}{
		{
			name: "not get benefit",
			args: args{context.Background(), []*domain.Account{
				{
					FreeTransfer: 1,
				},
			}, []uint64{20}},
			expec: []expected{
				{
					balance:      20,
					freeTransfer: 1,
				},
			},
		},
		{
			name: "get benefit",
			args: args{context.Background(), []*domain.Account{
				{
					FreeTransfer: 1,
				},
				{
					FreeTransfer: 2,
				},
			}, []uint64{100, 151}},
			expec: []expected{
				{
					balance:      100,
					freeTransfer: 5,
					thread2a:     true,
				},
				{
					balance:      151 + benefitBalance,
					freeTransfer: 2,
					thread2b:     true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := AccountUsecase{
				threadNum: new(thread),
			}
			for i := range tt.args.accounts {
				tt.args.accounts[i].SetBalance(tt.args.initialBalance[i])
			}
			au.CalculateBenefit(tt.args.ctx, tt.args.accounts)
			for i := range tt.args.accounts {
				account := tt.args.accounts[i]
				assert.Equal(t, tt.expec[i].balance, account.Balance())
				assert.Equal(t, tt.expec[i].freeTransfer, account.FreeTransfer)
				if tt.expec[i].thread2a {
					assert.NotZero(t, account.ThreadNum2a)
				} else if tt.expec[i].thread2b {
					assert.NotZero(t, account.ThreadNum2b)
				}
			}
		})
	}
}
