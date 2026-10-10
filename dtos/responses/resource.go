//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// ResourceResponse defines the Response Content for GET Resource DTOs.
type ResourceResponse struct {
	common.BaseResponse `json:",inline"`
	Resource            dtos.Resource `json:"resource"`
}

// NewResourceResponse creates resource response DTO with required fields
func NewResourceResponse(requestId string, message string, statusCode int, resource dtos.Resource) ResourceResponse {
	return ResourceResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Resource:     resource,
	}
}
