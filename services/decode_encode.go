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
	// if e, ok := response.(error); ok && e.error() != nil {
	// 	// Not a Go kit transport error, but a business-logic error.
	// 	// Provide those as HTTP errors.
	// 	encodeError(ctx, e.error(), w)
	// 	return nil
	// }
	fmt.Println("this is encode response", response.(Student))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodesetstudentrequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	fmt.Println("this is set student decode rerquest")

	var req setstudentrequest

	json.NewDecoder(r.Body).Decode(&req)

	return req, nil
}
