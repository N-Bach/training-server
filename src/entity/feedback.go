package entity

import (
	"time"
	"requestModel"
)

type Feedback struct {
	Id string `gorethink:"id,omitempty" json:"id"`
	AuthorId string `gorethink:"authorId" json:"authorId"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
	CreatedAt time.Time `gorethink:"createdAt" json:"createdAt"`
}

func NewFeedBack(req *requestModel.RequestFeedback) (*Feedback, error) {

	if err :=  req.IsValid(); err != nil {
		return nil, err
	}

	return &Feedback{
		AuthorId: req.AuthorId,
		Title: req.Title,
		Description: req.Description,
		CreatedAt: time.Now(),
	}, nil
}