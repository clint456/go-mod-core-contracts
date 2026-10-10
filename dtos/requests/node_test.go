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

var testNodeLabels = []string{"MODBUS", "TEMP"}
var testNodeLocation = "{40lat;45long}"
var testAutoEvents = []dtos.AutoEvent{
	{SourceName: "TestNode", Interval: "300ms", OnChange: true, OnChangeThreshold: 0.01},
}
var testAutoEventsWithInvalidFrequency = []dtos.AutoEvent{
	{SourceName: "TestNode", Interval: "300", OnChange: true},
}
var testProtocols = map[string]dtos.ProtocolProperties{
	"modbus-ip": {
		"Address": "localhost",
		"Port":    "1502",
		"UnitID":  "1",
	},
}

var testProperties = map[string]any{
	"LastScan": float64(1702275547),
}

var testParent = "ParentNode"

var testAddNode = AddNodeRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Node: dtos.Node{
		Name:           TestNodeName,
		ServiceName:    TestNodeServiceName,
		ProfileName:    TestProfileName,
		AdminState:     models.Locked,
		OperatingState: models.Up,
		Labels:         testNodeLabels,
		Location:       testNodeLocation,
		Tags:           testTags,
		AutoEvents:     testAutoEvents,
		Protocols:      testProtocols,
		Parent:         testParent,
		Properties:     make(map[string]any),
	},
}

var testUpdateNode = UpdateNodeRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Node: mockUpdateNode(),
}

func mockUpdateNode() dtos.UpdateNode {
	testId := ExampleUUID
	testName := TestNodeName
	testDescription := TestDescription
	testAdminState := models.Locked
	testOperatingState := models.Up
	testNodeServiceName := TestNodeServiceName
	testProfileName := TestProfileName
	d := dtos.UpdateNode{}
	d.Id = &testId
	d.Name = &testName
	d.Description = &testDescription
	d.AdminState = &testAdminState
	d.OperatingState = &testOperatingState
	d.ServiceName = &testNodeServiceName
	d.ProfileName = &testProfileName
	d.Labels = testNodeLabels
	d.Location = testNodeLocation
	d.AutoEvents = testAutoEvents
	d.Protocols = testProtocols
	d.Tags = testTags
	d.Properties = testProperties
	d.Parent = &testParent
	return d
}

func TestAddNodeRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := testAddNode
	invalidFrequency := testAddNode
	invalidFrequency.Node.AutoEvents = testAutoEventsWithInvalidFrequency
	noReqId := testAddNode
	noReqId.RequestId = ""
	invalidReqId := testAddNode
	invalidReqId.RequestId = "abc"
	noNodeName := testAddNode
	noNodeName.Node.Name = emptyString
	noServiceName := testAddNode
	noServiceName.Node.ServiceName = emptyString
	noProfileName := testAddNode
	noProfileName.Node.ProfileName = emptyString
	noProtocols := testAddNode
	noProtocols.Node.Protocols = map[string]dtos.ProtocolProperties{}
	noAutoEventFrequency := testAddNode
	noAutoEventFrequency.Node.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestNode", OnChange: true},
	}
	noAutoEventResource := testAddNode
	noAutoEventResource.Node.AutoEvents = []dtos.AutoEvent{
		{Interval: "300ms", OnChange: true},
	}
	invalidAutoEventOnChangeThreshold := testAddNode
	invalidAutoEventOnChangeThreshold.Node.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestNode", Interval: "300ms", OnChange: true, OnChangeThreshold: -0.01},
	}
	tests := []struct {
		name        string
		Node        AddNodeRequest
		expectError bool
	}{
		{"valid AddNodeRequest", valid, false},
		{"invalid AddNodeRequest, invalid autoEvent frequency", invalidFrequency, true},
		{"valid AddNodeRequest, no Request Id", noReqId, false},
		{"invalid AddNodeRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddNodeRequest, no NodeName", noNodeName, true},
		{"invalid AddNodeRequest, no ServiceName", noServiceName, true},
		{"valid AddNodeRequest, no ProfileName", noProfileName, false},
		{"valid AddNodeRequest, no Protocols", noProtocols, false},
		{"invalid AddNodeRequest, no AutoEvent frequency", noAutoEventFrequency, true},
		{"invalid AddNodeRequest, no AutoEvent resource", noAutoEventResource, true},
		{"invalid AddNodeRequest, OnChangeThreshold should greater than or equal 0", invalidAutoEventOnChangeThreshold, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Node.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addNodeRequest validation result: %v", err)
		})
	}

	type testForNameField struct {
		name        string
		Node        AddNodeRequest
		expectError bool
	}

	nodeNameWithUnreservedChars := testAddNode
	nodeNameWithUnreservedChars.Node.Name = nameWithUnreservedChars
	profileNameWithUnreservedChars := testAddNode
	profileNameWithUnreservedChars.Node.ProfileName = nameWithUnreservedChars
	serviceNameWithUnreservedChars := testAddNode
	serviceNameWithUnreservedChars.Node.ServiceName = nameWithUnreservedChars
	parentNameWithUnreservedChars := testAddNode
	parentNameWithUnreservedChars.Node.Parent = nameWithUnreservedChars

	// Following tests verify if name fields containing unreserved characters should pass edgex-dto-rfc3986-unreserved-chars check
	testsForNameFields := []testForNameField{
		{"Valid AddNodeRequest with node name containing unreserved chars", nodeNameWithUnreservedChars, false},
		{"Valid AddNodeRequest with profile name containing unreserved chars", profileNameWithUnreservedChars, false},
		{"Valid AddNodeRequest with service name containing unreserved chars", serviceNameWithUnreservedChars, false},
		{"Valid AddNodeRequest with parent name containing unreserved chars", parentNameWithUnreservedChars, false},
	}

	// Following tests verify if name fields containing reserved characters should not be detected with an error
	for _, n := range namesWithReservedChar {
		nodeNameWithReservedChar := testAddNode
		nodeNameWithReservedChar.Node.Name = n
		profileNameWithReservedChar := testAddNode
		profileNameWithReservedChar.Node.ProfileName = n
		serviceNameWithReservedChar := testAddNode
		serviceNameWithReservedChar.Node.ServiceName = n
		parentNameWithReservedChar := testAddNode
		parentNameWithReservedChar.Node.Parent = n

		testsForNameFields = append(testsForNameFields,
			testForNameField{"Valid AddNodeRequest with node name containing reserved char", nodeNameWithReservedChar, false},
			testForNameField{"Valid AddNodeRequest with profile name containing reserved char", profileNameWithReservedChar, false},
			testForNameField{"Valid AddNodeRequest with service name containing reserved char", serviceNameWithReservedChar, false},
			testForNameField{"Valid AddNodeRequest with parent name containing reserved char", parentNameWithReservedChar, false},
		)
	}

	for _, tt := range testsForNameFields {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Node.Validate()
			if tt.expectError {
				assert.Error(t, err, fmt.Sprintf("expect error but not : %s", tt.name))
			} else {
				assert.NoError(t, err, fmt.Sprintf("unexpected error occurs : %s", tt.name))
			}
		})
	}
}

func TestAddNode_UnmarshalJSON(t *testing.T) {
	valid := testAddNode
	resultTestBytes, _ := json.Marshal(testAddNode)
	nilNodeProperties := testAddNode
	nilNodeProperties.Node.Properties = nil
	bytesNilNodeProperties, _ := json.Marshal(nilNodeProperties)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		addNode AddNodeRequest
		args    args
		wantErr bool
	}{
		{"unmarshal AddNodeRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal AddNodeRequest with success, nil Node Properties", valid, args{bytesNilNodeProperties}, false},
		{"unmarshal invalid AddNodeRequest, empty data", AddNodeRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddNodeRequest, string data", AddNodeRequest{}, args{[]byte("Invalid AddNodeRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addNode
			err := tt.addNode.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addNode, "Unmarshal did not result in expected AddNodeRequest.")
			}
		})
	}
}

func Test_AddNodeReqToNodeModels(t *testing.T) {
	requests := []AddNodeRequest{testAddNode}
	expectedNodeModel := []models.Node{
		{
			Name:           TestNodeName,
			ServiceName:    TestNodeServiceName,
			ProfileName:    TestProfileName,
			AdminState:     models.Locked,
			OperatingState: models.Up,
			Labels:         testNodeLabels,
			Location:       testNodeLocation,
			Tags:           testTags,
			AutoEvents: []models.AutoEvent{
				{SourceName: "TestNode", Interval: "300ms", OnChange: true, OnChangeThreshold: 0.01},
			},
			Protocols: map[string]models.ProtocolProperties{
				"modbus-ip": {
					"Address": "localhost",
					"Port":    "1502",
					"UnitID":  "1",
				},
			},
			Parent:     testParent,
			Properties: make(map[string]any),
		},
	}
	resultModels := AddNodeReqToNodeModels(requests)
	assert.Equal(t, expectedNodeModel, resultModels, "AddNodeReqToNodeModels did not result in expected Node model.")
	for i := range requests {
		requests[i].Node.Properties = nil
	}
	resultModels = AddNodeReqToNodeModels(requests)
	for _, m := range resultModels {
		assert.NotNil(t, m.Properties)
	}
}

func TestUpdateNodeRequest_UnmarshalJSON(t *testing.T) {
	valid := testUpdateNode
	resultTestBytes, _ := json.Marshal(testUpdateNode)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		req     UpdateNodeRequest
		args    args
		wantErr bool
	}{
		{"unmarshal UpdateNodeRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateNodeRequest, empty data", UpdateNodeRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateNodeRequest, string data", UpdateNodeRequest{}, args{[]byte("Invalid UpdateNodeRequest")}, true},
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
				assert.Equal(t, expected, tt.req, "Unmarshal did not result in expected UpdateNodeRequest.", err)
			}
		})
	}
}

func TestUpdateNodeRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"

	valid := testUpdateNode
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	// id
	validOnlyId := valid
	validOnlyId.Node.Name = nil
	invalidId := valid
	invalidId.Node.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.Node.Id = nil
	nameAndEmptyId := valid
	nameAndEmptyId.Node.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.Node.Name = &emptyString
	// no id and name
	noIdAndName := valid
	noIdAndName.Node.Id = nil
	noIdAndName.Node.Name = nil
	// description
	validNilDescription := valid
	validNilDescription.Node.Description = nil
	invalidEmptyDescription := valid
	invalidEmptyDescription.Node.Description = &emptyString
	// ServiceName
	validNilServiceName := valid
	validNilServiceName.Node.ServiceName = nil
	invalidEmptyServiceName := valid
	invalidEmptyServiceName.Node.ServiceName = &emptyString
	// ProfileName
	validNilProfileName := valid
	validNilProfileName.Node.ProfileName = nil
	emptyProfileName := valid
	emptyProfileName.Node.ProfileName = &emptyString

	invalidState := "invalid state"
	invalidAdminState := valid
	invalidAdminState.Node.AdminState = &invalidState
	invalidOperatingState := valid
	invalidOperatingState.Node.OperatingState = &invalidState
	invalidFrequency := valid
	invalidFrequency.Node.AutoEvents = testAutoEventsWithInvalidFrequency
	emptyProtocols := valid
	emptyProtocols.Node.Protocols = map[string]dtos.ProtocolProperties{}

	tests := []struct {
		name        string
		req         UpdateNodeRequest
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

		{"valid, nil description", validNilDescription, false},

		{"valid, nil service name", validNilServiceName, false},
		{"invalid, empty service name", invalidEmptyServiceName, true},

		{"valid, nil profile name", validNilProfileName, false},
		{"valid, empty profile name", emptyProfileName, false},

		{"invalid, invalid admin state", invalidAdminState, true},
		{"invalid, invalid operating state", invalidOperatingState, true},
		{"invalid, invalid autoEvent frequency", invalidFrequency, true},

		{"valid, empty protocols", emptyProtocols, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateNodeRequest validation result.", err)
		})
	}
}

func TestUpdateNodeRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"node":{"apiVersion":"v3", "name":"TestNode"}
	}`
	var req UpdateNodeRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Node.Description)
	assert.Nil(t, req.Node.AdminState)
	assert.Nil(t, req.Node.OperatingState)
	assert.Nil(t, req.Node.ServiceName)
	assert.Nil(t, req.Node.ProfileName)
	assert.Nil(t, req.Node.Labels)
	assert.Nil(t, req.Node.Location)
	assert.Nil(t, req.Node.AutoEvents)
	assert.Nil(t, req.Node.Protocols)
	assert.Nil(t, req.Node.Tags)
	assert.Nil(t, req.Node.Properties)
}

func TestUpdateNodeRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"node":{
			"apiVersion":"v3",
			"name":"TestNode",
			"labels":[],
			"autoEvents":[]
		}
	}`
	var req UpdateNodeRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.Node.Labels)
	assert.NotNil(t, req.Node.AutoEvents)
}

func TestReplaceNodeModelFieldsWithDTO(t *testing.T) {
	node := models.Node{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "test node profile",
	}
	patch := mockUpdateNode()

	ReplaceNodeModelFieldsWithDTO(&node, patch)

	assert.Equal(t, TestDescription, node.Description)
	assert.Equal(t, models.AdminState(models.Locked), node.AdminState)
	assert.Equal(t, models.OperatingState(models.Up), node.OperatingState)
	assert.Equal(t, TestNodeServiceName, node.ServiceName)
	assert.Equal(t, TestProfileName, node.ProfileName)
	assert.Equal(t, testLabels, node.Labels)
	assert.Equal(t, testNodeLocation, node.Location)
	assert.Equal(t, testTags, node.Tags)
	assert.Equal(t, dtos.ToAutoEventModels(testAutoEvents), node.AutoEvents)
	assert.Equal(t, dtos.ToProtocolModels(testProtocols), node.Protocols)
	assert.Equal(t, testProperties, node.Properties)
}

func TestNewAddNodeRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewAddNodeRequest(dtos.Node{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}

func TestNewUpdateNodeRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewUpdateNodeRequest(dtos.UpdateNode{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}
