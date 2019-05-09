package routes

import (
	controllers "github.com/gamorvi/restapi2/app/controllers"
	"github.com/gamorvi/restapi2/app/controllers/auth"
	"github.com/gorilla/mux"
)

func ApiRoutes(prefix string, r *mux.Router) {

	s := r.PathPrefix(prefix).Subrouter()

	s.HandleFunc("/login", auth.Login).Methods("POST")
	s.HandleFunc("/users", auth.ValidateMiddleware(controllers.GetUsers)).Methods("GET")
	s.HandleFunc("/users/{id}", auth.ValidateMiddleware(controllers.GetUser)).Methods("GET")
	s.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	s.HandleFunc("/users/{id:[0-9]+}", auth.ValidateMiddleware(controllers.GetUser)).Methods("GET")
	s.HandleFunc("/users/{id:[0-9]+}", auth.ValidateMiddleware(controllers.UpdateUser)).Methods("PUT")
	s.HandleFunc("/users/{id}", auth.ValidateMiddleware(controllers.DeleteUser)).Methods("DELETE")
}
