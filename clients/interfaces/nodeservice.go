package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// NodeServiceClient defines the interface for interactions with the Node Service endpoint on the EdgeX Foundry core-metadata service.
type NodeServiceClient interface {
	// Add adds new node services.
	Add(ctx context.Context, reqs []requests.AddNodeServiceRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
	// Update updates node services.
	Update(ctx context.Context, reqs []requests.UpdateNodeServiceRequest) ([]common.BaseResponse, errors.EdgeX)
	// AllNodeServices returns all node services. Node services can also be filtered by labels.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset. The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllNodeServices(ctx context.Context, labels []string, offset int, limit int) (responses.MultiNodeServicesResponse, errors.EdgeX)
	// NodeServiceByName returns a node service by name.
	NodeServiceByName(ctx context.Context, name string) (responses.NodeServiceResponse, errors.EdgeX)
	// DeleteByName deletes a node service by name.
	DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX)
}
