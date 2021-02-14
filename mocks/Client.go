// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/v33/github"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateDeployment provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) CreateDeployment(_a0 context.Context, _a1 string, _a2 string, _a3 *github.DeploymentRequest) (*github.Deployment, *github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *github.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.DeploymentRequest) *github.Deployment); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Deployment)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.DeploymentRequest) *github.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.DeploymentRequest) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateDeploymentStatus provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Client) CreateDeploymentStatus(_a0 context.Context, _a1 string, _a2 string, _a3 int64, _a4 *github.DeploymentStatusRequest) (*github.DeploymentStatus, *github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 *github.DeploymentStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, *github.DeploymentStatusRequest) *github.DeploymentStatus); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.DeploymentStatus)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64, *github.DeploymentStatusRequest) *github.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, int64, *github.DeploymentStatusRequest) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
