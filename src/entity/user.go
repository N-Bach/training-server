package entity

type RequestUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id	string	`gorethink:"id,omitempty" json:"id"`
	Email string `gorethink:"email" json:"email"`
	Password string `gorethink:"password" json:"password"`
}