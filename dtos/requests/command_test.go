//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var emptyString = " "

var testCommand = AddCommandRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestProfileName,
	Command: dtos.Command{
		Name:      TestCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []dtos.ResourceOperation{{
			Resource: TestResourceName,
		}},
	},
}

var testUpdateCommand = UpdateCommandRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestProfileName,
	Command:     mockUpdateCommandDTO(),
}

func mockUpdateCommandDTO() dtos.UpdateCommand {
	testIsHidden := true
	testName := TestCommandName
	dc := dtos.UpdateCommand{}
	dc.Name = &testName
	dc.IsHidden = &testIsHidden
	return dc
}

func TestAddCommandRequest_Validate(t *testing.T) {
	valid := testCommand
	noProfileName := testCommand
	noProfileName.ProfileName = emptyString
	noCommandName := testCommand
	noCommandName.Command.Name = emptyString
	invalidReadWrite := testCommand
	invalidReadWrite.Command.ReadWrite = "invalid"
	noResourceOperations := testCommand
	noResourceOperations.Command.ResourceOperations = nil

	tests := []struct {
		name        string
		request     AddCommandRequest
		expectedErr bool
	}{
		{"valid AddCommandRequest", valid, false},
		{"invalid AddCommandRequest, no ProfileName", noProfileName, true},
		{"invalid AddCommandRequest, no Command Name", noCommandName, true},
		{"invalid AddCommandRequest, invalid ReadWrite", invalidReadWrite, true},
		{"invalid AddCommandRequest, no ResourceOperations", noResourceOperations, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddCommandRequest_UnmarshalJSON(t *testing.T) {
	valid := testCommand
	resultTestBytes, _ := json.Marshal(testCommand)
	type args struct {
		data []byte
	}

	tests := []struct {
		name        string
		request     AddCommandRequest
		args        args
		expectedErr bool
	}{
		{"unmarshal AddCommandRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddCommandRequest, empty data", AddCommandRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddCommandRequest, string data", AddCommandRequest{}, args{[]byte("Invalid AddCommandRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected AddCommandRequest.")
			}
		})
	}
}

func TestUpdateCommandRequest_Validate(t *testing.T) {
	valid := testUpdateCommand
	noProfileName := testUpdateCommand
	noProfileName.ProfileName = emptyString
	noCommandName := testUpdateCommand
	noCommandName.Command.Name = &emptyString

	tests := []struct {
		name        string
		request     UpdateCommandRequest
		expectedErr bool
	}{
		{"valid UpdateCommandRequest", valid, false},
		{"invalid UpdateCommandRequest, no profile name", noProfileName, true},
		{"invalid UpdateCommandRequest, no node command name", noCommandName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCommandRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
        "profileName": "TestProfile",
		"command":{"name":"TestCommand"}
	}`
	var req UpdateCommandRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Command.IsHidden)
}

func TestReplaceCommandModelFieldsWithDTO(t *testing.T) {
	command := models.Command{
		Name:      TestCommandName,
		ReadWrite: common.ReadWrite_R,
		IsHidden:  false,
		ResourceOperations: []models.ResourceOperation{{
			Resource: TestResourceName,
		}},
	}

	patch := mockUpdateCommandDTO()

	ReplaceCommandModelFieldsWithDTO(&command, patch)
	assert.Equal(t, common.ReadWrite_R, command.ReadWrite)
	assert.Equal(t, true, command.IsHidden)
	assert.Equal(t, TestResourceName, command.ResourceOperations[0].Resource)
}
