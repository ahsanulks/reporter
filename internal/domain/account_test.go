package domain

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount_IncBalance(t *testing.T) {
	account := new(Account)
	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			account.IncBalance(1)
		}(wg)
	}
	wg.Wait()
	assert.Equal(t, uint64(100), account.Balance())
}

func TestAccount_SetBalance(t *testing.T) {
	account := new(Account)
	account.SetBalance(100)
	assert.Equal(t, uint64(100), account.Balance())

	account.SetBalance(5)
	assert.Equal(t, uint64(5), account.Balance())
}

func TestAccount_Balance(t *testing.T) {
	account := &Account{
		balance: 2,
	}
	assert.Equal(t, uint64(2), account.Balance())
}
