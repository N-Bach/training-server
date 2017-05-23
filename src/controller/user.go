package controller

import (
	"encoding/json"
	"net/http"
	"entity"
)

type IUserRepo interface {
	GetByEmail(email string) (*entity.User, error)
	Save(user *entity.RequestUser) error
}

func (ctrl *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.RequestUser{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("Incorrect user request format",err).Excute(w)
		return
	}
	
	err := ctrl.UserRepo.Save(&option)
	if err != nil {
		ResponseInteralError("Cannot save user", err).Excute(w)
		return
	}

	ResponseOk("Saved success").Excute(w)
} 
