package models

import (
	"dbmysql/db"
)

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// Crear usuario e insertar a la bd
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insertar()
	return user
}

// Insertar Registro
func (usuario *User) insertar() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, usuario.Username, usuario.Password, usuario.Email)
	usuario.Id, _ = result.LastInsertId()

}

// Listar todos los registros
func ListUsers() Users {
	sql := "SELECT id, username, password , email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}
