package repo

import (
	"entity"
	rdb "github.com/GoRethink/gorethink"
	"errors"
)

var (
	USER_TABLE = "user"
	EMAIL_INDEX = "email"
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
	cursor, err := rdb.Table(USER_TABLE).
					GetAllByIndex(EMAIL_INDEX, email).
					Run(r.Session)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	if cursor.IsNil() {
		return nil, nil
	}

	user := entity.User{}
	if err = cursor.One(&user); err !=nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoRethink) Save(user *entity.User) error {
	res, err := rdb.Table(USER_TABLE).
				GetAllByIndex(EMAIL_INDEX, user.Email).
				Run(r.Session)
	defer res.Close()
	if err != nil {
		return err
	}
	if !res.IsNil() {
		return errors.New("User already exist in database")
	}

	_, erro := rdb.Table(USER_TABLE).Insert(user).RunWrite(r.Session)
	if erro != nil {
		panic(err)
		return err
	}

	return nil
}
