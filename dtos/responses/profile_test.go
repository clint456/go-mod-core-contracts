//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewProfileResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedProfile := dtos.Profile{ProfileBasicInfo: dtos.ProfileBasicInfo{
		Name: "test node profile"}}
	actual := NewProfileResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedProfile)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedProfile, actual.Profile)
}

func TestNewMultiProfilesResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedProfiles := []dtos.Profile{
		{ProfileBasicInfo: dtos.ProfileBasicInfo{Name: "test node profile1"}},
		{ProfileBasicInfo: dtos.ProfileBasicInfo{Name: "test node profile2"}},
	}
	expectedTotalCount := int64(2)
	actual := NewMultiProfilesResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedProfiles)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedProfiles, actual.Profiles)
}

func TestNewMultiProfileBasicInfosResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedProfileBasicInfos := []dtos.ProfileBasicInfo{
		{Name: "test node profile1"},
		{Name: "test node profile2"},
	}
	expectedTotalCount := int64(2)
	actual := NewMultiProfileBasicInfosResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedProfileBasicInfos)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedProfileBasicInfos, actual.Profiles)
}
