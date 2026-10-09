//
// @version: 0.0.1
// @contact.name: Clint
// @contact.email: clinton_luo@163.com
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

// NorthClient defines the interface for interactions with the North endpoint on the EdgeX Foundry core-metadata service.
type NorthClient interface {
	// Add adds new north channels
	Add(ctx context.Context, reqs []requests.NorthRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
	// Update updates north channels
	Update(ctx context.Context, reqs []requests.NorthRequest) ([]common.BaseResponse, errors.EdgeX)
	// DeleteByName deletes a north channel by name
	DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
	// NorthByName queries a north channel by name
	NorthByName(ctx context.Context, name string) (responses.NorthResponse, errors.EdgeX)
	// AllNorths queries all north channels with offset and limit
	AllNorths(ctx context.Context, offset int, limit int) (responses.MultiNorthsResponse, errors.EdgeX)
	// NorthsByServiceName queries north channels by service name with offset and limit
	NorthsByServiceName(ctx context.Context, name string, offset int, limit int) (responses.MultiNorthsResponse, errors.EdgeX)
}
