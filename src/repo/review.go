package repo

import (
	"entity"
	reqM "requestModel"
	"errors"
	"controller"

	rdb "github.com/GoRethink/gorethink"
)

var (
	REVIEW_TABLE = "review"
	
)

type ReviewRepoRethink struct {
	Session *rdb.Session
	UserRepo controller.IUserRepo
}

func NewReviewRepoRethink(session *rdb.Session) *ReviewRepoRethink {
	return &ReviewRepoRethink{
		Session: session,
	}
}

func (r *ReviewRepoRethink) Save(review *reqM.RequestReview) error {
	cursor, err :=  rdb.Table(USER_TABLE).
					GetAllByIndex(EMAIL_INDEX, review.For).
					Run(r.Session)
	defer cursor.Close()
	if err != nil {
		return err
	}
	if cursor.IsNil() {
		return errors.New("Reviewed user does not exist")
	}

	var reviewedUser = entity.User{}
	cursor.One(&reviewedUser)

	newReview := entity.NewReview(review)
	// update For field from email to Id
	review.For = reviewedUser.Id

	// add review data to reviewed user and update
	if err = reviewedUser.AddOneReview(newReview); err != nil {
		return err
	}
	_, err = rdb.Table(USER_TABLE).Get(reviewedUser.Id).Update(reviewedUser).RunWrite(r.Session)
	if err != nil {
		return err
	}

	_, err = rdb.Table(REVIEW_TABLE).Insert(newReview).RunWrite(r.Session)
	return err
}