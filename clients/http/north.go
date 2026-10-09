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

// NorthClient is the REST implementation of interfaces.NorthClient.
type NorthClient struct {
	baseUrlFunc           clients.ClientBaseUrlFunc
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewNorthClient creates an instance of NorthClient
func NewNorthClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NorthClient {
	return &NorthClient{
		baseUrlFunc:           clients.GetDefaultClientBaseUrlFunc(baseUrl),
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// NewNorthClientWithUrlCallback creates an instance of NorthClient with a ClientBaseUrlFunc.
func NewNorthClientWithUrlCallback(baseUrlFunc clients.ClientBaseUrlFunc, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NorthClient {
	return &NorthClient{
		baseUrlFunc:           baseUrlFunc,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new north channels.
func (client *NorthClient) Add(ctx context.Context, reqs []requests.NorthRequest) ([]dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	var res []dtoCommon.BaseWithIdResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PostRequestWithRawData(ctx, &res, baseUrl, common.ApiNorthRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// Update updates north channels.
func (client *NorthClient) Update(ctx context.Context, reqs []requests.NorthRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var res []dtoCommon.BaseResponse
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.PutRequest(ctx, &res, baseUrl, common.ApiNorthRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteByName deletes a north channel by name.
func (client *NorthClient) DeleteByName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNorthRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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

// NorthByName queries a north channel by name.
func (client *NorthClient) NorthByName(ctx context.Context, name string) (res responses.NorthResponse, edgexError errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNorthRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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

// AllNorths queries all north channels with offset and limit.
func (client *NorthClient) AllNorths(ctx context.Context, offset int, limit int) (res responses.MultiNorthsResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	baseUrl, err := clients.GetBaseUrl(client.baseUrlFunc)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &res, baseUrl, common.ApiAllNorthRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NorthsByServiceName queries north channels by service name with offset and limit.
func (client *NorthClient) NorthsByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiNorthsResponse, edgexError errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiNorthRoute).SetPath(common.Service).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
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
