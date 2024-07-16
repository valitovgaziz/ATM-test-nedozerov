package models

import (
	"fmt"
	"sync"

	"github.com/valitovgaziz/atm-test-nedozerov/util"
)

// Account реализует интерфейс BankAccount
type Account struct {
	ID      int
	Balance float64
	Lock    sync.Mutex
}

// Deposit позволяет вносить средства на баланс
func (acc *Account) Deposit(amount float64) error {
	acc.Balance += amount
	util.LogOperation("Deposit", acc.ID)
	return nil
}

// Withdraw позволяет снимать средства с баланса
func (acc *Account) Withdraw(amount float64) error {
	if amount > acc.Balance {
		return fmt.Errorf("insufficient funds")
	}
	acc.Balance -= amount
	util.LogOperation("Withdraw", acc.ID)
	return nil
}

// GetBalance возвращает текущий баланс
func (acc *Account) GetBalance() float64 {
	util.LogOperation("GetBalance", acc.ID)
	return acc.Balance
}
