package controller

import (
	"encoding/json"
	"github.com/noguchidaisuke/clean-architecture/interfaces/database"
	"github.com/noguchidaisuke/clean-architecture/usecase"
	"net/http"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

// usecaseを使う。controller => interactor
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := 1
	user, err := c.Interactor.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}