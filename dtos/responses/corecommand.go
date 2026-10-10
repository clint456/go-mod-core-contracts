//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// NodeCoreCommandResponse defines the Response Content for GET NodeCoreCommand DTO.
type NodeCoreCommandResponse struct {
	common.BaseResponse `json:",inline"`
	NodeCoreCommand     dtos.NodeCoreCommand `json:"nodeCoreCommand"`
}

func NewNodeCoreCommandResponse(requestId string, message string, statusCode int, nodeCoreCommand dtos.NodeCoreCommand) NodeCoreCommandResponse {
	return NodeCoreCommandResponse{
		BaseResponse:    common.NewBaseResponse(requestId, message, statusCode),
		NodeCoreCommand: nodeCoreCommand,
	}
}

// MultiNodeCoreCommandsResponse defines the Response Content for GET multiple NodeCoreCommand DTOs.
type MultiNodeCoreCommandsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	NodeCoreCommands                  []dtos.NodeCoreCommand `json:"nodeCoreCommands"`
}

func NewMultiNodeCoreCommandsResponse(requestId string, message string, statusCode int, totalCount int64, commands []dtos.NodeCoreCommand) MultiNodeCoreCommandsResponse {
	return MultiNodeCoreCommandsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		NodeCoreCommands:           commands,
	}
}
