// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"github.com/HaidarJbeily7/the-graph-go/option"
)

// NFTService contains methods and other services that help with interacting with
// the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTService] method instead.
type NFTService struct {
	Options     []option.RequestOption
	Ownerships  NFTOwnershipService
	Collections NFTCollectionService
	Items       NFTItemService
	Activities  NFTActivityService
	Holders     NFTHolderService
	Sales       NFTSaleService
}

// NewNFTService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNFTService(opts ...option.RequestOption) (r NFTService) {
	r = NFTService{}
	r.Options = opts
	r.Ownerships = NewNFTOwnershipService(opts...)
	r.Collections = NewNFTCollectionService(opts...)
	r.Items = NewNFTItemService(opts...)
	r.Activities = NewNFTActivityService(opts...)
	r.Holders = NewNFTHolderService(opts...)
	r.Sales = NewNFTSaleService(opts...)
	return
}
