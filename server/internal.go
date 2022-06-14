package main

import (
	boilerplate "github.com/ornlu-is/go-grpc-server-basic-boilerplate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add inserts a new entry in the database table.
// Fails if the given ID already exists in the table.
func Add(id int32, firstName, lastName string, age int32) (*boilerplate.AddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}

// Delete deletes an entry from the database table.
// Fails if an attempt to delete an inexisting ID.
func Delete(id int32) (*boilerplate.DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}

// Get returns the information corresponding to a given ID.
// In case nothing is found, an empty response is returned.
func Get(id int32) (*boilerplate.GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}

// Update updates the database entry corresponding to the
// given ID. Fails if the ID does not exist in the table.
func UpdateAge(id int32, age int32) (*boilerplate.UpdateAgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
