package empapp

import (
	"errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"	

	"github.com/gorilla/mux"
)

type UpdateRequest struct{
	id int
	emp *Employee
}

type ListRequest struct{
	Page int
	Limit int
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	empRequest:= &Employee{}
	if err := json.NewDecoder(r.Body).Decode(&empRequest); err != nil {
		return nil, err
	}
	return empRequest, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idstr, ok:= mux.Vars(r)["id"]
	if !ok{
		return nil, errors.New("Missing id path param in request")
	}
	id, _:= strconv.Atoi(idstr)
	return id, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idstr, ok:= mux.Vars(r)["id"]
	if !ok{
		return nil, errors.New("Missing id path param in request")
	}
	id, _:= strconv.Atoi(idstr)
	empBody:= &Employee{}
	if err := json.NewDecoder(r.Body).Decode(&empBody); err != nil {
		return nil, err
	}
	updRequest := UpdateRequest{}
	updRequest.id = id
	updRequest.emp = empBody
	return updRequest, nil
}

func decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	page, _:= strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	listRequest:= ListRequest{
		Page:page,
		Limit:limit,
	}
	return listRequest, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}