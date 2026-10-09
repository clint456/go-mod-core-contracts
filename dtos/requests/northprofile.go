//
// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// NorthProfileRequest defines the Request Content for POST/PUT NorthProfile DTO.
type NorthProfileRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Profile               dtos.NorthProfile `json:"profile"`
}

// NewNorthProfileRequest creates a NorthProfileRequest with a generated request id.
func NewNorthProfileRequest(dto dtos.NorthProfile) NorthProfileRequest {
	return NorthProfileRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Profile:     dto,
	}
}

// NorthProfileReqToNorthProfileModel transforms the NorthProfileRequest DTO to the NorthProfile model
func NorthProfileReqToNorthProfileModel(req NorthProfileRequest) models.NorthProfile {
	return dtos.ToNorthProfileModel(req.Profile)
}

// NorthProfileReqToNorthProfileModels transforms the NorthProfileRequest DTO array to the NorthProfile model array
func NorthProfileReqToNorthProfileModels(reqs []NorthProfileRequest) []models.NorthProfile {
	profiles := make([]models.NorthProfile, len(reqs))
	for i, req := range reqs {
		profiles[i] = NorthProfileReqToNorthProfileModel(req)
	}
	return profiles
}
