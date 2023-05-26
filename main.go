package main

import (
	"fmt"

	"github.com/Mariano-JR/go-api-rest/database"
	"github.com/Mariano-JR/go-api-rest/routes"
)

func main() {
	database.ConectaComDB()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
