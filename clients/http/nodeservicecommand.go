//
// Copyright (C) 2021-2024 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"

	"github.com/fxamacker/cbor/v2"
)

type nodeServiceCommandClient struct {
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewNodeServiceCommandClient creates an instance of nodeServiceCommandClient
func NewNodeServiceCommandClient(authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NodeServiceCommandClient {
	return &nodeServiceCommandClient{
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// GetCommand sends HTTP request to execute the Get command
func (client *nodeServiceCommandClient) GetCommand(ctx context.Context, baseUrl string, nodeName string, commandName string, queryParams string) (*responses.EventResponse, errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}
	res, contentType, edgeXerr := utils.GetRequestAndReturnBinaryRes(ctx, baseUrl, requestPath, params, client.authInjector)
	if edgeXerr != nil {
		return nil, errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	// If execute GetCommand with dsReturnEvent query parameter 'no', there will be no content returned in the http response.
	// So we can use the nil pointer to indicate that the HTTP response content is empty
	if len(res) == 0 {
		return nil, nil
	}
	response := &responses.EventResponse{}
	if contentType == common.ContentTypeCBOR {
		if err = cbor.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to decode the cbor response", err)
		}
	} else {
		if err = json.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to decode the json response", err)
		}
	}
	return response, nil
}

// SetCommand sends HTTP request to execute the Set command
func (client *nodeServiceCommandClient) SetCommand(ctx context.Context, baseUrl string, nodeName string, commandName string, queryParams string, settings map[string]string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &response, baseUrl, requestPath, params, settings, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// SetCommandWithObject invokes node service's set command API and the settings supports object value type
func (client *nodeServiceCommandClient) SetCommandWithObject(ctx context.Context, baseUrl string, nodeName string, commandName string, queryParams string, settings map[string]interface{}) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetNameFieldPath(commandName).BuildPath()
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &response, baseUrl, requestPath, params, settings, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *nodeServiceCommandClient) Discovery(ctx context.Context, baseUrl string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	err := utils.PostRequest(ctx, &response, baseUrl, common.ApiDiscoveryRoute, nil, "", client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// ProfileScan sends an HTTP POST request to the node service's profile scan API endpoint.
func (client *nodeServiceCommandClient) ProfileScan(ctx context.Context, baseUrl string, req requests.ProfileScanRequest) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	err := utils.PostRequestWithRawData(ctx, &response, baseUrl, common.ApiProfileScanRoute, nil, req, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *nodeServiceCommandClient) StopNodeDiscovery(ctx context.Context, baseUrl string, requestId string, queryParams map[string]string) (dtoCommon.BaseResponse, errors.EdgeX) {
	requestPath := common.ApiDiscoveryRoute
	if len(requestId) != 0 {
		requestPath = common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
			SetPath(common.ApiDiscoveryRoute).SetPath(common.RequestId).SetNameFieldPath(requestId).BuildPath()
	}
	response := dtoCommon.BaseResponse{}
	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v)
	}
	err := utils.DeleteRequestWithParams(ctx, &response, baseUrl, requestPath, params, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *nodeServiceCommandClient) StopProfileScan(ctx context.Context, baseUrl string, nodeName string, queryParams map[string]string) (dtoCommon.BaseResponse, errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiProfileScanRoute).SetPath(common.Node).SetPath(common.Name).SetNameFieldPath(nodeName).BuildPath()
	response := dtoCommon.BaseResponse{}
	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v)
	}
	err := utils.DeleteRequestWithParams(ctx, &response, baseUrl, requestPath, params, client.authInjector)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}
