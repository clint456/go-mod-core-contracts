//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// NodeResponse defines the Response Content for GET Node DTOs.
type NodeResponse struct {
	common.BaseResponse `json:",inline"`
	Node                dtos.Node `json:"node"`
}

func NewNodeResponse(requestId string, message string, statusCode int, node dtos.Node) NodeResponse {
	return NodeResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Node:         node,
	}
}

// MultiNodesResponse defines the Response Content for GET multiple Node DTOs.
type MultiNodesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Nodes                             []dtos.Node `json:"nodes"`
}

func NewMultiNodesResponse(requestId string, message string, statusCode int, totalCount int64, nodes []dtos.Node) MultiNodesResponse {
	return MultiNodesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Nodes:                      nodes,
	}
}
