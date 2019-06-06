package api

import (
	"app/api/controllers"
	"net/http"

	log "github.com/sirupsen/logrus"

	db "app/api/database"

	"github.com/gorilla/mux"
)

func Start(env map[string]string) {
	// connect to database
	err := db.Connect(env["DB_DATABASE"], env["DB_URI"])
	if err != nil {
		log.Panic(err)
	}
	db.Migrate()
	defer db.Close()

	// load middleware

	// init router && routes
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/auth/login", controllers.LoginUserHandler).Methods("POST")
	router.HandleFunc("/auth/check", controllers.CheckTokenHandler).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
