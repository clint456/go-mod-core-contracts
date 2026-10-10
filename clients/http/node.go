//
// Copyright (C) 2020-2021 Unknown author
// Copyright (C) 2023 Intel Corporation
// Copyright (C) 2024-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

type NodeClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewNodeClient creates an instance of NodeClient
func NewNodeClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NodeClient {
	return &NodeClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewNodeClientWithUrlCallback creates an instance of NodeClient with ClientBaseUrlFunc.
func NewNodeClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NodeClient {
	return &NodeClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

func (dc NodeClient) Add(ctx context.Context, reqs []requests.AddNodeRequest) (res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiNodeRoute, nil, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) AddWithQueryParams(ctx context.Context, reqs []requests.AddNodeRequest, queryParams map[string]string) (res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiNodeRoute, requestParams, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) Update(ctx context.Context, reqs []requests.UpdateNodeRequest) (res []dtoCommon.BaseResponse, err errors.EdgeX) {
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiNodeRoute, nil, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) UpdateWithQueryParams(ctx context.Context, reqs []requests.UpdateNodeRequest, queryParams map[string]string) (res []dtoCommon.BaseResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiNodeRoute, requestParams, reqs, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) UpdateNodeProperties(ctx context.Context, nodeName string, request requests.NodePropertiesRequest) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetPath(common.Properties).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, requestPath, nil, request, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) AllNodes(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiNodesResponse, err errors.EdgeX) {
	return dc.AllNodesWithQueryParams(ctx, labels, offset, limit, nil)
}

func (dc NodeClient) AllNodesWithQueryParams(ctx context.Context, labels []string, offset int, limit int, queryParams map[string]string) (res responses.MultiNodesResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllNodeRoute, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) AllNodesWithChildren(ctx context.Context, parent string, maxLevels uint, labels []string, offset int, limit int) (res responses.MultiNodesResponse, err errors.EdgeX) {
	return dc.AllNodesWithChildrenWithQueryParams(ctx, parent, maxLevels, labels, offset, limit, nil)
}

func (dc NodeClient) AllNodesWithChildrenWithQueryParams(ctx context.Context, parent string, maxLevels uint, labels []string, offset int, limit int, queryParams map[string]string) (res responses.MultiNodesResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.DescendantsOf, parent)
	requestParams.Set(common.MaxLevels, strconv.FormatUint(uint64(maxLevels), 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllNodeRoute, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) NodeNameExists(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Check).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) NodeByName(ctx context.Context, name string) (res responses.NodeResponse, err errors.EdgeX) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) DeleteNodeByName(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, path, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) NodesByProfileName(ctx context.Context, name string, offset int, limit int) (res responses.MultiNodesResponse, err errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Profile).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc NodeClient) NodesByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiNodesResponse, err errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(dc.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Service).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, requestParams, dc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
