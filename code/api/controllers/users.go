package controllers

import (
	db "app/api/database"
	"app/api/helpers/hash"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := db.GetUsers()
	sendStructAsJSON(w, users, http.StatusOK)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ok, user := db.GetUserById(id)

	if !ok {
		apiError := ApiError{
			ErrorMessage: "User not found",
		}
		sendStructAsJSON(w, apiError, http.StatusNotFound)
		return
	}

	sendStructAsJSON(w, user, http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validation
	validationErrors := make(map[string]string)
	if username == "" {
		validationErrors["username"] = "This field is required"
	}

	if email == "" {
		validationErrors["email"] = "This field is required"
	}

	if password == "" {
		validationErrors["password"] = "This field is required"
	}

	if len(validationErrors) > 0 {
		apiError := ApiError{
			ErrorMessage: "Missing required fields",
			Errors:       validationErrors,
		}
		sendStructAsJSON(w, apiError, http.StatusBadRequest)
		return
	}

	ok, _ := db.FindOneByEmailOrUsername(username, email, 0)
	if ok {
		apiError := ApiError{
			ErrorMessage: "There's already an user with this username and/or email",
		}
		sendStructAsJSON(w, apiError, http.StatusBadRequest)
		return
	}

	hash, _ := hash.Generate(password)

	// Create user
	err := db.CreateUser(username, email, hash)
	if err != nil {
		apiError := ApiError{
			ErrorMessage: "Error adding user to database",
		}
		sendStructAsJSON(w, apiError, http.StatusInternalServerError)
		return
	}

	sendJSON(w, []byte(""), http.StatusCreated)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	data := make(map[string]interface{})

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username != "" {
		data["username"] = username
	}
	if email != "" {
		data["email"] = email
	}

	if username != "" || email != "" {
		ok, _ := db.FindOneByEmailOrUsername(username, email, id)
		if ok {
			apiError := ApiError{
				ErrorMessage: "There's already an user with this username and/or email",
			}
			sendStructAsJSON(w, apiError, http.StatusBadRequest)
			return
		}
	}

	if password != "" {
		hash, _ := hash.Generate(password)
		data["password"] = hash
	}

	err := db.UpdateUser(id, data)
	if err != nil {
		apiError := ApiError{
			ErrorMessage: "Error updating user to database",
		}
		sendStructAsJSON(w, apiError, http.StatusInternalServerError)
		return
	}

	sendJSON(w, []byte(""), http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := db.DeleteUser(id)
	if err != nil {
		apiError := ApiError{
			ErrorMessage: "Error deleting user to database",
		}
		sendStructAsJSON(w, apiError, http.StatusInternalServerError)
		return
	}

	sendJSON(w, []byte(""), http.StatusOK)
}
