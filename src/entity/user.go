package entity

import (
	"errors"
)

type RequestUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id	string	`gorethink:"id,omitempty" json:"id"`
	Email string `gorethink:"email" json:"email"`
	Password string `gorethink:"password" json:"password"`
	RatingData
}

func (user *User) AddReview(review *Review) (*User, error) {
	if user.Id != review.For {
		return nil, errors.New("Id for user in review does not match")
	}
	
	err := user.AddOneReview(review)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func NewUser(req *RequestUser) (*User, error) {
	// TODO 
	// use regex to check valid email address
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("Empty request")
	}

	return &User{
		Email: req.Email,
		Password: req.Password,
	}, nil	
}