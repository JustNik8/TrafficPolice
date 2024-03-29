// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	domain "TrafficPolice/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ExpertRepo is an autogenerated mock type for the ExpertRepo type
type ExpertRepo struct {
	mock.Mock
}

// GetCaseFineDecisions provides a mock function with given fields: caseID, competenceSkill
func (_m *ExpertRepo) GetCaseFineDecisions(caseID string, competenceSkill int) (domain.FineDecisions, error) {
	ret := _m.Called(caseID, competenceSkill)

	if len(ret) == 0 {
		panic("no return value specified for GetCaseFineDecisions")
	}

	var r0 domain.FineDecisions
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (domain.FineDecisions, error)); ok {
		return rf(caseID, competenceSkill)
	}
	if rf, ok := ret.Get(0).(func(string, int) domain.FineDecisions); ok {
		r0 = rf(caseID, competenceSkill)
	} else {
		r0 = ret.Get(0).(domain.FineDecisions)
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(caseID, competenceSkill)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpertByUserID provides a mock function with given fields: userID
func (_m *ExpertRepo) GetExpertByUserID(userID string) (domain.Expert, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetExpertByUserID")
	}

	var r0 domain.Expert
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Expert, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Expert); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(domain.Expert)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpertsCountBySkill provides a mock function with given fields: competenceSkill
func (_m *ExpertRepo) GetExpertsCountBySkill(competenceSkill int) (int, error) {
	ret := _m.Called(competenceSkill)

	if len(ret) == 0 {
		panic("no return value specified for GetExpertsCountBySkill")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int, error)); ok {
		return rf(competenceSkill)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(competenceSkill)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(competenceSkill)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastNotSolvedCaseID provides a mock function with given fields: expertID
func (_m *ExpertRepo) GetLastNotSolvedCaseID(expertID string) (string, error) {
	ret := _m.Called(expertID)

	if len(ret) == 0 {
		panic("no return value specified for GetLastNotSolvedCaseID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(expertID)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(expertID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(expertID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotSolvedCase provides a mock function with given fields: expert
func (_m *ExpertRepo) GetNotSolvedCase(expert domain.Expert) (domain.Case, error) {
	ret := _m.Called(expert)

	if len(ret) == 0 {
		panic("no return value specified for GetNotSolvedCase")
	}

	var r0 domain.Case
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Expert) (domain.Case, error)); ok {
		return rf(expert)
	}
	if rf, ok := ret.Get(0).(func(domain.Expert) domain.Case); ok {
		r0 = rf(expert)
	} else {
		r0 = ret.Get(0).(domain.Case)
	}

	if rf, ok := ret.Get(1).(func(domain.Expert) error); ok {
		r1 = rf(expert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNotSolvedCase provides a mock function with given fields: solvedCase
func (_m *ExpertRepo) InsertNotSolvedCase(solvedCase domain.ExpertCase) error {
	ret := _m.Called(solvedCase)

	if len(ret) == 0 {
		panic("no return value specified for InsertNotSolvedCase")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.ExpertCase) error); ok {
		r0 = rf(solvedCase)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCaseDecision provides a mock function with given fields: decision
func (_m *ExpertRepo) SetCaseDecision(decision domain.Decision) error {
	ret := _m.Called(decision)

	if len(ret) == 0 {
		panic("no return value specified for SetCaseDecision")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Decision) error); ok {
		r0 = rf(decision)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewExpertRepo creates a new instance of ExpertRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExpertRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExpertRepo {
	mock := &ExpertRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
