// Code generated by mockery v2.20.0. DO NOT EDIT.

package test

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	info "yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
)

// IUserInfoQuery is an autogenerated mock type for the IUserInfoQuery type
type IUserInfoQuery struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, userInfo
func (_m *IUserInfoQuery) Create(ctx context.Context, userInfo *info.UserInfo) int {
	ret := _m.Called(ctx, userInfo)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *info.UserInfo) int); ok {
		r0 = rf(ctx, userInfo)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *IUserInfoQuery) Get(ctx context.Context, id string) (*info.UserInfo, int) {
	ret := _m.Called(ctx, id)

	var r0 *info.UserInfo
	var r1 int
	if rf, ok := ret.Get(0).(func(context.Context, string) (*info.UserInfo, int)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *info.UserInfo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.UserInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) int); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, userInfo
func (_m *IUserInfoQuery) Update(ctx context.Context, userInfo *info.UserInfo) int {
	ret := _m.Called(ctx, userInfo)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *info.UserInfo) int); ok {
		r0 = rf(ctx, userInfo)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
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
