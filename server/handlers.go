package main

import (
	"context"

	boilerplate "github.com/ornlu-is/go-grpc-server-basic-boilerplate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// ageUpperLimit is the highest possible age allowed.
	// TIL: the oldest person ever was 122 years old
	ageUpperLimit = 122
)

// Add implements boilerplate.Boiletplate.Add
func (s *server) Add(ctx context.Context, req *boilerplate.AddReq) (*boilerplate.AddResp, error) {
	if req.FirstName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "first name cannot be empty")
	}
	if req.LastName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "last name cannot be empty")
	}
	if req.Age < 0 || req.Age > ageUpperLimit {
		return nil, status.Errorf(codes.InvalidArgument, "invalid age")
	}

	resp, err := Add(req.Id, req.FirstName, req.LastName, req.Age)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Delete implements boilerplate.Boiletplate.Delete
func (s *server) Delete(ctx context.Context, req *boilerplate.DeleteReq) (*boilerplate.DeleteResp, error) {
	resp, err := Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get implements boilerplate.Boilerplate.Get
func (s *server) Get(ctx context.Context, req *boilerplate.GetReq) (*boilerplate.GetResp, error) {
	resp, err := Get(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateAge implements boilerplate.Boilerplate.UpdateAge
func (s *server) UpdateAge(ctx context.Context, req *boilerplate.UpdateAgeReq) (*boilerplate.UpdateAgeResp, error) {
	if req.Age < 0 || req.Age > ageUpperLimit {
		return nil, status.Errorf(codes.InvalidArgument, "invalid age")
	}

	resp, err := UpdateAge(req.Id, req.Age)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
