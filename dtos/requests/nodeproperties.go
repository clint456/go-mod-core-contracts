//
// Copyright (C) 2026 IOTech Ltd
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

// NodePropertiesRequest defines the Request Content for PATCH UpdateNodeProperties DTO.
type NodePropertiesRequest struct {
	dtoCommon.BaseRequest     `json:",inline"`
	dtos.UpdateNodeProperties `json:",inline"`
}

// Validate satisfies the Validator interface
func (dp *NodePropertiesRequest) Validate() error {
	err := common.Validate(dp)
	if err != nil {
		// The UpdateNodeProperties is the internal struct in Golang programming, not in the Node model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".UpdateNodeProperties", ""))
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "", err)
	}
	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the NodePropertiesRequest type
func (dp *NodePropertiesRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		dtos.UpdateNodeProperties
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*dp = NodePropertiesRequest(alias)

	// validate NodePropertiesRequest DTO
	if err := dp.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceNodeModelPropertiesWithDTO merges the properties patch into the existing Node.
// A property whose value is null is deleted rather than set (see mergeTags).
func ReplaceNodeModelPropertiesWithDTO(node *models.Node, patch dtos.UpdateNodeProperties) {
	node.Properties = mergeTags(node.Properties, patch.Properties)
}
