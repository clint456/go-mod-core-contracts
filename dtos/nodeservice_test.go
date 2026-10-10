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

var testNodeServiceDto = NodeService{
	DBTimestamp: DBTimestamp{Created: 123, Modified: 456},
	Id:          testId,
	Name:        testName,
	Description: testDescription,
	Labels:      []string{"label1", "label2"},
	BaseAddress: "http://localhost:49990",
	AdminState:  testAdminState,
	Properties:  map[string]any{"property": "value"},
}

var testNodeServiceModel = models.NodeService{
	DBTimestamp: models.DBTimestamp{Created: 123, Modified: 456},
	Id:          testId,
	Name:        testName,
	Description: testDescription,
	Labels:      []string{"label1", "label2"},
	BaseAddress: "http://localhost:49990",
	AdminState:  models.AdminState(testAdminState),
	Properties:  map[string]any{"property": "value"},
}

var testUpdateNodeServiceDto = UpdateNodeService{
	Id:          &testId,
	Name:        &testName,
	Description: &testDescription,
	BaseAddress: &testNodeServiceDto.BaseAddress,
	Labels:      []string{"label1", "label2"},
	AdminState:  &testAdminState,
	Properties:  map[string]any{"property": "value"},
}

func TestNodeServiceDTOtoModel(t *testing.T) {
	model := ToNodeServiceModel(testNodeServiceDto)
	// All fields should propagate except DBTimestamp
	testNodeServiceModelWithoutTimestamp := testNodeServiceModel
	testNodeServiceModelWithoutTimestamp.DBTimestamp = models.DBTimestamp{}
	assert.Equal(t, testNodeServiceModelWithoutTimestamp, model)
}

func TestNodeServiceDTOtoModelWithNilProperties(t *testing.T) {
	dto := testNodeServiceDto
	dto.Properties = nil
	model := ToNodeServiceModel(dto)
	assert.NotNil(t, model.Properties)
}

func TestNodeServiceModelToDTO(t *testing.T) {
	dto := FromNodeServiceModelToDTO(testNodeServiceModel)
	assert.Equal(t, testNodeServiceDto, dto)
}

func TestNodeServiceModelToDTOWithNilProperties(t *testing.T) {
	model := testNodeServiceModel
	model.Properties = nil
	dto := FromNodeServiceModelToDTO(model)
	assert.NotNil(t, dto.Properties)
}

func TestFromNodeServiceModelToUpdateDTO(t *testing.T) {
	model := testNodeServiceModel
	dto := FromNodeServiceModelToUpdateDTO(model)
	assert.Equal(t, testUpdateNodeServiceDto, dto)
}
