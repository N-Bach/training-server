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
	// lesson.CreateAt = time.Now()
	// lesson.UpdateAt = time.Now()
	_, err := rdb.Table(LESSON_TABLE).Insert(lesson).RunWrite(r.Session)
	return err
}