//
// Copyright (C) 2020-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// ProfileClient defines the interface for interactions with the Profile endpoint on the EdgeX Foundry core-metadata service.
type ProfileClient interface {
	// Add adds new profiles
	Add(ctx context.Context, reqs []requests.ProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
	// Update updates profiles
	Update(ctx context.Context, reqs []requests.ProfileRequest) ([]common.BaseResponse, errors.EdgeX)
	// AddByYaml adds new profile by uploading a file in YAML format
	AddByYaml(ctx context.Context, yamlFilePath string) (common.BaseWithIdResponse, errors.EdgeX)
	// UpdateByYaml updates profile by uploading a file in YAML format
	UpdateByYaml(ctx context.Context, yamlFilePath string) (common.BaseResponse, errors.EdgeX)
	// DeleteByName deletes profile by name
	DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
	// ProfileByName queries profile by name
	ProfileByName(ctx context.Context, name string) (responses.ProfileResponse, errors.EdgeX)
	// ProfileByNameWithQueryParams queries profile by name with query parameters
	ProfileByNameWithQueryParams(ctx context.Context, name string, queryParams map[string]string) (responses.ProfileResponse, errors.EdgeX)
	// AllProfiles queries all profiles
	AllProfiles(ctx context.Context, labels []string, offset int, limit int) (responses.MultiProfilesResponse, errors.EdgeX)
	// AllProfileBasicInfos queries all profile basic infos
	AllProfileBasicInfos(ctx context.Context, labels []string, offset int, limit int) (responses.MultiProfileBasicInfoResponse, errors.EdgeX)
	// ProfilesByModel queries profiles by model
	ProfilesByModel(ctx context.Context, model string, offset int, limit int) (responses.MultiProfilesResponse, errors.EdgeX)
	// ProfilesByManufacturer queries profiles by manufacturer
	ProfilesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (responses.MultiProfilesResponse, errors.EdgeX)
	// ProfilesByManufacturerAndModel queries profiles by manufacturer and model
	ProfilesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (responses.MultiProfilesResponse, errors.EdgeX)
	// ResourceByProfileNameAndResourceName queries the node resource by profileName and resourceName
	ResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (responses.ResourceResponse, errors.EdgeX)
	// UpdateProfileBasicInfo updates existing profile's basic info
	UpdateProfileBasicInfo(ctx context.Context, reqs []requests.ProfileBasicInfoRequest) ([]common.BaseResponse, errors.EdgeX)
	// UpdateProfileTags updates existing profile's node resources/commands tags
	UpdateProfileTags(ctx context.Context, profileName string, request requests.ProfileTagsRequest) (common.BaseResponse, errors.EdgeX)
	// AddProfileResource adds new node resource to an existing profile
	AddProfileResource(ctx context.Context, reqs []requests.AddResourceRequest) ([]common.BaseResponse, errors.EdgeX)
	// UpdateProfileResource updates existing node resource
	UpdateProfileResource(ctx context.Context, reqs []requests.UpdateResourceRequest) ([]common.BaseResponse, errors.EdgeX)
	// DeleteResourceByName deletes node resource by name
	DeleteResourceByName(ctx context.Context, profileName string, resourceName string) (common.BaseResponse, errors.EdgeX)
	// AddProfileCommand adds new node command to an existing profile
	AddProfileCommand(ctx context.Context, reqs []requests.AddCommandRequest) ([]common.BaseResponse, errors.EdgeX)
	// UpdateProfileCommand updates existing node command
	UpdateProfileCommand(ctx context.Context, reqs []requests.UpdateCommandRequest) ([]common.BaseResponse, errors.EdgeX)
	// DeleteCommandByName deletes node command by name
	DeleteCommandByName(ctx context.Context, profileName string, commandName string) (common.BaseResponse, errors.EdgeX)
}
