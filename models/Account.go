package models

import (
	"fmt"
)

// Account struct, implement the BankAccount interface
type Account struct {
	ID      int
	Balance float64 `json:"balance"`
}

// Deposits money to the account
func (acc *Account) Deposit(amount float64) error {
	// deposit money to the account balance
	acc.Balance += amount
	return nil
}

// Withdraws money from account
func (acc *Account) Withdraw(amount float64) error {
	// check is the amount more than the account balance, if so return error
	if amount > acc.Balance {
		return fmt.Errorf("insufficient funds")
	}
	// withdraw money from the account balance
	acc.Balance -= amount
	return nil
}

// GetBalance returns current account balance
func (acc *Account) GetBalance() float64 {
	// return account balance
	return acc.Balance
}
