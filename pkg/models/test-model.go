package models

import (
	"github.com/samber/lo"
)

func FilterStudentsUniqueByName(students []Student) []Student{
	uniqueMap := make(map[string]Student)

	lo.ForEach(students, func (student Student, _ int){
		uniqueMap[student.Name] = student
	})

	var uniques [] Student
	for name := range uniqueMap{
		uniques = append(uniques, uniqueMap[name])
	}

	return uniques
}
