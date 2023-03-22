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
	user := models.GetUser(2)
	fmt.Println(user)

	// Actualizar registro
	//user.Username = "Joan"
	//user.Password = "Joan123"
	//user.Email = "Joan@gmail.com"
	//user.Save()
	//fmt.Println(models.ListUsers())

	// Eliminar registro
	user.Delete()

	//db.TruncateTable("users")
	db.Close()
	//db.Ping()

}
