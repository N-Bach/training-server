package entity

import (
	"errors"
)

type RatingData struct {
	ReviewsNum int `gorethink:"reviewsNum" json:"reviewsNum"`
	AvrRating float64 `gorethink:"avrRating" json:"avrRating"`
}

func (data *RatingData) AddOneReview(review *Review) error{
	if review.Rated > 5 || review.Rated < 1 {
		return errors.New("Invalid rating, must be between 1 and 5")
	}
	data.AvrRating = ((data.AvrRating * float64(data.ReviewsNum)) + float64(review.Rated))  / float64(data.ReviewsNum +1)
	data.ReviewsNum++
	return nil
}