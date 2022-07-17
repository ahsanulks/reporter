package usecase

import (
	"context"
	"daily_report/internal/domain"
	"sync"
)

type accountRepo interface {
	GetReportBeforeEOD(context.Context) ([]*domain.Account, error)
	WriteReportAfterEOD(context.Context, []*domain.Account) error
}

type AccountUsecase struct {
	repo      accountRepo
	threadNum *thread
}

type thread struct {
	id uint16
	mu sync.Mutex
}

func NewAccountUsecase(repo accountRepo) *AccountUsecase {
	return &AccountUsecase{
		repo:      repo,
		threadNum: new(thread),
	}
}

// GetID get next thread ID
func (tc *thread) GetID() uint16 {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.id++
	return tc.id
}
