package controller

import (
	"encoding/json"
	"requestModel"
	"net/http"
	"util"
	"entity"
)

type IReviewRepo interface {
	Save(review *entity.Review) error
}

func (ctrl *Controller) AddReview(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	option := requestModel.RequestReview{}
	if err := decoder.Decode(&option); err != nil {
		ResponseBadRequest("cannot parse review request", err).Excute(w)
		return
	}

	claims:= util.GetClaimsFromRequest(r)
	option.ReviewerId = claims["id"].(string)

	review, err := entity.NewReview(&option)
	if err != nil {
		ResponseInteralError("Cannot create new review", err).Excute(w)
		return
	}
	err = ctrl.ReviewRepo.Save(review)
	if err != nil {
		ResponseInteralError("cannot save review", err).Excute(w)
		return 
	}

	ResponseOk("Saved review").Excute(w)
}