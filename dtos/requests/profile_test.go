//
// Copyright (C) 2020-2023 IOTech Ltd
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

var testLabels = []string{"MODBUS", "TEMP"}
var testAttributes = map[string]interface{}{
	"TestAttribute": "TestAttributeValue",
}
var testTags = map[string]any{
	"TestTagsKey": "TestTagsValue",
}

func profileData() ProfileRequest {
	var testResources = []dtos.Resource{{
		Name:        TestResourceName,
		Description: TestDescription,
		Attributes:  testAttributes,
		Properties: dtos.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_RW,
		},
		Tags: testTags,
	}}
	var testCommands = []dtos.Command{{
		Name:      TestCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []dtos.ResourceOperation{{
			Resource: TestResourceName,
		}},
		Tags: testTags,
	}}
	return ProfileRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		Profile: dtos.Profile{
			ApiVersion: common.ApiVersion,
			ProfileBasicInfo: dtos.ProfileBasicInfo{
				Name:         TestProfileName,
				Manufacturer: TestManufacturer,
				Description:  TestDescription,
				Model:        TestModel,
				Labels:       testLabels,
			},
			Resources: testResources,
			Commands:  testCommands,
		},
	}
}

var expectedProfile = models.Profile{
	ApiVersion:   common.ApiVersion,
	Name:         TestProfileName,
	Manufacturer: TestManufacturer,
	Description:  TestDescription,
	Model:        TestModel,
	Labels:       testLabels,
	Resources: []models.Resource{{
		Name:        TestResourceName,
		Description: TestDescription,
		Attributes:  testAttributes,
		Properties: models.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_RW,
		},
		Tags: testTags,
	}},
	Commands: []models.Command{{
		Name:      TestCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []models.ResourceOperation{{
			Resource: TestResourceName,
		}},
		Tags: testTags,
	}},
}

func TestProfileRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := profileData()
	noName := profileData()
	noName.Profile.Name = emptyString
	noResourceAndCommand := profileData()
	noResourceAndCommand.Profile.Resources = []dtos.Resource{}
	noResourceAndCommand.Profile.Commands = []dtos.Command{}
	noResourceName := profileData()
	noResourceName.Profile.Resources[0].Name = emptyString
	noResourcePropertyType := profileData()
	noResourcePropertyType.Profile.Resources[0].Properties.ValueType = emptyString
	invalidResourcePropertyType := profileData()
	invalidResourcePropertyType.Profile.Resources[0].Properties.ValueType = "BadType"
	noResourceReadWrite := profileData()
	noResourceReadWrite.Profile.Resources[0].Properties.ReadWrite = emptyString
	invalidResourceReadWrite := profileData()
	invalidResourceReadWrite.Profile.Resources[0].Properties.ReadWrite = "invalid"
	validResourceReadWrite := profileData()
	validResourceReadWrite.Profile.Resources[0].Properties.ReadWrite = common.ReadWrite_WR
	dismatchCommand := profileData()
	dismatchCommand.Profile.Resources = []dtos.Resource{}

	tests := []struct {
		name        string
		Profile     ProfileRequest
		expectError bool
	}{
		{"valid ProfileRequest", valid, false},
		{"invalid ProfileRequest, no name", noName, true},
		{"valid ProfileRequest, no resource and no command", noResourceAndCommand, false},
		{"invalid ProfileRequest, no resource name", noResourceName, true},
		{"invalid ProfileRequest, no resource property type", noResourcePropertyType, true},
		{"invalid ProfileRequest, invalid resource property type", invalidResourcePropertyType, true},
		{"invalid ProfileRequest, no resource property readWrite", noResourceReadWrite, true},
		{"invalid ProfileRequest, invalid resource property readWrite", invalidResourcePropertyType, true},
		{"valid ProfileRequest, valid resource property readWrite with value WR", validResourceReadWrite, false},
		{"invalid ProfileRequest, dismatch command", dismatchCommand, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Profile.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

	profileNameWithUnreservedChars := profileData()
	profileNameWithUnreservedChars.Profile.Name = nameWithUnreservedChars

	err := profileNameWithUnreservedChars.Validate()
	assert.NoError(t, err, fmt.Sprintf("ProfileRequest with profile name containing unreserved chars %s should pass validation", nameWithUnreservedChars))

	// Following tests verify if profile name containing reserved characters should not be detected with an error
	for _, n := range namesWithReservedChar {
		profileNameWithReservedChar := profileData()
		profileNameWithReservedChar.Profile.Name = n

		err := profileNameWithReservedChar.Validate()
		assert.NoError(t, err, fmt.Sprintf("ProfileRequest with profile name containing reserved char %s should not return error during validation", n))
	}
}

func TestAddProfile_UnmarshalJSON(t *testing.T) {
	expected := profileData()
	validData, err := json.Marshal(profileData())
	require.NoError(t, err)
	validValueTypeLowerCase := profileData()
	validValueTypeLowerCase.Profile.Resources[0].Properties.ValueType = "int16"
	validValueTypeLowerCaseData, err := json.Marshal(validValueTypeLowerCase)
	require.NoError(t, err)
	validValueTypeUpperCase := profileData()
	validValueTypeUpperCase.Profile.Resources[0].Properties.ValueType = "INT16"
	validValueTypeUpperCaseData, err := json.Marshal(validValueTypeUpperCase)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"unmarshal ProfileRequest with success", validData, false},
		{"unmarshal ProfileRequest with success, valid value type int16", validValueTypeLowerCaseData, false},
		{"unmarshal ProfileRequest with success, valid value type INT16", validValueTypeUpperCaseData, false},
		{"unmarshal invalid ProfileRequest, empty data", []byte{}, true},
		{"unmarshal invalid ProfileRequest, string data", []byte("Invalid ProfileRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp ProfileRequest
			err := dp.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, dp, "Unmarshal did not result in expected ProfileRequest.")
			}
		})
	}
}

func TestAddProfileReqToProfileModels(t *testing.T) {
	requests := []ProfileRequest{profileData()}
	expectedProfileModels := []models.Profile{expectedProfile}
	resultModels := ProfileReqToProfileModels(requests)
	assert.Equal(t, expectedProfileModels, resultModels, "ProfileReqToProfileModels did not result in expected Profile model.")
}

func TestNewProfileRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewProfileRequest(dtos.Profile{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}
