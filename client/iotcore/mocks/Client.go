// Copyright 2022 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	iotcore "github.com/mendersoftware/iot-manager/client/iotcore"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/iot-manager/model"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// DeleteDevice provides a mock function with given fields: ctx, creds, deviceID
func (_m *Client) DeleteDevice(ctx context.Context, creds model.AWSCredentials, deviceID string) error {
	ret := _m.Called(ctx, creds, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.AWSCredentials, string) error); ok {
		r0 = rf(ctx, creds, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: ctx, creds, deviceID
func (_m *Client) GetDevice(ctx context.Context, creds model.AWSCredentials, deviceID string) (*iotcore.Device, error) {
	ret := _m.Called(ctx, creds, deviceID)

	var r0 *iotcore.Device
	if rf, ok := ret.Get(0).(func(context.Context, model.AWSCredentials, string) *iotcore.Device); ok {
		r0 = rf(ctx, creds, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotcore.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AWSCredentials, string) error); ok {
		r1 = rf(ctx, creds, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceShadow provides a mock function with given fields: ctx, creds, id
func (_m *Client) GetDeviceShadow(ctx context.Context, creds model.AWSCredentials, id string) (*iotcore.DeviceShadow, error) {
	ret := _m.Called(ctx, creds, id)

	var r0 *iotcore.DeviceShadow
	if rf, ok := ret.Get(0).(func(context.Context, model.AWSCredentials, string) *iotcore.DeviceShadow); ok {
		r0 = rf(ctx, creds, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotcore.DeviceShadow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AWSCredentials, string) error); ok {
		r1 = rf(ctx, creds, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeviceShadow provides a mock function with given fields: ctx, creds, deviceID, update
func (_m *Client) UpdateDeviceShadow(ctx context.Context, creds model.AWSCredentials, deviceID string, update iotcore.DeviceShadowUpdate) (*iotcore.DeviceShadow, error) {
	ret := _m.Called(ctx, creds, deviceID, update)

	var r0 *iotcore.DeviceShadow
	if rf, ok := ret.Get(0).(func(context.Context, model.AWSCredentials, string, iotcore.DeviceShadowUpdate) *iotcore.DeviceShadow); ok {
		r0 = rf(ctx, creds, deviceID, update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotcore.DeviceShadow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AWSCredentials, string, iotcore.DeviceShadowUpdate) error); ok {
		r1 = rf(ctx, creds, deviceID, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertDevice provides a mock function with given fields: ctx, creds, deviceID, device, policy
func (_m *Client) UpsertDevice(ctx context.Context, creds model.AWSCredentials, deviceID string, device *iotcore.Device, policy string) (*iotcore.Device, error) {
	ret := _m.Called(ctx, creds, deviceID, device, policy)

	var r0 *iotcore.Device
	if rf, ok := ret.Get(0).(func(context.Context, model.AWSCredentials, string, *iotcore.Device, string) *iotcore.Device); ok {
		r0 = rf(ctx, creds, deviceID, device, policy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotcore.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AWSCredentials, string, *iotcore.Device, string) error); ok {
		r1 = rf(ctx, creds, deviceID, device, policy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}