package controller

import (
	"entity"
	"net/http"
	"encoding/json"
	"util"
)

type ILessonRepo interface {
	Save(lesson *entity.Lesson) error
	GetOne(id string) (*entity.Lesson, error)
	AddEnroll(lesson *entity.Lesson, userId string) error
}

func (ctrl *Controller) AddLesson(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.RequestLesson{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("Cannot parse from body", err).Excute(w)
		return
	}
	
	claims := util.GetClaimsFromRequest(r)
	lesson,err := entity.NewLesson(&option, claims["id"].(string))
	// lesson,err := entity.NewLesson(&option, claims.Id)
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
		ResponseBadRequest("Cannot parse from body", err).Excute(w)
		return
	}
	
	claims, err := util.ParseTokenWithClaims(r)
	if err != nil {
		ResponseBadRequest("Can not parse token", err).Excute(w)
		return
	}
	
	lesson, err := ctrl.LessonRepo.GetOne(option.Id)
	if err !=  nil {
		ResponseInteralError("Lesson does not exist", err).Excute(w)
		return
	}

	err = ctrl.LessonRepo.AddEnroll(lesson, claims.Id)
	if err != nil {
		ResponseInteralError("Cannot add enroll", err).Excute(w)
		return
	}

	ResponseOk("Enroll susscessfully").Excute(w)
}
