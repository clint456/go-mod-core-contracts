//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

var testDiscoveredNodeDto = DiscoveredNode{
	ProfileName: testProfileName,
	AdminState:  testAdminState,
	AutoEvents:  []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Properties:  map[string]any{"property": "value"},
}

var testDiscoveredNodeModel = models.DiscoveredNode{
	ProfileName: testProfileName,
	AdminState:  models.AdminState(testAdminState),
	AutoEvents:  []models.AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Properties:  map[string]any{"property": "value"},
}

var testUpdateDiscoveredNodeDto = UpdateDiscoveredNode{
	ProfileName: &testProfileName,
	AdminState:  &testAdminState,
	AutoEvents:  []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Properties:  map[string]any{"property": "value"},
}

func TestDiscoveredNodeDTOtoModel(t *testing.T) {
	model := ToDiscoveredNodeModel(testDiscoveredNodeDto)
	assert.Equal(t, testDiscoveredNodeModel, model)
}

func TestDiscoveredNodeDTOtoModelWithNilProperties(t *testing.T) {
	dto := testDiscoveredNodeDto
	dto.Properties = nil
	model := ToDiscoveredNodeModel(dto)
	assert.NotNil(t, model.Properties)
}

func TestDiscoveredNodeModelToDTO(t *testing.T) {
	dto := FromDiscoveredNodeModelToDTO(testDiscoveredNodeModel)
	assert.Equal(t, testDiscoveredNodeDto, dto)
}

func TestDiscoveredNodeModelToDTOWithNilProperties(t *testing.T) {
	model := testDiscoveredNodeModel
	model.Properties = nil
	dto := FromDiscoveredNodeModelToDTO(model)
	assert.NotNil(t, dto.Properties)
}

func TestFromDiscoveredNodeModelToUpdateDTO(t *testing.T) {
	dto := FromDiscoveredNodeModelToUpdateDTO(testDiscoveredNodeModel)
	assert.Equal(t, testUpdateDiscoveredNodeDto, dto)
}
