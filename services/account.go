package services

import (
	"fmt"
	"sync"

	"github.com/valitovgaziz/atm-test-nedozerov/models"
	"github.com/valitovgaziz/atm-test-nedozerov/util"
)

// init variabels
var (
	accounts = make(map[int]*models.Account)
	accMutex = sync.RWMutex{}
	nextID   = 1
)

// Create account
func CreateAccount(resultChan chan<- models.Account) {
	// lock mutex
	accMutex.Lock()
	defer accMutex.Unlock()
	// create account
	id := nextID
	nextID++
	accounts[id] = &models.Account{ID: id, Balance: 0.0}
	resultChan <- *accounts[id]
	// log operation
	util.LogOperation("create account", id)
}

// ckeck is account exists
func IsAccountExist(id int) bool {
	// lock mutex
	accMutex.RLock()
	// check if account exists
	_, ok := accounts[id]
	// unlock mutex
	accMutex.RUnlock()
	return ok
}

// deposite to account by id
func DepositToAccount(id int, amount float64, resultChan chan<- float64, errorChan chan<- error) {
	// lock mutex and unlock mutex after end function
	accMutex.Lock()
	defer accMutex.Unlock()
	// check is account exists
	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}
	// deposit account balance
	account.Balance += amount
	resultChan <- account.Balance
	// log operation
	util.LogOperation("deposit", id)
}

// withdraw from account by id
func WithdrawFromAccount(id int, amount float64, resultChan chan<- float64, errorChan chan<- error) {
	// lock mutex and unlock mutex after end function
	accMutex.Lock()
	defer accMutex.Unlock()
	// check is account exists by this id
	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}
	// check not negative amount
	if account.Balance < amount {
		errorChan <- fmt.Errorf("insufficient funds")
		return
	}
	// withdraw from the account balance
	account.Balance -= amount
	resultChan <- account.Balance
	// log operation
	util.LogOperation("withdraw", id)
}

// get account balance by id
func GetAccountBalance(id int, resultChan chan<- float64, errorChan chan<- error) {
	// lock mutex and unlock mutex after end function
	accMutex.RLock()
	defer accMutex.RUnlock()
	// check is account exists
	account, exists := accounts[id]
	if !exists {
		errorChan <- fmt.Errorf("account not found")
		return
	}
	// get account balance
	resultChan <- account.Balance
	// log operation
	util.LogOperation("get balance", id)
}
