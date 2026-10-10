//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewNodeServiceResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedNodeService := dtos.NodeService{Name: "test node service"}
	actual := NewNodeServiceResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedNodeService)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedNodeService, actual.Service)
}

func TestNewMultiNodeServicesResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedNodeServices := []dtos.NodeService{
		{Name: "test node service1"},
		{Name: "test node service2"},
	}
	expectedTotalCount := int64(2)
	actual := NewMultiNodeServicesResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedNodeServices)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedNodeServices, actual.Services)
}
