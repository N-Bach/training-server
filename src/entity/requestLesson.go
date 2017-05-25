package entity

import (
	"time"
	"errors"
)

type RequestLesson struct {
	Date time.Time `gorethink:"date" json:"date"`
	Location string `gorethink:"location" json:"location"`
	Period Period `gorethink:"period" json:"period"`
	Description string `gorethink:"description" json:"description"`
	AuthorId string `gorethink:"authorId" json:"authorId"`
}

func (req *RequestLesson) IsValid() error {
	if req.Location == "" {
		return errors.New("Empty location")
	}

	if req.Description == "" {
		return errors.New("Empty description")
	}

	return nil
}
