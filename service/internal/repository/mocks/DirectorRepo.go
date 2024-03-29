// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	domain "TrafficPolice/internal/domain"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// DirectorRepo is an autogenerated mock type for the DirectorRepo type
type DirectorRepo struct {
	mock.Mock
}

// GetCase provides a mock function with given fields: caseID
func (_m *DirectorRepo) GetCase(caseID string) (domain.CaseStatus, error) {
	ret := _m.Called(caseID)

	if len(ret) == 0 {
		panic("no return value specified for GetCase")
	}

	var r0 domain.CaseStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.CaseStatus, error)); ok {
		return rf(caseID)
	}
	if rf, ok := ret.Get(0).(func(string) domain.CaseStatus); ok {
		r0 = rf(caseID)
	} else {
		r0 = ret.Get(0).(domain.CaseStatus)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(caseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpertIntervalCases provides a mock function with given fields: expertID, startDate, endDate
func (_m *DirectorRepo) GetExpertIntervalCases(expertID string, startDate time.Time, endDate time.Time) (map[domain.Date][]domain.IntervalCase, error) {
	ret := _m.Called(expertID, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for GetExpertIntervalCases")
	}

	var r0 map[domain.Date][]domain.IntervalCase
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) (map[domain.Date][]domain.IntervalCase, error)); ok {
		return rf(expertID, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) map[domain.Date][]domain.IntervalCase); ok {
		r0 = rf(expertID, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[domain.Date][]domain.IntervalCase)
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Time, time.Time) error); ok {
		r1 = rf(expertID, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExpertSkill provides a mock function with given fields: expertID, skill
func (_m *DirectorRepo) UpdateExpertSkill(expertID string, skill int) error {
	ret := _m.Called(expertID, skill)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExpertSkill")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(expertID, skill)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDirectorRepo creates a new instance of DirectorRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDirectorRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *DirectorRepo {
	mock := &DirectorRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
