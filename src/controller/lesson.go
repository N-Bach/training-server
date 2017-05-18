package controller

import (
	"entity"
	"net/http"
	"encoding/json"
)

type ILessonRepo interface {
	Save(lesson *entity.Lesson) error
}

func (ctrl *Controller) AddLesson(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := entity.Lesson{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("Bad Request", err).Excute(w)
	}

	err := ctrl.LessonRepo.Save(&option)
	if err != nil {
		ResponseInteralError("Cannot insert new lesson", err).Excute(w)
	}
	ResponseOk(option).Excute(w)
}