// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/HaidarJbeily7/the-graph-go/internal/apijson"
	"github.com/HaidarJbeily7/the-graph-go/internal/apiquery"
	"github.com/HaidarJbeily7/the-graph-go/internal/requestconfig"
	"github.com/HaidarJbeily7/the-graph-go/option"
	"github.com/HaidarJbeily7/the-graph-go/packages/respjson"
)

// NFTHolderService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTHolderService] method instead.
type NFTHolderService struct {
	Options []option.RequestOption
}

// NewNFTHolderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNFTHolderService(opts ...option.RequestOption) (r NFTHolderService) {
	r = NFTHolderService{}
	r.Options = opts
	return
}

// Returns wallet addresses holding NFT collection tokens with quantity and
// percentage distribution.
func (r *NFTHolderService) GetByContract(ctx context.Context, contract string, query NFTHolderGetByContractParams, opts ...option.RequestOption) (res *NFTHolderGetByContractResponse, err error) {
	opts = append(r.Options[:], opts...)
	if contract == "" {
		err = errors.New("missing required contract parameter")
		return
	}
	path := fmt.Sprintf("nft/holders/evm/%s", contract)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NFTHolderGetByContractResponse struct {
	Data         []NFTHolderGetByContractResponseData     `json:"data,required"`
	DurationMs   float64                                  `json:"duration_ms,required"`
	Pagination   NFTHolderGetByContractResponsePagination `json:"pagination,required"`
	RequestTime  string                                   `json:"request_time,required"`
	Results      float64                                  `json:"results,required"`
	Statistics   NFTHolderGetByContractResponseStatistics `json:"statistics,required"`
	TotalResults float64                                  `json:"total_results,required"`
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
func (r NFTHolderGetByContractResponse) RawJSON() string { return r.JSON.raw }
func (r *NFTHolderGetByContractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTHolderGetByContractResponseData struct {
	// Filter by address
	Address string `json:"address,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID string `json:"network_id,required"`
	// Percentage of total supply held by this address
	Percentage float64 `json:"percentage,required"`
	// Number of tokens held by this address
	Quantity      float64 `json:"quantity,required"`
	TokenStandard string  `json:"token_standard,required"`
	// Number of unique token IDs held by this address
	UniqueTokens float64 `json:"unique_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address       respjson.Field
		NetworkID     respjson.Field
		Percentage    respjson.Field
		Quantity      respjson.Field
		TokenStandard respjson.Field
		UniqueTokens  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTHolderGetByContractResponseData) RawJSON() string { return r.JSON.raw }
func (r *NFTHolderGetByContractResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTHolderGetByContractResponsePagination struct {
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
func (r NFTHolderGetByContractResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *NFTHolderGetByContractResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTHolderGetByContractResponseStatistics struct {
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
func (r NFTHolderGetByContractResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *NFTHolderGetByContractResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTHolderGetByContractParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID NFTHolderGetByContractParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [NFTHolderGetByContractParams]'s query parameters as
// `url.Values`.
func (r NFTHolderGetByContractParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type NFTHolderGetByContractParamsNetworkID string

const (
	NFTHolderGetByContractParamsNetworkIDArbitrumOne NFTHolderGetByContractParamsNetworkID = "arbitrum-one"
	NFTHolderGetByContractParamsNetworkIDAvalanche   NFTHolderGetByContractParamsNetworkID = "avalanche"
	NFTHolderGetByContractParamsNetworkIDBase        NFTHolderGetByContractParamsNetworkID = "base"
	NFTHolderGetByContractParamsNetworkIDBsc         NFTHolderGetByContractParamsNetworkID = "bsc"
	NFTHolderGetByContractParamsNetworkIDMainnet     NFTHolderGetByContractParamsNetworkID = "mainnet"
	NFTHolderGetByContractParamsNetworkIDMatic       NFTHolderGetByContractParamsNetworkID = "matic"
	NFTHolderGetByContractParamsNetworkIDOptimism    NFTHolderGetByContractParamsNetworkID = "optimism"
	NFTHolderGetByContractParamsNetworkIDUnichain    NFTHolderGetByContractParamsNetworkID = "unichain"
)
