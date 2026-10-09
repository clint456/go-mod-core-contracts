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

// NorthRequest defines the Request Content for POST/PUT North DTO.
type NorthRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	North                 dtos.North `json:"north"`
}

// NewNorthRequest creates a NorthRequest with a generated request id.
func NewNorthRequest(dto dtos.North) NorthRequest {
	return NorthRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		North:       dto,
	}
}

// NorthReqToNorthModel transforms the NorthRequest DTO to the North model
func NorthReqToNorthModel(req NorthRequest) models.North {
	return dtos.ToNorthModel(req.North)
}

// NorthReqToNorthModels transforms the NorthRequest DTO array to the North model array
func NorthReqToNorthModels(reqs []NorthRequest) []models.North {
	norths := make([]models.North, len(reqs))
	for i, req := range reqs {
		norths[i] = NorthReqToNorthModel(req)
	}
	return norths
}
