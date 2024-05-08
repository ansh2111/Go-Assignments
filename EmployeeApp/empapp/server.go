package empapp

import (
    "context"
    "net/http"

    "github.com/hashicorp/go-memdb"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context) http.Handler {
    	r := mux.NewRouter()
	
	// Create a new data base
	db, err := memdb.NewMemDB(GetEmployeeSchema())
	if err != nil {
		panic(err)
	}
    	svc := NewEmployeeSvc(db)
	
	createHandler := httptransport.NewServer(
		createEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)

    	r.Methods("POST").Path("/v1/employees").Handler(createHandler)

	getByIDHandler := httptransport.NewServer(
		getByIdEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)

    	r.Methods("GET").Path("/v1/employees/{id}").Handler(getByIDHandler)

	updateHandler := httptransport.NewServer(
		updateEndpoint(svc),
		decodeUpdateRequest,
		encodeResponse,
	)

    	r.Methods("PUT").Path("/v1/employees/{id}").Handler(updateHandler)

	deleteHandler := httptransport.NewServer(
		deleteEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)

    	r.Methods("DELETE").Path("/v1/employees/{id}").Handler(deleteHandler)

	listHandler := httptransport.NewServer(
		listEndpoint(svc),
		decodeListRequest,
		encodeResponse,
	)

    	r.Methods("GET").Path("/v1/employees").Handler(listHandler)

    	return r
}
