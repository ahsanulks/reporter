package domain

import "sync"

type Account struct {
	ID   uint64
	Name string
	Age  uint8

	balance     uint64
	PrevBalance uint64
	AvgBalance  float64

	FreeTransfer uint8

	ThreadNum1  uint16
	ThreadNum2a uint16
	ThreadNum2b uint16
	ThreadNum3  uint16

	mu sync.RWMutex
}

// IncBalance will be increment the current balance
func (acc *Account) IncBalance(inc uint64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += inc
}

// balance will get current value of balance
func (acc *Account) Balance() uint64 {
	acc.mu.RLock()
	defer acc.mu.RUnlock()
	return acc.balance
}

// set balance will replace current balance to the new balance
func (acc *Account) SetBalance(balance uint64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance = balance
}
