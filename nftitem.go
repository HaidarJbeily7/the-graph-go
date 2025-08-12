// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"github.com/HaidarJbeily7/the-graph-go/option"
)

// NFTItemService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTItemService] method instead.
type NFTItemService struct {
	Options []option.RequestOption
	Evm     NFTItemEvmService
}

// NewNFTItemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNFTItemService(opts ...option.RequestOption) (r NFTItemService) {
	r = NFTItemService{}
	r.Options = opts
	r.Evm = NewNFTItemEvmService(opts...)
	return
}
