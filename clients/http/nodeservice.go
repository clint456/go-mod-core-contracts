//
// Copyright (C) 2020-2021 Unknown author
// Copyright (C) 2023 Intel Corporation
// Copyright (C) 2025 IOTech Ltd
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

type NodeServiceClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewNodeServiceClient creates an instance of NodeServiceClient
func NewNodeServiceClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NodeServiceClient {
	return &NodeServiceClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewNodeServiceClientWithUrlCallback creates an instance of NodeServiceClient with ClientBaseUrlFunc.
func NewNodeServiceClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NodeServiceClient {
	return &NodeServiceClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

func (dsc NodeServiceClient) Add(ctx context.Context, reqs []requests.AddNodeServiceRequest) (
	res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	baseUrl, goErr := clients.GetBaseUrl(dsc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiNodeServiceRoute, nil, reqs, dsc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc NodeServiceClient) Update(ctx context.Context, reqs []requests.UpdateNodeServiceRequest) (
	res []dtoCommon.BaseResponse, err errors.EdgeX) {
	baseUrl, goErr := clients.GetBaseUrl(dsc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.PatchRequest(ctx, &res, baseUrl, common.ApiNodeServiceRoute, nil, reqs, dsc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc NodeServiceClient) AllNodeServices(ctx context.Context, labels []string, offset int, limit int) (
	res responses.MultiNodeServicesResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, goErr := clients.GetBaseUrl(dsc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllNodeServiceRoute, requestParams, dsc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc NodeServiceClient) NodeServiceByName(ctx context.Context, name string) (
	res responses.NodeServiceResponse, err errors.EdgeX) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dsc.enableNameFieldEscape).
		SetPath(common.ApiNodeServiceRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dsc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, path, nil, dsc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc NodeServiceClient) DeleteByName(ctx context.Context, name string) (
	res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := common.NewPathBuilder().EnableNameFieldEscape(dsc.enableNameFieldEscape).
		SetPath(common.ApiNodeServiceRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, goErr := clients.GetBaseUrl(dsc.baseUrlFunc)
	if goErr != nil {
		return res, errors.NewCommonEdgeXWrapper(goErr)
	}
	err = utils.DeleteRequest(ctx, &res, baseUrl, path, dsc.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
