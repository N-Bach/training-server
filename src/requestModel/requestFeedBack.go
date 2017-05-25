package requestModel

import (
	"errors"
)

type RequestFeedback struct {
	Id string `gorethink:"id,omitempty" json:"id"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
	AuthorId string `gorethink:"description" json:"description"`
}

func (req *RequestFeedback) IsValid() error {
	if req.Title == "" {
		return errors.New("Empty title")
	}

	return nil
}