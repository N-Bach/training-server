package controller

import (
	"entity"
	"util"
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func (ctrl *Controller) ValidateUser(u entity.User, ru entity.RequestUser) bool {
	return u.Email == ru.Email && u.Password == ru.Password
}

func (ctrl *Controller) CreateToken(user *entity.User) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,

	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(util.Secret)

	// fmt.Println(tokenString, err)
	return tokenString, err
}

func (ctrl *Controller) UserSignin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.RequestUser{}
	if err := decoder.Decode(&option); err != nil {
		panic(err)
	}

	user,err := ctrl.UserRepo.GetByEmail(option.Email)
	if err != nil {
		ResponseInteralError("Something went wrong", err).Excute(w)
		return
	}

	if ctrl.ValidateUser(*user, option) {
		token, e := ctrl.CreateToken(user)
		if e != nil {
			ResponseInteralError("cannot create token", e).Excute(w)
			return
		}
		ResponseOk(token).Excute(w)
		return

	}
	
	ResponseOk("Incorrect Password").Excute(w)
	return
	
}
