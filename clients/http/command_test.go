//
// Copyright (C) 2021-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/require"
)

func TestQueryNodeCoreCommands(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllNodeRoute, responses.MultiNodeCoreCommandsResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllNodeCoreCommands(context.Background(), 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNodeCoreCommandsResponse{}, res)
}

func TestQueryNodeCoreCommandsByNodeName(t *testing.T) {
	nodeName := "Simple-Node01"
	path := common.NewPathBuilder().EnableNameFieldEscape(false).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).BuildPath()
	ts := newTestServer(http.MethodGet, path, responses.NodeCoreCommandResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.NodeCoreCommandsByNodeName(context.Background(), nodeName)
	require.NoError(t, err)
	require.IsType(t, responses.NodeCoreCommandResponse{}, res)
}

func TestIssueGetCommandByName(t *testing.T) {
	nodeName := "Simple-Node01"
	cmdName := "SwitchButton"
	path := common.NewPathBuilder().EnableNameFieldEscape(false).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetNameFieldPath(cmdName).BuildPath()
	ts := newTestServer(http.MethodGet, path, &responses.EventResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewNullAuthenticationInjector(), false)
	pushEvent, err := strconv.ParseBool(common.ValueTrue)
	require.NoError(t, err)
	notReturnEvent, err := strconv.ParseBool(common.ValueFalse)
	require.NoError(t, err)
	res, err := client.IssueGetCommandByName(context.Background(), nodeName, cmdName, pushEvent, notReturnEvent)
	require.NoError(t, err)
	require.IsType(t, &responses.EventResponse{}, res)
}

func TestIssueGetCommandByNameWithQueryParams(t *testing.T) {
	nodeName := "Simple-Node01"
	cmdName := "SwitchButton"
	testQueryParams := map[string]string{"foo": "bar", "key": "value"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range testQueryParams {
			if r.URL.Query().Get(k) != v {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(responses.EventResponse{})
		_, _ = w.Write(b)
	}))
	defer ts.Close()

	client := NewCommandClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.IssueGetCommandByNameWithQueryParams(context.Background(), nodeName, cmdName, testQueryParams)
	require.NoError(t, err)
	require.IsType(t, &responses.EventResponse{}, res)
}

func TestIssueIssueSetCommandByName(t *testing.T) {
	nodeName := "Simple-Node01"
	cmdName := "SwitchButton"
	settings := map[string]any{
		"SwitchButton": "true",
	}
	path := common.NewPathBuilder().EnableNameFieldEscape(false).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetNameFieldPath(cmdName).BuildPath()
	ts := newTestServer(http.MethodPut, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.IssueSetCommandByName(context.Background(), nodeName, cmdName, settings)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestIssueIssueSetCommandByNameWithObject(t *testing.T) {
	nodeName := "Simple-Node01"
	cmdName := "SwitchButton"
	settings := map[string]any{
		"SwitchButton": map[string]any{
			"kind":  "button",
			"value": "on",
		},
	}
	path := common.NewPathBuilder().EnableNameFieldEscape(false).
		SetPath(common.ApiNodeRoute).SetPath(common.Name).SetNameFieldPath(nodeName).SetNameFieldPath(cmdName).BuildPath()
	ts := newTestServer(http.MethodPut, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.IssueSetCommandByName(context.Background(), nodeName, cmdName, settings)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}
