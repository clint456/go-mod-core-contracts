//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// ProfileScanRequest is the struct for requesting a profile for a specified node.
type ProfileScanRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	NodeName              string `json:"nodeName" validate:"required"`
	ProfileName           string `json:"profileName,omitempty"`
	Options               any    `json:"options,omitempty"`
}

// Validate satisfies the Validator interface
func (request ProfileScanRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddCommandRequest type
func (psr *ProfileScanRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		NodeName    string
		ProfileName string
		Options     any
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*psr = ProfileScanRequest(alias)

	if err := psr.Validate(); err != nil {
		return err
	}

	return nil
}
