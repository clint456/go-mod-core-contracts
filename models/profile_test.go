// Copyright (C) 2025 IOTech Ltd

package models

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfile_Clone(t *testing.T) {
	testMinimum := -1.123
	testMaximum := 1.123
	testProfile := Profile{
		DBTimestamp:  DBTimestamp{},
		ApiVersion:   common.ApiVersion,
		Description:  "test description",
		Id:           "ca93c8fa-9919-4ec5-85d3-f81b2b6a7bc1",
		Name:         "TestProfile",
		Manufacturer: "testManufacturer",
		Model:        "testModel",
		Labels:       []string{"label1", "label2"},
		Resources: []Resource{{
			Description: "test description",
			Name:        "TestResource",
			IsHidden:    false,
			Properties: ResourceProperties{
				ValueType: common.ValueTypeString, Minimum: &testMinimum, Maximum: &testMaximum},
			Attributes: map[string]any{
				"foo": "bar",
			},
			Tags: map[string]any{
				"tag1": "val1",
			},
		}},
		Commands: []Command{{
			Name:      "TestCommand",
			IsHidden:  false,
			ReadWrite: "RW",
			ResourceOperations: []ResourceOperation{{
				Resource:     "TestResource1",
				DefaultValue: "",
				Mappings: map[string]string{
					"on": "true",
				},
			}, {
				Resource:     "TestResource2",
				DefaultValue: "",
				Mappings: map[string]string{
					"off": "false",
				},
			}},
			Tags: map[string]any{
				"tag3": "val3",
			},
		}},
	}
	clone := testProfile.Clone()
	assert.Equal(t, testProfile, clone)
}
