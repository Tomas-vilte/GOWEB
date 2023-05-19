package controllers

import (
	"gorm/db"
	"gorm/models"
	"net/http"
)

// GetUsers obtiene todos los usuarios y los devuelve en formato JSON.
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.User{}
	db.Database.Find(&users)
	models.SendData(rw, users, http.StatusOK)
}

/*
func GetUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener registro
	if user, err := getUserByRequest(r); err != nil {
		fmt.Println(err)
	} else if user == nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user, "Usuario encontrado con exito")
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
		err := userPersistence.Save(&user)
		if err != nil {
			fmt.Println(err)
		}
		models.SendData(rw, user, "Usuario creado con exito")
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
		err := userPersistence.Update(&user)
		if err != nil {
			fmt.Println(err)
		}
		models.SendData(rw, user, "Usuario actualizado con exito")
	}
}

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

func getUserByRequest(r *http.Request) (*models.User, error) {
	// Obtener Id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	userPersistence := db.NewUserPersistence()
	if user, err := userPersistence.GetUser(userId); err != nil {
		return nil, err
	} else if user == nil {
		return nil, nil
	} else {
		return user, nil
	}
}
*/
