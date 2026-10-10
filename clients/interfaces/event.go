//
// Copyright (C) 2021-2025 IOTech Ltd
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

// EventClient defines the interface for interactions with the Event endpoint on the EdgeX Foundry core-data service.
type EventClient interface {
	// Add adds new event.
	Add(ctx context.Context, serviceName string, req requests.AddEventRequest) (common.BaseWithIdResponse, errors.EdgeX)
	// AllEvents returns all events sorted in descending order of created time.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllEvents(ctx context.Context, offset, limit int) (responses.MultiEventsResponse, errors.EdgeX)
	// AllEventsWithQueryParams returns all events with specified query parameters. Events are sorted in descending order of created time.
	AllEventsWithQueryParams(ctx context.Context, offset, limit int, queryParams map[string]string) (responses.MultiEventsResponse, errors.EdgeX)
	// EventCount returns a count of all of events currently stored in the database.
	EventCount(ctx context.Context) (common.CountResponse, errors.EdgeX)
	// EventCountByNodeName returns a count of all of events currently stored in the database, sourced from the specified node.
	EventCountByNodeName(ctx context.Context, name string) (common.CountResponse, errors.EdgeX)
	// EventsByNodeName returns a portion of the entire events according to the node name, offset and limit parameters. Events are sorted in descending order of created time.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	EventsByNodeName(ctx context.Context, name string, offset, limit int) (responses.MultiEventsResponse, errors.EdgeX)
	// EventsByNodeNameWithQueryParams returns a portion of the entire events according to the node name and specified query parameters. Events are sorted in descending order of created time.
	EventsByNodeNameWithQueryParams(ctx context.Context, name string, offset, limit int, queryParams map[string]string) (responses.MultiEventsResponse, errors.EdgeX)
	// DeleteByNodeName deletes all events for the specified node.
	DeleteByNodeName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
	// EventsByTimeRange returns events between a given start and end date/time. Events are sorted in descending order of created time.
	// start, end: Unix timestamp, indicating the date/time range.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	EventsByTimeRange(ctx context.Context, start, end int64, offset, limit int) (responses.MultiEventsResponse, errors.EdgeX)
	// EventsByTimeRangeWithQueryParams returns events between a given start, end date/time, and specified query parameters. Events are sorted in descending order of created time.
	EventsByTimeRangeWithQueryParams(ctx context.Context, start, end int64, offset, limit int, queryParams map[string]string) (responses.MultiEventsResponse, errors.EdgeX)
	// DeleteByAge deletes events that are older than the given age. Age is supposed in milliseconds from created timestamp.
	DeleteByAge(ctx context.Context, age int) (common.BaseResponse, errors.EdgeX)
	// DeleteById deletes an event by its id
	DeleteById(ctx context.Context, id string) (common.BaseResponse, errors.EdgeX)
}
