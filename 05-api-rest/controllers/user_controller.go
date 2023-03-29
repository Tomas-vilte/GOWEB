package controllers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetUsers obtiene todos los usuarios y los devuelve en formato JSON.
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	userPersistence := db.NewUserPersistence()

	if users, err := userPersistence.GetUsers(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	userPersistence := db.NewUserPersistence()

	if err := decoder.Decode(&user); err != nil {
		models.SendUnproccesableEntity(rw)
	} else {
		userPersistence.Save(&user)
		models.SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	// Obtener registro
	var userId int64
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	userPersistence := db.NewUserPersistence()
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnproccesableEntity(rw)
	} else {
		user.Id = userId
		userPersistence.Save(&user)
		models.SendData(rw, user)
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	userPersistence := db.NewUserPersistence()
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		userPersistence.Delete(int(user.Id))
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	// Obtener Id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	userPersistence := db.NewUserPersistence()
	if user, err := userPersistence.GetUser(userId); err != nil {
		return models.User{}, err
	} else if user != nil {
		return *user, nil
	} else {
		return models.User{}, fmt.Errorf("Usuario no encontrado %d", userId)
	}
}
