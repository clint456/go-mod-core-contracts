//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/assert"
)

func TestNewNodeCoreCommandResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	testParameters := []dtos.CoreCommandParameter{
		{ResourceName: "resource", ValueType: common.ValueTypeInt8},
	}
	expectedNodeCoreCommand := dtos.NodeCoreCommand{
		NodeName:    "testNode1",
		ProfileName: "testProfile",
		CoreCommands: []dtos.CoreCommand{
			{Name: "testCommand1", Get: true, Set: false, Path: "/node/name/testNode1/command/testCommand1", Url: "http://127.0.0.1:48082", Parameters: testParameters},
			{Name: "testCommand2", Get: false, Set: true, Path: "/node/name/testNode1/command/testCommand2", Url: "http://127.0.0.1:48082", Parameters: testParameters},
		},
	}
	actual := NewNodeCoreCommandResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedNodeCoreCommand)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedNodeCoreCommand, actual.NodeCoreCommand)
}

func TestNewMultiNodeCoreCommandsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	testParameters := []dtos.CoreCommandParameter{
		{ResourceName: "resource", ValueType: common.ValueTypeInt8},
	}
	expectedNodeCoreCommands := []dtos.NodeCoreCommand{
		dtos.NodeCoreCommand{
			NodeName:    "testNode1",
			ProfileName: "testProfile",
			CoreCommands: []dtos.CoreCommand{
				{Name: "testCommand1", Get: true, Set: false, Path: "/node/name/testNode1/command/testCommand1", Url: "http://127.0.0.1:48082", Parameters: testParameters},
				{Name: "testCommand2", Get: false, Set: true, Path: "/node/name/testNode1/command/testCommand2", Url: "http://127.0.0.1:48082", Parameters: testParameters},
			},
		},
		dtos.NodeCoreCommand{
			NodeName:    "testNode2",
			ProfileName: "testProfile",
			CoreCommands: []dtos.CoreCommand{
				{Name: "testCommand3", Get: true, Set: false, Path: "/node/name/testNode1/command/testCommand1", Url: "http://127.0.0.1:48082", Parameters: testParameters},
				{Name: "testCommand4", Get: false, Set: true, Path: "/node/name/testNode1/command/testCommand2", Url: "http://127.0.0.1:48082", Parameters: testParameters},
			},
		},
	}
	expectedTotalCount := int64(2)
	actual := NewMultiNodeCoreCommandsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedNodeCoreCommands)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedNodeCoreCommands, actual.NodeCoreCommands)
}
