//
// Copyright (C) 2025 IOTech Ltd
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

// ProfileTagsRequest defines the Request Content for PATCH UpdateProfileTags DTO.
type ProfileTagsRequest struct {
	dtoCommon.BaseRequest  `json:",inline"`
	dtos.UpdateProfileTags `json:",inline"`
}

// Validate satisfies the Validator interface
func (dt *ProfileTagsRequest) Validate() error {
	err := common.Validate(dt)
	if err != nil {
		// The UpdateProfileTags is the internal struct in Golang programming, not in the Profile model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".UpdateProfileTags", ""))
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "", err)
	}
	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the ProfileTagsRequest type
func (dt *ProfileTagsRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		dtos.UpdateProfileTags
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*dt = ProfileTagsRequest(alias)

	// validate ProfileTagsRequest DTO
	if err := dt.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceProfileModelTagsWithDTO replaces existing Profile's node resources/commands tags fields with the DTO patch.
// A tag whose value is null is deleted rather than set (see mergeTags).
func ReplaceProfileModelTagsWithDTO(dp *models.Profile, patch dtos.UpdateProfileTags) {
	// Convert UpdateTags list to maps for fast lookup
	resourceUpdate := convertUpdateTagsToMap(patch.Resources)
	commandUpdate := convertUpdateTagsToMap(patch.Commands)

	// Update Node Resource Tags
	for i := range dp.Resources {
		r := &dp.Resources[i]
		if tags, ok := resourceUpdate[r.Name]; ok && tags != nil {
			r.Tags = mergeTags(r.Tags, tags)
		}
	}

	// Update Node Command Tags
	for i := range dp.Commands {
		c := &dp.Commands[i]
		if tags, ok := commandUpdate[c.Name]; ok && tags != nil {
			c.Tags = mergeTags(c.Tags, tags)
		}
	}
}

func convertUpdateTagsToMap(list []dtos.UpdateTags) map[string]map[string]any {
	tagsMap := make(map[string]map[string]any)
	for _, u := range list {
		tagsMap[u.Name] = u.Tags
	}

	return tagsMap
}

// mergeTags merges src into dest, where a nil value deletes the key from dest.
func mergeTags(dest, src map[string]any) map[string]any {
	if dest == nil {
		dest = make(map[string]any)
	}
	for key, value := range src {
		if value == nil {
			delete(dest, key)
			continue
		}

		if sv, srcOk := value.(map[string]any); srcOk {
			dv, _ := dest[key].(map[string]any)
			dest[key] = mergeTags(dv, sv)
			continue
		}

		dest[key] = value
	}
	return dest
}
