//
// Copyright (C) 2020-2021 Unknown author
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddNodeServices(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiNodeServiceRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()

	client := NewNodeServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddNodeServiceRequest{})

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestPatchNodeServices(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiNodeServiceRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewNodeServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateNodeServiceRequest{})
	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestQueryAllNodeServices(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllNodeServiceRoute, responses.MultiNodeServicesResponse{})
	defer ts.Close()

	client := NewNodeServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllNodeServices(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiNodeServicesResponse{}, res)
}

func TestQueryNodeServiceByName(t *testing.T) {
	nodeServiceName := "nodeService"
	path := path.Join(common.ApiNodeServiceRoute, common.Name, nodeServiceName)

	ts := newTestServer(http.MethodGet, path, responses.NodeResponse{})
	defer ts.Close()

	client := NewNodeServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.NodeServiceByName(context.Background(), nodeServiceName)
	require.NoError(t, err)
	assert.IsType(t, responses.NodeServiceResponse{}, res)
}

func TestDeleteNodeServiceByName(t *testing.T) {
	nodeServiceName := "nodeService"
	path := path.Join(common.ApiNodeServiceRoute, common.Name, nodeServiceName)

	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewNodeServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteByName(context.Background(), nodeServiceName)
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}
