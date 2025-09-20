// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/HaidarJbeily7/the-graph-go/internal/apijson"
	"github.com/HaidarJbeily7/the-graph-go/internal/apiquery"
	"github.com/HaidarJbeily7/the-graph-go/internal/requestconfig"
	"github.com/HaidarJbeily7/the-graph-go/option"
	"github.com/HaidarJbeily7/the-graph-go/packages/param"
	"github.com/HaidarJbeily7/the-graph-go/packages/respjson"
)

// NFTOwnershipService contains methods and other services that help with
// interacting with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTOwnershipService] method instead.
type NFTOwnershipService struct {
	Options []option.RequestOption
}

// NewNFTOwnershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNFTOwnershipService(opts ...option.RequestOption) (r NFTOwnershipService) {
	r = NFTOwnershipService{}
	r.Options = opts
	return
}

// Returns NFT tokens owned by a wallet address with metadata and ownership
// information.
func (r *NFTOwnershipService) GetByAddress(ctx context.Context, address string, query NFTOwnershipGetByAddressParams, opts ...option.RequestOption) (res *NFTOwnershipGetByAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("nft/ownerships/evm/%s", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NFTOwnershipGetByAddressResponse struct {
	Data         []NFTOwnershipGetByAddressResponseData     `json:"data,required"`
	DurationMs   float64                                    `json:"duration_ms,required"`
	Pagination   NFTOwnershipGetByAddressResponsePagination `json:"pagination,required"`
	RequestTime  string                                     `json:"request_time,required"`
	Results      float64                                    `json:"results,required"`
	Statistics   NFTOwnershipGetByAddressResponseStatistics `json:"statistics,required"`
	TotalResults float64                                    `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data         respjson.Field
		DurationMs   respjson.Field
		Pagination   respjson.Field
		RequestTime  respjson.Field
		Results      respjson.Field
		Statistics   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTOwnershipGetByAddressResponse) RawJSON() string { return r.JSON.raw }
func (r *NFTOwnershipGetByAddressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTOwnershipGetByAddressResponseData struct {
	Contract string `json:"contract,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID string `json:"network_id,required"`
	Owner     string `json:"owner,required"`
	TokenID   string `json:"token_id,required"`
	// Any of "", "ERC721", "ERC1155".
	TokenStandard string `json:"token_standard,required"`
	Description   string `json:"description"`
	Image         string `json:"image"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Uri           string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contract      respjson.Field
		NetworkID     respjson.Field
		Owner         respjson.Field
		TokenID       respjson.Field
		TokenStandard respjson.Field
		Description   respjson.Field
		Image         respjson.Field
		Name          respjson.Field
		Symbol        respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTOwnershipGetByAddressResponseData) RawJSON() string { return r.JSON.raw }
func (r *NFTOwnershipGetByAddressResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTOwnershipGetByAddressResponsePagination struct {
	CurrentPage  int64 `json:"current_page,required"`
	NextPage     int64 `json:"next_page,required"`
	PreviousPage int64 `json:"previous_page,required"`
	TotalPages   int64 `json:"total_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentPage  respjson.Field
		NextPage     respjson.Field
		PreviousPage respjson.Field
		TotalPages   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTOwnershipGetByAddressResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *NFTOwnershipGetByAddressResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTOwnershipGetByAddressResponseStatistics struct {
	BytesRead float64 `json:"bytes_read"`
	Elapsed   float64 `json:"elapsed"`
	RowsRead  float64 `json:"rows_read"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BytesRead   respjson.Field
		Elapsed     respjson.Field
		RowsRead    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTOwnershipGetByAddressResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *NFTOwnershipGetByAddressResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTOwnershipGetByAddressParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID NFTOwnershipGetByAddressParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by address
	Contract param.Opt[string] `query:"contract,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Any of "", "ERC721", "ERC1155".
	TokenStandard NFTOwnershipGetByAddressParamsTokenStandard `query:"token_standard,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NFTOwnershipGetByAddressParams]'s query parameters as
// `url.Values`.
func (r NFTOwnershipGetByAddressParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type NFTOwnershipGetByAddressParamsNetworkID string

const (
	NFTOwnershipGetByAddressParamsNetworkIDArbitrumOne NFTOwnershipGetByAddressParamsNetworkID = "arbitrum-one"
	NFTOwnershipGetByAddressParamsNetworkIDAvalanche   NFTOwnershipGetByAddressParamsNetworkID = "avalanche"
	NFTOwnershipGetByAddressParamsNetworkIDBase        NFTOwnershipGetByAddressParamsNetworkID = "base"
	NFTOwnershipGetByAddressParamsNetworkIDBsc         NFTOwnershipGetByAddressParamsNetworkID = "bsc"
	NFTOwnershipGetByAddressParamsNetworkIDMainnet     NFTOwnershipGetByAddressParamsNetworkID = "mainnet"
	NFTOwnershipGetByAddressParamsNetworkIDMatic       NFTOwnershipGetByAddressParamsNetworkID = "matic"
	NFTOwnershipGetByAddressParamsNetworkIDOptimism    NFTOwnershipGetByAddressParamsNetworkID = "optimism"
	NFTOwnershipGetByAddressParamsNetworkIDUnichain    NFTOwnershipGetByAddressParamsNetworkID = "unichain"
)

type NFTOwnershipGetByAddressParamsTokenStandard string

const (
	NFTOwnershipGetByAddressParamsTokenStandardEmpty   NFTOwnershipGetByAddressParamsTokenStandard = ""
	NFTOwnershipGetByAddressParamsTokenStandardErc721  NFTOwnershipGetByAddressParamsTokenStandard = "ERC721"
	NFTOwnershipGetByAddressParamsTokenStandardErc1155 NFTOwnershipGetByAddressParamsTokenStandard = "ERC1155"
)
