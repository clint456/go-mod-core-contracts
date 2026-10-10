//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAddNodeService = AddNodeServiceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Service: dtos.NodeService{
		Name:        TestNodeServiceName,
		BaseAddress: TestBaseAddress,
		Labels:      []string{"MODBUS", "TEMP"},
		AdminState:  models.Locked,
		Properties:  make(map[string]any),
	},
}

var testUpdateNodeService = UpdateNodeServiceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Service: mockNodeServiceDTO(),
}

func mockNodeServiceDTO() dtos.UpdateNodeService {
	testUUID := ExampleUUID
	testName := TestNodeServiceName
	testBaseAddress := TestBaseAddress
	testAdminState := models.Locked
	ds := dtos.UpdateNodeService{}
	ds.Id = &testUUID
	ds.Name = &testName
	ds.BaseAddress = &testBaseAddress
	ds.AdminState = &testAdminState
	ds.Labels = testLabels
	return ds
}

func TestAddNodeServiceRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := testAddNodeService
	noReqId := testAddNodeService
	noReqId.RequestId = ""
	invalidReqId := testAddNodeService
	invalidReqId.RequestId = "jfdw324"
	noName := testAddNodeService
	noName.Service.Name = emptyString
	noAdminState := testAddNodeService
	noAdminState.Service.AdminState = emptyString
	invalidAdminState := testAddNodeService
	invalidAdminState.Service.AdminState = "invalid"
	noBaseAddress := testAddNodeService
	noBaseAddress.Service.BaseAddress = emptyString
	invalidBaseAddress := testAddNodeService
	invalidBaseAddress.Service.BaseAddress = "invalid"
	tests := []struct {
		name        string
		NodeService AddNodeServiceRequest
		expectError bool
	}{
		{"valid AddNodeServiceRequest", valid, false},
		{"valid AddNodeServiceRequest, no Request Id", noReqId, false},
		{"invalid AddNodeServiceRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddNodeServiceRequest, no Name", noName, true},
		{"invalid AddNodeServiceRequest, no AdminState", noAdminState, true},
		{"invalid AddNodeServiceRequest, invalid AdminState", invalidAdminState, true},
		{"invalid AddNodeServiceRequest, no BaseAddress", noBaseAddress, true},
		{"invalid AddNodeServiceRequest, no BaseAddress", invalidBaseAddress, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.NodeService.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addNodeServiceRequest validation result.", err)
		})
	}

	serviceNameWithUnreservedChars := testAddNodeService
	serviceNameWithUnreservedChars.Service.Name = nameWithUnreservedChars

	err := serviceNameWithUnreservedChars.Validate()
	assert.NoError(t, err, fmt.Sprintf("AddNodeServiceRequest with service name containing unreserved chars %s should pass validation", nameWithUnreservedChars))

	// Following tests verify if service name containing reserved characters should be detected without an error
	for _, n := range namesWithReservedChar {
		serviceNameWithReservedChar := testAddNodeService
		serviceNameWithReservedChar.Service.Name = n

		err := serviceNameWithReservedChar.Validate()
		assert.NoError(t, err, fmt.Sprintf("AddNodeServiceRequest with service name containing reserved char %s should nout return error during validation", n))
	}
}

func TestAddNodeService_UnmarshalJSON(t *testing.T) {
	valid := testAddNodeService
	resultTestBytes, _ := json.Marshal(testAddNodeService)
	nilNodeServiceProperties := testAddNodeService
	nilNodeServiceProperties.Service.Properties = nil
	bytesNilNodeServiceProperties, _ := json.Marshal(nilNodeServiceProperties)
	type args struct {
		data []byte
	}
	tests := []struct {
		name           string
		addNodeService AddNodeServiceRequest
		args           args
		wantErr        bool
	}{
		{"unmarshal AddNodeServiceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal AddNodeServiceRequest with success, nil Service Properties", valid, args{bytesNilNodeServiceProperties}, false},
		{"unmarshal invalid AddNodeServiceRequest, empty data", AddNodeServiceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddNodeServiceRequest, string data", AddNodeServiceRequest{}, args{[]byte("Invalid AddNodeServiceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addNodeService
			err := tt.addNodeService.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addNodeService, "Unmarshal did not result in expected AddNodeServiceRequest.")
			}
		})
	}
}

func TestAddNodeServiceReqToNodeServiceModels(t *testing.T) {
	requests := []AddNodeServiceRequest{testAddNodeService}
	expectedNodeServiceModel := []models.NodeService{{
		Name:        TestNodeServiceName,
		BaseAddress: TestBaseAddress,
		Labels:      []string{"MODBUS", "TEMP"},
		AdminState:  models.Locked,
		Properties:  make(map[string]any),
	}}
	resultModels := AddNodeServiceReqToNodeServiceModels(requests)
	assert.Equal(t, expectedNodeServiceModel, resultModels, "AddNodeServiceReqToNodeServiceModels did not result in expected NodeService model.")
	for i := range requests {
		requests[i].Service.Properties = nil
	}
	resultModels = AddNodeServiceReqToNodeServiceModels(requests)
	for _, ds := range resultModels {
		assert.NotNil(t, ds.Properties)
	}
}

func TestUpdateNodeService_UnmarshalJSON(t *testing.T) {
	valid := testUpdateNodeService
	resultTestBytes, _ := json.Marshal(testUpdateNodeService)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		req     UpdateNodeServiceRequest
		args    args
		wantErr bool
	}{
		{"unmarshal UpdateNodeServiceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateNodeServiceRequest, empty data", UpdateNodeServiceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateNodeServiceRequest, string data", UpdateNodeServiceRequest{}, args{[]byte("Invalid UpdateNodeServiceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.req
			err := tt.req.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.req, "Unmarshal did not result in expected UpdateNodeServiceRequest.", err)
			}
		})
	}
}

func TestUpdateNodeServiceRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"
	invalidUrl := "http127.0.0.1"

	valid := testUpdateNodeService
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	// id
	validOnlyId := valid
	validOnlyId.Service.Name = nil
	invalidId := valid
	invalidId.Service.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.Service.Id = nil
	invalidEmptyName := valid
	invalidEmptyName.Service.Name = &emptyString
	nameAndEmptyId := valid
	nameAndEmptyId.Service.Id = &emptyString
	// no id and name
	noIdAndName := valid
	noIdAndName.Service.Id = nil
	noIdAndName.Service.Name = nil
	// baseAddress
	validNilBaseAddress := valid
	validNilBaseAddress.Service.BaseAddress = nil
	invalidBaseAddress := valid
	invalidBaseAddress.Service.BaseAddress = &invalidUrl

	invalid := "invalid"
	invalidAdminState := valid
	invalidAdminState.Service.AdminState = &invalid
	tests := []struct {
		name        string
		req         UpdateNodeServiceRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no Request Id", noReqId, false},
		{"invalid, Request Id is not an uuid", invalidReqId, true},

		{"valid, only id", validOnlyId, false},
		{"invalid, invalid Id", invalidId, true},
		{"valid, only name", validOnlyName, false},
		{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},

		{"invalid, no Id and name", noIdAndName, true},

		{"valid, nil baseAddress", validNilBaseAddress, false},
		{"invalid, invalid baseAddress", invalidBaseAddress, true},

		{"invalid, invalid AdminState", invalidAdminState, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateNodeServiceRequest validation result.", err)
		})
	}

	serviceNameWithUnreservedChars := testUpdateNodeService
	serviceNameWithUnreservedChars.Service.Name = &nameWithUnreservedChars

	err := serviceNameWithUnreservedChars.Validate()
	assert.NoError(t, err, fmt.Sprintf("UpdateNodeServiceRequest with service name containing unreserved chars %s should pass validation", nameWithUnreservedChars))

	// Following tests verify if service name containing reserved characters should not be detected with an error
	for i, n := range namesWithReservedChar {
		serviceNameWithReservedChar := testUpdateNodeService
		serviceNameWithReservedChar.Service.Name = &namesWithReservedChar[i]

		err := serviceNameWithReservedChar.Validate()
		assert.NoError(t, err, fmt.Sprintf("UpdateNodeServiceRequest with service name containing reserved char %s should not return error during validation", n))
	}
}

func TestUpdateNodeServiceRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"service":{
			"name":"test node service"
		}
	}`
	var req UpdateNodeServiceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, nil, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Service.BaseAddress)
	assert.Nil(t, req.Service.AdminState)
	assert.Nil(t, req.Service.Labels)
}

func TestUpdateNodeServiceRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"service":{
			"apiVersion":"v3",
			"name":"TestNode",
			"labels":[]
		}
	}`
	var req UpdateNodeServiceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.Service.Labels)
}

func TestReplaceNodeServiceModelFieldsWithDTO(t *testing.T) {
	properties := testProperties
	ds := models.NodeService{
		Id:         "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name:       "test node service",
		Properties: properties,
	}
	patch := mockNodeServiceDTO()

	ReplaceNodeServiceModelFieldsWithDTO(&ds, patch)

	assert.Equal(t, TestBaseAddress, ds.BaseAddress)
	assert.Equal(t, models.Locked, string(ds.AdminState))
	assert.Equal(t, testLabels, ds.Labels)
	assert.Equal(t, properties, ds.Properties)
}

func TestNewAddNodeServiceRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewAddNodeServiceRequest(dtos.NodeService{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}

func TestNewUpdateNodeServiceRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewUpdateNodeServiceRequest(dtos.UpdateNodeService{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}
