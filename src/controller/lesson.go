package controller

import (
	"entity"
	"net/http"
	"encoding/json"
	"util"
	jwt "github.com/dgrijalva/jwt-go"
)

type ILessonRepo interface {
	Save(lesson *entity.Lesson) error
	GetOne(id string) (* entity.Lesson, error)
	AddEnroll(lesson *entity.Lesson, userId string) error
}

func (ctrl *Controller) AddLesson(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.RequestLesson{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("Cannot parse from body", err).Excute(w)
		return
	}
	tokenString, err := util.FromAuthHeader(r)
	if err != nil {
		ResponseBadRequest("Request header do not have token",err).Excute(w)
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &entity.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(util.Secret), nil
    })
	if err != nil {
		ResponseBadRequest("Request do not have token", err).Excute(w)
		return
	}

	claims, ok := token.Claims.(*entity.TokenClaims)
    if !(ok && token.Valid) {
		ResponseBadRequest("Invalid claims", err).Excute(w)
		return
    } 

	lesson,err := entity.NewLesson(&option, claims.Id)
	if err != nil {
		ResponseBadRequest("Cannot create new lesson", err).Excute(w)
		return
	}

	err = ctrl.LessonRepo.Save(lesson)
	if err != nil {
		ResponseInteralError("Cannot save new lesson", err).Excute(w)
		return
	}
	ResponseOk("Lesson created").Excute(w)
}

func (ctrl *Controller) AddLessonEnroll(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.DId{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("Body does not contain id", err).Excute(w)
		return
	}
	
	claims, err := util.ParseTokenWithClaims(r)
	if err != nil {
		ResponseInteralError("Can not parse token", err).Excute(w)
		return
	}
	
	lesson, err := ctrl.LessonRepo.GetOne(option.Id)
	if err !=  nil {
		util.PrintObj(err)
		ResponseInteralError("Lesson does not exist", err).Excute(w)
		return
	}

	if util.Contains(lesson.Enrolled, claims.Id) {
		ResponseInteralError("Already enrolled", err).Excute(w)
		return
	}

	if lesson.AuthorId == claims.Id {
		ResponseInteralError("Author cannot self-enroll", err).Excute(w)
		return
	}

	err = ctrl.LessonRepo.AddEnroll(lesson, claims.Id)
	if err != nil {
		ResponseInteralError("Cannot add enroll", err).Excute(w)
		return
	}

	ResponseOk("Enroll susscessfully").Excute(w)
}
