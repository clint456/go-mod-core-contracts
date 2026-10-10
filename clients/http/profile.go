//
// Copyright (C) 2020-2025 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

type ProfileClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	resourcesCache        map[string]responses.ResourceResponse
	mux                   sync.RWMutex
	enableNameFieldEscape bool
}

// NewProfileClient creates an instance of ProfileClient
func NewProfileClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ProfileClient {
	return &ProfileClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		resourcesCache:        make(map[string]responses.ResourceResponse),
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewProfileClientWithUrlCallback creates an instance of ProfileClient with ClientBaseUrlFunc.
func NewProfileClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ProfileClient {
	return &ProfileClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		resourcesCache:        make(map[string]responses.ResourceResponse),
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new node profile
func (client *ProfileClient) Add(ctx context.Context, reqs []requests.ProfileRequest) ([]dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	var res []dtoCommon.BaseWithIdResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiProfileRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// Update updates node profile
func (client *ProfileClient) Update(ctx context.Context, reqs []requests.ProfileRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &res, baseUrl, common.ApiProfileRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AddByYaml adds new node profile by uploading a yaml file
func (client *ProfileClient) AddByYaml(ctx context.Context, yamlFilePath string) (dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	var res dtoCommon.BaseWithIdResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PostByFileRequest(ctx, &res, baseUrl, common.ApiProfileUploadFileRoute, yamlFilePath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// UpdateByYaml updates node profile by uploading a yaml file
func (client *ProfileClient) UpdateByYaml(ctx context.Context, yamlFilePath string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var res dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutByFileRequest(ctx, &res, baseUrl, common.ApiProfileUploadFileRoute, yamlFilePath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteByName deletes the node profile by name
func (client *ProfileClient) DeleteByName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &response, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// ProfileByName queries the node profile by name
func (client *ProfileClient) ProfileByName(ctx context.Context, name string) (res responses.ProfileResponse, edgexError errors.EdgeX) {
	return client.ProfileByNameWithQueryParams(ctx, name, nil)
}

// ProfileByNameWithQueryParams queries the node profile by name with query parameters
func (client *ProfileClient) ProfileByNameWithQueryParams(ctx context.Context, name string, queryParams map[string]string) (res responses.ProfileResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllProfiles queries the node profiles with offset, and limit
func (client *ProfileClient) AllProfiles(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiProfilesResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllProfileRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllProfileBasicInfos queries the node profile basic infos with offset, and limit
func (client *ProfileClient) AllProfileBasicInfos(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiProfileBasicInfoResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllProfileBasicInfoRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ProfilesByModel queries the node profiles with offset, limit and model
func (client *ProfileClient) ProfilesByModel(ctx context.Context, model string, offset int, limit int) (res responses.MultiProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiProfileRoute, common.Model, model)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ProfilesByManufacturer queries the node profiles with offset, limit and manufacturer
func (client *ProfileClient) ProfilesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (res responses.MultiProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiProfileRoute, common.Manufacturer, manufacturer)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ProfilesByManufacturerAndModel queries the node profiles with offset, limit, manufacturer and model
func (client *ProfileClient) ProfilesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (res responses.MultiProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiProfileRoute, common.Manufacturer, manufacturer, common.Model, model)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ResourceByProfileNameAndResourceName queries the node resource by profileName and resourceName
func (client *ProfileClient) ResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (res responses.ResourceResponse, edgexError errors.EdgeX) {
	resourceMapKey := fmt.Sprintf("%s:%s", profileName, resourceName)
	res, exists := client.resourceByMapKey(resourceMapKey)
	if exists {
		return res, nil
	}
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiResourceRoute).SetPath(common.Profile).SetNameFieldPath(profileName).SetPath(common.Resource).SetNameFieldPath(resourceName).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	client.setResourceWithMapKey(res, resourceMapKey)
	return res, nil
}

func (client *ProfileClient) resourceByMapKey(key string) (res responses.ResourceResponse, exists bool) {
	client.mux.RLock()
	defer client.mux.RUnlock()
	res, exists = client.resourcesCache[key]
	return
}

func (client *ProfileClient) setResourceWithMapKey(res responses.ResourceResponse, key string) {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache[key] = res
}

func (client *ProfileClient) CleanResourcesCache() {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache = make(map[string]responses.ResourceResponse)
}

// UpdateProfileBasicInfo updates existing profile's basic info
func (client *ProfileClient) UpdateProfileBasicInfo(ctx context.Context, reqs []requests.ProfileBasicInfoRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiProfileBasicInfoRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// UpdateProfileTags updates existing profile's node resources/commands tags
func (client *ProfileClient) UpdateProfileTags(ctx context.Context, profileName string, request requests.ProfileTagsRequest) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileRoute).SetPath(common.Name).SetNameFieldPath(profileName).SetPath(common.Tags).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PatchRequest(ctx, &response, baseUrl, requestPath, nil, request, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// AddProfileResource adds new node resource to an existing profile
func (client *ProfileClient) AddProfileResource(ctx context.Context, reqs []requests.AddResourceRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiProfileResourceRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// UpdateProfileResource updates existing node resource
func (client *ProfileClient) UpdateProfileResource(ctx context.Context, reqs []requests.UpdateResourceRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiProfileResourceRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteResourceByName deletes node resource by name
func (client *ProfileClient) DeleteResourceByName(ctx context.Context, profileName string, resourceName string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileRoute).SetPath(common.Name).SetNameFieldPath(profileName).SetPath(common.Resource).SetNameFieldPath(resourceName).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &response, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// AddProfileCommand adds new node command to an existing profile
func (client *ProfileClient) AddProfileCommand(ctx context.Context, reqs []requests.AddCommandRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiProfileCommandRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// UpdateProfileCommand updates existing node command
func (client *ProfileClient) UpdateProfileCommand(ctx context.Context, reqs []requests.UpdateCommandRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiProfileCommandRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteCommandByName deletes node command by name
func (client *ProfileClient) DeleteCommandByName(ctx context.Context, profileName string, commandName string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileRoute).SetPath(common.Name).SetNameFieldPath(profileName).SetPath(common.Command).SetNameFieldPath(commandName).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.DeleteRequest(ctx, &response, baseUrl, requestPath, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}
