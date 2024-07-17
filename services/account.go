package services

import (
	"fmt"
	"sync"

	"github.com/valitovgaziz/atm-test-nedozerov/models"
	"github.com/valitovgaziz/atm-test-nedozerov/util"
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
	util.LogOperation("create account", id)
}

func IsAccountExist(id int) bool {
	accMutex.RLock()
	_, ok := accounts[id]
	accMutex.RUnlock()
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
	util.LogOperation("deposit", id)
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
	util.LogOperation("withdraw", id)
}

func GetAccountBalance(id int, resultChan chan<- float64, errorChan chan<- error) {
	accMutex.RLock()
	defer accMutex.RUnlock()

	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}

	resultChan <- account.Balance
	util.LogOperation("get balance", id)
}
