// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TransportRepo is an autogenerated mock type for the TransportRepo type
type TransportRepo struct {
	mock.Mock
}

// GetTransportID provides a mock function with given fields: chars, num, region
func (_m *TransportRepo) GetTransportID(chars string, num string, region string) (string, error) {
	ret := _m.Called(chars, num, region)

	if len(ret) == 0 {
		panic("no return value specified for GetTransportID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(chars, num, region)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(chars, num, region)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(chars, num, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransportRepo creates a new instance of TransportRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransportRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransportRepo {
	mock := &TransportRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
