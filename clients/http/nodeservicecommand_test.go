//
// Copyright (C) 2021-2024 IOTech Ltd
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
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TestUUID        = "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	TestTimestamp   = 1594963842
	TestCommandName = "TestCommand"
	TestNodeName    = "TestNode"
	TestProfileName = "TestProfileName"
)

var testEventDTO = dtos.Event{
	Versionable: dtoCommon.Versionable{ApiVersion: common.ApiVersion},
	Id:          TestUUID,
	NodeName:    TestNodeName,
	ProfileName: TestProfileName,
	Origin:      TestTimestamp,
	Tags: map[string]interface{}{
		"GatewayID": "Houston-0001",
		"Latitude":  "29.630771",
		"Longitude": "-95.377603",
	},
	Readings:   make([]dtos.BaseReading, 0),
	Extensions: make(map[string]any),
}

func TestGetCommand(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := responses.NewEventResponse(requestId, "", http.StatusOK, testEventDTO)
	ts := newTestServer(http.MethodGet, common.ApiNodeRoute+"/"+common.Name+"/"+TestNodeName+"/"+TestCommandName, expectedResponse)
	defer ts.Close()

	client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
	res, err := client.GetCommand(context.Background(), ts.URL, TestNodeName, TestCommandName, "")

	require.NoError(t, err)
	assert.Equal(t, expectedResponse, *res)
}

func TestSetCommand(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiNodeRoute+"/"+common.Name+"/"+TestNodeName+"/"+TestCommandName, expectedResponse)
	defer ts.Close()

	client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
	res, err := client.SetCommand(context.Background(), ts.URL, TestNodeName, TestCommandName, "", nil)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestSetCommandWithObject(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiNodeRoute+"/"+common.Name+"/"+TestNodeName+"/"+TestCommandName, expectedResponse)
	defer ts.Close()
	settings := map[string]interface{}{
		"SwitchButton": map[string]interface{}{
			"kind":  "button",
			"value": "on",
		},
	}

	client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
	res, err := client.SetCommandWithObject(context.Background(), ts.URL, TestNodeName, TestCommandName, "", settings)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestDiscovery(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusAccepted)
	ts := newTestServer(http.MethodPost, common.ApiDiscoveryRoute, expectedResponse)
	defer ts.Close()

	client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
	res, err := client.Discovery(context.Background(), ts.URL)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestProfileScan(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusAccepted)
	ts := newTestServer(http.MethodPost, common.ApiProfileScanRoute, expectedResponse)
	defer ts.Close()

	client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
	res, err := client.ProfileScan(context.Background(), ts.URL, requests.ProfileScanRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestStopNodeDiscovery(t *testing.T) {
	id := uuid.New().String()
	requestRoute := path.Join(common.ApiDiscoveryRoute, common.RequestId, id)

	tests := []struct {
		name      string
		requestId string
		route     string
	}{
		{"stop node discovery", "", common.ApiDiscoveryRoute},
		{"stop node discovery with request id", id, requestRoute},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			ts := newTestServer(http.MethodDelete, testCase.route, dtoCommon.BaseResponse{})
			//defer ts.Close()

			client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
			res, err := client.StopNodeDiscovery(context.Background(), ts.URL, testCase.requestId, nil)

			require.NoError(t, err)
			assert.IsType(t, dtoCommon.BaseResponse{}, res)
			ts.Close()
		})
	}
}

func TestStopProfileScan(t *testing.T) {
	route := path.Join(common.ApiProfileScanRoute, common.Node, common.Name, TestNodeName)
	ts := newTestServer(http.MethodDelete, route, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewNodeServiceCommandClient(NewNullAuthenticationInjector(), false)
	res, err := client.StopProfileScan(context.Background(), ts.URL, TestNodeName, nil)

	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}
