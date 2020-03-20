package services

import (
	"context"
	"fmt"
)

type Service interface {
	SetStudent(ctx context.Context, student Student) error
	GetStudent(ctx context.Context, id int) (Student, error)
	DeleteStudent(ctx context.Context, id int) error
	GetStudents(ctx context.Context) ([]Student, error)
}
type Studentservice struct {
	Students []Student
}

type Student struct {
	ID             int    `json:"Id,omitempty"`
	FirstName      string `json:"firstname,omitempty"`
	LastName       string `json:"lastname,omitempty"`
	Age            int    `json:"age,omitempty"`
	Degree         string `json:"degree,omitempty"`
	Major          string `json:"major,omitempty"`
	Country        string `json:"country,omitempty"`
	Enrollmenttype string `json:"enrillmenttype,omitempty"`
}

func NewStudentservice() *Studentservice {

	var stu []Student

	Student1 := Student{
		ID:             1,
		FirstName:      "phani",
		LastName:       "gurram",
		Age:            25,
		Degree:         "masters",
		Major:          "ACS",
		Country:        "India",
		Enrollmenttype: "full time",
	}
	stu = append(stu, Student1)
	Student2 := Student{
		ID:             2,
		FirstName:      "sindhu",
		LastName:       "gurram",
		Age:            26,
		Degree:         "masters",
		Major:          "ES",
		Country:        "India",
		Enrollmenttype: "full time",
	}
	stu = append(stu, Student2)

	return &Studentservice{
		Students: stu,
	}

}

func (s *Studentservice) SetStudent(ctx context.Context, student Student) error {
	s.Students = append(s.Students, student)
	fmt.Println(s.Students)
	return nil
}

func (s *Studentservice) GetStudent(ctx context.Context, id int) (Student, error) {
	var student Student
	for _, v := range s.Students {

		if v.ID == id {

			fmt.Println(v)
			student = v
		}
	}
	return student, nil
}

func (s *Studentservice) DeleteStudent(ctx context.Context, id int) error {
	fmt.Println("delete student", s.Students)
	return nil

}

func (s *Studentservice) GetStudents(ctx context.Context) ([]Student, error) {
	fmt.Println(s.Students)

	return s.Students, nil
}
