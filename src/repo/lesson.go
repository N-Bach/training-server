package repo

import (
	rdb "github.com/GoRethink/gorethink"
	"entity"
)

var (
	LESSON_TABLE = "lesson"
)

type LessonRepoRethink struct {
	Session *rdb.Session 
}

func NewLessonRepoRethink(session *rdb.Session) *LessonRepoRethink{
	return &LessonRepoRethink{
		Session: session,
	}
}

func (r *LessonRepoRethink) Save(lesson *entity.Lesson) error {
	_, err := rdb.Table(lesson).Insert(lesson).RunWrite(r.Session)
	if err != nil {
		return err
	}
	return nil
}