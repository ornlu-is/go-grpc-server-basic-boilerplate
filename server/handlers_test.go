package main

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	boilerplate "github.com/ornlu-is/go-grpc-server-basic-boilerplate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		name              string
		givenAddReq       *boilerplate.AddReq
		expectedAddResp   *boilerplate.AddResp
		expectedErrorCode codes.Code
	}{
		{
			name:              "fails for missing first name",
			givenAddReq:       &boilerplate.AddReq{},
			expectedAddResp:   &boilerplate.AddResp{},
			expectedErrorCode: codes.InvalidArgument,
		},
		{
			name: "fails for missing last name",
			givenAddReq: &boilerplate.AddReq{
				FirstName: "Guy",
			},
			expectedAddResp:   &boilerplate.AddResp{},
			expectedErrorCode: codes.InvalidArgument,
		},
		{
			name: "fails for age below zero",
			givenAddReq: &boilerplate.AddReq{
				FirstName: "Guy",
				LastName:  "Fawkes",
				Age:       -10,
			},
			expectedAddResp:   &boilerplate.AddResp{},
			expectedErrorCode: codes.InvalidArgument,
		},
		{
			name: "fails for age above the oldest person ever",
			givenAddReq: &boilerplate.AddReq{
				FirstName: "Guy",
				LastName:  "Fawkes",
				Age:       123,
			},
			expectedAddResp:   &boilerplate.AddResp{},
			expectedErrorCode: codes.InvalidArgument,
		},
		{
			name: "returns not implemented yet",
			givenAddReq: &boilerplate.AddReq{
				FirstName: "Guy",
				LastName:  "Fawkes",
			},
			expectedAddResp:   &boilerplate.AddResp{},
			expectedErrorCode: codes.Unimplemented,
		},
	} {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			ctx := context.Background()
			srv := server{}

			// when
			resp, err := srv.Add(ctx, tc.givenAddReq)

			// then
			if tc.expectedErrorCode != codes.OK {
				if err == nil {
					tt.Fatalf("expected error but got nil")
				}

				if code := status.Code(err); code != tc.expectedErrorCode {
					tt.Fatalf("expected error code %s but got %s", tc.expectedErrorCode.String(), code.String())
				}
				return
			}

			if err != nil {
				tt.Fatalf("expected no error but got: %s", err.Error())
			}

			if diff := cmp.Diff(tc.expectedAddResp, resp); diff != "" {
				tt.Fatalf("unexpected response, diff: %s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	for _, tc := range []struct {
		name               string
		givenDeleteReq     *boilerplate.DeleteReq
		expectedDeleteResp *boilerplate.DeleteResp
		expectedErrorCode  codes.Code
	}{
		{
			name:               "returns not implemented yet",
			givenDeleteReq:     &boilerplate.DeleteReq{},
			expectedDeleteResp: &boilerplate.DeleteResp{},
			expectedErrorCode:  codes.Unimplemented,
		},
	} {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			ctx := context.Background()
			srv := server{}

			// when
			resp, err := srv.Delete(ctx, tc.givenDeleteReq)

			// then
			if tc.expectedErrorCode != codes.OK {
				if err == nil {
					tt.Fatalf("expected error but got nil")
				}

				if code := status.Code(err); code != tc.expectedErrorCode {
					tt.Fatalf("expected error code %s but got %s", tc.expectedErrorCode.String(), code.String())
				}
				return
			}

			if err != nil {
				tt.Fatalf("expected no error but got: %s", err.Error())
			}

			if diff := cmp.Diff(tc.expectedDeleteResp, resp); diff != "" {
				tt.Fatalf("unexpected response, diff: %s", diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	for _, tc := range []struct {
		name              string
		givenGetReq       *boilerplate.GetReq
		expectedGetResp   *boilerplate.GetResp
		expectedErrorCode codes.Code
	}{
		{
			name:              "returns not implemented yet",
			givenGetReq:       &boilerplate.GetReq{},
			expectedGetResp:   &boilerplate.GetResp{},
			expectedErrorCode: codes.Unimplemented,
		},
	} {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			ctx := context.Background()
			srv := server{}

			// when
			resp, err := srv.Get(ctx, tc.givenGetReq)

			// then
			if tc.expectedErrorCode != codes.OK {
				if err == nil {
					tt.Fatalf("expected error but got nil")
				}

				if code := status.Code(err); code != tc.expectedErrorCode {
					tt.Fatalf("expected error code %s but got %s", tc.expectedErrorCode.String(), code.String())
				}
				return
			}

			if err != nil {
				tt.Fatalf("expected no error but got: %s", err.Error())
			}

			if diff := cmp.Diff(tc.expectedGetResp, resp); diff != "" {
				tt.Fatalf("unexpected response, diff: %s", diff)
			}
		})
	}
}

func TestUpdateAge(t *testing.T) {
	for _, tc := range []struct {
		name                  string
		givenUpdateAgeReq     *boilerplate.UpdateAgeReq
		expectedUpdateAgeResp *boilerplate.UpdateAgeResp
		expectedErrorCode     codes.Code
	}{
		{
			name:                  "returns not implemented yet",
			givenUpdateAgeReq:     &boilerplate.UpdateAgeReq{},
			expectedUpdateAgeResp: &boilerplate.UpdateAgeResp{},
			expectedErrorCode:     codes.Unimplemented,
		},
	} {
		t.Run(tc.name, func(tt *testing.T) {
			// given
			ctx := context.Background()
			srv := server{}

			// when
			resp, err := srv.UpdateAge(ctx, tc.givenUpdateAgeReq)

			// then
			if tc.expectedErrorCode != codes.OK {
				if err == nil {
					tt.Fatalf("expected error but got nil")
				}

				if code := status.Code(err); code != tc.expectedErrorCode {
					tt.Fatalf("expected error code %s but got %s", tc.expectedErrorCode.String(), code.String())
				}
				return
			}

			if err != nil {
				tt.Fatalf("expected no error but got: %s", err.Error())
			}

			if diff := cmp.Diff(tc.expectedUpdateAgeResp, resp); diff != "" {
				tt.Fatalf("unexpected response, diff: %s", diff)
			}
		})
	}
}
