package models

// BankAccount определяет интерфейс банковского счета
type BankAccount interface {
	// deposite to account balance
	Deposit(amount float64) error
	// withdraw from account balance
	Withdraw(amount float64) error
	// get balance 
	GetBalance() (float64, error)
}
