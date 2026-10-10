//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// NodeServiceResponse defines the Response Content for GET NodeService DTOs.
type NodeServiceResponse struct {
	common.BaseResponse `json:",inline"`
	Service             dtos.NodeService `json:"service"`
}

func NewNodeServiceResponse(requestId string, message string, statusCode int, nodeService dtos.NodeService) NodeServiceResponse {
	return NodeServiceResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Service:      nodeService,
	}
}

// MultiNodeServicesResponse defines the Response Content for GET multiple NodeService DTOs.
type MultiNodeServicesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Services                          []dtos.NodeService `json:"services"`
}

func NewMultiNodeServicesResponse(requestId string, message string, statusCode int, totalCount int64, nodeServices []dtos.NodeService) MultiNodeServicesResponse {
	return MultiNodeServicesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Services:                   nodeServices,
	}
}
