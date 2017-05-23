package controller

import (
	"encoding/json"
	"requestModel"
	"net/http"
)

type IReviewRepo interface {
	Save(review *requestModel.RequestReview) error
}

func (ctrl *Controller) AddReview(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := requestModel.RequestReview{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("cannot parse review request", err).Excute(w)
		return
	}

	err := ctrl.ReviewRepo.Save(&option)
	if err != nil {
		ResponseInteralError("cannot save review", err).Excute(w)
		return 
	}

	ResponseOk("Saved review").Excute(w)
}