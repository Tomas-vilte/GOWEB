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

	users := models.ListUsers()
	fmt.Println(users)

	//db.TruncateTable("users")
	db.Close()
	//db.Ping()

}
