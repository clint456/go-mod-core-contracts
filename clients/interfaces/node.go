//
// Copyright (C) 2020-2024 IOTech Ltd
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

// NodeClient defines the interface for interactions with the Node endpoint on the EdgeX Foundry core-metadata service.
type NodeClient interface {
	// Add adds new nodes.
	Add(ctx context.Context, reqs []requests.AddNodeRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
	// AddWithQueryParams adds new nodes with query parameters.
	AddWithQueryParams(ctx context.Context, reqs []requests.AddNodeRequest, queryParams map[string]string) ([]common.BaseWithIdResponse, errors.EdgeX)
	// Update updates nodes.
	Update(ctx context.Context, reqs []requests.UpdateNodeRequest) ([]common.BaseResponse, errors.EdgeX)
	// UpdateWithQueryParams updates nodes with query parameters.
	UpdateWithQueryParams(ctx context.Context, reqs []requests.UpdateNodeRequest, queryParams map[string]string) ([]common.BaseResponse, errors.EdgeX)
	// UpdateNodeProperties merges the given properties into an existing node.
	// A property whose value is nil is deleted rather than set.
	UpdateNodeProperties(ctx context.Context, nodeName string, request requests.NodePropertiesRequest) (common.BaseResponse, errors.EdgeX)
	// AllNodes returns all nodes. Nodes can also be filtered by labels.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllNodes(ctx context.Context, labels []string, offset int, limit int) (responses.MultiNodesResponse, errors.EdgeX)
	// AllNodesWithQueryParams returns all nodes with the specified query parameters, which are
	// applied on top of the labels, offset and limit ones.
	AllNodesWithQueryParams(ctx context.Context, labels []string, offset int, limit int, queryParams map[string]string) (responses.MultiNodesResponse, errors.EdgeX)
	// AllNodesWithChildren returns all nodes who have parent, grandparent, etc. of the
	// given node name. Nodes can also be filtered by labels.
	// Node tree is descended at most maxLevels. If maxLevels is 0, there is no limit.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllNodesWithChildren(ctx context.Context, parent string, maxLevels uint, labels []string, offset int, limit int) (responses.MultiNodesResponse, errors.EdgeX)
	// AllNodesWithChildrenWithQueryParams returns the node tree with the specified query
	// parameters, which are applied on top of the ones AllNodesWithChildren already sets.
	AllNodesWithChildrenWithQueryParams(ctx context.Context, parent string, maxLevels uint, labels []string, offset int, limit int, queryParams map[string]string) (responses.MultiNodesResponse, errors.EdgeX)
	// NodeNameExists checks whether the node exists.
	NodeNameExists(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
	// NodeByName returns a node by node name.
	NodeByName(ctx context.Context, name string) (responses.NodeResponse, errors.EdgeX)
	// DeleteByName deletes a node by node name.
	DeleteNodeByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
	// NodesByProfileName returns nodes associated with the specified node profile.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	NodesByProfileName(ctx context.Context, name string, offset int, limit int) (responses.MultiNodesResponse, errors.EdgeX)
	// NodesByServiceName returns nodes associated with the specified node service.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	NodesByServiceName(ctx context.Context, name string, offset int, limit int) (responses.MultiNodesResponse, errors.EdgeX)
}
