// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package ethereum is a generated GoMock package.
package ethereum

import (
	context "context"
	big "math/big"
	reflect "reflect"

	common "github.com/autonity/autonity/common"
	types "github.com/autonity/autonity/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockSubscription is a mock of Subscription interface.
type MockSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionMockRecorder
}

// MockSubscriptionMockRecorder is the mock recorder for MockSubscription.
type MockSubscriptionMockRecorder struct {
	mock *MockSubscription
}

// NewMockSubscription creates a new mock instance.
func NewMockSubscription(ctrl *gomock.Controller) *MockSubscription {
	mock := &MockSubscription{ctrl: ctrl}
	mock.recorder = &MockSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscription) EXPECT() *MockSubscriptionMockRecorder {
	return m.recorder
}

// Err mocks base method.
func (m *MockSubscription) Err() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSubscriptionMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSubscription)(nil).Err))
}

// Unsubscribe mocks base method.
func (m *MockSubscription) Unsubscribe() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unsubscribe")
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscriptionMockRecorder) Unsubscribe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscription)(nil).Unsubscribe))
}

// MockChainReader is a mock of ChainReader interface.
type MockChainReader struct {
	ctrl     *gomock.Controller
	recorder *MockChainReaderMockRecorder
}

// MockChainReaderMockRecorder is the mock recorder for MockChainReader.
type MockChainReaderMockRecorder struct {
	mock *MockChainReader
}

// NewMockChainReader creates a new mock instance.
func NewMockChainReader(ctrl *gomock.Controller) *MockChainReader {
	mock := &MockChainReader{ctrl: ctrl}
	mock.recorder = &MockChainReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainReader) EXPECT() *MockChainReaderMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockChainReader) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockChainReaderMockRecorder) BlockByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockChainReader)(nil).BlockByHash), ctx, hash)
}

// BlockByNumber mocks base method.
func (m *MockChainReader) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockChainReaderMockRecorder) BlockByNumber(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockChainReader)(nil).BlockByNumber), ctx, number)
}

// HeaderByHash mocks base method.
func (m *MockChainReader) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByHash indicates an expected call of HeaderByHash.
func (mr *MockChainReaderMockRecorder) HeaderByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByHash", reflect.TypeOf((*MockChainReader)(nil).HeaderByHash), ctx, hash)
}

// HeaderByNumber mocks base method.
func (m *MockChainReader) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", ctx, number)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockChainReaderMockRecorder) HeaderByNumber(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockChainReader)(nil).HeaderByNumber), ctx, number)
}

// SubscribeNewHead mocks base method.
func (m *MockChainReader) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewHead", ctx, ch)
	ret0, _ := ret[0].(Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewHead indicates an expected call of SubscribeNewHead.
func (mr *MockChainReaderMockRecorder) SubscribeNewHead(ctx, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewHead", reflect.TypeOf((*MockChainReader)(nil).SubscribeNewHead), ctx, ch)
}

// TransactionCount mocks base method.
func (m *MockChainReader) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionCount", ctx, blockHash)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionCount indicates an expected call of TransactionCount.
func (mr *MockChainReaderMockRecorder) TransactionCount(ctx, blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCount", reflect.TypeOf((*MockChainReader)(nil).TransactionCount), ctx, blockHash)
}

// TransactionInBlock mocks base method.
func (m *MockChainReader) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionInBlock", ctx, blockHash, index)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionInBlock indicates an expected call of TransactionInBlock.
func (mr *MockChainReaderMockRecorder) TransactionInBlock(ctx, blockHash, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionInBlock", reflect.TypeOf((*MockChainReader)(nil).TransactionInBlock), ctx, blockHash, index)
}

// MockTransactionReader is a mock of TransactionReader interface.
type MockTransactionReader struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionReaderMockRecorder
}

// MockTransactionReaderMockRecorder is the mock recorder for MockTransactionReader.
type MockTransactionReaderMockRecorder struct {
	mock *MockTransactionReader
}

// NewMockTransactionReader creates a new mock instance.
func NewMockTransactionReader(ctrl *gomock.Controller) *MockTransactionReader {
	mock := &MockTransactionReader{ctrl: ctrl}
	mock.recorder = &MockTransactionReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionReader) EXPECT() *MockTransactionReaderMockRecorder {
	return m.recorder
}

// TransactionByHash mocks base method.
func (m *MockTransactionReader) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, txHash)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockTransactionReaderMockRecorder) TransactionByHash(ctx, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockTransactionReader)(nil).TransactionByHash), ctx, txHash)
}

// TransactionReceipt mocks base method.
func (m *MockTransactionReader) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockTransactionReaderMockRecorder) TransactionReceipt(ctx, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockTransactionReader)(nil).TransactionReceipt), ctx, txHash)
}

// MockChainStateReader is a mock of ChainStateReader interface.
type MockChainStateReader struct {
	ctrl     *gomock.Controller
	recorder *MockChainStateReaderMockRecorder
}

// MockChainStateReaderMockRecorder is the mock recorder for MockChainStateReader.
type MockChainStateReaderMockRecorder struct {
	mock *MockChainStateReader
}

// NewMockChainStateReader creates a new mock instance.
func NewMockChainStateReader(ctrl *gomock.Controller) *MockChainStateReader {
	mock := &MockChainStateReader{ctrl: ctrl}
	mock.recorder = &MockChainStateReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainStateReader) EXPECT() *MockChainStateReaderMockRecorder {
	return m.recorder
}

// BalanceAt mocks base method.
func (m *MockChainStateReader) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceAt", ctx, account, blockNumber)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceAt indicates an expected call of BalanceAt.
func (mr *MockChainStateReaderMockRecorder) BalanceAt(ctx, account, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceAt", reflect.TypeOf((*MockChainStateReader)(nil).BalanceAt), ctx, account, blockNumber)
}

// CodeAt mocks base method.
func (m *MockChainStateReader) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, account, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockChainStateReaderMockRecorder) CodeAt(ctx, account, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockChainStateReader)(nil).CodeAt), ctx, account, blockNumber)
}

// NonceAt mocks base method.
func (m *MockChainStateReader) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NonceAt", ctx, account, blockNumber)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NonceAt indicates an expected call of NonceAt.
func (mr *MockChainStateReaderMockRecorder) NonceAt(ctx, account, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NonceAt", reflect.TypeOf((*MockChainStateReader)(nil).NonceAt), ctx, account, blockNumber)
}

// StorageAt mocks base method.
func (m *MockChainStateReader) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAt", ctx, account, key, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAt indicates an expected call of StorageAt.
func (mr *MockChainStateReaderMockRecorder) StorageAt(ctx, account, key, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAt", reflect.TypeOf((*MockChainStateReader)(nil).StorageAt), ctx, account, key, blockNumber)
}

// MockChainSyncReader is a mock of ChainSyncReader interface.
type MockChainSyncReader struct {
	ctrl     *gomock.Controller
	recorder *MockChainSyncReaderMockRecorder
}

// MockChainSyncReaderMockRecorder is the mock recorder for MockChainSyncReader.
type MockChainSyncReaderMockRecorder struct {
	mock *MockChainSyncReader
}

// NewMockChainSyncReader creates a new mock instance.
func NewMockChainSyncReader(ctrl *gomock.Controller) *MockChainSyncReader {
	mock := &MockChainSyncReader{ctrl: ctrl}
	mock.recorder = &MockChainSyncReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainSyncReader) EXPECT() *MockChainSyncReaderMockRecorder {
	return m.recorder
}

// SyncProgress mocks base method.
func (m *MockChainSyncReader) SyncProgress(ctx context.Context) (*SyncProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncProgress", ctx)
	ret0, _ := ret[0].(*SyncProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncProgress indicates an expected call of SyncProgress.
func (mr *MockChainSyncReaderMockRecorder) SyncProgress(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProgress", reflect.TypeOf((*MockChainSyncReader)(nil).SyncProgress), ctx)
}

// MockContractCaller is a mock of ContractCaller interface.
type MockContractCaller struct {
	ctrl     *gomock.Controller
	recorder *MockContractCallerMockRecorder
}

// MockContractCallerMockRecorder is the mock recorder for MockContractCaller.
type MockContractCallerMockRecorder struct {
	mock *MockContractCaller
}

// NewMockContractCaller creates a new mock instance.
func NewMockContractCaller(ctrl *gomock.Controller) *MockContractCaller {
	mock := &MockContractCaller{ctrl: ctrl}
	mock.recorder = &MockContractCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractCaller) EXPECT() *MockContractCallerMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractCaller) CallContract(ctx context.Context, call CallMsg, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, call, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractCallerMockRecorder) CallContract(ctx, call, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractCaller)(nil).CallContract), ctx, call, blockNumber)
}

// MockLogFilterer is a mock of LogFilterer interface.
type MockLogFilterer struct {
	ctrl     *gomock.Controller
	recorder *MockLogFiltererMockRecorder
}

// MockLogFiltererMockRecorder is the mock recorder for MockLogFilterer.
type MockLogFiltererMockRecorder struct {
	mock *MockLogFilterer
}

// NewMockLogFilterer creates a new mock instance.
func NewMockLogFilterer(ctrl *gomock.Controller) *MockLogFilterer {
	mock := &MockLogFilterer{ctrl: ctrl}
	mock.recorder = &MockLogFiltererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFilterer) EXPECT() *MockLogFiltererMockRecorder {
	return m.recorder
}

// FilterLogs mocks base method.
func (m *MockLogFilterer) FilterLogs(ctx context.Context, q FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterLogs", ctx, q)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterLogs indicates an expected call of FilterLogs.
func (mr *MockLogFiltererMockRecorder) FilterLogs(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterLogs", reflect.TypeOf((*MockLogFilterer)(nil).FilterLogs), ctx, q)
}

// SubscribeFilterLogs mocks base method.
func (m *MockLogFilterer) SubscribeFilterLogs(ctx context.Context, q FilterQuery, ch chan<- types.Log) (Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeFilterLogs", ctx, q, ch)
	ret0, _ := ret[0].(Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeFilterLogs indicates an expected call of SubscribeFilterLogs.
func (mr *MockLogFiltererMockRecorder) SubscribeFilterLogs(ctx, q, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeFilterLogs", reflect.TypeOf((*MockLogFilterer)(nil).SubscribeFilterLogs), ctx, q, ch)
}

// MockTransactionSender is a mock of TransactionSender interface.
type MockTransactionSender struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionSenderMockRecorder
}

// MockTransactionSenderMockRecorder is the mock recorder for MockTransactionSender.
type MockTransactionSenderMockRecorder struct {
	mock *MockTransactionSender
}

// NewMockTransactionSender creates a new mock instance.
func NewMockTransactionSender(ctrl *gomock.Controller) *MockTransactionSender {
	mock := &MockTransactionSender{ctrl: ctrl}
	mock.recorder = &MockTransactionSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionSender) EXPECT() *MockTransactionSenderMockRecorder {
	return m.recorder
}

// SendTransaction mocks base method.
func (m *MockTransactionSender) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockTransactionSenderMockRecorder) SendTransaction(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockTransactionSender)(nil).SendTransaction), ctx, tx)
}

// MockGasPricer is a mock of GasPricer interface.
type MockGasPricer struct {
	ctrl     *gomock.Controller
	recorder *MockGasPricerMockRecorder
}

// MockGasPricerMockRecorder is the mock recorder for MockGasPricer.
type MockGasPricerMockRecorder struct {
	mock *MockGasPricer
}

// NewMockGasPricer creates a new mock instance.
func NewMockGasPricer(ctrl *gomock.Controller) *MockGasPricer {
	mock := &MockGasPricer{ctrl: ctrl}
	mock.recorder = &MockGasPricerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGasPricer) EXPECT() *MockGasPricerMockRecorder {
	return m.recorder
}

// SuggestGasPrice mocks base method.
func (m *MockGasPricer) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasPrice", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice.
func (mr *MockGasPricerMockRecorder) SuggestGasPrice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockGasPricer)(nil).SuggestGasPrice), ctx)
}

// MockPendingStateReader is a mock of PendingStateReader interface.
type MockPendingStateReader struct {
	ctrl     *gomock.Controller
	recorder *MockPendingStateReaderMockRecorder
}

// MockPendingStateReaderMockRecorder is the mock recorder for MockPendingStateReader.
type MockPendingStateReaderMockRecorder struct {
	mock *MockPendingStateReader
}

// NewMockPendingStateReader creates a new mock instance.
func NewMockPendingStateReader(ctrl *gomock.Controller) *MockPendingStateReader {
	mock := &MockPendingStateReader{ctrl: ctrl}
	mock.recorder = &MockPendingStateReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPendingStateReader) EXPECT() *MockPendingStateReaderMockRecorder {
	return m.recorder
}

// PendingBalanceAt mocks base method.
func (m *MockPendingStateReader) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingBalanceAt", ctx, account)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingBalanceAt indicates an expected call of PendingBalanceAt.
func (mr *MockPendingStateReaderMockRecorder) PendingBalanceAt(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingBalanceAt", reflect.TypeOf((*MockPendingStateReader)(nil).PendingBalanceAt), ctx, account)
}

// PendingCodeAt mocks base method.
func (m *MockPendingStateReader) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingCodeAt", ctx, account)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingCodeAt indicates an expected call of PendingCodeAt.
func (mr *MockPendingStateReaderMockRecorder) PendingCodeAt(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingCodeAt", reflect.TypeOf((*MockPendingStateReader)(nil).PendingCodeAt), ctx, account)
}

// PendingNonceAt mocks base method.
func (m *MockPendingStateReader) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingNonceAt", ctx, account)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingNonceAt indicates an expected call of PendingNonceAt.
func (mr *MockPendingStateReaderMockRecorder) PendingNonceAt(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingNonceAt", reflect.TypeOf((*MockPendingStateReader)(nil).PendingNonceAt), ctx, account)
}

// PendingStorageAt mocks base method.
func (m *MockPendingStateReader) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingStorageAt", ctx, account, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingStorageAt indicates an expected call of PendingStorageAt.
func (mr *MockPendingStateReaderMockRecorder) PendingStorageAt(ctx, account, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingStorageAt", reflect.TypeOf((*MockPendingStateReader)(nil).PendingStorageAt), ctx, account, key)
}

// PendingTransactionCount mocks base method.
func (m *MockPendingStateReader) PendingTransactionCount(ctx context.Context) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingTransactionCount", ctx)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingTransactionCount indicates an expected call of PendingTransactionCount.
func (mr *MockPendingStateReaderMockRecorder) PendingTransactionCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingTransactionCount", reflect.TypeOf((*MockPendingStateReader)(nil).PendingTransactionCount), ctx)
}

// MockPendingContractCaller is a mock of PendingContractCaller interface.
type MockPendingContractCaller struct {
	ctrl     *gomock.Controller
	recorder *MockPendingContractCallerMockRecorder
}

// MockPendingContractCallerMockRecorder is the mock recorder for MockPendingContractCaller.
type MockPendingContractCallerMockRecorder struct {
	mock *MockPendingContractCaller
}

// NewMockPendingContractCaller creates a new mock instance.
func NewMockPendingContractCaller(ctrl *gomock.Controller) *MockPendingContractCaller {
	mock := &MockPendingContractCaller{ctrl: ctrl}
	mock.recorder = &MockPendingContractCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPendingContractCaller) EXPECT() *MockPendingContractCallerMockRecorder {
	return m.recorder
}

// PendingCallContract mocks base method.
func (m *MockPendingContractCaller) PendingCallContract(ctx context.Context, call CallMsg) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingCallContract", ctx, call)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingCallContract indicates an expected call of PendingCallContract.
func (mr *MockPendingContractCallerMockRecorder) PendingCallContract(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingCallContract", reflect.TypeOf((*MockPendingContractCaller)(nil).PendingCallContract), ctx, call)
}

// MockGasEstimator is a mock of GasEstimator interface.
type MockGasEstimator struct {
	ctrl     *gomock.Controller
	recorder *MockGasEstimatorMockRecorder
}

// MockGasEstimatorMockRecorder is the mock recorder for MockGasEstimator.
type MockGasEstimatorMockRecorder struct {
	mock *MockGasEstimator
}

// NewMockGasEstimator creates a new mock instance.
func NewMockGasEstimator(ctrl *gomock.Controller) *MockGasEstimator {
	mock := &MockGasEstimator{ctrl: ctrl}
	mock.recorder = &MockGasEstimatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGasEstimator) EXPECT() *MockGasEstimatorMockRecorder {
	return m.recorder
}

// EstimateGas mocks base method.
func (m *MockGasEstimator) EstimateGas(ctx context.Context, call CallMsg) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", ctx, call)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas.
func (mr *MockGasEstimatorMockRecorder) EstimateGas(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockGasEstimator)(nil).EstimateGas), ctx, call)
}

// MockPendingStateEventer is a mock of PendingStateEventer interface.
type MockPendingStateEventer struct {
	ctrl     *gomock.Controller
	recorder *MockPendingStateEventerMockRecorder
}

// MockPendingStateEventerMockRecorder is the mock recorder for MockPendingStateEventer.
type MockPendingStateEventerMockRecorder struct {
	mock *MockPendingStateEventer
}

// NewMockPendingStateEventer creates a new mock instance.
func NewMockPendingStateEventer(ctrl *gomock.Controller) *MockPendingStateEventer {
	mock := &MockPendingStateEventer{ctrl: ctrl}
	mock.recorder = &MockPendingStateEventerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPendingStateEventer) EXPECT() *MockPendingStateEventerMockRecorder {
	return m.recorder
}

// SubscribePendingTransactions mocks base method.
func (m *MockPendingStateEventer) SubscribePendingTransactions(ctx context.Context, ch chan<- *types.Transaction) (Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribePendingTransactions", ctx, ch)
	ret0, _ := ret[0].(Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribePendingTransactions indicates an expected call of SubscribePendingTransactions.
func (mr *MockPendingStateEventerMockRecorder) SubscribePendingTransactions(ctx, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribePendingTransactions", reflect.TypeOf((*MockPendingStateEventer)(nil).SubscribePendingTransactions), ctx, ch)
}

// MockPeer is a mock of Peer interface.
type MockPeer struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMockRecorder
}

// MockPeerMockRecorder is the mock recorder for MockPeer.
type MockPeerMockRecorder struct {
	mock *MockPeer
}

// NewMockPeer creates a new mock instance.
func NewMockPeer(ctrl *gomock.Controller) *MockPeer {
	mock := &MockPeer{ctrl: ctrl}
	mock.recorder = &MockPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeer) EXPECT() *MockPeerMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockPeer) Send(msgcode uint64, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", msgcode, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockPeerMockRecorder) Send(msgcode, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPeer)(nil).Send), msgcode, data)
}