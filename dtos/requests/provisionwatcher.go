//
// Copyright (C) 2021-2024 IOTech Ltd
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

// AddProvisionWatcherRequest defines the Request Content for POST ProvisionWatcher DTO.
type AddProvisionWatcherRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProvisionWatcher      dtos.ProvisionWatcher `json:"provisionWatcher"`
}

// Validate satisfies the Validator interface
func (pw *AddProvisionWatcherRequest) Validate() error {
	err := common.Validate(pw)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddProvisionWatcherRequest type
func (pw *AddProvisionWatcherRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		ProvisionWatcher dtos.ProvisionWatcher
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	if alias.ProvisionWatcher.DiscoveredNode.Properties == nil {
		alias.ProvisionWatcher.DiscoveredNode.Properties = make(map[string]any)
	}

	*pw = AddProvisionWatcherRequest(alias)

	// validate AddNodeRequest DTO
	if err := pw.Validate(); err != nil {
		return err
	}
	return nil
}

// AddProvisionWatcherReqToProvisionWatcherModels transforms the AddProvisionWatcherRequest DTO array to the ProvisionWatcher model array
func AddProvisionWatcherReqToProvisionWatcherModels(addRequests []AddProvisionWatcherRequest) (ProvisionWatchers []models.ProvisionWatcher) {
	for _, req := range addRequests {
		d := dtos.ToProvisionWatcherModel(req.ProvisionWatcher)
		ProvisionWatchers = append(ProvisionWatchers, d)
	}
	return ProvisionWatchers
}

// UpdateProvisionWatcherRequest defines the Request Content for PUT event as pushed DTO.
type UpdateProvisionWatcherRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProvisionWatcher      dtos.UpdateProvisionWatcher `json:"provisionWatcher"`
}

// Validate satisfies the Validator interface
func (pw *UpdateProvisionWatcherRequest) Validate() error {
	err := common.Validate(pw)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateProvisionWatcherRequest type
func (pw *UpdateProvisionWatcherRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		ProvisionWatcher dtos.UpdateProvisionWatcher
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*pw = UpdateProvisionWatcherRequest(alias)

	// validate UpdateNodeRequest DTO
	if err := pw.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceProvisionWatcherModelFieldsWithDTO replace existing ProvisionWatcher's fields with DTO patch
func ReplaceProvisionWatcherModelFieldsWithDTO(pw *models.ProvisionWatcher, patch dtos.UpdateProvisionWatcher) {
	if patch.Labels != nil {
		pw.Labels = patch.Labels
	}
	if patch.Identifiers != nil {
		pw.Identifiers = patch.Identifiers
	}
	if patch.BlockingIdentifiers != nil {
		pw.BlockingIdentifiers = patch.BlockingIdentifiers
	}
	if patch.AdminState != nil {
		pw.AdminState = models.AdminState(*patch.AdminState)
	}
	if patch.DiscoveredNode.ProfileName != nil {
		pw.DiscoveredNode.ProfileName = *patch.DiscoveredNode.ProfileName
	}
	if patch.ServiceName != nil {
		pw.ServiceName = *patch.ServiceName
	}
	if patch.DiscoveredNode.AdminState != nil {
		pw.DiscoveredNode.AdminState = models.AdminState(*patch.DiscoveredNode.AdminState)
	}
	if patch.DiscoveredNode.AutoEvents != nil {
		pw.DiscoveredNode.AutoEvents = dtos.ToAutoEventModels(patch.DiscoveredNode.AutoEvents)
	}
	if patch.DiscoveredNode.Properties != nil {
		pw.DiscoveredNode.Properties = patch.DiscoveredNode.Properties
	}
}

func NewAddProvisionWatcherRequest(dto dtos.ProvisionWatcher) AddProvisionWatcherRequest {
	return AddProvisionWatcherRequest{
		BaseRequest:      dtoCommon.NewBaseRequest(),
		ProvisionWatcher: dto,
	}
}

func NewUpdateProvisionWatcherRequest(dto dtos.UpdateProvisionWatcher) UpdateProvisionWatcherRequest {
	return UpdateProvisionWatcherRequest{
		BaseRequest:      dtoCommon.NewBaseRequest(),
		ProvisionWatcher: dto,
	}
}
