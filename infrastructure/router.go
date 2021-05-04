package infrastructure

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/clean-architecture/interfaces/controller"
)

var Router *mux.Router

func init() {
	router := mux.NewRouter()
	userController := controller.NewUserController(NewSqlHandler())
	router.HandleFunc("/", userController.GetUserById).Methods("GET")

	Router = router
}