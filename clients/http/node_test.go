//
// Copyright (C) 2020-2021 Unknown author
// Copyright (C) 2023 Intel Corporation
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddNodes(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiNodeRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddNodeRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestAddNodesWithQueryParams(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiNodeRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AddWithQueryParams(context.Background(), []requests.AddNodeRequest{}, map[string]string{"foo": "bar"})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestPatchNodes(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiNodeRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateNodeRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestPatchNodesWithQueryParams(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiNodeRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateWithQueryParams(context.Background(), []requests.UpdateNodeRequest{}, map[string]string{"foo": "bar"})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestUpdateNodeProperties(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(common.ApiNodeRoute, common.Name, testName, common.Properties)
	ts := newTestServer(http.MethodPatch, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateNodeProperties(context.Background(), testName, requests.NodePropertiesRequest{})
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryAllNodes(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllNodeRoute, responses.MultiNodesResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllNodes(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodesResponse{}, res)
}

func TestNodeNameExists(t *testing.T) {
	nodeName := "node"
	path := path.Join(common.ApiNodeRoute, common.Check, common.Name, nodeName)
	ts := newTestServer(http.MethodGet, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.NodeNameExists(context.Background(), nodeName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryNodeByName(t *testing.T) {
	nodeName := "node"
	path := path.Join(common.ApiNodeRoute, common.Name, nodeName)
	ts := newTestServer(http.MethodGet, path, responses.NodeResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.NodeByName(context.Background(), nodeName)
	require.NoError(t, err)
	require.IsType(t, responses.NodeResponse{}, res)
}

func TestDeleteNodeByName(t *testing.T) {
	nodeName := "node"
	path := path.Join(common.ApiNodeRoute, common.Name, nodeName)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteNodeByName(context.Background(), nodeName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryNodesByProfileName(t *testing.T) {
	profileName := "profile"
	urlPath := path.Join(common.ApiNodeRoute, common.Profile, common.Name, profileName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNodesResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.NodesByProfileName(context.Background(), profileName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodesResponse{}, res)
}

func TestQueryNodesByServiceName(t *testing.T) {
	serviceName := "service"
	urlPath := path.Join(common.ApiNodeRoute, common.Service, common.Name, serviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNodesResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.NodesByServiceName(context.Background(), serviceName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodesResponse{}, res)
}

func TestQueryNodeTree(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllNodeRoute, responses.MultiNodesResponse{})
	defer ts.Close()
	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllNodesWithChildren(context.Background(), "MyRoot", 3, []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodesResponse{}, res)
}

// newQueryCapturingServer records the query string of the request it serves, so tests can assert
// what actually reached the wire rather than only that the call succeeded.
func newQueryCapturingServer(t *testing.T, apiRoute string, captured *url.Values) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.EscapedPath() != apiRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		q := r.URL.Query()
		*captured = q
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(responses.MultiNodesResponse{})
		_, _ = w.Write(b)
	}))
}

func TestQueryAllNodesWithQueryParams(t *testing.T) {
	var captured url.Values
	ts := newQueryCapturingServer(t, common.ApiAllNodeRoute, &captured)
	defer ts.Close()

	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllNodesWithQueryParams(context.Background(), []string{"label1", "label2"}, 1, 10,
		map[string]string{common.BasicInfoOnly: common.ValueTrue})
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodesResponse{}, res)

	assert.Equal(t, common.ValueTrue, captured.Get(common.BasicInfoOnly), "extra query param did not reach the request")
	assert.Equal(t, "1", captured.Get(common.Offset))
	assert.Equal(t, "10", captured.Get(common.Limit))
	assert.Equal(t, "label1,label2", captured.Get(common.Labels))
}

// AllNodes must keep sending the offset/limit/labels params it always has, so that delegating to
// the WithQueryParams variant stays behaviour-preserving.
func TestQueryAllNodes_NoQueryParams(t *testing.T) {
	var captured url.Values
	ts := newQueryCapturingServer(t, common.ApiAllNodeRoute, &captured)
	defer ts.Close()

	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.AllNodes(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)

	assert.Equal(t, "1", captured.Get(common.Offset))
	assert.Equal(t, "10", captured.Get(common.Limit))
	assert.Equal(t, "label1,label2", captured.Get(common.Labels))
	assert.NotContains(t, captured, common.BasicInfoOnly, "no extra param must be sent")
}

func TestQueryNodeTreeWithQueryParams(t *testing.T) {
	var captured url.Values
	ts := newQueryCapturingServer(t, common.ApiAllNodeRoute, &captured)
	defer ts.Close()

	client := NewNodeClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllNodesWithChildrenWithQueryParams(context.Background(), "parent", 3, []string{"label1"}, 0, 20,
		map[string]string{common.BasicInfoOnly: common.ValueTrue})
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodesResponse{}, res)

	assert.Equal(t, common.ValueTrue, captured.Get(common.BasicInfoOnly), "extra query param did not reach the request")
	assert.Equal(t, "parent", captured.Get(common.DescendantsOf))
	assert.Equal(t, "3", captured.Get(common.MaxLevels))
	assert.Equal(t, "label1", captured.Get(common.Labels))
}
