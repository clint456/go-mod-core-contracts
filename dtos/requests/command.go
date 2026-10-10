//
// Copyright (C) 2022 IOTech Ltd
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

// AddCommandRequest defines the Request Content for POST Command DTO.
type AddCommandRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string       `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	Command               dtos.Command `json:"command"`
}

// Validate satisfies the Validator interface
func (request AddCommandRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddCommandRequest type
func (dc *AddCommandRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Command     dtos.Command
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dc = AddCommandRequest(alias)

	if err := dc.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateCommandRequest defines the Request Content for PATCH Command DTO.
type UpdateCommandRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string             `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	Command               dtos.UpdateCommand `json:"command"`
}

// Validate satisfies the Validator interface
func (request UpdateCommandRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateCommandRequest type
func (dc *UpdateCommandRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Command     dtos.UpdateCommand
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dc = UpdateCommandRequest(alias)

	if err := dc.Validate(); err != nil {
		return err
	}

	return nil
}

// ReplaceCommandModelFieldsWithDTO replace existing Command's fields with DTO patch
func ReplaceCommandModelFieldsWithDTO(dc *models.Command, patch dtos.UpdateCommand) {
	if patch.IsHidden != nil {
		dc.IsHidden = *patch.IsHidden
	}
}
