package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:root@tcp(localhost:3306)/goweb"

var db *sql.DB

// Realiza la conexion 
func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = connection
}


// Cerrar la conexion
func Close() {
	db.Close()
}

// Verificar la conexion
func Ping(){
	if err := db.Ping(); err != nil{
		panic(err)
	}
}

// Crea una tabla
func CreateTable(schema string){
	db.Exec(schema)
}