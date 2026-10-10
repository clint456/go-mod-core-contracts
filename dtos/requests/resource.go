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

// AddResourceRequest defines the Request Content for POST Resource DTO.
type AddResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string        `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	Resource              dtos.Resource `json:"resource"`
}

func (request AddResourceRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddResourceReques type
func (dr *AddResourceRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Resource    dtos.Resource
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dr = AddResourceRequest(alias)

	if err := dr.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateResourceRequest defines the Request Content for PATCH Resource DTO.
type UpdateResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string              `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	Resource              dtos.UpdateResource `json:"resource"`
}

func (request UpdateResourceRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateResourceRequest type
func (dr *UpdateResourceRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Resource    dtos.UpdateResource
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dr = UpdateResourceRequest(alias)

	if err := dr.Validate(); err != nil {
		return err
	}

	return nil
}

// ReplaceResourceModelFieldsWithDTO replace existing Resource's fields with DTO patch
func ReplaceResourceModelFieldsWithDTO(dr *models.Resource, patch dtos.UpdateResource) {
	if patch.Description != nil {
		dr.Description = *patch.Description
	}
	if patch.IsHidden != nil {
		dr.IsHidden = *patch.IsHidden
	}
}
