package empapp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func createEndpoint(svc EmployeeIface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*Employee)
		v, err := svc.Create(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func getByIdEndpoint(svc EmployeeIface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetByID(request.(int))
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func updateEndpoint(svc EmployeeIface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		v, err := svc.Update(req.id, req.emp)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func deleteEndpoint(svc EmployeeIface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(int)
		err := svc.Delete(req)
		if err != nil {
			return nil, err
		}
		return "ok", nil
	}
}

func listEndpoint(svc EmployeeIface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		res, err := svc.ListEmployees(req.Page, req.Limit)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
