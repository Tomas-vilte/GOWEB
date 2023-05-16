package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:root@tcp(localhost:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()
