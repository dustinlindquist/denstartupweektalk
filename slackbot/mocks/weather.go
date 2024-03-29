// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import weather "github.com/dustinlindquist/denstartupweektalk/slackbot/weather"

// Weather is an autogenerated mock type for the Weather type
type Weather struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx
func (_m *Weather) Get(ctx context.Context) (weather.Data, error) {
	ret := _m.Called(ctx)

	var r0 weather.Data
	if rf, ok := ret.Get(0).(func(context.Context) weather.Data); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(weather.Data)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
