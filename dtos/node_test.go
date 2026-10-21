//
// Copyright (C) 2021-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

var testId = "8b0ee7cb-7a94-4e21-bebf-55961071b060"
var testName = "NodeName"
var testParent = "ParentName"
var testDescription = "Describe Me"
var testAdminState = "LOCKED"
var testOperatingState = "UP"
var testDirection = "SOUTH"
var testLocation = "Location"
var testServiceName = "ServiceName"
var testProfileName = "ProfileName"

var testNodeDto = Node{
	DBTimestamp:    DBTimestamp{Created: 123, Modified: 456},
	Id:             testId,
	Name:           testName,
	Parent:         testParent,
	Description:    testDescription,
	AdminState:     testAdminState,
	OperatingState: testOperatingState,
	Labels:         []string{"label1", "label2"},
	Direction:      testDirection,
	Location:       testLocation,
	ServiceName:    testServiceName,
	ProfileName:    testProfileName,
	AutoEvents:     []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Protocols:      map[string]ProtocolProperties{"protocol": {"key": "value"}},
	Tags:           map[string]any{"tag": "value"},
	Properties:     map[string]any{"property": "value"},
}

var testNodeModel = models.Node{
	DBTimestamp:    models.DBTimestamp{Created: 123, Modified: 456},
	Id:             testId,
	Name:           testName,
	Parent:         testParent,
	Description:    testDescription,
	AdminState:     models.AdminState(testAdminState),
	OperatingState: models.OperatingState(testOperatingState),
	Labels:         []string{"label1", "label2"},
	Direction:      testDirection,
	Location:       testLocation,
	ServiceName:    testServiceName,
	ProfileName:    testProfileName,
	AutoEvents:     []models.AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Protocols:      map[string]models.ProtocolProperties{"protocol": {"key": "value"}},
	Tags:           map[string]any{"tag": "value"},
	Properties:     map[string]any{"property": "value"},
}

var testUpdateDto = UpdateNode{
	Id:             &testId,
	Name:           &testName,
	Parent:         &testParent,
	Description:    &testDescription,
	AdminState:     &testAdminState,
	OperatingState: &testOperatingState,
	Labels:         []string{"label1", "label2"},
	Direction:      &testDirection,
	Location:       testLocation,
	ServiceName:    &testServiceName,
	ProfileName:    &testProfileName,
	AutoEvents:     []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Protocols:      map[string]ProtocolProperties{"protocol": {"key": "value"}},
	Tags:           map[string]any{"tag": "value"},
	Properties:     map[string]any{"property": "value"},
}

func TestNodeDTOtoModel(t *testing.T) {
	dto := ToNodeModel(testNodeDto)
	// All fields should propagate except DBTimestamp
	testNodeModelWithoutTime := testNodeModel
	testNodeModelWithoutTime.DBTimestamp = models.DBTimestamp{}
	assert.Equal(t, testNodeModelWithoutTime, dto)
}

func TestNodeDTOtoModelWithNilProperties(t *testing.T) {
	dto := testNodeDto
	dto.Properties = nil
	model := ToNodeModel(dto)
	assert.NotNil(t, model.Properties)
}

func TestNodeModeltoDTO(t *testing.T) {
	model := testNodeModel
	dto := FromNodeModelToDTO(model)
	assert.Equal(t, testNodeDto, dto)
}

func TestNodeModeltoDTOWithNilProperties(t *testing.T) {
	model := testNodeModel
	model.Properties = nil
	dto := FromNodeModelToDTO(model)
	assert.NotNil(t, dto.Properties)
}

func TestFromNodeModelToUpdateDTO(t *testing.T) {
	model := testNodeModel
	dto := FromNodeModelToUpdateDTO(model)
	assert.Equal(t, testUpdateDto, dto)
}
