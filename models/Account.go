package models

import (
	"fmt"

	"github.com/valitovgaziz/atm-test-nedozerov/util"
)

// Account implement the BankAccount interface
type Account struct {
	ID      int
	Balance float64 `json:"balance"`
}

// Deposits money to the account
func (acc *Account) Deposit(amount float64) error {
	acc.Balance += amount
	util.LogOperation("Deposit", acc.ID)
	return nil
}

// Withdraws money from account
func (acc *Account) Withdraw(amount float64) error {
	if amount > acc.Balance {
		return fmt.Errorf("insufficient funds")
	}
	acc.Balance -= amount
	util.LogOperation("Withdraw", acc.ID)
	return nil
}

// GetBalance returns current account balance
func (acc *Account) GetBalance() float64 {
	util.LogOperation("GetBalance", acc.ID)
	return acc.Balance
}
