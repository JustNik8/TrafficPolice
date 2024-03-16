// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	domain "TrafficPolice/internal/domain"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// DirectorService is an autogenerated mock type for the DirectorService type
type DirectorService struct {
	mock.Mock
}

// GetCases provides a mock function with given fields:
func (_m *DirectorService) GetCases() ([]domain.CaseStatus, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCases")
	}

	var r0 []domain.CaseStatus
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.CaseStatus, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.CaseStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CaseStatus)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpertAnalytics provides a mock function with given fields: expertID, startTime, endTime
func (_m *DirectorService) GetExpertAnalytics(expertID string, startTime time.Time, endTime time.Time) ([]domain.AnalyticsInterval, error) {
	ret := _m.Called(expertID, startTime, endTime)

	if len(ret) == 0 {
		panic("no return value specified for GetExpertAnalytics")
	}

	var r0 []domain.AnalyticsInterval
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) ([]domain.AnalyticsInterval, error)); ok {
		return rf(expertID, startTime, endTime)
	}
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) []domain.AnalyticsInterval); ok {
		r0 = rf(expertID, startTime, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.AnalyticsInterval)
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Time, time.Time) error); ok {
		r1 = rf(expertID, startTime, endTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDirectorService creates a new instance of DirectorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDirectorService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DirectorService {
	mock := &DirectorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}