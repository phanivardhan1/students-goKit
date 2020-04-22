package services

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	setstudent    endpoint.Endpoint
	getstudent    endpoint.Endpoint
	deletestudent endpoint.Endpoint
	getStudents   endpoint.Endpoint
}

func MakeserverEndpoints(s Service) *Endpoints {
	return &Endpoints{
		setstudent:    MakesetstudentEnpoint(s),
		getstudent:    MakegetstudentEnpoint(s),
		deletestudent: MakedeletestudentEnpoint(s),
		getStudents:   MakegetStudentsEnpoint(s),
	}
}

func MakesetstudentEnpoint(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(setstudentrequest)
		fmt.Println("this is set student endpoint request", req)
		students := s.SetStudent(ctx, req.student)
		fmt.Println("this is set student end point result")
		return students, nil
	}

}

func MakegetstudentEnpoint(s Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getstudentrequest)

		stu, err := s.GetStudent(ctx, req.ID)

		return stu, nil
	}
}

func MakedeletestudentEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deletestudentrequest)

		s.DeleteStudent(ctx, req.ID)

		return nil, nil
	}

}
func MakegetStudentsEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		response, _ = s.GetStudents(ctx)

		return response, nil
	}

}

type setstudentrequest struct {
	student Student
}

type setstudentresponse struct {
	students []Student
}

type getstudentrequest struct {
	ID int
}

type getstudentresponse struct {
	student Student
}

type deletestudentrequest struct {
	ID int
}

type getstudentsresponse struct {
	students []Student
}
