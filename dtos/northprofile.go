//
// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// NorthProfile is the DTO for a northbound profile (the northbound equivalent of a DeviceProfile).
type NorthProfile struct {
	DBTimestamp    `json:",inline"`
	ApiVersion     string          `json:"apiVersion,omitempty"`
	Description    string          `json:"description,omitempty"`
	Id             string          `json:"id,omitempty"`
	Name           string          `json:"name" validate:"required,edgex-dto-none-empty-string"`
	Manufacturer   string          `json:"manufacturer,omitempty"`
	Model          string          `json:"model,omitempty"`
	Labels         []string        `json:"labels,omitempty"`
	NorthResources []NorthResource `json:"northResources,omitempty" validate:"dive"`
	NorthCommands  []NorthCommand  `json:"northCommands,omitempty" validate:"dive"`
}

// NorthResource is the DTO for a northbound resource.
type NorthResource struct {
	Description string             `json:"description,omitempty"`
	Name        string             `json:"name" validate:"required,edgex-dto-none-empty-string"`
	IsHidden    bool               `json:"isHidden,omitempty"`
	Properties  ResourceProperties `json:"properties" validate:"required"`
	Attributes  map[string]any     `json:"attributes,omitempty"`
	Tags        map[string]any     `json:"tags,omitempty"`
}

// NorthCommand is the DTO for a northbound command.
type NorthCommand struct {
	Name               string              `json:"name" validate:"required,edgex-dto-none-empty-string"`
	IsHidden           bool                `json:"isHidden,omitempty"`
	ReadWrite          string              `json:"readWrite" validate:"required,oneof='R' 'W' 'RW' 'WR'"`
	ResourceOperations []ResourceOperation `json:"resourceOperations,omitempty" validate:"dive"`
	Tags               map[string]any      `json:"tags,omitempty"`
}

// ToNorthResourceModel transforms the NorthResource DTO to the NorthResource model
func ToNorthResourceModel(dto NorthResource) models.NorthResource {
	return models.NorthResource{
		Description: dto.Description,
		Name:        dto.Name,
		IsHidden:    dto.IsHidden,
		Properties:  ToResourcePropertiesModel(dto.Properties),
		Attributes:  dto.Attributes,
		Tags:        dto.Tags,
	}
}

// ToNorthResourceModels transforms the NorthResource DTO array to the NorthResource model array
func ToNorthResourceModels(dtos []NorthResource) []models.NorthResource {
	models := make([]models.NorthResource, len(dtos))
	for i, dto := range dtos {
		models[i] = ToNorthResourceModel(dto)
	}
	return models
}

// FromNorthResourceModelToDTO transforms the NorthResource model to the NorthResource DTO
func FromNorthResourceModelToDTO(resource models.NorthResource) NorthResource {
	return NorthResource{
		Description: resource.Description,
		Name:        resource.Name,
		IsHidden:    resource.IsHidden,
		Properties:  FromResourcePropertiesModelToDTO(resource.Properties),
		Attributes:  resource.Attributes,
		Tags:        resource.Tags,
	}
}

// FromNorthResourceModelsToDTOs transforms the NorthResource model array to the NorthResource DTO array
func FromNorthResourceModelsToDTOs(resources []models.NorthResource) []NorthResource {
	dtos := make([]NorthResource, len(resources))
	for i, resource := range resources {
		dtos[i] = FromNorthResourceModelToDTO(resource)
	}
	return dtos
}

// ToNorthCommandModel transforms the NorthCommand DTO to the NorthCommand model
func ToNorthCommandModel(dto NorthCommand) models.NorthCommand {
	return models.NorthCommand{
		Name:               dto.Name,
		IsHidden:           dto.IsHidden,
		ReadWrite:          dto.ReadWrite,
		ResourceOperations: ToResourceOperationModels(dto.ResourceOperations),
		Tags:               dto.Tags,
	}
}

// ToNorthCommandModels transforms the NorthCommand DTO array to the NorthCommand model array
func ToNorthCommandModels(dtos []NorthCommand) []models.NorthCommand {
	models := make([]models.NorthCommand, len(dtos))
	for i, dto := range dtos {
		models[i] = ToNorthCommandModel(dto)
	}
	return models
}

// FromNorthCommandModelToDTO transforms the NorthCommand model to the NorthCommand DTO
func FromNorthCommandModelToDTO(command models.NorthCommand) NorthCommand {
	resourceOperations := make([]ResourceOperation, len(command.ResourceOperations))
	for i, ro := range command.ResourceOperations {
		resourceOperations[i] = FromResourceOperationModelToDTO(ro)
	}
	return NorthCommand{
		Name:               command.Name,
		IsHidden:           command.IsHidden,
		ReadWrite:          command.ReadWrite,
		ResourceOperations: resourceOperations,
		Tags:               command.Tags,
	}
}

// FromNorthCommandModelsToDTOs transforms the NorthCommand model array to the NorthCommand DTO array
func FromNorthCommandModelsToDTOs(commands []models.NorthCommand) []NorthCommand {
	dtos := make([]NorthCommand, len(commands))
	for i, command := range commands {
		dtos[i] = FromNorthCommandModelToDTO(command)
	}
	return dtos
}

// ToNorthProfileModel transforms the NorthProfile DTO to the NorthProfile model
func ToNorthProfileModel(dto NorthProfile) models.NorthProfile {
	return models.NorthProfile{
		DBTimestamp:    models.DBTimestamp(dto.DBTimestamp),
		ApiVersion:     dto.ApiVersion,
		Description:    dto.Description,
		Id:             dto.Id,
		Name:           dto.Name,
		Manufacturer:   dto.Manufacturer,
		Model:          dto.Model,
		Labels:         dto.Labels,
		NorthResources: ToNorthResourceModels(dto.NorthResources),
		NorthCommands:  ToNorthCommandModels(dto.NorthCommands),
	}
}

// FromNorthProfileModelToDTO transforms the NorthProfile model to the NorthProfile DTO
func FromNorthProfileModelToDTO(profile models.NorthProfile) NorthProfile {
	if profile.ApiVersion == "" {
		profile.ApiVersion = common.ApiVersion
	}
	return NorthProfile{
		DBTimestamp:    DBTimestamp(profile.DBTimestamp),
		ApiVersion:     profile.ApiVersion,
		Description:    profile.Description,
		Id:             profile.Id,
		Name:           profile.Name,
		Manufacturer:   profile.Manufacturer,
		Model:          profile.Model,
		Labels:         profile.Labels,
		NorthResources: FromNorthResourceModelsToDTOs(profile.NorthResources),
		NorthCommands:  FromNorthCommandModelsToDTOs(profile.NorthCommands),
	}
}
