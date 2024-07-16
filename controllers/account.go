package controllers

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/atm-test-nedozerov/models"
	"github.com/valitovgaziz/atm-test-nedozerov/util"
)

// инициализация глобальных переменных
var (
	accounts = make(map[int]*models.Account)
	AccMutex = sync.Mutex{}
	NextID   = 1
)

// create new account and return account id
func createAccount(ctx *gin.Context) {
	id := NextID
	NextID++
	accounts[id] = &models.Account{ID: id, Balance: 0.0}

	// response
	ctx.JSON(201, gin.H{
		"message":    "Account created",
		"account_id": id,
	})

	// log operation
	util.LogOperation("Crate Account", NextID-1)
}

// deposit to account and return new balance
func depositToAccount(c *gin.Context) {
	// convert id string type param to int
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	var json struct {
		Amount float64 `json:"amount"`
	}
	// get amount from json body
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	// check if account is exists
	account, exists := accounts[accountId]
	if !exists {
		c.JSON(404, gin.H{"error": "account not found"})
		return
	}
	// deposit to account
	if err := account.Deposit(json.Amount); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// response
	c.JSON(200, gin.H{
		"message":     "Deposit successful",
		"new_balance": account.GetBalance(),
	})
}

// withdraw from account and return new balance
func withdrawFromAccount(c *gin.Context) {
	// convert id string type param to int and check error
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	// get amount from json body
	var json struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	// check if account is exists
	account, exists := accounts[accountId]
	if !exists {
		c.JSON(404, gin.H{"error": "account not found"})
		return
	}
	// withdrop from account
	if err := account.Withdraw(json.Amount); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// response
	c.JSON(200, gin.H{
		"message":     "Withdraw successful",
		"new_balance": account.GetBalance(),
	})
}

// get account's balance
func getAccountBalance(c *gin.Context) {
	// convert id string type param to int and check error
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	// check if account is exists
	account, exists := accounts[accountId]
	if !exists {
		c.JSON(404, gin.H{"error": "account not found"})
		return
	}
	// response
	c.JSON(200, gin.H{
		"account_id": accountId,
		"balance":    account.GetBalance(),
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/accounts", createAccount)
	r.POST("/accounts/:id/deposit", depositToAccount)
	r.POST("/accounts/:id/withdraw", withdrawFromAccount)
	r.GET("/accounts/:id/balance", getAccountBalance)
	return r
}
