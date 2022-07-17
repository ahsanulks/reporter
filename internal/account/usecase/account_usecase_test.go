package usecase

import (
	"context"
	"daily_report/internal/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeRepo struct{}

func (fr fakeRepo) GetReportBeforeEOD(ctx context.Context) ([]*domain.Account, error) {
	account := []*domain.Account{}
	var err error
	if ctx.Value("need_error") != nil && ctx.Value("need_error").(string) == "read" {
		err = errors.New("can't get the data")
	} else if ctx.Value("need_error") != nil && ctx.Value("need_error").(string) == "write" {
		account = append(account, &domain.Account{
			ID: 1,
		})
		return account, nil
	}
	return account, err
}

func (fr fakeRepo) WriteReportAfterEOD(ctx context.Context, accounts []*domain.Account) error {
	if len(accounts) == 1 {
		return errors.New("can't write data")
	}
	return nil
}

func TestNewAccountUsecase(t *testing.T) {
	repo := new(fakeRepo)
	got := NewAccountUsecase(repo)
	assert.Equal(t, &AccountUsecase{
		repo:      repo,
		threadNum: new(thread),
	}, got)
}

func Test_thread_GetID(t *testing.T) {
	tc := new(thread)
	tests := []struct {
		name string
		want uint16
	}{
		{
			name: "inc to 1",
			want: 1,
		},
		{
			name: "inc to 2",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tc.GetID(); got != tt.want {
				t.Errorf("thread.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}
