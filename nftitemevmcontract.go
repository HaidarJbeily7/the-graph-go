// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/the-graph-go/internal/apijson"
	"github.com/stainless-sdks/the-graph-go/internal/apiquery"
	"github.com/stainless-sdks/the-graph-go/internal/requestconfig"
	"github.com/stainless-sdks/the-graph-go/option"
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// NFTItemEvmContractService contains methods and other services that help with
// interacting with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTItemEvmContractService] method instead.
type NFTItemEvmContractService struct {
	Options []option.RequestOption
}

// NewNFTItemEvmContractService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNFTItemEvmContractService(opts ...option.RequestOption) (r NFTItemEvmContractService) {
	r = NFTItemEvmContractService{}
	r.Options = opts
	return
}

// Returns NFT token metadata, attributes, current owner, and media URIs.
func (r *NFTItemEvmContractService) GetByToken(ctx context.Context, tokenID string, params NFTItemEvmContractGetByTokenParams, opts ...option.RequestOption) (res *NFTItemEvmContractGetByTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.Contract == "" {
		err = errors.New("missing required contract parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("nft/items/evm/contract/%s/token_id/%s", params.Contract, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type NFTItemEvmContractGetByTokenResponse struct {
	Data         []NFTItemEvmContractGetByTokenResponseData     `json:"data,required"`
	DurationMs   float64                                        `json:"duration_ms,required"`
	Pagination   NFTItemEvmContractGetByTokenResponsePagination `json:"pagination,required"`
	RequestTime  string                                         `json:"request_time,required"`
	Results      float64                                        `json:"results,required"`
	Statistics   NFTItemEvmContractGetByTokenResponseStatistics `json:"statistics,required"`
	TotalResults float64                                        `json:"total_results,required"`
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
func (r NFTItemEvmContractGetByTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *NFTItemEvmContractGetByTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTItemEvmContractGetByTokenResponseData struct {
	Contract string `json:"contract,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID string `json:"network_id,required"`
	Owner     string `json:"owner,required"`
	TokenID   string `json:"token_id,required"`
	// Any of "", "ERC721", "ERC1155".
	TokenStandard string                                              `json:"token_standard,required"`
	Attributes    []NFTItemEvmContractGetByTokenResponseDataAttribute `json:"attributes"`
	Description   string                                              `json:"description"`
	Image         string                                              `json:"image"`
	Name          string                                              `json:"name"`
	Uri           string                                              `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contract      respjson.Field
		NetworkID     respjson.Field
		Owner         respjson.Field
		TokenID       respjson.Field
		TokenStandard respjson.Field
		Attributes    respjson.Field
		Description   respjson.Field
		Image         respjson.Field
		Name          respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTItemEvmContractGetByTokenResponseData) RawJSON() string { return r.JSON.raw }
func (r *NFTItemEvmContractGetByTokenResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTItemEvmContractGetByTokenResponseDataAttribute struct {
	TraitType   string `json:"trait_type,required"`
	Value       string `json:"value,required"`
	DisplayType string `json:"display_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TraitType   respjson.Field
		Value       respjson.Field
		DisplayType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTItemEvmContractGetByTokenResponseDataAttribute) RawJSON() string { return r.JSON.raw }
func (r *NFTItemEvmContractGetByTokenResponseDataAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTItemEvmContractGetByTokenResponsePagination struct {
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
func (r NFTItemEvmContractGetByTokenResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *NFTItemEvmContractGetByTokenResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTItemEvmContractGetByTokenResponseStatistics struct {
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
func (r NFTItemEvmContractGetByTokenResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *NFTItemEvmContractGetByTokenResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTItemEvmContractGetByTokenParams struct {
	// Filter by NFT contract address
	Contract string `path:"contract,required" json:"-"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID NFTItemEvmContractGetByTokenParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [NFTItemEvmContractGetByTokenParams]'s query parameters as
// `url.Values`.
func (r NFTItemEvmContractGetByTokenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type NFTItemEvmContractGetByTokenParamsNetworkID string

const (
	NFTItemEvmContractGetByTokenParamsNetworkIDArbitrumOne NFTItemEvmContractGetByTokenParamsNetworkID = "arbitrum-one"
	NFTItemEvmContractGetByTokenParamsNetworkIDAvalanche   NFTItemEvmContractGetByTokenParamsNetworkID = "avalanche"
	NFTItemEvmContractGetByTokenParamsNetworkIDBase        NFTItemEvmContractGetByTokenParamsNetworkID = "base"
	NFTItemEvmContractGetByTokenParamsNetworkIDBsc         NFTItemEvmContractGetByTokenParamsNetworkID = "bsc"
	NFTItemEvmContractGetByTokenParamsNetworkIDMainnet     NFTItemEvmContractGetByTokenParamsNetworkID = "mainnet"
	NFTItemEvmContractGetByTokenParamsNetworkIDMatic       NFTItemEvmContractGetByTokenParamsNetworkID = "matic"
	NFTItemEvmContractGetByTokenParamsNetworkIDOptimism    NFTItemEvmContractGetByTokenParamsNetworkID = "optimism"
	NFTItemEvmContractGetByTokenParamsNetworkIDUnichain    NFTItemEvmContractGetByTokenParamsNetworkID = "unichain"
)
