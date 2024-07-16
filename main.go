package main

import (
	"github.com/valitovgaziz/atm-test-nedozerov/controllers"
)

func main() {
	router := controllers.SetupRouter() // создание маршрутизатора
	router.Run(":8080")                 // запуск сервера на порту 8080
}
