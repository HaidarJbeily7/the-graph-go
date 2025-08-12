// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"github.com/stainless-sdks/the-graph-go/option"
)

// NFTItemEvmService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTItemEvmService] method instead.
type NFTItemEvmService struct {
	Options  []option.RequestOption
	Contract NFTItemEvmContractService
}

// NewNFTItemEvmService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNFTItemEvmService(opts ...option.RequestOption) (r NFTItemEvmService) {
	r = NFTItemEvmService{}
	r.Options = opts
	r.Contract = NewNFTItemEvmContractService(opts...)
	return
}
