package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func decodegetstudentreq(_ context.Context, r *http.Request) (request interface{}, err error) {
	fmt.Println("this is decode request")
	var req getstudentrequest
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	req = getstudentrequest{ID: id}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	fmt.Println("this is encode response", response.(Student))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodesetstudentrequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	var req setstudentrequest

	json.NewDecoder(r.Body).Decode(&req.student)
	fmt.Println("this is set student decode rerquest", req)

	return req, nil
}

func encodesetStudentsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	fmt.Println("this is encode setstudent response")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response.([]Student))
}
