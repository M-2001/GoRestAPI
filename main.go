package main

import (
	"log"

	"github.com/M-2001/GoRestAPI/database"
)

func main() {

	databaseConnection := database.ConnectDB()

	if databaseConnection == nil {
		defer databaseConnection.Close()
		log.Fatal("Error al establecer conexion con la base de datos")
		return
	}
}
