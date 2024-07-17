package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/atm-test-nedozerov/controllers"
)

func main() {
	r := gin.Default()
	r.POST("/accounts", controllers.CreateAccount)
	r.POST("/accounts/:id/deposit", controllers.DepositToAccount)
	r.POST("/accounts/:id/withdraw", controllers.WithdrawFromAccount)
	r.GET("/accounts/:id/balance", controllers.GetAccountBalance)
	r.Run(":8080")                 // start server on port 8080
}
