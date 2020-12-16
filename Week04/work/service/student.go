package service

import "Go-000/Week04/work/models"

func SaveStudent(param models.Student)(err error){
	a:=models.StudentDaoImpl{}
	err = a.Save(param)
	return
}