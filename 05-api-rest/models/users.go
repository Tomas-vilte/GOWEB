package models

import (
	"apirest/db"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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

// Obtener un registro
func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

// Actualizar registro
func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)

}

// Guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insertar()
	} else {
		user.update()
	}
}

// Eliminar usuario
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
