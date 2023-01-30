// Code generated by mockery v2.16.0. DO NOT EDIT.

package test

import (
	mock "github.com/stretchr/testify/mock"
	info "yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

// IUserInfoQuery is an autogenerated mock type for the IUserInfoQuery type
type IUserInfoQuery struct {
	mock.Mock
}

// Create provides a mock function with given fields: userInfo
func (_m *IUserInfoQuery) Create(userInfo *info.UserInfo) int {
	ret := _m.Called(userInfo)

	var r0 int
	if rf, ok := ret.Get(0).(func(*info.UserInfo) int); ok {
		r0 = rf(userInfo)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *IUserInfoQuery) Get(id string) (*info.UserInfo, int) {
	ret := _m.Called(id)

	var r0 *info.UserInfo
	if rf, ok := ret.Get(0).(func(string) *info.UserInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.UserInfo)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserInfoQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserInfoQuery creates a new instance of IUserInfoQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserInfoQuery(t mockConstructorTestingTNewIUserInfoQuery) *IUserInfoQuery {
	mock := &IUserInfoQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}