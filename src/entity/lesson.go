package entity

import (
	"time"
)

type Lesson struct {
	Id string `gorethink:"id,omitempty" json:"id"`
	Date time.Time `gorethink:"date" json:"date"`
	Location string `gorethink:"location" json:"location"`
	Period Period `gorethink:"period" json:"period"`
	Description string `gorethink:"description" json:"description"`
	Enrolled []string `gorethink:"enrolled" json:"enrolled"`
	AuthorId string `gorethink:"authorId,omitempty" json:"authorId"`
	TimeStamp
}

func NewLesson(req *RequestLesson, authId string) *Lesson {
	return &Lesson{
		Date: req.Date,
		Location: req.Location,
		Period: req.Period,
		Description: req.Description,
		Enrolled: []string{},
		AuthorId: authId,
		TimeStamp: TimeStamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}