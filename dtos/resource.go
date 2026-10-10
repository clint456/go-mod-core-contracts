//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

type Resource struct {
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Name        string                 `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string"`
	IsHidden    bool                   `json:"isHidden" yaml:"isHidden"`
	Properties  ResourceProperties     `json:"properties" yaml:"properties"`
	Attributes  map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Tags        map[string]any         `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Validate satisfies the Validator interface
func (dr *Resource) Validate() error {
	err := common.Validate(dr)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid Resource.", err)
	}

	return nil
}

type UpdateResource struct {
	Description *string `json:"description"`
	Name        *string `json:"name" validate:"required,edgex-dto-none-empty-string"`
	IsHidden    *bool   `json:"isHidden"`
}

// ToResourceModel transforms the Resource DTO to the Resource model
func ToResourceModel(d Resource) models.Resource {
	return models.Resource{
		Description: d.Description,
		Name:        d.Name,
		IsHidden:    d.IsHidden,
		Properties:  ToResourcePropertiesModel(d.Properties),
		Attributes:  d.Attributes,
		Tags:        d.Tags,
	}
}

// ToResourceModels transforms the Resource DTOs to the Resource models
func ToResourceModels(resourceDTOs []Resource) []models.Resource {
	resourceModels := make([]models.Resource, len(resourceDTOs))
	for i, d := range resourceDTOs {
		resourceModels[i] = ToResourceModel(d)
	}
	return resourceModels
}

// FromResourceModelToDTO transforms the Resource model to the Resource DTO
func FromResourceModelToDTO(d models.Resource) Resource {
	return Resource{
		Description: d.Description,
		Name:        d.Name,
		IsHidden:    d.IsHidden,
		Properties:  FromResourcePropertiesModelToDTO(d.Properties),
		Attributes:  d.Attributes,
		Tags:        d.Tags,
	}
}

// FromResourceModelsToDTOs transforms the Resource models to the Resource DTOs
func FromResourceModelsToDTOs(resourceModels []models.Resource) []Resource {
	resourceDTOs := make([]Resource, len(resourceModels))
	for i, d := range resourceModels {
		resourceDTOs[i] = FromResourceModelToDTO(d)
	}
	return resourceDTOs
}
