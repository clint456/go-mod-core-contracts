//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

type NodeService struct {
	DBTimestamp `json:",inline"`
	Id          string         `json:"id,omitempty" validate:"omitempty,uuid"`
	Name        string         `json:"name" validate:"required,edgex-dto-none-empty-string"`
	Description string         `json:"description,omitempty"`
	Labels      []string       `json:"labels,omitempty"`
	BaseAddress string         `json:"baseAddress" validate:"required,uri"`
	AdminState  string         `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	Properties  map[string]any `json:"properties" yaml:"properties"`
}

type UpdateNodeService struct {
	Id          *string        `json:"id" validate:"required_without=Name,edgex-dto-uuid"`
	Name        *string        `json:"name" validate:"required_without=Id,edgex-dto-none-empty-string"`
	Description *string        `json:"description"`
	BaseAddress *string        `json:"baseAddress" validate:"omitempty,uri"`
	Labels      []string       `json:"labels"`
	AdminState  *string        `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	Properties  map[string]any `json:"properties"`
}

// ToNodeServiceModel transforms the NodeService DTO to the NodeService Model
func ToNodeServiceModel(dto NodeService) models.NodeService {
	var ds models.NodeService
	ds.Id = dto.Id
	ds.Name = dto.Name
	ds.Description = dto.Description
	ds.BaseAddress = dto.BaseAddress
	ds.Labels = dto.Labels
	ds.AdminState = models.AdminState(dto.AdminState)
	ds.Properties = dto.Properties
	if ds.Properties == nil {
		ds.Properties = make(map[string]any)
	}
	return ds
}

// FromNodeServiceModelToDTO transforms the NodeService Model to the NodeService DTO
func FromNodeServiceModelToDTO(ds models.NodeService) NodeService {
	var dto NodeService
	dto.DBTimestamp = DBTimestamp(ds.DBTimestamp)
	dto.Id = ds.Id
	dto.Name = ds.Name
	dto.Description = ds.Description
	dto.BaseAddress = ds.BaseAddress
	dto.Labels = ds.Labels
	dto.AdminState = string(ds.AdminState)
	dto.Properties = ds.Properties
	if dto.Properties == nil {
		dto.Properties = make(map[string]any)
	}
	return dto
}

// FromNodeServiceModelToUpdateDTO transforms the NodeService Model to the UpdateNodeService DTO
func FromNodeServiceModelToUpdateDTO(ds models.NodeService) UpdateNodeService {
	adminState := string(ds.AdminState)
	dto := UpdateNodeService{
		Id:          &ds.Id,
		Name:        &ds.Name,
		Description: &ds.Description,
		Labels:      ds.Labels,
		BaseAddress: &ds.BaseAddress,
		AdminState:  &adminState,
		Properties:  ds.Properties,
	}
	return dto
}
