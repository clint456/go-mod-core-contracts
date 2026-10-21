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

// AddNodeRequest defines the Request Content for POST Node DTO.
type AddNodeRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Node                  dtos.Node `json:"node"`
}

// Validate satisfies the Validator interface
func (d AddNodeRequest) Validate() error {
	err := common.Validate(d)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddNodeRequest type
func (d *AddNodeRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Node dtos.Node
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	if alias.Node.Properties == nil {
		alias.Node.Properties = make(map[string]any)
	}

	*d = AddNodeRequest(alias)

	// validate AddNodeRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

// AddNodeReqToNodeModels transforms the AddNodeRequest DTO array to the Node model array
func AddNodeReqToNodeModels(addRequests []AddNodeRequest) (Nodes []models.Node) {
	for _, req := range addRequests {
		d := dtos.ToNodeModel(req.Node)
		Nodes = append(Nodes, d)
	}
	return Nodes
}

// UpdateNodeRequest defines the Request Content for PUT event as pushed DTO.
type UpdateNodeRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Node                  dtos.UpdateNode `json:"node"`
}

// Validate satisfies the Validator interface
func (d UpdateNodeRequest) Validate() error {
	err := common.Validate(d)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateNodeRequest type
func (d *UpdateNodeRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Node dtos.UpdateNode
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*d = UpdateNodeRequest(alias)

	// validate UpdateNodeRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceNodeModelFieldsWithDTO replace existing Node's fields with DTO patch
func ReplaceNodeModelFieldsWithDTO(node *models.Node, patch dtos.UpdateNode) {
	if patch.Description != nil {
		node.Description = *patch.Description
	}
	if patch.Parent != nil {
		node.Parent = *patch.Parent
	}
	if patch.AdminState != nil {
		node.AdminState = models.AdminState(*patch.AdminState)
	}
	if patch.OperatingState != nil {
		node.OperatingState = models.OperatingState(*patch.OperatingState)
	}
	if patch.Direction != nil {
		node.Direction = *patch.Direction
	}
	if patch.ServiceName != nil {
		node.ServiceName = *patch.ServiceName
	}
	if patch.ProfileName != nil {
		node.ProfileName = *patch.ProfileName
	}
	if patch.Labels != nil {
		node.Labels = patch.Labels
	}
	if patch.Location != nil {
		node.Location = patch.Location
	}
	if patch.AutoEvents != nil {
		node.AutoEvents = dtos.ToAutoEventModels(patch.AutoEvents)
	}
	if patch.Protocols != nil {
		node.Protocols = dtos.ToProtocolModels(patch.Protocols)
	}
	if patch.Tags != nil {
		node.Tags = patch.Tags
	}
	if patch.Properties != nil {
		node.Properties = patch.Properties
	}
}

func NewAddNodeRequest(dto dtos.Node) AddNodeRequest {
	return AddNodeRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Node:        dto,
	}
}

func NewUpdateNodeRequest(dto dtos.UpdateNode) UpdateNodeRequest {
	return UpdateNodeRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Node:        dto,
	}
}
