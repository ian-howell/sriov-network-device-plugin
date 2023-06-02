// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	types "github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// DeviceSelector is an autogenerated mock type for the DeviceSelector type
type DeviceSelector struct {
	mock.Mock
}

// Filter provides a mock function with given fields: _a0
func (_m *DeviceSelector) Filter(_a0 []types.HostDevice) []types.HostDevice {
	ret := _m.Called(_a0)

	var r0 []types.HostDevice
	if rf, ok := ret.Get(0).(func([]types.HostDevice) []types.HostDevice); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.HostDevice)
		}
	}

	return r0
}

type mockConstructorTestingTNewDeviceSelector interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeviceSelector creates a new instance of DeviceSelector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeviceSelector(t mockConstructorTestingTNewDeviceSelector) *DeviceSelector {
	mock := &DeviceSelector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
