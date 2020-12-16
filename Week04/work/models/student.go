package models

import "fmt"
type Scores struct {
	English int
	Chinese int
	Math int
}

type Student struct {
	Name string  `json:"name" orm:"column(name)"`
    Grade int    `json:"grade" orm:"column(grade)"`
	Scores Scores    `json:"score" orm:"column(score)"`
}

type StudentDao interface {
	Save(stu Student)(err error)
}

type StudentDaoImpl struct {
}

func (s *StudentDaoImpl)Save(stu Student)(err error)  {
	fmt.Print("Student Info ",stu)
	//todo
	fmt.Print("Student Saved ",stu)
	return
}

type ServiceStudentSave struct {
	Student StudentDao
}