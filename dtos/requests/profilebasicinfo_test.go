//
// Copyright (C) 2022-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

var testBasicInfoRequest = ProfileBasicInfoRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	BasicInfo: mockUpdateProfileBasicInfo(),
}

func mockUpdateProfileBasicInfo() dtos.UpdateProfileBasicInfo {
	testId := ExampleUUID
	testName := TestNodeName
	testDescription := TestDescription
	testManufacturer := TestManufacturer
	TestModel := TestModel
	d := dtos.UpdateProfileBasicInfo{}
	d.Id = &testId
	d.Name = &testName
	d.Description = &testDescription
	d.Manufacturer = &testManufacturer
	d.Model = &TestModel
	d.Labels = testLabels
	return d
}
func TestProfileBasicInfoRequest_Validate(t *testing.T) {
	invalidUUID := "invalidUUID"
	valid := testBasicInfoRequest
	// id
	validOnlyId := valid
	validOnlyId.BasicInfo.Name = nil
	invalidId := valid
	invalidId.BasicInfo.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.BasicInfo.Id = nil
	nameAndEmptyId := valid
	nameAndEmptyId.BasicInfo.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.BasicInfo.Name = &emptyString
	// no id and name
	noIdAndName := valid
	noIdAndName.BasicInfo.Id = nil
	noIdAndName.BasicInfo.Name = nil
	// description
	validNilDescription := valid
	validNilDescription.BasicInfo.Description = nil
	validEmptyDescription := valid
	validEmptyDescription.BasicInfo.Description = &emptyString
	// manufacturer
	validNilManufacturer := valid
	validNilManufacturer.BasicInfo.Manufacturer = nil
	validEmptyManufacturer := valid
	validEmptyManufacturer.BasicInfo.Manufacturer = &emptyString
	// model
	validNilModel := valid
	validNilModel.BasicInfo.Model = nil
	validEmptyModel := valid
	validEmptyModel.BasicInfo.Model = &emptyString
	// labels
	validNilLabels := valid
	validNilLabels.BasicInfo.Labels = nil
	validEmptyLabels := valid
	validEmptyLabels.BasicInfo.Labels = []string{}

	tests := []struct {
		name        string
		request     ProfileBasicInfoRequest
		expectedErr bool
	}{
		{"valid ProfileBasicInfoRequest", valid, false},
		{"valid ProfileBasicInfoRequest, only id", validOnlyId, false},
		{"invalid ProfileBasicInfoRequest, invalid id ", invalidId, true},
		{"valid ProfileBasicInfoRequest, only name", validOnlyName, false},
		{"valid ProfileBasicInfoRequest, name and empty id", nameAndEmptyId, false},
		{"invalid ProfileBasicInfoRequest, empty name", invalidEmptyName, true},
		{"invalid ProfileBasicInfoRequest, no name and no id ", noIdAndName, true},
		{"valid ProfileBasicInfoRequest, nil description", validNilDescription, false},
		{"valid ProfileBasicInfoRequest, empty description", validEmptyDescription, false},
		{"valid ProfileBasicInfoRequest, nil manufacturer", validNilManufacturer, false},
		{"valid ProfileBasicInfoRequest, empty manufacturer", validEmptyManufacturer, false},
		{"valid ProfileBasicInfoRequest, nil model", validNilModel, false},
		{"valid ProfileBasicInfoRequest, empty model", validEmptyModel, false},
		{"valid ProfileBasicInfoRequest, nil labels", validNilLabels, false},
		{"valid ProfileBasicInfoRequest, empty lables", validEmptyLabels, false},
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

func TestProfileBasicInfoRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"basicinfo": {
          "name": "TestProfile"
		}
	}`
	var req ProfileBasicInfoRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.BasicInfo.Manufacturer)
	assert.Nil(t, req.BasicInfo.Description)
	assert.Nil(t, req.BasicInfo.Model)
	assert.Nil(t, req.BasicInfo.Labels)
}

func TestProfileBasicInfoRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"basicinfo": {
          "name": "TestProfile",
		  "labels":[]
		}
	}`
	var req ProfileBasicInfoRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.BasicInfo.Labels)
}

func TestReplaceProfileModelBasicInfoFieldsWithDTO(t *testing.T) {
	profile := models.Profile{
		Name:         TestProfileName,
		Manufacturer: "",
		Description:  "",
		Model:        "",
		Labels:       []string{},
		Resources: []models.Resource{{
			Name: TestResourceName,
			Properties: models.ResourceProperties{
				ValueType: common.ValueTypeInt16,
				ReadWrite: common.ReadWrite_RW,
			},
		}},
	}

	patch := mockUpdateProfileBasicInfo()

	ReplaceProfileModelBasicInfoFieldsWithDTO(&profile, patch)

	assert.Equal(t, TestDescription, profile.Description)
	assert.Equal(t, TestManufacturer, profile.Manufacturer)
	assert.Equal(t, TestModel, profile.Model)
	assert.Equal(t, testLabels, profile.Labels)
}
