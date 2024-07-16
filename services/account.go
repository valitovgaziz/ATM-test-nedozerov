package services

import (
	"sync"

	"github.com/valitovgaziz/atm-test-nedozerov/controllers"
	"github.com/valitovgaziz/atm-test-nedozerov/models"
)

// инициализация глобальных переменных
var (
	accounts = make(map[int]*models.Account)
	accMutex = sync.Mutex{}
	nextID   = 1
)

func CreateAccount(account *models.Account) *models.Account {
	accMutex.Lock()
	id := nextID
	nextID++
	accounts[id] = &models.Account{ID: id, Balance: 0.0}
	account = accounts[id]
	accMutex.Unlock()
	return accounts[id]
}

func IsAccountExist(id int) bool {
	accMutex.Lock()
	_, ok := accounts[id]
	accMutex.Unlock()
	return ok
}

func DepositToAccount(id int, amount float64) {
	accMutex.Lock()
	accounts[id].Balance += amount
	controllers.NewBalance <- accounts[id].Balance
	accMutex.Unlock()
}
