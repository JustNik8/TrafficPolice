// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	domain "TrafficPolice/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// AuthRepo is an autogenerated mock type for the AuthRepo type
type AuthRepo struct {
	mock.Mock
}

// CheckUserExists provides a mock function with given fields: username
func (_m *AuthRepo) CheckUserExists(username string) bool {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserExists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ConfirmExpert provides a mock function with given fields: data
func (_m *AuthRepo) ConfirmExpert(data domain.ConfirmExpert) error {
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

// InsertCamera provides a mock function with given fields: camera, userID
func (_m *AuthRepo) InsertCamera(camera domain.Camera, userID uuid.UUID) (string, error) {
	ret := _m.Called(camera, userID)

	if len(ret) == 0 {
		panic("no return value specified for InsertCamera")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Camera, uuid.UUID) (string, error)); ok {
		return rf(camera, userID)
	}
	if rf, ok := ret.Get(0).(func(domain.Camera, uuid.UUID) string); ok {
		r0 = rf(camera, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(domain.Camera, uuid.UUID) error); ok {
		r1 = rf(camera, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertDirector provides a mock function with given fields: director
func (_m *AuthRepo) InsertDirector(director domain.Director) error {
	ret := _m.Called(director)

	if len(ret) == 0 {
		panic("no return value specified for InsertDirector")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Director) error); ok {
		r0 = rf(director)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertExpert provides a mock function with given fields: expert
func (_m *AuthRepo) InsertExpert(expert domain.Expert) error {
	ret := _m.Called(expert)

	if len(ret) == 0 {
		panic("no return value specified for InsertExpert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Expert) error); ok {
		r0 = rf(expert)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertUser provides a mock function with given fields: user
func (_m *AuthRepo) InsertUser(user domain.UserInfo) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for InsertUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.UserInfo) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignIn provides a mock function with given fields: username
func (_m *AuthRepo) SignIn(username string) (domain.UserInfo, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for SignIn")
	}

	var r0 domain.UserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.UserInfo, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) domain.UserInfo); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.UserInfo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthRepo creates a new instance of AuthRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthRepo {
	mock := &AuthRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
