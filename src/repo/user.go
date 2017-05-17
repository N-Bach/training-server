package repo

import (
	"entity"
	rdb "github.com/GoRethink/gorethink"
)

type UserRepoRethink struct {
	Session *rdb.Session
}

func NewUserRepoRethink(session *rdb.Session) *UserRepoRethink {
	return &UserRepoRethink{
		Session: session,
	}
}

func (r *UserRepoRethink) GetByEmail(email string) (*entity.User, error) {
	// TODO
	return nil, nil
}

func (r *UserRepoRethink) Save(user *entity.User) error {
	// TODO
	return nil
}