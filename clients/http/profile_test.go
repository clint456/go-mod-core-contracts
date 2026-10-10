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
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	edgexErrors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProfileClientWithUrlCallback(t *testing.T) {
	baseUrlFunc := func() (string, error) {
		return "", nil
	}
	client := NewProfileClientWithUrlCallback(baseUrlFunc, &emptyAuthenticationInjector{}, true)
	require.NotNil(t, client)
	require.NotNil(t, client.(*ProfileClient).resourcesCache)
}

func TestAddProfiles(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != common.ApiProfileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusMultiStatus)
		br := dtoCommon.NewBaseWithIdResponse(requestId, "", http.StatusMultiStatus, uuid.New().String())
		res, _ := json.Marshal([]dtoCommon.BaseWithIdResponse{br})
		_, _ = w.Write(res)
	}))
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.ProfileRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, requestId, res[0].RequestId)
}

func TestPutProfiles(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != common.ApiProfileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusMultiStatus)
		br := dtoCommon.NewBaseResponse(requestId, "", http.StatusMultiStatus)
		res, _ := json.Marshal([]dtoCommon.BaseResponse{br})
		_, _ = w.Write(res)
	}))
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.ProfileRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, requestId, res[0].RequestId)
}

func TestAddProfileByYaml(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != common.ApiProfileUploadFileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		br := dtoCommon.NewBaseWithIdResponse(requestId, "", http.StatusCreated, uuid.New().String())
		res, _ := json.Marshal(br)
		_, _ = w.Write(res)
	}))
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, b, _, _ := runtime.Caller(0)

	tests := []struct {
		name          string
		filePath      string
		errorExpected bool
	}{
		{name: "Add node profile by yaml file", filePath: filepath.Dir(b) + "/data/sample-profile.yaml", errorExpected: false},
		{name: "file not found error", filePath: filepath.Dir(b) + "/data/file-not-found.yaml", errorExpected: true},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := client.AddByYaml(context.Background(), testCase.filePath)
			if testCase.errorExpected {
				require.True(t, errors.Is(err, os.ErrNotExist))
				assert.Equal(t, edgexErrors.KindIOError, edgexErrors.Kind(err))
			} else {
				require.NoError(t, err)
				assert.Equal(t, requestId, res.RequestId)
				assert.Equal(t, http.StatusCreated, res.StatusCode)
			}
		})
	}
}

func TestUpdateProfileByYaml(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != common.ApiProfileUploadFileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		br := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
		res, _ := json.Marshal(br)
		_, _ = w.Write(res)
	}))
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, b, _, _ := runtime.Caller(0)

	tests := []struct {
		name          string
		filePath      string
		errorExpected bool
	}{
		{name: "Update node profile by yaml file", filePath: filepath.Dir(b) + "/data/sample-profile.yaml", errorExpected: false},
		{name: "file not found error", filePath: filepath.Dir(b) + "/data/file-not-found.yaml", errorExpected: true},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := client.UpdateByYaml(context.Background(), testCase.filePath)
			if testCase.errorExpected {
				require.True(t, errors.Is(err, os.ErrNotExist))
				assert.Equal(t, edgexErrors.KindIOError, edgexErrors.Kind(err))
			} else {
				require.NoError(t, err)
				assert.Equal(t, requestId, res.RequestId)
				assert.Equal(t, http.StatusOK, res.StatusCode)
			}
		})
	}
}

func TestDeleteProfileByName(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(common.ApiProfileRoute, common.Name, testName)
	ts := newTestServer(http.MethodDelete, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteByName(context.Background(), testName)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestQueryProfileByName(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(common.ApiProfileRoute, common.Name, testName)
	ts := newTestServer(http.MethodGet, urlPath, responses.ProfileResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.ProfileByName(context.Background(), testName)
	require.NoError(t, err)
}

func TestQueryProfileByNameWithQueryParams(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(common.ApiProfileRoute, common.Name, testName)

	// newTestServer ignores the query string, so capture it here to prove the params are actually sent
	var rawQuery string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawQuery = r.URL.RawQuery
		if r.Method != http.MethodGet || r.URL.EscapedPath() != urlPath {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, _ := json.Marshal(responses.ProfileResponse{})
		_, _ = w.Write(b)
	}))
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)

	tests := []struct {
		name             string
		queryParams      map[string]string
		expectedRawQuery string
	}{
		{"existence only", map[string]string{common.ExistenceOnly: common.ValueTrue}, "existenceOnly=true"},
		{"basic info only", map[string]string{common.BasicInfoOnly: common.ValueTrue}, "basicInfoOnly=true"},
		{"no params", nil, ""},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			rawQuery = ""

			res, err := client.ProfileByNameWithQueryParams(context.Background(), testName, testCase.queryParams)

			require.NoError(t, err)
			assert.IsType(t, responses.ProfileResponse{}, res)
			assert.Equal(t, testCase.expectedRawQuery, rawQuery)
		})
	}

	// the legacy method must stay byte-identical on the wire
	rawQuery = ""
	_, err := client.ProfileByName(context.Background(), testName)
	require.NoError(t, err)
	assert.Empty(t, rawQuery)
}

func TestQueryAllProfiles(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllProfileRoute, responses.MultiProfilesResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.AllProfiles(context.Background(), []string{"testLabel1", "testLabel2"}, 1, 10)
	require.NoError(t, err)
}

func TestQueryAllProfileBasicInfos(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllProfileBasicInfoRoute, responses.MultiProfilesResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.AllProfileBasicInfos(context.Background(), []string{"testLabel1", "testLabel2"}, 1, 10)
	require.NoError(t, err)
}

func TestQueryProfilesByModel(t *testing.T) {
	testModel := "testModel"
	urlPath := path.Join(common.ApiProfileRoute, common.Model, testModel)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProfilesResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.ProfilesByModel(context.Background(), testModel, 1, 10)
	require.NoError(t, err)
}

func TestQueryProfilesByManufacturer(t *testing.T) {
	testManufacturer := "testManufacturer"
	urlPath := path.Join(common.ApiProfileRoute, common.Manufacturer, testManufacturer)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProfilesResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.ProfilesByManufacturer(context.Background(), testManufacturer, 1, 10)
	require.NoError(t, err)
}

func TestQueryProfilesByManufacturerAndModel(t *testing.T) {
	testManufacturer := "testManufacturer"
	testModel := "testModel"
	urlPath := path.Join(common.ApiProfileRoute, common.Manufacturer, testManufacturer, common.Model, testModel)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProfilesResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.ProfilesByManufacturerAndModel(context.Background(), testManufacturer, testModel, 1, 10)
	require.NoError(t, err)
}

func TestResourceByProfileNameAndResourceName(t *testing.T) {
	profileName := "testProfile"
	resourceName := "testResource"
	urlPath := path.Join(common.ApiResourceRoute, common.Profile, profileName, common.Resource, resourceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.ResourceResponse{})
	defer ts.Close()
	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.ResourceByProfileNameAndResourceName(context.Background(), profileName, resourceName)

	require.NoError(t, err)
	assert.IsType(t, responses.ResourceResponse{}, res)
}

func TestUpdateProfileBasicInfo(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiProfileBasicInfoRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateProfileBasicInfo(context.Background(), []requests.ProfileBasicInfoRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestUpdateProfileTags(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(common.ApiProfileRoute, common.Name, testName, common.Tags)
	ts := newTestServer(http.MethodPatch, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateProfileTags(context.Background(), testName, requests.ProfileTagsRequest{})
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestAddProfileResource(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiProfileResourceRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AddProfileResource(context.Background(), []requests.AddResourceRequest{})

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestUpdateProfileResource(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiProfileResourceRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateProfileResource(context.Background(), []requests.UpdateResourceRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestDeleteResourceByName(t *testing.T) {
	profileName := "testProfile"
	resourceName := "testResource"
	urlPath := path.Join(common.ApiProfileRoute, common.Name, profileName, common.Resource, resourceName)
	ts := newTestServer(http.MethodDelete, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteResourceByName(context.Background(), profileName, resourceName)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestAddProfileCommand(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiProfileCommandRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AddProfileCommand(context.Background(), []requests.AddCommandRequest{})

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestUpdateProfileCommand(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiProfileCommandRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateProfileCommand(context.Background(), []requests.UpdateCommandRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestDeleteCommandByName(t *testing.T) {
	profileName := "testProfile"
	commandName := "testCommand"
	urlPath := path.Join(common.ApiProfileRoute, common.Name, profileName, common.Command, commandName)
	ts := newTestServer(http.MethodDelete, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProfileClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteCommandByName(context.Background(), profileName, commandName)
	require.NoError(t, err)
	require.NotNil(t, res)
}
