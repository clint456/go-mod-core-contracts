//
// Copyright (C) 2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

type UpdateNodeProperties struct {
	Properties map[string]any `json:"properties" validate:"required,gt=0,dive"`
}
