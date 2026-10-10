//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v4/models"

type Command struct {
	Name               string              `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string"`
	IsHidden           bool                `json:"isHidden" yaml:"isHidden"`
	ReadWrite          string              `json:"readWrite" yaml:"readWrite" validate:"required,oneof='R' 'W' 'RW' 'WR'"`
	ResourceOperations []ResourceOperation `json:"resourceOperations" yaml:"resourceOperations" validate:"gt=0,dive"`
	Tags               map[string]any      `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type UpdateCommand struct {
	Name     *string `json:"name" validate:"required,edgex-dto-none-empty-string"`
	IsHidden *bool   `json:"isHidden"`
}

// ToCommandModel transforms the Command DTO to the Command model
func ToCommandModel(dto Command) models.Command {
	resourceOperations := ToResourceOperationModels(dto.ResourceOperations)

	return models.Command{
		Name:               dto.Name,
		IsHidden:           dto.IsHidden,
		ReadWrite:          dto.ReadWrite,
		ResourceOperations: resourceOperations,
		Tags:               dto.Tags,
	}
}

// ToCommandModels transforms the Command DTOs to the Command models
func ToCommandModels(commandDTOs []Command) []models.Command {
	commandModels := make([]models.Command, len(commandDTOs))
	for i, p := range commandDTOs {
		commandModels[i] = ToCommandModel(p)
	}
	return commandModels
}

// FromCommandModelToDTO transforms the Command model to the Command DTO
func FromCommandModelToDTO(d models.Command) Command {
	resourceOperations := make([]ResourceOperation, len(d.ResourceOperations))
	for i, ro := range d.ResourceOperations {
		resourceOperations[i] = FromResourceOperationModelToDTO(ro)
	}
	return Command{
		Name:               d.Name,
		IsHidden:           d.IsHidden,
		ReadWrite:          d.ReadWrite,
		ResourceOperations: resourceOperations,
		Tags:               d.Tags,
	}
}

// FromCommandModelsToDTOs transforms the Command models to the Command DTOs
func FromCommandModelsToDTOs(commandModels []models.Command) []Command {
	commandDTOs := make([]Command, len(commandModels))
	for i, p := range commandModels {
		commandDTOs[i] = FromCommandModelToDTO(p)
	}
	return commandDTOs
}
