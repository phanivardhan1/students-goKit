package services

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandlers(s Service) http.Handler {
	r := mux.NewRouter()
	e := MakeserverEndpoints(s)
	r.Methods("GET").Path("/getstudent/{id}").Handler(httptransport.NewServer(e.getstudent, decodegetstudentreq, encodeResponse))
	r.Methods("POST").Path("/setstudent").Handler(httptransport.NewServer(e.setstudent, decodesetstudentrequest, encodeResponse))
	return r
}
