package controller

import (
	"encoding/json"
	"net/http"

	"entity"
	"util"
)

type IUserRepo interface {
	GetByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
}

func (ctrl *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.RequestUser{}
	if err := decoder.Decode(&option); err != nil {
		panic(err)
	}

	util.PrintStr("Body received")
	util.PrintObj(option)

	

	json.NewEncoder(w).Encode(option.Email)
} 