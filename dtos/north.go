//
// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// North is the DTO for a northbound channel.
type North struct {
	DBTimestamp    `json:",inline"`
	Id             string                        `json:"id,omitempty"`
	Name           string                        `json:"name" validate:"required,edgex-dto-none-empty-string"`
	Description    string                        `json:"description,omitempty"`
	AdminState     string                        `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	OperatingState string                        `json:"operatingState" validate:"oneof='UP' 'DOWN'"`
	Protocols      map[string]ProtocolProperties `json:"protocols" validate:"required"`
	Labels         []string                      `json:"labels,omitempty"`
	ServiceName    string                        `json:"serviceName"`
	ProfileName    string                        `json:"profileName"`
	Tags           map[string]any                `json:"tags,omitempty"`
	Properties     map[string]any                `json:"properties,omitempty"`
}

// ToNorthModel transforms the North DTO to the North model
func ToNorthModel(dto North) models.North {
	return models.North{
		DBTimestamp:    models.DBTimestamp(dto.DBTimestamp),
		Id:             dto.Id,
		Name:           dto.Name,
		Description:    dto.Description,
		AdminState:     models.AdminState(dto.AdminState),
		OperatingState: models.OperatingState(dto.OperatingState),
		Protocols:      ToProtocolModels(dto.Protocols),
		Labels:         dto.Labels,
		ServiceName:    dto.ServiceName,
		ProfileName:    dto.ProfileName,
		Tags:           dto.Tags,
		Properties:     dto.Properties,
	}
}

// ToNorthModels transforms the North DTO array to the North model array
func ToNorthModels(dtos []North) []models.North {
	models := make([]models.North, len(dtos))
	for i, dto := range dtos {
		models[i] = ToNorthModel(dto)
	}
	return models
}

// FromNorthModelToDTO transforms the North model to the North DTO
func FromNorthModelToDTO(north models.North) North {
	var protocols map[string]ProtocolProperties
	if north.Protocols != nil {
		protocols = make(map[string]ProtocolProperties, len(north.Protocols))
		for name, properties := range north.Protocols {
			protocols[name] = ProtocolProperties(properties)
		}
	}
	return North{
		DBTimestamp:    DBTimestamp(north.DBTimestamp),
		Id:             north.Id,
		Name:           north.Name,
		Description:    north.Description,
		AdminState:     string(north.AdminState),
		OperatingState: string(north.OperatingState),
		Protocols:      protocols,
		Labels:         north.Labels,
		ServiceName:    north.ServiceName,
		ProfileName:    north.ProfileName,
		Tags:           north.Tags,
		Properties:     north.Properties,
	}
}

// FromNorthModelsToDTOs transforms the North model array to the North DTO array
func FromNorthModelsToDTOs(norths []models.North) []North {
	dtos := make([]North, len(norths))
	for i, north := range norths {
		dtos[i] = FromNorthModelToDTO(north)
	}
	return dtos
}
