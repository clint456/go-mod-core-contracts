//
// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// NorthResponse defines the Response Content for GET North DTOs.
type NorthResponse struct {
	common.BaseResponse `json:",inline"`
	North               dtos.North `json:"north"`
}

// NewNorthResponse creates a NorthResponse.
func NewNorthResponse(requestId string, message string, statusCode int, north dtos.North) NorthResponse {
	return NorthResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		North:        north,
	}
}

// MultiNorthsResponse defines the Response Content for GET multiple North DTOs.
type MultiNorthsResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Norths                            []dtos.North `json:"norths"`
}

// NewMultiNorthsResponse creates a MultiNorthsResponse.
func NewMultiNorthsResponse(requestId string, message string, statusCode int, totalCount int64, norths []dtos.North) MultiNorthsResponse {
	return MultiNorthsResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Norths:                     norths,
	}
}
