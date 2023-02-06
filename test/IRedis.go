// Code generated by mockery v2.16.0. DO NOT EDIT.

package test

import (
	redis "github.com/go-redis/redis/v8"
	mock "github.com/stretchr/testify/mock"
)

// IRedis is an autogenerated mock type for the IRedis type
type IRedis struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *IRedis) Client() *redis.Client {
	ret := _m.Called()

	var r0 *redis.Client
	if rf, ok := ret.Get(0).(func() *redis.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.Client)
		}
	}

	return r0
}

type mockConstructorTestingTNewIRedis interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRedis creates a new instance of IRedis. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRedis(t mockConstructorTestingTNewIRedis) *IRedis {
	mock := &IRedis{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}