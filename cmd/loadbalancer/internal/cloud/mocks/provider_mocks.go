// /*
// Copyright ApeCloud, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apecloud/kubeblocks/cmd/loadbalancer/internal/cloud (interfaces: Provider)

// Package mock_cloud is a generated GoMock package.
package mock_cloud

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	cloud "github.com/apecloud/kubeblocks/cmd/loadbalancer/internal/cloud"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AllocIPAddresses mocks base method.
func (m *MockProvider) AllocIPAddresses(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocIPAddresses", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocIPAddresses indicates an expected call of AllocIPAddresses.
func (mr *MockProviderMockRecorder) AllocIPAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocIPAddresses", reflect.TypeOf((*MockProvider)(nil).AllocIPAddresses), arg0)
}

// AssignPrivateIPAddresses mocks base method.
func (m *MockProvider) AssignPrivateIPAddresses(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignPrivateIPAddresses", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignPrivateIPAddresses indicates an expected call of AssignPrivateIPAddresses.
func (mr *MockProviderMockRecorder) AssignPrivateIPAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignPrivateIPAddresses", reflect.TypeOf((*MockProvider)(nil).AssignPrivateIPAddresses), arg0, arg1)
}

// AttachENI mocks base method.
func (m *MockProvider) AttachENI(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachENI", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachENI indicates an expected call of AttachENI.
func (mr *MockProviderMockRecorder) AttachENI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachENI", reflect.TypeOf((*MockProvider)(nil).AttachENI), arg0, arg1)
}

// CreateENI mocks base method.
func (m *MockProvider) CreateENI(arg0, arg1 string, arg2 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateENI", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateENI indicates an expected call of CreateENI.
func (mr *MockProviderMockRecorder) CreateENI(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateENI", reflect.TypeOf((*MockProvider)(nil).CreateENI), arg0, arg1, arg2)
}

// DeallocIPAddresses mocks base method.
func (m *MockProvider) DeallocIPAddresses(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeallocIPAddresses", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeallocIPAddresses indicates an expected call of DeallocIPAddresses.
func (mr *MockProviderMockRecorder) DeallocIPAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeallocIPAddresses", reflect.TypeOf((*MockProvider)(nil).DeallocIPAddresses), arg0, arg1)
}

// DeleteENI mocks base method.
func (m *MockProvider) DeleteENI(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteENI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteENI indicates an expected call of DeleteENI.
func (mr *MockProviderMockRecorder) DeleteENI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteENI", reflect.TypeOf((*MockProvider)(nil).DeleteENI), arg0)
}

// DescribeAllENIs mocks base method.
func (m *MockProvider) DescribeAllENIs() (map[string]*cloud.ENIMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAllENIs")
	ret0, _ := ret[0].(map[string]*cloud.ENIMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAllENIs indicates an expected call of DescribeAllENIs.
func (mr *MockProviderMockRecorder) DescribeAllENIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAllENIs", reflect.TypeOf((*MockProvider)(nil).DescribeAllENIs))
}

// FindLeakedENIs mocks base method.
func (m *MockProvider) FindLeakedENIs(arg0 string) ([]*cloud.ENIMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLeakedENIs", arg0)
	ret0, _ := ret[0].([]*cloud.ENIMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLeakedENIs indicates an expected call of FindLeakedENIs.
func (mr *MockProviderMockRecorder) FindLeakedENIs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLeakedENIs", reflect.TypeOf((*MockProvider)(nil).FindLeakedENIs), arg0)
}

// FreeENI mocks base method.
func (m *MockProvider) FreeENI(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreeENI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FreeENI indicates an expected call of FreeENI.
func (mr *MockProviderMockRecorder) FreeENI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeENI", reflect.TypeOf((*MockProvider)(nil).FreeENI), arg0)
}

// GetENIIPv4Limit mocks base method.
func (m *MockProvider) GetENIIPv4Limit() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetENIIPv4Limit")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetENIIPv4Limit indicates an expected call of GetENIIPv4Limit.
func (mr *MockProviderMockRecorder) GetENIIPv4Limit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetENIIPv4Limit", reflect.TypeOf((*MockProvider)(nil).GetENIIPv4Limit))
}

// GetENILimit mocks base method.
func (m *MockProvider) GetENILimit() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetENILimit")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetENILimit indicates an expected call of GetENILimit.
func (mr *MockProviderMockRecorder) GetENILimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetENILimit", reflect.TypeOf((*MockProvider)(nil).GetENILimit))
}

// GetInstanceInfo mocks base method.
func (m *MockProvider) GetInstanceInfo() *cloud.InstanceInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceInfo")
	ret0, _ := ret[0].(*cloud.InstanceInfo)
	return ret0
}

// GetInstanceInfo indicates an expected call of GetInstanceInfo.
func (mr *MockProviderMockRecorder) GetInstanceInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceInfo", reflect.TypeOf((*MockProvider)(nil).GetInstanceInfo))
}

// ModifySourceDestCheck mocks base method.
func (m *MockProvider) ModifySourceDestCheck(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifySourceDestCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifySourceDestCheck indicates an expected call of ModifySourceDestCheck.
func (mr *MockProviderMockRecorder) ModifySourceDestCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifySourceDestCheck", reflect.TypeOf((*MockProvider)(nil).ModifySourceDestCheck), arg0, arg1)
}

// WaitForENIAttached mocks base method.
func (m *MockProvider) WaitForENIAttached(arg0 string) (cloud.ENIMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForENIAttached", arg0)
	ret0, _ := ret[0].(cloud.ENIMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForENIAttached indicates an expected call of WaitForENIAttached.
func (mr *MockProviderMockRecorder) WaitForENIAttached(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForENIAttached", reflect.TypeOf((*MockProvider)(nil).WaitForENIAttached), arg0)
}