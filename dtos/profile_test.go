//
// Copyright (C) 2021-2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testLabels = []string{"MODBUS", "TEMP"}
var testAttributes = map[string]interface{}{
	"TestAttribute": "TestAttributeValue",
}
var testMappings = map[string]string{"0": "off", "1": "on"}
var testTags = map[string]any{"TestTagsKey": "TestTagsValue"}
var testOptional = map[string]any{"isVirtual": false}

var testProfile = models.Profile{
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
			Optional:  testOptional,
		},
		Tags: testTags,
	}},
	Commands: []models.Command{{
		Name:      TestCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []models.ResourceOperation{{
			Resource: TestResourceName,
			Mappings: testMappings,
		}},
		Tags: testTags,
	}},
}

func profileData() Profile {
	return Profile{
		ApiVersion: common.ApiVersion,
		ProfileBasicInfo: ProfileBasicInfo{
			Name:         TestProfileName,
			Manufacturer: TestManufacturer,
			Description:  TestDescription,
			Model:        TestModel,
			Labels:       testLabels,
		},
		Resources: []Resource{{
			Name:        TestResourceName,
			Description: TestDescription,
			Attributes:  testAttributes,
			Properties: ResourceProperties{
				ValueType: common.ValueTypeInt16,
				ReadWrite: common.ReadWrite_RW,
				Optional:  testOptional,
			},
			Tags: testTags,
		}},
		Commands: []Command{{
			Name:      TestCommandName,
			ReadWrite: common.ReadWrite_RW,
			ResourceOperations: []ResourceOperation{{
				Resource: TestResourceName,
				Mappings: testMappings,
			}},
			Tags: testTags,
		}},
	}
}

func TestFromProfileModelToDTO(t *testing.T) {
	result := FromProfileModelToDTO(testProfile)
	assert.Equal(t, profileData(), result, "FromProfileModelToDTO did not result in expected node profile DTO.")
}

func TestProfileBasicInfo_LinkedNodeCount_JSON(t *testing.T) {
	tests := []struct {
		name            string
		linkedNodeCount int64
	}{
		{"non-zero LinkedNodeCount", 5},
		{"zero LinkedNodeCount present without omitempty", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := ProfileBasicInfo{
				Name:            TestProfileName,
				LinkedNodeCount: tt.linkedNodeCount,
			}

			data, err := json.Marshal(info)
			require.NoError(t, err)
			assert.Contains(t, string(data), "linkedNodeCount")

			var result ProfileBasicInfo
			err = json.Unmarshal(data, &result)
			require.NoError(t, err)
			assert.Equal(t, tt.linkedNodeCount, result.LinkedNodeCount)
		})
	}
}

func TestProfileBasicInfo_LinkedNodeCount_YAML(t *testing.T) {
	tests := []struct {
		name            string
		linkedNodeCount int64
	}{
		{"non-zero LinkedNodeCount", 10},
		{"zero LinkedNodeCount", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := ProfileBasicInfo{
				Name:            TestProfileName,
				LinkedNodeCount: tt.linkedNodeCount,
			}

			data, err := yaml.Marshal(info)
			require.NoError(t, err)

			var result ProfileBasicInfo
			err = yaml.Unmarshal(data, &result)
			require.NoError(t, err)
			assert.Equal(t, tt.linkedNodeCount, result.LinkedNodeCount)
		})
	}
}

func TestFromProfileModelToBasicInfoDTO(t *testing.T) {
	result := FromProfileModelToBasicInfoDTO(testProfile)
	assert.Equal(t, profileData().ProfileBasicInfo, result, "FromProfileModelToBasicInfoDTO did not result in expected node profile basic info DTO.")
}

func TestProfileDTOValidation(t *testing.T) {
	valid := profileData()
	duplicatedResource := profileData()
	duplicatedResource.Resources = append(
		duplicatedResource.Resources, Resource{Name: TestResourceName})
	duplicatedCommand := profileData()
	duplicatedCommand.Commands = append(
		duplicatedCommand.Commands, Command{Name: TestCommandName})
	mismatchedResource := profileData()
	mismatchedResource.Commands[0].ResourceOperations = append(
		mismatchedResource.Commands[0].ResourceOperations, ResourceOperation{Resource: "missMatchedResource"})
	invalidReadWrite := profileData()
	invalidReadWrite.Resources[0].Properties.ReadWrite = common.ReadWrite_R
	binaryWithWritePermission := profileData()
	binaryWithWritePermission.Resources[0].Properties.ValueType = common.ValueTypeBinary
	binaryWithWritePermission.Resources[0].Properties.ReadWrite = common.ReadWrite_RW

	tests := []struct {
		name        string
		profile     Profile
		expectError bool
	}{
		{"valid node profile", valid, false},
		{"duplicated node resource", duplicatedResource, true},
		{"duplicated node command", duplicatedCommand, true},
		{"mismatched resource", mismatchedResource, true},
		{"invalid ReadWrite permission", invalidReadWrite, true},
		{"write permission not support Binary value type", binaryWithWritePermission, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.profile.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddProfile_UnmarshalYAML(t *testing.T) {
	valid := profileData()
	resultTestBytes, _ := yaml.Marshal(profileData())
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected Profile
		args     args
		wantErr  bool
	}{
		{"valid", valid, args{resultTestBytes}, false},
		{"invalid", Profile{}, args{[]byte("Invalid Profile")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp Profile
			err := yaml.Unmarshal(tt.args.data, &dp)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, dp, "Unmarshal did not result in expected ProfileRequest.")
			}
		})
	}
}
