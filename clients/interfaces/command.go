//
// Copyright (C) 2021-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// CommandClient defines the interface for interactions with the command endpoints on the EdgeX Foundry core-command service.
type CommandClient interface {
	// AllNodeCoreCommands returns a paginated list of MultiNodeCoreCommandsResponse. The list contains all of the commands in the system associated with their respective node.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllNodeCoreCommands(ctx context.Context, offset int, limit int) (responses.MultiNodeCoreCommandsResponse, errors.EdgeX)
	// NodeCoreCommandsByNodeName returns all commands associated with the specified node name.
	NodeCoreCommandsByNodeName(ctx context.Context, nodeName string) (responses.NodeCoreCommandResponse, errors.EdgeX)
	// IssueGetCommandByName issues the specified read command referenced by the command name to the node/sensor that is also referenced by name.
	// dsPushEvent: If set to true, a successful GET will result in an event being pushed to the EdgeX system. Default value is false.
	// dsReturnEvent: If set to false, there will be no Event returned in the http response. Default value is true.
	IssueGetCommandByName(ctx context.Context, nodeName string, commandName string, dsPushEvent bool, dsReturnEvent bool) (*responses.EventResponse, errors.EdgeX)
	// IssueGetCommandByNameWithQueryParams issues the specified read command by nodeName and commandName with additional query parameters.
	IssueGetCommandByNameWithQueryParams(ctx context.Context, nodeName string, commandName string, queryParams map[string]string) (*responses.EventResponse, errors.EdgeX)
	// IssueSetCommandByName issues the specified write command referenced by the command name to the node/sensor that is also referenced by name.
	IssueSetCommandByName(ctx context.Context, nodeName string, commandName string, settings map[string]any) (common.BaseResponse, errors.EdgeX)
}
