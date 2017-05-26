package requestModel

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	invalid_title = "Empty title"
	invalid_rated = "Rating is out of range (from 1 to 5)"
	invalid_for = "Missing subject"
	invalid_self_review = "Cannot write a review for yourself"
)

func TestIsValid_Review_Test1(t *testing.T) {
	req := RequestReview{
		Title: "Test title",
		Description: "Test description",
		Rated: 3,
		ReviewerId: "1",
		For: "2",
	}
	result := req.IsValid()

	assert.Nil(t, result)
}

func TestIsValid_Review_Test2(t *testing.T) {
	req := RequestReview{
		Description: "Test description",
		Rated: 3,
		ReviewerId: "1",
		For: "2",
	}
	result := req.IsValid()

	assert.EqualError(t, result, invalid_title)
}

func TestIsValid_Review_Test3(t *testing.T) {
	req := RequestReview{
		Title: "Test title",
		Description: "Test description",
		ReviewerId: "1",
		For: "2",
	}

	req.Rated = 0
	result := req.IsValid()
	assert.EqualError(t, result, invalid_rated)

	req.Rated = 6
	result = req.IsValid()
	assert.EqualError(t, result, invalid_rated)
	
	req.Rated = 1
	result = req.IsValid()
	assert.Nil(t, result)
	
	req.Rated = 5
	result = req.IsValid()
	assert.Nil(t, result)
	
	req.Rated = 3
	result = req.IsValid()
	assert.Nil(t, result)

}

func TestIsValid_Review_Test4(t *testing.T) {
	req := RequestReview{
		Title: "Test title",
		Description: "Test description",
		Rated: 3,
		ReviewerId: "1",
		For: "",
	}
	result := req.IsValid()

	assert.EqualError(t, result, invalid_for)
}

func TestIsValid_Review_Test5(t *testing.T) {
	req := RequestReview{
		Title: "Test title",
		Description: "Test description",
		Rated: 3,
		ReviewerId: "1",
		For: "1",
	}
	result := req.IsValid()

	assert.EqualError(t, result, invalid_self_review)
}
