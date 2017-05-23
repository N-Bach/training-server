package repo

import (
	rdb "github.com/GoRethink/gorethink"
	reqM "requestModel"
	"entity"
	"errors"
)

type FeedbackRepoRethink struct {
	Session *rdb.Session
}

var (
	FEEDBACK_TABLE = "feedback"
)

func NewFeedbackRepoRethink(session *rdb.Session) *FeedbackRepoRethink {
	return &FeedbackRepoRethink{
		Session: session,
	}
}

func (r *FeedbackRepoRethink) Save(feedback *entity.Feedback) error {
	_, err := rdb.Table(FEEDBACK_TABLE).Insert(object).RunWrite(r.Session)
	return err
}