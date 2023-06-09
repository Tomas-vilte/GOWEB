package controllers

import (
	"database/sql"
	"fmt"
	"gorm/models"
)

type UserPersistence interface {
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	Save(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}

func NewUserPersistence() UserPersistence {
	connection, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)

	}
	fmt.Println("Conexion exitosa")
	db := connection
	return &userPersintence{db}
}

type userPersintence struct {
	db *sql.DB
}

func (persistence *userPersintence) GetUser(id int) (*models.User, error) {
	fmt.Printf("The value is: %p", persistence.db)
	user := &models.User{}
	err := persistence.db.QueryRow("SELECT id, username, password, email FROM users WHERE id=?", id).Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("Usuarios:", user)
	return user, nil

}

func (persistence *userPersintence) GetUsers() ([]*models.User, error) {
	fmt.Printf("The value is: %p", persistence.db)
	rows, err := persistence.db.Query("SELECT id, username, password, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		fmt.Println("Usuarios:", users)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (persistence *userPersintence) Save(user *models.User) error {
	result, err := persistence.db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}
	user.Id = int64(id)
	return nil

}

func (persistence *userPersintence) Update(user *models.User) error {
	_, err := persistence.db.Exec("UPDATE users SET username=?, password=?, email=? WHERE id=?", user.Username, user.Password, user.Email, user.Id)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (persistence *userPersintence) Delete(id int) error {
	_, err := persistence.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
