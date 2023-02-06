// Code generated by mockery v2.16.0. DO NOT EDIT.

package test

import (
	mock "github.com/stretchr/testify/mock"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
)

// IUserAccountQuery is an autogenerated mock type for the IUserAccountQuery type
type IUserAccountQuery struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *IUserAccountQuery) Create(input *account.SignUpForm) (string, int) {
	ret := _m.Called(input)

	var r0 string
	if rf, ok := ret.Get(0).(func(*account.SignUpForm) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*account.SignUpForm) int); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *IUserAccountQuery) GetByUsername(username string) (*account.UserAccount, int) {
	ret := _m.Called(username)

	var r0 *account.UserAccount
	if rf, ok := ret.Get(0).(func(string) *account.UserAccount); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.UserAccount)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserAccountQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserAccountQuery creates a new instance of IUserAccountQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserAccountQuery(t mockConstructorTestingTNewIUserAccountQuery) *IUserAccountQuery {
	mock := &IUserAccountQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
