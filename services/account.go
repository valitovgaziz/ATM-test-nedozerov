package services

import (
	"fmt"
	"sync"

	"github.com/valitovgaziz/atm-test-nedozerov/models"
)

// инициализация глобальных переменных
var (
	accounts = make(map[int]*models.Account)
	accMutex = sync.RWMutex{}
	nextID   = 1
)

// Create account
func CreateAccount(resultChan chan<- models.Account) {
	accMutex.Lock()
	defer accMutex.Unlock()
	id := nextID
	nextID++
	accounts[id] = &models.Account{ID: id, Balance: 0.0}
	resultChan <- *accounts[id]
}

func IsAccountExist(id int) bool {
	accMutex.Lock()
	_, ok := accounts[id]
	accMutex.Unlock()
	return ok
}

func DepositToAccount(id int, amount float64, resultChan chan<- float64, errorChan chan<- error) {
	accMutex.Lock()
	defer accMutex.Unlock()

	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}

	account.Balance += amount
	resultChan <- account.Balance
}

func WithdrawFromAccount(id int, amount float64, resultChan chan<- float64, errorChan chan<- error) {
	accMutex.Lock()
	defer accMutex.Unlock()

	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}

	if account.Balance < amount {
		errorChan <- fmt.Errorf("insufficient funds")
		return
	}

	account.Balance -= amount
	resultChan <- account.Balance
}

func GetAccountBalance(id int, resultChan chan<- float64, errorChan chan<- error) {
	accMutex.Lock()
	defer accMutex.Unlock()

	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}

	resultChan <- account.Balance

}
