// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/server/catalog (interfaces: Catalog)

// Package mock_catalog is a generated GoMock package.
package mock_catalog

import (
	gomock "github.com/golang/mock/gomock"
	catalog "github.com/spiffe/spire/pkg/common/catalog"
	ca "github.com/spiffe/spire/proto/server/ca"
	datastore "github.com/spiffe/spire/proto/server/datastore"
	nodeattestor "github.com/spiffe/spire/proto/server/nodeattestor"
	noderesolver "github.com/spiffe/spire/proto/server/noderesolver"
	upstreamca "github.com/spiffe/spire/proto/server/upstreamca"
	reflect "reflect"
)

// MockCatalog is a mock of Catalog interface
type MockCatalog struct {
	ctrl     *gomock.Controller
	recorder *MockCatalogMockRecorder
}

// MockCatalogMockRecorder is the mock recorder for MockCatalog
type MockCatalogMockRecorder struct {
	mock *MockCatalog
}

// NewMockCatalog creates a new mock instance
func NewMockCatalog(ctrl *gomock.Controller) *MockCatalog {
	mock := &MockCatalog{ctrl: ctrl}
	mock.recorder = &MockCatalogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCatalog) EXPECT() *MockCatalogMockRecorder {
	return m.recorder
}

// CAs mocks base method
func (m *MockCatalog) CAs() []ca.ControlPlaneCa {
	ret := m.ctrl.Call(m, "CAs")
	ret0, _ := ret[0].([]ca.ControlPlaneCa)
	return ret0
}

// CAs indicates an expected call of CAs
func (mr *MockCatalogMockRecorder) CAs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CAs", reflect.TypeOf((*MockCatalog)(nil).CAs))
}

// DataStores mocks base method
func (m *MockCatalog) DataStores() []datastore.DataStore {
	ret := m.ctrl.Call(m, "DataStores")
	ret0, _ := ret[0].([]datastore.DataStore)
	return ret0
}

// DataStores indicates an expected call of DataStores
func (mr *MockCatalogMockRecorder) DataStores() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataStores", reflect.TypeOf((*MockCatalog)(nil).DataStores))
}

// NodeAttestors mocks base method
func (m *MockCatalog) NodeAttestors() []nodeattestor.NodeAttestor {
	ret := m.ctrl.Call(m, "NodeAttestors")
	ret0, _ := ret[0].([]nodeattestor.NodeAttestor)
	return ret0
}

// NodeAttestors indicates an expected call of NodeAttestors
func (mr *MockCatalogMockRecorder) NodeAttestors() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeAttestors", reflect.TypeOf((*MockCatalog)(nil).NodeAttestors))
}

// NodeResolvers mocks base method
func (m *MockCatalog) NodeResolvers() []noderesolver.NodeResolver {
	ret := m.ctrl.Call(m, "NodeResolvers")
	ret0, _ := ret[0].([]noderesolver.NodeResolver)
	return ret0
}

// NodeResolvers indicates an expected call of NodeResolvers
func (mr *MockCatalogMockRecorder) NodeResolvers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeResolvers", reflect.TypeOf((*MockCatalog)(nil).NodeResolvers))
}

// Plugins mocks base method
func (m *MockCatalog) Plugins() []*catalog.ManagedPlugin {
	ret := m.ctrl.Call(m, "Plugins")
	ret0, _ := ret[0].([]*catalog.ManagedPlugin)
	return ret0
}

// Plugins indicates an expected call of Plugins
func (mr *MockCatalogMockRecorder) Plugins() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Plugins", reflect.TypeOf((*MockCatalog)(nil).Plugins))
}

// Reload mocks base method
func (m *MockCatalog) Reload() error {
	ret := m.ctrl.Call(m, "Reload")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload
func (mr *MockCatalogMockRecorder) Reload() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockCatalog)(nil).Reload))
}

// Run mocks base method
func (m *MockCatalog) Run() error {
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockCatalogMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCatalog)(nil).Run))
}

// Stop mocks base method
func (m *MockCatalog) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockCatalogMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCatalog)(nil).Stop))
}

// UpstreamCAs mocks base method
func (m *MockCatalog) UpstreamCAs() []upstreamca.UpstreamCa {
	ret := m.ctrl.Call(m, "UpstreamCAs")
	ret0, _ := ret[0].([]upstreamca.UpstreamCa)
	return ret0
}

// UpstreamCAs indicates an expected call of UpstreamCAs
func (mr *MockCatalogMockRecorder) UpstreamCAs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpstreamCAs", reflect.TypeOf((*MockCatalog)(nil).UpstreamCAs))
}
