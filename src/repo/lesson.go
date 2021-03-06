package repo

import (
	rdb "github.com/GoRethink/gorethink"
	"entity"
	"util"
	"errors"
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
	_, err := rdb.Table(LESSON_TABLE).Insert(lesson).RunWrite(r.Session)
	return err
}

func (r *LessonRepoRethink) GetOne(id string) (* entity.Lesson, error) {
	cursor, err := rdb.Table(LESSON_TABLE).Get(id).Run(r.Session)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	if cursor.IsNil() {
		return nil, nil
	}

	lesson := entity.Lesson{}
	if err = cursor.One(&lesson); err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *LessonRepoRethink) AddEnroll(lesson *entity.Lesson, userId string) error {
	if util.Contains(lesson.Enrolled, userId) {
		return errors.New("Already enrolled")
	}

	if lesson.AuthorId == userId {
		return errors.New("Author cannot self-enroll")
	}

	_, err := rdb.Table(LESSON_TABLE).Get(lesson.Id).Update(map[string]interface{}{
    	"enrolled":  append(lesson.Enrolled, userId) ,
	}).RunWrite(r.Session)
	if err != nil {
		return err
	}
	return nil
}
