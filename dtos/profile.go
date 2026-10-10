//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"errors"
	"fmt"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	edgexErrors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

type Profile struct {
	ProfileBasicInfo `json:",inline" yaml:",inline"`
	Resources        []Resource `json:"resources" yaml:"resources" validate:"dive"`
	Commands         []Command  `json:"commands" yaml:"commands" validate:"dive"`
	ApiVersion       string     `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
}

// Validate satisfies the Validator interface
func (dp *Profile) Validate() error {
	err := common.Validate(dp)
	if err != nil {
		// The ProfileBasicInfo is the internal struct in Golang programming, not in the Profile model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".ProfileBasicInfo", ""))
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "Invalid Profile.", err)
	}
	return ValidateProfileDTO(*dp)
}

// UnmarshalYAML implements the Unmarshaler interface for the Profile type
func (dp *Profile) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var alias struct {
		ProfileBasicInfo `yaml:",inline"`
		Resources        []Resource `yaml:"resources"`
		Commands         []Command  `yaml:"commands"`
		ApiVersion       string     `yaml:"apiVersion"`
	}
	if err := unmarshal(&alias); err != nil {
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "failed to unmarshal request body as YAML.", err)
	}
	*dp = Profile(alias)

	if err := dp.Validate(); err != nil {
		return edgexErrors.NewCommonEdgeXWrapper(err)
	}

	// Normalize resource's value type
	for i, resource := range dp.Resources {
		valueType, err := common.NormalizeValueType(resource.Properties.ValueType)
		if err != nil {
			return edgexErrors.NewCommonEdgeXWrapper(err)
		}
		dp.Resources[i].Properties.ValueType = valueType
	}
	return nil
}

// ToProfileModel transforms the Profile DTO to the Profile model
func ToProfileModel(profileDTO Profile) models.Profile {
	return models.Profile{
		DBTimestamp:  models.DBTimestamp(profileDTO.DBTimestamp),
		Id:           profileDTO.Id,
		Name:         profileDTO.Name,
		Description:  profileDTO.Description,
		Manufacturer: profileDTO.Manufacturer,
		Model:        profileDTO.Model,
		Labels:       profileDTO.Labels,
		Resources:    ToResourceModels(profileDTO.Resources),
		Commands:     ToCommandModels(profileDTO.Commands),
		ApiVersion:   profileDTO.ApiVersion,
	}
}

// FromProfileModelToDTO transforms the Profile Model to the Profile DTO
func FromProfileModelToDTO(profile models.Profile) Profile {
	if profile.ApiVersion == "" {
		profile.ApiVersion = common.ApiVersion
	}
	return Profile{
		ProfileBasicInfo: ProfileBasicInfo{
			DBTimestamp:  DBTimestamp(profile.DBTimestamp),
			Id:           profile.Id,
			Name:         profile.Name,
			Description:  profile.Description,
			Manufacturer: profile.Manufacturer,
			Model:        profile.Model,
			Labels:       profile.Labels,
		},
		Resources:  FromResourceModelsToDTOs(profile.Resources),
		Commands:   FromCommandModelsToDTOs(profile.Commands),
		ApiVersion: profile.ApiVersion,
	}
}

// FromProfileModelToBasicInfoDTO transforms the Profile Model to the ProfileBasicInfo DTO
func FromProfileModelToBasicInfoDTO(profile models.Profile) ProfileBasicInfo {
	return ProfileBasicInfo{
		DBTimestamp:  DBTimestamp(profile.DBTimestamp),
		Id:           profile.Id,
		Name:         profile.Name,
		Description:  profile.Description,
		Manufacturer: profile.Manufacturer,
		Model:        profile.Model,
		Labels:       profile.Labels,
	}
}

func ValidateProfileDTO(profile Profile) error {
	// resources validation
	dupCheck := make(map[string]bool)
	for _, resource := range profile.Resources {
		if resource.Properties.ValueType == common.ValueTypeBinary &&
			strings.Contains(resource.Properties.ReadWrite, common.ReadWrite_W) {
			return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, fmt.Sprintf("write permission not support %s value type for resource '%s'", common.ValueTypeBinary, resource.Name), nil)
		}
		// resource name should not duplicated
		if dupCheck[resource.Name] {
			return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, fmt.Sprintf("node resource %s is duplicated", resource.Name), nil)
		}
		dupCheck[resource.Name] = true
	}
	// commands validation
	dupCheck = make(map[string]bool)
	for _, command := range profile.Commands {
		// command name should not duplicated
		if dupCheck[command.Name] {
			return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, fmt.Sprintf("node command %s is duplicated", command.Name), nil)
		}
		dupCheck[command.Name] = true

		resourceOperations := command.ResourceOperations
		for _, ro := range resourceOperations {
			// ResourceOperations referenced in commands must exist
			if !resourcesContains(profile.Resources, ro.Resource) {
				return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, fmt.Sprintf("node command's resource %s doesn't match any node resource", ro.Resource), nil)
			}
			// Check the ReadWrite whether is align to the resource
			if !validReadWritePermission(profile.Resources, ro.Resource, command.ReadWrite) {
				return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, fmt.Sprintf("node command's ReadWrite permission '%s' doesn't align the node resource", command.ReadWrite), nil)
			}
		}
	}
	return nil
}

func resourcesContains(resources []Resource, name string) bool {
	contains := false
	for _, resource := range resources {
		if resource.Name == name {
			contains = true
			break
		}
	}
	return contains
}

func validReadWritePermission(resources []Resource, name string, readWrite string) bool {
	valid := true
	for _, resource := range resources {
		if resource.Name == name {
			if resource.Properties.ReadWrite != common.ReadWrite_RW && resource.Properties.ReadWrite != common.ReadWrite_WR &&
				resource.Properties.ReadWrite != readWrite {
				valid = false
				break
			}
		}
	}
	return valid
}
