//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// ReadingClient defines the interface for interactions with the Reading endpoint on the EdgeX Foundry core-data service.
type ReadingClient interface {
	// AllReadings returns all readings sorted in descending order of created time.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllReadings(ctx context.Context, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// AllReadingsWithQueryParams returns all readings with specified query parameters. Readings are sorted in descending order of created time.
	AllReadingsWithQueryParams(ctx context.Context, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingCount returns a count of all readings currently stored in the database.
	ReadingCount(ctx context.Context) (common.CountResponse, errors.EdgeX)
	// ReadingCountByNodeName returns a count of all readings currently stored in the database, sourced from the specified node.
	ReadingCountByNodeName(ctx context.Context, name string) (common.CountResponse, errors.EdgeX)
	// ReadingsByNodeName returns a portion of the entire readings according to the node name, offset and limit parameters. Readings are sorted in descending order of created time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByNodeName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameWithQueryParams returns a portion of the entire readings according to the node name, specified query parameters. Readings are sorted in descending order of created time.
	ReadingsByNodeNameWithQueryParams(ctx context.Context, name string, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByResourceName returns a portion of the entire readings according to the node resource name, offset and limit parameters. Readings are sorted in descending order of created time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByResourceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByResourceNameWithQueryParams returns a portion of the entire readings according to the node resource name, specified query parameters. Readings are sorted in descending order of created time.
	ReadingsByResourceNameWithQueryParams(ctx context.Context, name string, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByTimeRange returns readings between a given start and end date/time. Readings are sorted in descending order of created time.
	// start, end: Unix timestamp, indicating the date/time range.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByTimeRangeWithQueryParams returns readings between a given start, end date/time, and specified query parameters. Readings are sorted in descending order of created time.
	ReadingsByTimeRangeWithQueryParams(ctx context.Context, start, end int64, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByResourceNameAndTimeRange returns readings by resource name and specified time range. Readings are sorted in descending order of origin time.
	// start, end: Unix timestamp, indicating the date/time range
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByResourceNameAndTimeRange(ctx context.Context, name string, start, end int64, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByResourceNameAndTimeRangeWithQueryParams returns readings by resource name and specified time range, query parameters. Readings are sorted in descending order of origin time.
	ReadingsByResourceNameAndTimeRangeWithQueryParams(ctx context.Context, name string, start, end int64, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameAndResourceName returns readings by node name and resource name. Readings are sorted in descending order of origin time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByNodeNameAndResourceName(ctx context.Context, nodeName, resourceName string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameAndResourceNameWithQueryParams returns readings by node name, resource name, and specified query parameters. Readings are sorted in descending order of origin time.
	ReadingsByNodeNameAndResourceNameWithQueryParams(ctx context.Context, nodeName, resourceName string, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameAndResourceNameAndTimeRange returns readings by node name, resource name and specified time range. Readings are sorted in descending order of origin time.
	// start, end: Unix timestamp, indicating the date/time range
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByNodeNameAndResourceNameAndTimeRange(ctx context.Context, nodeName, resourceName string, start, end int64, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameAndResourceNameAndTimeRangeWithQueryParams returns readings by node name, resource name and specified time range, query parameters. Readings are sorted in descending order of origin time.
	ReadingsByNodeNameAndResourceNameAndTimeRangeWithQueryParams(ctx context.Context, nodeName, resourceName string, start, end int64, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameAndResourceNamesAndTimeRange returns readings by node name, multiple resource names and specified time range. Readings are sorted in descending order of origin time.
	// If none of resourceNames is specified, return all Readings under specified nodeName and within specified time range
	// start, end: Unix timestamp, indicating the date/time range
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	ReadingsByNodeNameAndResourceNamesAndTimeRange(ctx context.Context, nodeName string, resourceNames []string, start, end int64, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX)
	// ReadingsByNodeNameAndResourceNamesAndTimeRangeWithQueryParams returns readings by node name, multiple resource names and specified time range, query parameters. Readings are sorted in descending order of origin time.
	ReadingsByNodeNameAndResourceNamesAndTimeRangeWithQueryParams(ctx context.Context, nodeName string, resourceNames []string, start, end int64, offset, limit int, queryParams map[string]string) (responses.MultiReadingsResponse, errors.EdgeX)
}
