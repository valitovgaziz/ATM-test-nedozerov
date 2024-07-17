package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valitovgaziz/atm-test-nedozerov/controllers"
)

func main() {
	// set up default gin engine
	r := gin.Default()
	// set up routings
	r.POST("/accounts", controllers.CreateAccount)
	r.POST("/accounts/:id/deposit", controllers.DepositToAccount)
	r.POST("/accounts/:id/withdraw", controllers.WithdrawFromAccount)
	r.GET("/accounts/:id/balance", controllers.GetAccountBalance)
	// run gin engine and listen on 8080
	r.Run(":8080")
}
