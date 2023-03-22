package main

import (
	"dbmysql/db"
	"dbmysql/models"
	"fmt"
)

func main() {
	db.Connect()

	fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	//user := models.CreateUser("Alexis", "Alexis123", "Alexis@mail.com")
	//fmt.Println(user)

	// Obtener varios usuarios
	//users := models.ListUsers()
	//fmt.Println(users)

	// Obtener usuario proporcionado el id
	user := models.GetUser(3)
	fmt.Println(user)

	//db.TruncateTable("users")
	db.Close()
	//db.Ping()

}
