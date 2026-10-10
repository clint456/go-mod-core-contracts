//
// Copyright (C) 2021-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// NodeServiceCommandClient defines the interface for interactions with the node command endpoints on the EdgeX Foundry node services.
type NodeServiceCommandClient interface {
	// GetCommand invokes node service's command API for issuing get(read) command
	GetCommand(ctx context.Context, baseUrl string, nodeName string, commandName string, queryParams string) (*responses.EventResponse, errors.EdgeX)
	// SetCommand invokes node service's command API for issuing set(write) command
	SetCommand(ctx context.Context, baseUrl string, nodeName string, commandName string, queryParams string, settings map[string]string) (common.BaseResponse, errors.EdgeX)
	// SetCommandWithObject invokes node service's set command API and the settings supports object value type
	SetCommandWithObject(ctx context.Context, baseUrl string, nodeName string, commandName string, queryParams string, settings map[string]interface{}) (common.BaseResponse, errors.EdgeX)
	// Discovery invokes node service's discovery API
	Discovery(ctx context.Context, baseUrl string) (common.BaseResponse, errors.EdgeX)
	// ProfileScan sends an HTTP POST request to the node service's profile scan API endpoint.
	ProfileScan(ctx context.Context, baseUrl string, req requests.ProfileScanRequest) (common.BaseResponse, errors.EdgeX)
	// StopNodeDiscovery invokes node service's stop node discovery API
	StopNodeDiscovery(ctx context.Context, baseUrl string, requestId string, queryParams map[string]string) (common.BaseResponse, errors.EdgeX)
	// StopProfileScan invokes node service's stop profile scan API
	StopProfileScan(ctx context.Context, baseUrl string, nodeName string, queryParams map[string]string) (common.BaseResponse, errors.EdgeX)
}
