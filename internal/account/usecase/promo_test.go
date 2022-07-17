package usecase

import (
	"context"
	"daily_report/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountUsecase_CalculatePromo(t *testing.T) {
	type args struct {
		ctx            context.Context
		accounts       []*domain.Account
		initialBalance []uint64
	}
	type expected struct {
		balance uint64
		thread3 bool
	}
	tests := []struct {
		name  string
		args  args
		expec []expected
	}{
		{
			name: "not eligible for promo",
			args: args{context.Background(), []*domain.Account{
				{
					ID: 101,
				},
			}, []uint64{20}},
			expec: []expected{
				{
					balance: 20,
				},
			},
		},
		{
			name: "eligible for promo",
			args: args{context.Background(), []*domain.Account{
				{
					ID: 1,
				},
				{
					ID: 100,
				},
			}, []uint64{20, 50}},
			expec: []expected{
				{
					balance: 30,
					thread3: true,
				},
				{
					balance: 60,
					thread3: true,
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
			au.CalculatePromo(tt.args.ctx, tt.args.accounts)
			for i := range tt.args.accounts {
				account := tt.args.accounts[i]
				assert.Equal(t, tt.expec[i].balance, account.Balance())
				if tt.expec[i].thread3 {
					assert.NotZero(t, account.ThreadNum3)
				} else {
					assert.Zero(t, account.ThreadNum3)
				}
			}
		})
	}
}
