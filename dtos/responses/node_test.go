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

func TestNewNodeResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedNode := dtos.Node{Name: "test node"}
	actual := NewNodeResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedNode)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedNode, actual.Node)
}

func TestNewMultiNodesResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedNodes := []dtos.Node{
		{Name: "test node1"},
		{Name: "test node2"},
	}
	expectedTotalCount := int64(2)
	actual := NewMultiNodesResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedNodes)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedNodes, actual.Nodes)
}
