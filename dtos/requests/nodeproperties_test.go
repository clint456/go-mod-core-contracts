//
// Copyright (C) 2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testNodePropertiesRequest = NodePropertiesRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	UpdateNodeProperties: dtos.UpdateNodeProperties{
		Properties: testTags,
	},
}

func TestNodePropertiesRequest_Validate(t *testing.T) {
	valid := testNodePropertiesRequest
	noProperties := valid
	noProperties.Properties = nil
	emptyProperties := valid
	emptyProperties.Properties = map[string]any{}

	tests := []struct {
		name           string
		NodeProperties NodePropertiesRequest
		expectError    bool
	}{
		{"valid NodePropertiesRequest", valid, false},
		{"invalid NodePropertiesRequest, no properties", noProperties, true},
		{"invalid NodePropertiesRequest, empty properties", emptyProperties, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.NodeProperties.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// The validator reports the embedded struct in its field path, which is a Go implementation
// detail rather than part of the JSON contract (see NodePropertiesRequest.Validate).
func TestNodePropertiesRequest_Validate_HidesInternalStructName(t *testing.T) {
	invalid := testNodePropertiesRequest
	invalid.Properties = nil

	err := invalid.Validate()
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "UpdateNodeProperties")
}

func TestNodePropertiesRequest_UnmarshalJSON(t *testing.T) {
	expected := testNodePropertiesRequest
	validData, err := json.Marshal(testNodePropertiesRequest)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"unmarshal NodePropertiesRequest with success", validData, false},
		{"unmarshal invalid NodePropertiesRequest, empty data", []byte{}, true},
		{"unmarshal invalid NodePropertiesRequest, string data", []byte("Invalid NodePropertiesRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp NodePropertiesRequest
			err := dp.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, dp, "Unmarshal did not result in expected NodePropertiesRequest.")
			}
		})
	}
}

func TestReplaceNodeModelPropertiesWithDTO(t *testing.T) {
	const keptKey = "KeptPropertyKey"
	const keptValue = "KeptPropertyValue"
	const newKey = "NewPropertyKey"
	const newValue = "NewPropertyValue"

	node := models.Node{
		Name:       TestNodeName,
		Properties: map[string]any{keptKey: keptValue},
	}
	patch := dtos.UpdateNodeProperties{Properties: map[string]any{newKey: newValue}}

	ReplaceNodeModelPropertiesWithDTO(&node, patch)

	assert.Equal(t, map[string]any{keptKey: keptValue, newKey: newValue}, node.Properties,
		"properties not merged; keys absent from the patch must be preserved")
}

func TestReplaceNodeModelPropertiesWithDTO_DeleteProperty(t *testing.T) {
	const deletedKey = "DeletedPropertyKey"
	const keptKey = "KeptPropertyKey"
	const keptValue = "KeptPropertyValue"

	node := models.Node{
		Name:       TestNodeName,
		Properties: map[string]any{deletedKey: "DeletedPropertyValue", keptKey: keptValue},
	}
	patch := dtos.UpdateNodeProperties{Properties: map[string]any{deletedKey: nil}}

	ReplaceNodeModelPropertiesWithDTO(&node, patch)

	assert.NotContains(t, node.Properties, deletedKey, "property with a nil value was not deleted")
	assert.Equal(t, keptValue, node.Properties[keptKey], "sibling property must survive the delete")
}

func TestReplaceNodeModelPropertiesWithDTO_NestedDeleteOnNewKey(t *testing.T) {
	body := []byte(`{
		"apiVersion": "v3",
		"properties": {
			"property1": null,
			"nested1": {"field3": "field3Value", "field4ToRemove": null}
		}
	}`)
	var req NodePropertiesRequest
	require.NoError(t, req.UnmarshalJSON(body))

	node := models.Node{Name: TestNodeName, Properties: map[string]any{"keep": "me"}}
	ReplaceNodeModelPropertiesWithDTO(&node, req.UpdateNodeProperties)

	nested, ok := node.Properties["nested1"].(map[string]any)
	require.True(t, ok, "nested1 must be stored as a map")
	assert.NotContains(t, nested, "field4ToRemove", "null inside a brand-new nested map must delete, not persist as null")
	assert.Equal(t, "field3Value", nested["field3"])
	assert.Equal(t, "me", node.Properties["keep"], "unrelated property must survive")
	assert.NotContains(t, node.Properties, "property1", "top-level null must delete")
}

// A node with no properties yet must still accept a patch (see mergeTags for the nil-dest case).
func TestReplaceNodeModelPropertiesWithDTO_NilProperties(t *testing.T) {
	const newKey = "NewPropertyKey"
	const newValue = "NewPropertyValue"

	node := models.Node{Name: TestNodeName}
	patch := dtos.UpdateNodeProperties{Properties: map[string]any{newKey: newValue}}

	ReplaceNodeModelPropertiesWithDTO(&node, patch)

	assert.Equal(t, map[string]any{newKey: newValue}, node.Properties)
}
