//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	edgexErrors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// ProfileRequest defines the Request Content for POST Profile DTO.
type ProfileRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Profile               dtos.Profile `json:"profile"`
}

// Validate satisfies the Validator interface
func (dp ProfileRequest) Validate() error {
	err := common.Validate(dp)
	if err != nil {
		// The ProfileBasicInfo is the internal struct in Golang programming, not in the Profile model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".ProfileBasicInfo", ""))
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "", err)
	}
	return dp.Profile.Validate()
}

// UnmarshalJSON implements the Unmarshaler interface for the ProfileRequest type
func (dp *ProfileRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Profile dtos.Profile
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*dp = ProfileRequest(alias)

	// validate ProfileRequest DTO
	if err := dp.Validate(); err != nil {
		return err
	}

	// Normalize resource's value type
	for i, resource := range dp.Profile.Resources {
		valueType, err := common.NormalizeValueType(resource.Properties.ValueType)
		if err != nil {
			return edgexErrors.NewCommonEdgeXWrapper(err)
		}
		dp.Profile.Resources[i].Properties.ValueType = valueType
	}
	return nil
}

// ProfileReqToProfileModel transforms the ProfileRequest DTO to the Profile model
func ProfileReqToProfileModel(addReq ProfileRequest) (Profiles models.Profile) {
	return dtos.ToProfileModel(addReq.Profile)
}

// ProfileReqToProfileModels transforms the ProfileRequest DTO array to the Profile model array
func ProfileReqToProfileModels(addRequests []ProfileRequest) (Profiles []models.Profile) {
	for _, req := range addRequests {
		dp := ProfileReqToProfileModel(req)
		Profiles = append(Profiles, dp)
	}
	return Profiles
}

func NewProfileRequest(dto dtos.Profile) ProfileRequest {
	return ProfileRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Profile:     dto,
	}
}
