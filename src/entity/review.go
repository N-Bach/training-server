package entity

import (
	"requestModel"
	"time"
)

type Review struct {
	Id string `gorethink:"id,omitempty" json:"id"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
	Rated int `gorethink:"rated" json:"rated"`
	For string `gorethink:"for" json:"for"`
	TimeStamp
	ReviewerId string `gorethink:"reviewerid" json:"reviewerid"`
}

func NewReview(req *requestModel.RequestReview) *Review {
	return &Review {
		ReviewerId: req.ReviewerId,
		Title: req.Title,
		Description: req.Description,
		Rated: req.Rated,
		For: req.For,
		TimeStamp: TimeStamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}