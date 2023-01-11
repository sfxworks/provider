// Code generated by mockery 2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	v1beta2 "github.com/akash-network/node/x/deployment/types/v1beta2"

	v2beta1 "github.com/akash-network/node/manifest/v2beta1"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// IsActive provides a mock function with given fields: _a0, _a1
func (_m *Client) IsActive(_a0 context.Context, _a1 v1beta2.DeploymentID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, v1beta2.DeploymentID) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1beta2.DeploymentID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Submit provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) Submit(_a0 context.Context, _a1 v1beta2.DeploymentID, _a2 v2beta1.Manifest) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1beta2.DeploymentID, v2beta1.Manifest) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t testing.TB) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
