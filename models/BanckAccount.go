package models

// BankAccount определяет интерфейс банковского счета
type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() (float64, error)
}
