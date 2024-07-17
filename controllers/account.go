package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/atm-test-nedozerov/models"
	"github.com/valitovgaziz/atm-test-nedozerov/services"
)

// create new account and return account id
func CreateAccount(ctx *gin.Context) {
	resultChan := make(chan models.Account, 1)
	go services.CreateAccount(resultChan)
	resultValue := <-resultChan
	// response
	ctx.JSON(201, gin.H{
		"message":    "Account created",
		"account_id": &resultValue,
	})
}

// deposit to account and return new balance
func DepositToAccount(c *gin.Context) {
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
	// check for amount not negative, if so return error
	if json.Amount < 0 {
		c.JSON(400, gin.H{"error": "amount must be positive"})
		return
	}
	// check if account is exists
	exists := services.IsAccountExist(accountId)
	if !exists {
		c.JSON(404, gin.H{"error": "account not found"})
		return
	}

	// deposit to account
	resultChan := make(chan float64, 1)
	errorChan := make(chan error, 1)
	defer close(resultChan)
	defer close(errorChan)
	go services.DepositToAccount(accountId, json.Amount, resultChan, errorChan)

	// return new balance or error
	select {
	case newBalance := <-resultChan:
		c.JSON(200, gin.H{
			"message":     "Deposit successful",
			"new_balance": newBalance,
		})
	case err := <-errorChan:
		c.JSON(404, gin.H{"error": err.Error()})
	}
}

// withdraw from account and return new balance
func WithdrawFromAccount(c *gin.Context) {
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
	// check for amount not negative, if so return error
	if json.Amount < 0 {
		c.JSON(400, gin.H{"error": "amount must be positive"})
		return
	}
	// check if account is exists
	exists := services.IsAccountExist(accountId)
	if !exists {
		c.JSON(404, gin.H{"error": "account not found"})
		return
	}
	// withdrop from account
	resultChan := make(chan float64, 1)
	errorChan := make(chan error, 1)
	defer close(resultChan)
	defer close(errorChan)
	go services.WithdrawFromAccount(accountId, json.Amount, resultChan, errorChan)
	// return new balance or error
	select {
	case newBalance := <-resultChan:
		c.JSON(200, gin.H{
			"message":     "Withdraw successful",
			"new_balance": newBalance,
		})
	case err := <-errorChan:
		c.JSON(404, gin.H{"error": err.Error()})

	}
}

// get account's balance
func GetAccountBalance(c *gin.Context) {
	// convert id string type param to int and check error
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	// check if account is exists
	exists := services.IsAccountExist(accountId)
	if !exists {
		c.JSON(404, gin.H{"error": "account not found"})
		return
	}
	// Get balance
	resultChan := make(chan float64, 1)
	errorChan := make(chan error, 1)
	defer close(resultChan)
	defer close(errorChan)
	go services.GetAccountBalance(accountId, resultChan, errorChan)
	// return balance or error
	select {
	case balance := <-resultChan:
		c.JSON(200, gin.H{
			"balance": balance,
		})
	case err := <-errorChan:
		c.JSON(404, gin.H{"error": err.Error()})
	}
}
