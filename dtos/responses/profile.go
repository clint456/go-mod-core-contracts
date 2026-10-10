//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

// ProfileResponse defines the Response Content for GET Profile DTOs.
type ProfileResponse struct {
	common.BaseResponse `json:",inline"`
	Profile             dtos.Profile `json:"profile"`
}

func NewProfileResponse(requestId string, message string, statusCode int, profile dtos.Profile) ProfileResponse {
	return ProfileResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Profile:      profile,
	}
}

// MultiProfilesResponse defines the Response Content for GET multiple Profile DTOs.
type MultiProfilesResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Profiles                          []dtos.Profile `json:"profiles"`
}

func NewMultiProfilesResponse(requestId string, message string, statusCode int, totalCount int64, profiles []dtos.Profile) MultiProfilesResponse {
	return MultiProfilesResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Profiles:                   profiles,
	}
}

// ProfileBasicInfoResponse defines the Response Content for GET ProfileBasicInfo DTOs.
type ProfileBasicInfoResponse struct {
	common.BaseResponse `json:",inline"`
	Profile             dtos.ProfileBasicInfo `json:"profile"`
}

func NewProfileBasicInfoResponse(requestId string, message string, statusCode int, basicInfo dtos.ProfileBasicInfo) ProfileBasicInfoResponse {
	return ProfileBasicInfoResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Profile:      basicInfo,
	}
}

// MultiProfileBasicInfoResponse defines the Response Content for GET multiple ProfileBasicInfo DTOs.
type MultiProfileBasicInfoResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Profiles                          []dtos.ProfileBasicInfo `json:"profiles"`
}

func NewMultiProfileBasicInfosResponse(requestId string, message string, statusCode int, totalCount int64, profileBasicInfos []dtos.ProfileBasicInfo) MultiProfileBasicInfoResponse {
	return MultiProfileBasicInfoResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Profiles:                   profileBasicInfos,
	}
}
