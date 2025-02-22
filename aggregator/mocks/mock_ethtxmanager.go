// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethtxmanager "github.com/0xPolygonHermez/zkevm-node/ethtxmanager"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"
)

// EthTxManager is an autogenerated mock type for the ethTxManager type
type EthTxManager struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, owner, id, from, to, value, data, dbTx
func (_m *EthTxManager) Add(ctx context.Context, owner string, id string, from common.Address, to *common.Address, value *big.Int, data []byte, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, owner, id, from, to, value, data, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, common.Address, *common.Address, *big.Int, []byte, pgx.Tx) error); ok {
		r0 = rf(ctx, owner, id, from, to, value, data, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessPendingMonitoredTxs provides a mock function with given fields: ctx, owner, failedResultHandler, dbTx
func (_m *EthTxManager) ProcessPendingMonitoredTxs(ctx context.Context, owner string, failedResultHandler ethtxmanager.ResultHandler, dbTx pgx.Tx) {
	_m.Called(ctx, owner, failedResultHandler, dbTx)
}

// Result provides a mock function with given fields: ctx, owner, id, dbTx
func (_m *EthTxManager) Result(ctx context.Context, owner string, id string, dbTx pgx.Tx) (ethtxmanager.MonitoredTxResult, error) {
	ret := _m.Called(ctx, owner, id, dbTx)

	var r0 ethtxmanager.MonitoredTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, pgx.Tx) (ethtxmanager.MonitoredTxResult, error)); ok {
		return rf(ctx, owner, id, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, pgx.Tx) ethtxmanager.MonitoredTxResult); ok {
		r0 = rf(ctx, owner, id, dbTx)
	} else {
		r0 = ret.Get(0).(ethtxmanager.MonitoredTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, pgx.Tx) error); ok {
		r1 = rf(ctx, owner, id, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResultsByStatus provides a mock function with given fields: ctx, owner, statuses, dbTx
func (_m *EthTxManager) ResultsByStatus(ctx context.Context, owner string, statuses []ethtxmanager.MonitoredTxStatus, dbTx pgx.Tx) ([]ethtxmanager.MonitoredTxResult, error) {
	ret := _m.Called(ctx, owner, statuses, dbTx)

	var r0 []ethtxmanager.MonitoredTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) ([]ethtxmanager.MonitoredTxResult, error)); ok {
		return rf(ctx, owner, statuses, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) []ethtxmanager.MonitoredTxResult); ok {
		r0 = rf(ctx, owner, statuses, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethtxmanager.MonitoredTxResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []ethtxmanager.MonitoredTxStatus, pgx.Tx) error); ok {
		r1 = rf(ctx, owner, statuses, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEthTxManager creates a new instance of EthTxManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthTxManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthTxManager {
	mock := &EthTxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
