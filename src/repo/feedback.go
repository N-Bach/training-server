package repo

import (
	rdb "github.com/GoRethink/gorethink"
	"entity"
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
	_, err := rdb.Table(FEEDBACK_TABLE).Insert(feedback).RunWrite(r.Session)
	return err
}