package controllers

import (
	"apirest/db"
	"apirest/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers obtiene todos los usuarios y los devuelve en formato JSON.
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// Establecer tipo de contenido de la respuesta.
	rw.Header().Set("Content-Type", "application/json")

	// Conectarse a la base de datos.
	db.Connect()

	// Obtener todos los usuarios.

	userPersistencece := db.NewUserPersistence(&sql.DB{})
	users, err := userPersistencece.GetUsers()
	if err != nil {
		println(err)
	}

	// Cerrar la conexión con la base de datos.
	db.Close()

	// Convertir los usuarios a formato JSON.
	output, _ := json.Marshal(users)

	// Devolver los usuarios en formato JSON.
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	userPersistencece := db.NewUserPersistence(&sql.DB{})
	user, _ := userPersistencece.GetUser(userId)
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		db.Close()
	}
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener registro
	user := models.User{}
	userPersistencece := db.NewUserPersistence(&sql.DB{})
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		userPersistencece.Save(&user)
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Println(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	userPersistencece := db.NewUserPersistence(&sql.DB{})
	user, _ := userPersistencece.GetUser(userId)
	userPersistencece.Delete(userId)
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}
