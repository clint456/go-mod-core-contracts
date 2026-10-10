//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// AddNodeServiceRequest defines the Request Content for POST NodeService DTO.
type AddNodeServiceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Service               dtos.NodeService `json:"service"`
}

// Validate satisfies the Validator interface
func (ds AddNodeServiceRequest) Validate() error {
	err := common.Validate(ds)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddNodeServiceRequest type
func (ds *AddNodeServiceRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Service dtos.NodeService
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	if alias.Service.Properties == nil {
		alias.Service.Properties = make(map[string]any)
	}

	*ds = AddNodeServiceRequest(alias)

	// validate AddNodeServiceRequest DTO
	if err := ds.Validate(); err != nil {
		return err
	}
	return nil
}

// AddNodeServiceReqToNodeServiceModels transforms the AddNodeServiceRequest DTO array to the NodeService model array
func AddNodeServiceReqToNodeServiceModels(addRequests []AddNodeServiceRequest) (NodeServices []models.NodeService) {
	for _, req := range addRequests {
		ds := dtos.ToNodeServiceModel(req.Service)
		NodeServices = append(NodeServices, ds)
	}
	return NodeServices
}

// UpdateNodeServiceRequest defines the Request Content for PUT event as pushed DTO.
type UpdateNodeServiceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Service               dtos.UpdateNodeService `json:"service"`
}

// Validate satisfies the Validator interface
func (ds UpdateNodeServiceRequest) Validate() error {
	err := common.Validate(ds)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateNodeServiceRequest type
func (ds *UpdateNodeServiceRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Service dtos.UpdateNodeService
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*ds = UpdateNodeServiceRequest(alias)

	// validate UpdateNodeServiceRequest DTO
	if err := ds.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceNodeServiceModelFieldsWithDTO replace existing NodeService's fields with DTO patch
func ReplaceNodeServiceModelFieldsWithDTO(ds *models.NodeService, patch dtos.UpdateNodeService) {
	if patch.Description != nil {
		ds.Description = *patch.Description
	}
	if patch.AdminState != nil {
		ds.AdminState = models.AdminState(*patch.AdminState)
	}
	if patch.Labels != nil {
		ds.Labels = patch.Labels
	}
	if patch.BaseAddress != nil {
		ds.BaseAddress = *patch.BaseAddress
	}
	if patch.Properties != nil {
		ds.Properties = patch.Properties
	}
}

func NewAddNodeServiceRequest(dto dtos.NodeService) AddNodeServiceRequest {
	return AddNodeServiceRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Service:     dto,
	}
}

func NewUpdateNodeServiceRequest(dto dtos.UpdateNodeService) UpdateNodeServiceRequest {
	return UpdateNodeServiceRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Service:     dto,
	}
}
