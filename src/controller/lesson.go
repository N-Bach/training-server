package controller

import (
	"entity"
	"net/http"
	"encoding/json"
	"util"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

type ILessonRepo interface {
	Save(lesson *entity.Lesson) error
}

func (ctrl *Controller) AddLesson(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.RequestLesson{}
	if err := decoder.Decode(&option); err != nil {
		util.PrintObj(err)
		ResponseBadRequest("Bad Request", err).Excute(w)
		return
	}
	tokenString, err := util.FromAuthHeader(r)
	token, err := jwt.ParseWithClaims(tokenString, &entity.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("AllYourBase"), nil
    })

	claims, ok := token.Claims.(*entity.TokenClaims)
    if ok && token.Valid {
        fmt.Printf("%v", claims.Id)
    } else {
        fmt.Println(err)
    }

	lesson := entity.NewLesson(&option, claims.Id)
	err = ctrl.LessonRepo.Save(lesson)
	if err != nil {
		util.PrintObj(err)
		ResponseInteralError("Cannot insert new lesson", err).Excute(w)
		return
	}
	ResponseOk("Lesson created").Excute(w)
}

