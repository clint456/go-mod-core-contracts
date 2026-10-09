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

// NorthProfileResponse defines the Response Content for GET NorthProfile DTOs.
type NorthProfileResponse struct {
	common.BaseResponse `json:",inline"`
	Profile             dtos.NorthProfile `json:"profile"`
}

// NewNorthProfileResponse creates a NorthProfileResponse.
func NewNorthProfileResponse(requestId string, message string, statusCode int, profile dtos.NorthProfile) NorthProfileResponse {
	return NorthProfileResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Profile:      profile,
	}
}

// MultiNorthProfilesResponse defines the Response Content for GET multiple NorthProfile DTOs.
type MultiNorthProfilesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Profiles                          []dtos.NorthProfile `json:"profiles"`
}

// NewMultiNorthProfilesResponse creates a MultiNorthProfilesResponse.
func NewMultiNorthProfilesResponse(requestId string, message string, statusCode int, totalCount int64, profiles []dtos.NorthProfile) MultiNorthProfilesResponse {
	return MultiNorthProfilesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Profiles:                   profiles,
	}
}
