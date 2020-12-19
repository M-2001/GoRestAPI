package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//ConnectString es el string de conexio a la DB
var ConnectString = "root:admin@tcp(localhost:3306)/castlem"

//ConnectDB servira para abrir coneccion con la base  de datos en mariadb
func ConnectDB() *sql.DB {
	dbConnect, err := sql.Open("mysql", ConnectString)
	if err != nil {
		log.Fatal(err.Error()) //Error Handling
	}
	log.Println("Conexion exitosa con la DB")
	return dbConnect
}
