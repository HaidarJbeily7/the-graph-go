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

// NFTCollectionService contains methods and other services that help with
// interacting with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTCollectionService] method instead.
type NFTCollectionService struct {
	Options []option.RequestOption
}

// NewNFTCollectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNFTCollectionService(opts ...option.RequestOption) (r NFTCollectionService) {
	r = NFTCollectionService{}
	r.Options = opts
	return
}

// Returns NFT collection metadata, supply statistics, owner count, and transfer
// history.
func (r *NFTCollectionService) GetByContract(ctx context.Context, contract string, query NFTCollectionGetByContractParams, opts ...option.RequestOption) (res *NFTCollectionGetByContractResponse, err error) {
	opts = append(r.Options[:], opts...)
	if contract == "" {
		err = errors.New("missing required contract parameter")
		return
	}
	path := fmt.Sprintf("nft/collections/evm/%s", contract)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NFTCollectionGetByContractResponse struct {
	Data         []NFTCollectionGetByContractResponseData     `json:"data,required"`
	DurationMs   float64                                      `json:"duration_ms,required"`
	Pagination   NFTCollectionGetByContractResponsePagination `json:"pagination,required"`
	RequestTime  string                                       `json:"request_time,required"`
	Results      float64                                      `json:"results,required"`
	Statistics   NFTCollectionGetByContractResponseStatistics `json:"statistics,required"`
	TotalResults float64                                      `json:"total_results,required"`
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
func (r NFTCollectionGetByContractResponse) RawJSON() string { return r.JSON.raw }
func (r *NFTCollectionGetByContractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTCollectionGetByContractResponseData struct {
	// Filter by address
	Contract         string `json:"contract,required"`
	ContractCreation string `json:"contract_creation,required"`
	// Filter by address
	ContractCreator string `json:"contract_creator,required"`
	Name            string `json:"name,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID         string  `json:"network_id,required"`
	Owners            float64 `json:"owners,required"`
	Symbol            string  `json:"symbol,required"`
	TotalSupply       float64 `json:"total_supply,required"`
	TotalTransfers    float64 `json:"total_transfers,required"`
	TotalUniqueSupply float64 `json:"total_unique_supply,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contract          respjson.Field
		ContractCreation  respjson.Field
		ContractCreator   respjson.Field
		Name              respjson.Field
		NetworkID         respjson.Field
		Owners            respjson.Field
		Symbol            respjson.Field
		TotalSupply       respjson.Field
		TotalTransfers    respjson.Field
		TotalUniqueSupply respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTCollectionGetByContractResponseData) RawJSON() string { return r.JSON.raw }
func (r *NFTCollectionGetByContractResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTCollectionGetByContractResponsePagination struct {
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
func (r NFTCollectionGetByContractResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *NFTCollectionGetByContractResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTCollectionGetByContractResponseStatistics struct {
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
func (r NFTCollectionGetByContractResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *NFTCollectionGetByContractResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTCollectionGetByContractParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID NFTCollectionGetByContractParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [NFTCollectionGetByContractParams]'s query parameters as
// `url.Values`.
func (r NFTCollectionGetByContractParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type NFTCollectionGetByContractParamsNetworkID string

const (
	NFTCollectionGetByContractParamsNetworkIDArbitrumOne NFTCollectionGetByContractParamsNetworkID = "arbitrum-one"
	NFTCollectionGetByContractParamsNetworkIDAvalanche   NFTCollectionGetByContractParamsNetworkID = "avalanche"
	NFTCollectionGetByContractParamsNetworkIDBase        NFTCollectionGetByContractParamsNetworkID = "base"
	NFTCollectionGetByContractParamsNetworkIDBsc         NFTCollectionGetByContractParamsNetworkID = "bsc"
	NFTCollectionGetByContractParamsNetworkIDMainnet     NFTCollectionGetByContractParamsNetworkID = "mainnet"
	NFTCollectionGetByContractParamsNetworkIDMatic       NFTCollectionGetByContractParamsNetworkID = "matic"
	NFTCollectionGetByContractParamsNetworkIDOptimism    NFTCollectionGetByContractParamsNetworkID = "optimism"
	NFTCollectionGetByContractParamsNetworkIDUnichain    NFTCollectionGetByContractParamsNetworkID = "unichain"
)
