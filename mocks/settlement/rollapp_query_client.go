// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/dymensionxyz/dymension/x/rollapp/types"
)

// RollAppQueryClient is an autogenerated mock type for the QueryClient type
type RollAppQueryClient struct {
	mock.Mock
}

// BlockHeightToFinalizationQueue provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) BlockHeightToFinalizationQueue(ctx context.Context, in *types.QueryGetBlockHeightToFinalizationQueueRequest, opts ...grpc.CallOption) (*types.QueryGetBlockHeightToFinalizationQueueResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetBlockHeightToFinalizationQueueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetBlockHeightToFinalizationQueueRequest, ...grpc.CallOption) *types.QueryGetBlockHeightToFinalizationQueueResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetBlockHeightToFinalizationQueueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetBlockHeightToFinalizationQueueRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockHeightToFinalizationQueueAll provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) BlockHeightToFinalizationQueueAll(ctx context.Context, in *types.QueryAllBlockHeightToFinalizationQueueRequest, opts ...grpc.CallOption) (*types.QueryAllBlockHeightToFinalizationQueueResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllBlockHeightToFinalizationQueueResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllBlockHeightToFinalizationQueueRequest, ...grpc.CallOption) *types.QueryAllBlockHeightToFinalizationQueueResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllBlockHeightToFinalizationQueueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllBlockHeightToFinalizationQueueRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateInfoByHeight provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) GetStateInfoByHeight(ctx context.Context, in *types.QueryGetStateInfoByHeightRequest, opts ...grpc.CallOption) (*types.QueryGetStateInfoByHeightResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetStateInfoByHeightResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetStateInfoByHeightRequest, ...grpc.CallOption) *types.QueryGetStateInfoByHeightResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetStateInfoByHeightResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetStateInfoByHeightRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestFinalizedStateInfo provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) LatestFinalizedStateInfo(ctx context.Context, in *types.QueryGetLatestFinalizedStateInfoRequest, opts ...grpc.CallOption) (*types.QueryGetLatestFinalizedStateInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetLatestFinalizedStateInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetLatestFinalizedStateInfoRequest, ...grpc.CallOption) *types.QueryGetLatestFinalizedStateInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetLatestFinalizedStateInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetLatestFinalizedStateInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestStateInfoIndex provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) LatestStateInfoIndex(ctx context.Context, in *types.QueryGetLatestStateInfoIndexRequest, opts ...grpc.CallOption) (*types.QueryGetLatestStateInfoIndexResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetLatestStateInfoIndexResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetLatestStateInfoIndexRequest, ...grpc.CallOption) *types.QueryGetLatestStateInfoIndexResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetLatestStateInfoIndexResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetLatestStateInfoIndexRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestStateInfoIndexAll provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) LatestStateInfoIndexAll(ctx context.Context, in *types.QueryAllLatestStateInfoIndexRequest, opts ...grpc.CallOption) (*types.QueryAllLatestStateInfoIndexResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllLatestStateInfoIndexResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllLatestStateInfoIndexRequest, ...grpc.CallOption) *types.QueryAllLatestStateInfoIndexResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllLatestStateInfoIndexResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllLatestStateInfoIndexRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rollapp provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) Rollapp(ctx context.Context, in *types.QueryGetRollappRequest, opts ...grpc.CallOption) (*types.QueryGetRollappResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetRollappResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetRollappRequest, ...grpc.CallOption) *types.QueryGetRollappResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetRollappResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetRollappRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RollappAll provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) RollappAll(ctx context.Context, in *types.QueryAllRollappRequest, opts ...grpc.CallOption) (*types.QueryAllRollappResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllRollappResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllRollappRequest, ...grpc.CallOption) *types.QueryAllRollappResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllRollappResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllRollappRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInfo provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) StateInfo(ctx context.Context, in *types.QueryGetStateInfoRequest, opts ...grpc.CallOption) (*types.QueryGetStateInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetStateInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetStateInfoRequest, ...grpc.CallOption) *types.QueryGetStateInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetStateInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetStateInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInfoAll provides a mock function with given fields: ctx, in, opts
func (_m *RollAppQueryClient) StateInfoAll(ctx context.Context, in *types.QueryAllStateInfoRequest, opts ...grpc.CallOption) (*types.QueryAllStateInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllStateInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllStateInfoRequest, ...grpc.CallOption) *types.QueryAllStateInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllStateInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllStateInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRollAppQueryClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewRollAppQueryClient creates a new instance of RollAppQueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRollAppQueryClient(t mockConstructorTestingTNewRollAppQueryClient) *RollAppQueryClient {
	mock := &RollAppQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
