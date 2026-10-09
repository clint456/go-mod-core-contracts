//
// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// NorthProfileClient is the REST implementation of interfaces.NorthProfileClient.
type NorthProfileClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewNorthProfileClient creates an instance of NorthProfileClient
func NewNorthProfileClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NorthProfileClient {
	return &NorthProfileClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewNorthProfileClientWithUrlCallback creates an instance of NorthProfileClient with a ClientBaseUrlFunc.
func NewNorthProfileClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NorthProfileClient {
	return &NorthProfileClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new north profiles.
func (client *NorthProfileClient) Add(ctx context.Context, reqs []requests.NorthProfileRequest) ([]dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	var res []dtoCommon.BaseWithIdResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiNorthProfileRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// Update updates north profiles.
func (client *NorthProfileClient) Update(ctx context.Context, reqs []requests.NorthProfileRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &res, baseUrl, common.ApiNorthProfileRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteByName deletes a north profile by name.
func (client *NorthProfileClient) DeleteByName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNorthProfileRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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

// NorthProfileByName queries a north profile by name.
func (client *NorthProfileClient) NorthProfileByName(ctx context.Context, name string) (res responses.NorthProfileResponse, edgexError errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNorthProfileRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, requestPath, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllNorthProfiles queries all north profiles with offset and limit.
func (client *NorthProfileClient) AllNorthProfiles(ctx context.Context, offset int, limit int) (res responses.MultiNorthProfilesResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllNorthProfileRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
