package main

import (
	"github.com/valitovgaziz/atm-test-nedozerov/controllers"
)

func main() {
	router := controllers.SetupRouter() // create routing
	router.Run(":8080")                 // start server on port 8080
}
