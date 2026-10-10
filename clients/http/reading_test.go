//
// Copyright (C) 2020-2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
)

func TestQueryAllReadings(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllReadingRoute, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.AllReadings(context.Background(), 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.AllReadingsWithQueryParams(context.Background(), 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingCount(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiReadingCountRoute, dtoCommon.CountResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingCount(context.Background())
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.CountResponse{}, res)
}

func TestQueryReadingCountByNodeName(t *testing.T) {
	nodeName := "node"
	path := path.Join(common.ApiReadingCountRoute, common.Node, common.Name, nodeName)
	ts := newTestServer(http.MethodGet, path, dtoCommon.CountResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingCountByNodeName(context.Background(), nodeName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.CountResponse{}, res)
}

func TestQueryReadingsByNodeName(t *testing.T) {
	nodeName := "node"
	urlPath := path.Join(common.ApiReadingRoute, common.Node, common.Name, nodeName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ReadingsByNodeName(context.Background(), nodeName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByNodeNameWithQueryParams(context.Background(), nodeName, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByResourceName(t *testing.T) {
	resourceName := "resource"
	urlPath := path.Join(common.ApiReadingRoute, common.ResourceName, resourceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ReadingsByResourceName(context.Background(), resourceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByResourceNameWithQueryParams(context.Background(), resourceName, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByTimeRange(t *testing.T) {
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ReadingsByTimeRange(context.Background(), start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByTimeRangeWithQueryParams(context.Background(), start, end, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByResourceNameAndTimeRange(t *testing.T) {
	resourceName := "resource"
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.ResourceName, resourceName, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ReadingsByResourceNameAndTimeRange(context.Background(), resourceName, start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByResourceNameAndTimeRangeWithQueryParams(context.Background(), resourceName, start, end, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByNodeNameAndResourceName(t *testing.T) {
	nodeName := "node"
	resourceName := "resource"
	urlPath := path.Join(common.ApiReadingRoute, common.Node, common.Name, nodeName, common.ResourceName, resourceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ReadingsByNodeNameAndResourceName(context.Background(), nodeName, resourceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByNodeNameAndResourceNameWithQueryParams(context.Background(), nodeName, resourceName, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByNodeNameAndResourceNameAndTimeRange(t *testing.T) {
	nodeName := "node"
	resourceName := "resource"
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.Node, common.Name, nodeName, common.ResourceName, resourceName, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ReadingsByNodeNameAndResourceNameAndTimeRange(context.Background(), nodeName, resourceName, start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByNodeNameAndResourceNameAndTimeRangeWithQueryParams(context.Background(), nodeName, resourceName, start, end, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByNodeNameAndResourceNamesAndTimeRange(t *testing.T) {
	nodeName := "node"
	resourceNames := []string{"resource01", "resource02"}
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.Node, common.Name, nodeName, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByNodeNameAndResourceNamesAndTimeRange(context.Background(), nodeName, resourceNames, start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)

	res, err = client.ReadingsByNodeNameAndResourceNamesAndTimeRangeWithQueryParams(context.Background(), nodeName, resourceNames, start, end, 1, 10, queryParameters)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}
