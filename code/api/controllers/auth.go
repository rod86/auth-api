package controllers

import (
	db "app/api/database"
	"app/api/helpers/hash"
	"app/api/helpers/jwttoken"
	"net/http"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validation
	validationErrors := make(map[string]string)
	if username == "" {
		validationErrors["username"] = "This field is required"
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

	ok, user := db.FindOneByEmailOrUsername(username, username, 0)
	if !ok {
		sendUnauthorizedError(w)
		return
	}

	// compare password
	ok = hash.VerifyHash(password, user.Password)
	if !ok {
		sendUnauthorizedError(w)
		return
	}

	authUser := jwttoken.User{user.ID, username}

	token, err := jwttoken.GenerateJWTToken(authUser)
	if err != nil {
		apiError := ApiError{
			ErrorMessage: "Error generating authentication token",
		}
		sendStructAsJSON(w, apiError, http.StatusInternalServerError)
		return
	}

	// generate jwt token
	jwt := make(map[string]string)
	jwt["token"] = token
	sendStructAsJSON(w, jwt, http.StatusOK)
	return
}

func CheckTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	if token == "" {
		apiError := ApiError{
			ErrorMessage: "Missing token field",
		}
		sendStructAsJSON(w, apiError, http.StatusBadRequest)
		return
	}

	ok, user := jwttoken.GetUserFromJWTToken(token)
	if !ok {
		apiError := ApiError{
			ErrorMessage: "Invalid token",
		}
		sendStructAsJSON(w, apiError, http.StatusUnauthorized)
		return
	}

	sendStructAsJSON(w, user, http.StatusOK)
	return
}

func sendUnauthorizedError(w http.ResponseWriter) {
	apiError := ApiError{
		ErrorMessage: "Invalid Credentials",
	}
	sendStructAsJSON(w, apiError, http.StatusUnauthorized)
	return
}
