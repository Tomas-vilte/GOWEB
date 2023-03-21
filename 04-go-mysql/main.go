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
	user := models.CreateUser("Thomas", "Thomas123", "Thomas@mail.com")
	fmt.Println(user)
	//db.TruncateTable("users")
	db.Close()
	//db.Ping()

}
