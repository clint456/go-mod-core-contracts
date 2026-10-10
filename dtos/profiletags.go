//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

type UpdateProfileTags struct {
	Resources []UpdateTags `json:"resources,omitempty" validate:"dive"`
	Commands  []UpdateTags `json:"commands,omitempty" validate:"dive"`
}

type UpdateTags struct {
	Name string         `json:"name" validate:"required,edgex-dto-none-empty-string"`
	Tags map[string]any `json:"tags" validate:"required,gt=0,dive"`
}
