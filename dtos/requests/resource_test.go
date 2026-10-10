//
// Copyright (C) 2022-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testResource = AddResourceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestProfileName,
	Resource: dtos.Resource{
		Name:        TestCommandName,
		Description: TestDescription,
		Attributes:  testAttributes,
		Properties: dtos.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_RW,
		},
	},
}

var testUpdateResource = UpdateResourceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestProfileName,
	Resource:    mockUpdateResourceDTO(),
}

func mockUpdateResourceDTO() dtos.UpdateResource {
	testName := TestCommandName
	testIsHidden := true
	testDescription := TestDescription

	dr := dtos.UpdateResource{}
	dr.Name = &testName
	dr.IsHidden = &testIsHidden
	dr.Description = &testDescription

	return dr
}

func TestAddResourceRequest_Validate(t *testing.T) {
	valid := testResource
	noProfileName := testResource
	noProfileName.ProfileName = emptyString
	noResourceName := testResource
	noResourceName.Resource.Name = emptyString

	tests := []struct {
		name        string
		request     AddResourceRequest
		expectedErr bool
	}{
		{"valid AddResourceRequest", valid, false},
		{"invalid AddResourceRequest, no ProfileName", noProfileName, true},
		{"invalid AddResourceRequest, no Resource Name", noResourceName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddResourceRequest_UnmarshalJSON(t *testing.T) {
	valid := testResource
	resultTestBytes, _ := json.Marshal(testResource)
	type args struct {
		data []byte
	}

	tests := []struct {
		name        string
		request     AddResourceRequest
		args        args
		expectedErr bool
	}{
		{"unmarshal AddResourceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddResourceRequest, empty data", AddResourceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddResourceRequest, string data", AddResourceRequest{}, args{[]byte("Invalid AddResourceRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected AddResourceRequest.")
			}
		})
	}
}

func TestUpdateResourceRequest_Validate(t *testing.T) {
	valid := testUpdateResource
	noProfileName := testUpdateResource
	noProfileName.ProfileName = emptyString
	noResourceName := testUpdateResource
	noResourceName.Resource.Name = &emptyString

	tests := []struct {
		name        string
		request     UpdateResourceRequest
		expectedErr bool
	}{
		{"valid UpdateResourceRequest", valid, false},
		{"invalid UpdateResourceRequest, no ProfileName", noProfileName, true},
		{"invalid UpdateResourceRequest, no Resource Name", noResourceName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateResourceRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
        "profileName": "TestProfile",
		"resource":{"name":"TestResource"}
	}`
	var req UpdateResourceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Resource.Description)
	assert.Nil(t, req.Resource.IsHidden)
}

func TestReplaceResourceModelFieldsWithDTO(t *testing.T) {
	resource := models.Resource{
		Name:        TestResourceName,
		Description: emptyString,
		Attributes:  testAttributes,
		Properties: models.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_R,
		},
		Tags: testTags,
	}

	patch := mockUpdateResourceDTO()

	ReplaceResourceModelFieldsWithDTO(&resource, patch)

	assert.Equal(t, TestDescription, resource.Description)
	assert.Equal(t, true, resource.IsHidden)
	assert.Equal(t, testAttributes, resource.Attributes)
	assert.Equal(t, common.ReadWrite_R, resource.Properties.ReadWrite)
	assert.Equal(t, testTags, resource.Tags)
}
