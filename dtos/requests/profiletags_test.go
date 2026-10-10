//
// Copyright (C) 2025 IOTech Ltd
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

var testProfileModel = models.Profile{
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
	}},
	Commands: []models.Command{{
		Name:      TestCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []models.ResourceOperation{{
			Resource: TestResourceName,
		}},
	}},
}

var testProfileTagsRequest = ProfileTagsRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	UpdateProfileTags: dtos.UpdateProfileTags{
		Resources: []dtos.UpdateTags{testUpdateTags(TestResourceName, testTags)},
		Commands:  []dtos.UpdateTags{testUpdateTags(TestCommandName, testTags)},
	},
}

func testUpdateTags(name string, tags map[string]any) dtos.UpdateTags {
	return dtos.UpdateTags{Name: name, Tags: tags}
}

func TestProfileTagsRequest_Validate(t *testing.T) {
	valid := testProfileTagsRequest
	emptyUpdateProfileTags := valid
	emptyUpdateProfileTags.UpdateProfileTags = dtos.UpdateProfileTags{}
	emptyDR := valid
	emptyDR.Resources = []dtos.UpdateTags{}
	emptyDC := valid
	emptyDC.Commands = []dtos.UpdateTags{}

	noDRName := valid
	noDRName.Resources = []dtos.UpdateTags{testUpdateTags("", testTags)}
	noDCName := valid
	noDCName.Commands = []dtos.UpdateTags{testUpdateTags("", testTags)}
	noDRTags := valid
	noDRTags.Resources = []dtos.UpdateTags{testUpdateTags(TestResourceName, nil)}
	noDCTags := valid
	noDCTags.Commands = []dtos.UpdateTags{testUpdateTags(TestCommandName, nil)}
	emptyDRTags := valid
	emptyDRTags.Resources = []dtos.UpdateTags{testUpdateTags(TestResourceName, map[string]any{})}
	emptyDCTags := valid
	emptyDCTags.Commands = []dtos.UpdateTags{testUpdateTags(TestCommandName, map[string]any{})}

	tests := []struct {
		name        string
		ProfileTags ProfileTagsRequest
		expectError bool
	}{
		{"valid ProfileTagsRequest", valid, false},
		{"valid ProfileTagsRequest, empty UpdateProfileTags", emptyUpdateProfileTags, false},
		{"valid ProfileTagsRequest, empty resources", emptyDR, false},
		{"valid ProfileTagsRequest, empty commands", emptyDC, false},
		{"invalid ProfileTagsRequest, no resource name", noDRName, true},
		{"invalid ProfileTagsRequest, no command name", noDCName, true},
		{"invalid ProfileTagsRequest, no resource tags", noDRTags, true},
		{"invalid ProfileTagsRequest, no command tags", noDCTags, true},
		{"invalid ProfileTagsRequest, empty resource tags", emptyDRTags, true},
		{"invalid ProfileTagsRequest, empty command tags", emptyDCTags, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ProfileTags.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddProfileTags_UnmarshalJSON(t *testing.T) {
	expected := testProfileTagsRequest
	validData, err := json.Marshal(testProfileTagsRequest)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"unmarshal ProfileTagsRequest with success", validData, false},
		{"unmarshal invalid ProfileTagsRequest, empty data", []byte{}, true},
		{"unmarshal invalid ProfileTagsRequest, string data", []byte("Invalid ProfileTagsRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp ProfileTagsRequest
			err := dp.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, dp, "Unmarshal did not result in expected ProfileTagsRequest.")
			}
		})
	}
}

func TestReplaceProfileModelTagsWithDTO(t *testing.T) {
	profile := testProfileModel
	patch := testProfileTagsRequest.UpdateProfileTags
	expectedProfileModels := expectedProfile
	ReplaceProfileModelTagsWithDTO(&profile, patch)
	assert.Equal(t, expectedProfileModels, profile, "ReplaceProfileModelTagsWithDTO did not result in expected Profile model.")
}

func TestReplaceProfileModelTagsWithDTO_DeleteTag(t *testing.T) {
	const deletedKey = "TestTagsKey"
	const keptKey = "KeptTagsKey"
	// the shared fixtures are package-level, so give this profile its own maps to mutate
	startingTags := func() map[string]any {
		return map[string]any{deletedKey: "TestTagsValue", keptKey: "KeptTagsValue"}
	}

	profile := testProfileModel
	profile.Resources = []models.Resource{{
		Name: TestResourceName,
		Tags: startingTags(),
	}}
	profile.Commands = []models.Command{{
		Name: TestCommandName,
		Tags: startingTags(),
	}}

	patch := dtos.UpdateProfileTags{
		Resources: []dtos.UpdateTags{testUpdateTags(TestResourceName, map[string]any{deletedKey: nil})},
		Commands:  []dtos.UpdateTags{testUpdateTags(TestCommandName, map[string]any{deletedKey: nil})},
	}
	ReplaceProfileModelTagsWithDTO(&profile, patch)

	expected := map[string]any{keptKey: "KeptTagsValue"}
	assert.Equal(t, expected, profile.Resources[0].Tags, "node resource tag was not deleted")
	assert.Equal(t, expected, profile.Commands[0].Tags, "node command tag was not deleted")
}

func TestMergeTags(t *testing.T) {
	testTagsKey := "TestTagsKey"
	testTagsValue := "TestTagsValue"
	testNewTagsKey := "TestNewTagsKey"
	testNewTagsValue := "TestNewTagsValue"

	emptyDest := make(map[string]any)
	emptySrc := make(map[string]any)
	existKeyDest := map[string]any{testTagsKey: testTagsValue}
	updateTestTags := map[string]any{testTagsKey: testNewTagsValue}
	noExistKeyDest := map[string]any{testTagsKey: testTagsValue}
	newTestTags := map[string]any{testNewTagsKey: testNewTagsValue}
	expectedMergeMap := map[string]any{testTagsKey: testTagsValue, testNewTagsKey: testNewTagsValue}
	nestedDest := map[string]any{testTagsKey: noExistKeyDest}
	nestedSrc := map[string]any{testTagsKey: newTestTags}
	expectedNestedMergeMap := map[string]any{testTagsKey: expectedMergeMap}

	// the delete cases keep a sibling key so that wiping the whole map is distinguishable from
	// deleting the requested one
	deleteDest := map[string]any{testTagsKey: testTagsValue, testNewTagsKey: testNewTagsValue}
	deleteSrc := map[string]any{testTagsKey: nil}
	expectedDeleteMap := map[string]any{testNewTagsKey: testNewTagsValue}

	nestedDeleteDest := map[string]any{testTagsKey: map[string]any{testTagsKey: testTagsValue, testNewTagsKey: testNewTagsValue}}
	nestedDeleteSrc := map[string]any{testTagsKey: map[string]any{testTagsKey: nil}}
	expectedNestedDeleteMap := map[string]any{testTagsKey: map[string]any{testNewTagsKey: testNewTagsValue}}

	nestedNewKeyDest := map[string]any{testTagsKey: testTagsValue}
	nestedNewKeySrc := map[string]any{testNewTagsKey: map[string]any{"keep": "v", "drop": nil}}
	expectedNestedNewKeyMap := map[string]any{testTagsKey: testTagsValue, testNewTagsKey: map[string]any{"keep": "v"}}

	deleteNoExistKeyDest := map[string]any{testTagsKey: testTagsValue}
	deleteNoExistKeySrc := map[string]any{testNewTagsKey: nil}
	expectedDeleteNoExistKeyMap := map[string]any{testTagsKey: testTagsValue}

	deleteAndUpdateDest := map[string]any{testTagsKey: testTagsValue}
	deleteAndUpdateSrc := map[string]any{testTagsKey: nil, testNewTagsKey: testNewTagsValue}
	expectedDeleteAndUpdateMap := map[string]any{testNewTagsKey: testNewTagsValue}

	tests := []struct {
		name     string
		dest     map[string]any
		src      map[string]any
		expected map[string]any
	}{
		{"merge tags with nil dest", nil, testTags, testTags},
		{"merge tags with nil src", testTags, nil, testTags},
		{"merge tags with empty dest", emptyDest, testTags, testTags},
		{"merge tags with empty src", testTags, emptySrc, testTags},
		{"merge tags with existing key", existKeyDest, updateTestTags, updateTestTags},
		{"merge tags with new key", noExistKeyDest, newTestTags, expectedMergeMap},
		{"merge tags with nested struct", nestedDest, nestedSrc, expectedNestedMergeMap},
		{"merge tags with nil value deletes the key", deleteDest, deleteSrc, expectedDeleteMap},
		{"merge tags with nil value deletes the nested key", nestedDeleteDest, nestedDeleteSrc, expectedNestedDeleteMap},
		{"merge tags with nil value in a nested map on a new key", nestedNewKeyDest, nestedNewKeySrc, expectedNestedNewKeyMap},
		{"merge tags with nil value on non-existing key", deleteNoExistKeyDest, deleteNoExistKeySrc, expectedDeleteNoExistKeyMap},
		{"merge tags with nil value mixed with an update", deleteAndUpdateDest, deleteAndUpdateSrc, expectedDeleteAndUpdateMap},
		{"merge tags with nil value and nil dest", nil, map[string]any{testTagsKey: nil}, map[string]any{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeTags(tt.dest, tt.src)
			assert.Equal(t, tt.expected, result, "mergeTags did not result in expected map")
		})
	}
}
