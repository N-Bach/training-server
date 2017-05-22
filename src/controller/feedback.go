package controller

import (
	"entity"
	"encoding/json"
	"net/http"
	"requestModel"
)

type IFeedbackRepo interface {
	Save(feedback *entity.Feedback) error
}

func (ctrl *Controller) AddFeedback(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := requestModel.RequestFeedback{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("Cannot parse feedback from body", err).Excute(w)
		return  
	}

	feedback := entity.NewFeedBack(&option)
	err := ctrl.FeedbackRepo.Save(feedback)
	if err != nil {
		ResponseInteralError("Cannot save feedback", err).Excute(w)
		return
	}

	ResponseOk("Feedback added").Excute(w)
}