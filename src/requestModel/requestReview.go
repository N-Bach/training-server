package requestModel

import (
	"errors"
) 

type RequestReview struct {
	ReviewerId string `gorethink:"reviewerid" json:"reviewerid"`
	Title string `gorethink:"title" json:"title"`
	Description string `gorethink:"description" json:"description"`
	Rated int `gorethink:"rated" json:"rated"`
	For string `gorethink:"for" json:"for"`
}

func (req *RequestReview) IsValid() error {

	if req.Title == "" {
		return errors.New("Empty title") 
	}

	if req.Rated > 5 || req.Rated < 1 {
		return errors.New("Rating is out of range (from 1 to 5)")
	}

	if req.For == "" {
		return errors.New("Missing subject")
	}

	return nil
}
