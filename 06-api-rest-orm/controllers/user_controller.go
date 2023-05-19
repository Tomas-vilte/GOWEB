package controllers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers obtiene todos los usuarios y los devuelve en formato JSON.
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.User{}
	db.Database.Find(&users)
	models.SendData(rw, users, http.StatusOK)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	user := getUserById(r)
	models.SendData(rw, user, http.StatusOK)
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnproccesableEntity(rw)
	} else {
		db.Database.Save(&user)
		models.SendData(rw, user, http.StatusOK)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener registro
	var userId int

	user_ant := getUserById(r)

	userId = int(user_ant.Id)

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnproccesableEntity(rw)
	} else {
		user.Id = int64(userId)
		db.Database.Save(&user)
		models.SendData(rw, user, http.StatusOK)
	}
}

/*
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	userPersistence := db.NewUserPersistence()
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		err := userPersistence.Delete(int(user.Id))
		if err != nil {
			fmt.Println(err)
		}
		models.SendData(rw, user, "Usuario eliminado con exito")
	}
}
*/

func getUserById(r *http.Request) models.User {
	// Obtener registro
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	db.Database.First(&user, userId)
	return user
}
