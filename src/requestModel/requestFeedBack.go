package requestModel

import (
	"errors"
)

type RequestFeedback struct {
	Id string `gorethink:"id,omitempty" json:"id"`
	AuthorId string `gorethink:"authorId" json:"authorId"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
}

func (req *RequestFeedback) IsValid() error {
	if req.AuthorId == "" {
		return errors.New("Missing sender ID")
	}

	if req.Title == "" {
		return errors.New("Empty title")
	}

	return nil
}