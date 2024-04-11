// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// URLDeletter is an autogenerated mock type for the URLDeletter type
type URLDeletter struct {
	mock.Mock
}

// DeleteURL provides a mock function with given fields: alias
func (_m *URLDeletter) DeleteURL(alias string) error {
	ret := _m.Called(alias)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewURLDeletter interface {
	mock.TestingT
	Cleanup(func())
}

// NewURLDeletter creates a new instance of URLDeletter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewURLDeletter(t mockConstructorTestingTNewURLDeletter) *URLDeletter {
	mock := &URLDeletter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
