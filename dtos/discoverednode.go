//
// Copyright (C) 2023-2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v4/models"

type DiscoveredNode struct {
	ProfileName string         `json:"profileName,omitempty" yaml:"profileName,omitempty" validate:"omitempty,edgex-dto-none-empty-string"`
	AdminState  string         `json:"adminState" yaml:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents  []AutoEvent    `json:"autoEvents,omitempty" yaml:"autoEvents,omitempty" validate:"dive"`
	Properties  map[string]any `json:"properties" yaml:"properties"`
}

type UpdateDiscoveredNode struct {
	ProfileName *string        `json:"profileName,omitempty" validate:"omitempty,len=0|edgex-dto-none-empty-string"`
	AdminState  *string        `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents  []AutoEvent    `json:"autoEvents" validate:"dive"`
	Properties  map[string]any `json:"properties"`
}

func ToDiscoveredNodeModel(dto DiscoveredNode) models.DiscoveredNode {
	m := models.DiscoveredNode{
		ProfileName: dto.ProfileName,
		AdminState:  models.AdminState(dto.AdminState),
		AutoEvents:  ToAutoEventModels(dto.AutoEvents),
		Properties:  dto.Properties,
	}
	if m.Properties == nil {
		m.Properties = make(map[string]any)
	}
	return m
}

func FromDiscoveredNodeModelToDTO(d models.DiscoveredNode) DiscoveredNode {
	dto := DiscoveredNode{
		ProfileName: d.ProfileName,
		AdminState:  string(d.AdminState),
		AutoEvents:  FromAutoEventModelsToDTOs(d.AutoEvents),
		Properties:  d.Properties,
	}
	if dto.Properties == nil {
		dto.Properties = make(map[string]any)
	}
	return dto
}

func FromDiscoveredNodeModelToUpdateDTO(d models.DiscoveredNode) UpdateDiscoveredNode {
	adminState := string(d.AdminState)
	return UpdateDiscoveredNode{
		ProfileName: &d.ProfileName,
		AdminState:  &adminState,
		AutoEvents:  FromAutoEventModelsToDTOs(d.AutoEvents),
		Properties:  d.Properties,
	}
}
