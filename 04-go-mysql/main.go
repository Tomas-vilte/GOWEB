package main

import (
	"dbmysql/db"
	"fmt"
)

func main() {
	db.Connect()

	fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")

	db.TruncateTable("users")
	db.Close()
	//db.Ping()

}
