// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	domain "TrafficPolice/internal/domain"

	mock "github.com/stretchr/testify/mock"

	tokens "TrafficPolice/internal/tokens"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// ConfirmExpert provides a mock function with given fields: data
func (_m *AuthService) ConfirmExpert(data domain.ConfirmExpert) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmExpert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.ConfirmExpert) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParseAccessToken provides a mock function with given fields: accessToken
func (_m *AuthService) ParseAccessToken(accessToken string) (tokens.TokenInfo, error) {
	ret := _m.Called(accessToken)

	if len(ret) == 0 {
		panic("no return value specified for ParseAccessToken")
	}

	var r0 tokens.TokenInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (tokens.TokenInfo, error)); ok {
		return rf(accessToken)
	}
	if rf, ok := ret.Get(0).(func(string) tokens.TokenInfo); ok {
		r0 = rf(accessToken)
	} else {
		r0 = ret.Get(0).(tokens.TokenInfo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCamera provides a mock function with given fields: info
func (_m *AuthService) RegisterCamera(info domain.RegisterCamera) (string, error) {
	ret := _m.Called(info)

	if len(ret) == 0 {
		panic("no return value specified for RegisterCamera")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.RegisterCamera) (string, error)); ok {
		return rf(info)
	}
	if rf, ok := ret.Get(0).(func(domain.RegisterCamera) string); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.RegisterCamera) error); ok {
		r1 = rf(info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDirectors provides a mock function with given fields: users
func (_m *AuthService) RegisterDirectors(users []domain.UserInfo) error {
	ret := _m.Called(users)

	if len(ret) == 0 {
		panic("no return value specified for RegisterDirectors")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]domain.UserInfo) error); ok {
		r0 = rf(users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterExpert provides a mock function with given fields: input
func (_m *AuthService) RegisterExpert(input domain.UserInfo) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for RegisterExpert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.UserInfo) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignIn provides a mock function with given fields: input
func (_m *AuthService) SignIn(input domain.UserInfo) (domain.Tokens, error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for SignIn")
	}

	var r0 domain.Tokens
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.UserInfo) (domain.Tokens, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(domain.UserInfo) domain.Tokens); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(domain.Tokens)
	}

	if rf, ok := ret.Get(1).(func(domain.UserInfo) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}