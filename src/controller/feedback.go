package controller

import (
	"encoding/json"
	"net/http"
	"requestModel"
	"entity"
	"util"
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

	claims := util.GetClaimsFromRequest(r)
	option.AuthorId = claims["id"].(string)
	feedback, err := entity.NewFeedBack(&option)
	if err != nil {
		ResponseBadRequest("Cannot create new feedback", err).Excute(w)
		return
	}

	err = ctrl.FeedbackRepo.Save(feedback)
	if err != nil {
		ResponseInteralError("Cannot save feedback", err).Excute(w)
		return
	}

	ResponseOk("Feedback added").Excute(w)
}