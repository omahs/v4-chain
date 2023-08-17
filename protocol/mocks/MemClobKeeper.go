// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	indexer_manager "github.com/dydxprotocol/v4-chain/protocol/indexer/indexer_manager"
	clobtypes "github.com/dydxprotocol/v4-chain/protocol/x/clob/types"

	mock "github.com/stretchr/testify/mock"

	subaccountstypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"

	time "time"

	types "github.com/cosmos/cosmos-sdk/types"
)

// MemClobKeeper is an autogenerated mock type for the MemClobKeeper type
type MemClobKeeper struct {
	mock.Mock
}

// AddOrderToOrderbookCollatCheck provides a mock function with given fields: ctx, clobPairId, subaccountOpenOrders
func (_m *MemClobKeeper) AddOrderToOrderbookCollatCheck(ctx types.Context, clobPairId clobtypes.ClobPairId, subaccountOpenOrders map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) (bool, map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult) {
	ret := _m.Called(ctx, clobPairId, subaccountOpenOrders)

	var r0 bool
	var r1 map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId, map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) (bool, map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult)); ok {
		return rf(ctx, clobPairId, subaccountOpenOrders)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId, map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) bool); ok {
		r0 = rf(ctx, clobPairId, subaccountOpenOrders)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.ClobPairId, map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, clobPairId, subaccountOpenOrders)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult)
		}
	}

	return r0, r1
}

// CancelShortTermOrder provides a mock function with given fields: ctx, msgCancelOrder
func (_m *MemClobKeeper) CancelShortTermOrder(ctx types.Context, msgCancelOrder *clobtypes.MsgCancelOrder) error {
	ret := _m.Called(ctx, msgCancelOrder)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r0 = rf(ctx, msgCancelOrder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIndexerEventManager provides a mock function with given fields:
func (_m *MemClobKeeper) GetIndexerEventManager() indexer_manager.IndexerEventManager {
	ret := _m.Called()

	var r0 indexer_manager.IndexerEventManager
	if rf, ok := ret.Get(0).(func() indexer_manager.IndexerEventManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(indexer_manager.IndexerEventManager)
		}
	}

	return r0
}

// GetLongTermOrderPlacement provides a mock function with given fields: ctx, orderId
func (_m *MemClobKeeper) GetLongTermOrderPlacement(ctx types.Context, orderId clobtypes.OrderId) (clobtypes.LongTermOrderPlacement, bool) {
	ret := _m.Called(ctx, orderId)

	var r0 clobtypes.LongTermOrderPlacement
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (clobtypes.LongTermOrderPlacement, bool)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) clobtypes.LongTermOrderPlacement); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(clobtypes.LongTermOrderPlacement)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) bool); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetOrderFillAmount provides a mock function with given fields: ctx, orderId
func (_m *MemClobKeeper) GetOrderFillAmount(ctx types.Context, orderId clobtypes.OrderId) (bool, subaccountstypes.BaseQuantums, uint32) {
	ret := _m.Called(ctx, orderId)

	var r0 bool
	var r1 subaccountstypes.BaseQuantums
	var r2 uint32
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (bool, subaccountstypes.BaseQuantums, uint32)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) bool); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) subaccountstypes.BaseQuantums); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(2).(func(types.Context, clobtypes.OrderId) uint32); ok {
		r2 = rf(ctx, orderId)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	return r0, r1, r2
}

// GetStatePosition provides a mock function with given fields: ctx, subaccountId, clobPairId
func (_m *MemClobKeeper) GetStatePosition(ctx types.Context, subaccountId subaccountstypes.SubaccountId, clobPairId clobtypes.ClobPairId) *big.Int {
	ret := _m.Called(ctx, subaccountId, clobPairId)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, clobtypes.ClobPairId) *big.Int); ok {
		r0 = rf(ctx, subaccountId, clobPairId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// IsLiquidatable provides a mock function with given fields: ctx, subaccountId
func (_m *MemClobKeeper) IsLiquidatable(ctx types.Context, subaccountId subaccountstypes.SubaccountId) (bool, error) {
	ret := _m.Called(ctx, subaccountId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) (bool, error)); ok {
		return rf(ctx, subaccountId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) bool); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId) error); ok {
		r1 = rf(ctx, subaccountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MustAddOrderToStatefulOrdersTimeSlice provides a mock function with given fields: ctx, goodTilBlockTime, orderId
func (_m *MemClobKeeper) MustAddOrderToStatefulOrdersTimeSlice(ctx types.Context, goodTilBlockTime time.Time, orderId clobtypes.OrderId) {
	_m.Called(ctx, goodTilBlockTime, orderId)
}

// OffsetSubaccountPerpetualPosition provides a mock function with given fields: ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal
func (_m *MemClobKeeper) OffsetSubaccountPerpetualPosition(ctx types.Context, liquidatedSubaccountId subaccountstypes.SubaccountId, perpetualId uint32, deltaQuantumsTotal *big.Int) ([]clobtypes.MatchPerpetualDeleveraging_Fill, *big.Int) {
	ret := _m.Called(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal)

	var r0 []clobtypes.MatchPerpetualDeleveraging_Fill
	var r1 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) ([]clobtypes.MatchPerpetualDeleveraging_Fill, *big.Int)); ok {
		return rf(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) []clobtypes.MatchPerpetualDeleveraging_Fill); ok {
		r0 = rf(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.MatchPerpetualDeleveraging_Fill)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) *big.Int); ok {
		r1 = rf(ctx, liquidatedSubaccountId, perpetualId, deltaQuantumsTotal)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	return r0, r1
}

// ProcessSingleMatch provides a mock function with given fields: ctx, matchWithOrders
func (_m *MemClobKeeper) ProcessSingleMatch(ctx types.Context, matchWithOrders *clobtypes.MatchWithOrders) (bool, subaccountstypes.UpdateResult, subaccountstypes.UpdateResult, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, matchWithOrders)

	var r0 bool
	var r1 subaccountstypes.UpdateResult
	var r2 subaccountstypes.UpdateResult
	var r3 *clobtypes.OffchainUpdates
	var r4 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MatchWithOrders) (bool, subaccountstypes.UpdateResult, subaccountstypes.UpdateResult, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, matchWithOrders)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MatchWithOrders) bool); ok {
		r0 = rf(ctx, matchWithOrders)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MatchWithOrders) subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, matchWithOrders)
	} else {
		r1 = ret.Get(1).(subaccountstypes.UpdateResult)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MatchWithOrders) subaccountstypes.UpdateResult); ok {
		r2 = rf(ctx, matchWithOrders)
	} else {
		r2 = ret.Get(2).(subaccountstypes.UpdateResult)
	}

	if rf, ok := ret.Get(3).(func(types.Context, *clobtypes.MatchWithOrders) *clobtypes.OffchainUpdates); ok {
		r3 = rf(ctx, matchWithOrders)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(4).(func(types.Context, *clobtypes.MatchWithOrders) error); ok {
		r4 = rf(ctx, matchWithOrders)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// ReplayPlaceOrder provides a mock function with given fields: ctx, msg
func (_m *MemClobKeeper) ReplayPlaceOrder(ctx types.Context, msg *clobtypes.MsgPlaceOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, msg)

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 *clobtypes.OffchainUpdates
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MsgPlaceOrder) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MsgPlaceOrder) *clobtypes.OffchainUpdates); ok {
		r2 = rf(ctx, msg)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(3).(func(types.Context, *clobtypes.MsgPlaceOrder) error); ok {
		r3 = rf(ctx, msg)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// SetLongTermOrderPlacement provides a mock function with given fields: ctx, order, blockHeight
func (_m *MemClobKeeper) SetLongTermOrderPlacement(ctx types.Context, order clobtypes.Order, blockHeight uint32) {
	_m.Called(ctx, order, blockHeight)
}

type mockConstructorTestingTNewMemClobKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewMemClobKeeper creates a new instance of MemClobKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMemClobKeeper(t mockConstructorTestingTNewMemClobKeeper) *MemClobKeeper {
	mock := &MemClobKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
