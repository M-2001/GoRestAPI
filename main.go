package main

import (
	"fmt"

	"github.com/M-2001/GoRestAPI/database"
	"github.com/M-2001/GoRestAPI/handler"
)

func main() {

	databaseConnection := database.ConnectDB()
	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
	handler.Handler()

}
