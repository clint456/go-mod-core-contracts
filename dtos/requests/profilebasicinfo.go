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

// ProfileBasicInfoRequest defines the Request Content for PATCH UpdateProfileBasicInfo DTO.
type ProfileBasicInfoRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	BasicInfo             dtos.UpdateProfileBasicInfo `json:"basicinfo"`
}

// Validate satisfies the Validator interface
func (d ProfileBasicInfoRequest) Validate() error {
	err := common.Validate(d)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateNodeRequest type
func (d *ProfileBasicInfoRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		BasicInfo dtos.UpdateProfileBasicInfo
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*d = ProfileBasicInfoRequest(alias)

	// validate ProfileBasicInfoRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceProfileModelBasicInfoFieldsWithDTO replace existing Profile's basic info fields with DTO patch
func ReplaceProfileModelBasicInfoFieldsWithDTO(dp *models.Profile, patch dtos.UpdateProfileBasicInfo) {
	if patch.Description != nil {
		dp.Description = *patch.Description
	}
	if patch.Manufacturer != nil {
		dp.Manufacturer = *patch.Manufacturer
	}
	if patch.Model != nil {
		dp.Model = *patch.Model
	}
	if patch.Labels != nil {
		dp.Labels = patch.Labels
	}
}
