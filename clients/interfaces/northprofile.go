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

// NorthProfileClient defines the interface for interactions with the NorthProfile endpoint on the EdgeX Foundry core-metadata service.
type NorthProfileClient interface {
	// Add adds new north profiles
	Add(ctx context.Context, reqs []requests.NorthProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
	// Update updates north profiles
	Update(ctx context.Context, reqs []requests.NorthProfileRequest) ([]common.BaseResponse, errors.EdgeX)
	// DeleteByName deletes a north profile by name
	DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
	// NorthProfileByName queries a north profile by name
	NorthProfileByName(ctx context.Context, name string) (responses.NorthProfileResponse, errors.EdgeX)
	// AllNorthProfiles queries all north profiles with offset and limit
	AllNorthProfiles(ctx context.Context, offset int, limit int) (responses.MultiNorthProfilesResponse, errors.EdgeX)
}
