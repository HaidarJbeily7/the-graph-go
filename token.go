// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/the-graph-go/internal/apijson"
	"github.com/stainless-sdks/the-graph-go/internal/apiquery"
	"github.com/stainless-sdks/the-graph-go/internal/requestconfig"
	"github.com/stainless-sdks/the-graph-go/option"
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// TokenService contains methods and other services that help with interacting with
// the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r TokenService) {
	r = TokenService{}
	r.Options = opts
	return
}

// Returns ERC-20 token metadata including supply, holder count, and price data.
func (r *TokenService) GetMetadata(ctx context.Context, contract string, query TokenGetMetadataParams, opts ...option.RequestOption) (res *TokenGetMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if contract == "" {
		err = errors.New("missing required contract parameter")
		return
	}
	path := fmt.Sprintf("tokens/evm/%s", contract)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TokenGetMetadataResponse struct {
	Data         []TokenGetMetadataResponseData     `json:"data,required"`
	DurationMs   float64                            `json:"duration_ms,required"`
	Pagination   TokenGetMetadataResponsePagination `json:"pagination,required"`
	RequestTime  string                             `json:"request_time,required"`
	Results      float64                            `json:"results,required"`
	Statistics   TokenGetMetadataResponseStatistics `json:"statistics,required"`
	TotalResults float64                            `json:"total_results,required"`
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
func (r TokenGetMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenGetMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGetMetadataResponseData struct {
	BlockNum          float64 `json:"block_num,required"`
	CirculatingSupply float64 `json:"circulating_supply,required"`
	// Filter by address
	Contract string                           `json:"contract,required"`
	Datetime time.Time                        `json:"datetime,required" format:"date-time"`
	Holders  float64                          `json:"holders,required"`
	Icon     TokenGetMetadataResponseDataIcon `json:"icon,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID    string  `json:"network_id,required"`
	Timestamp    float64 `json:"timestamp,required"`
	TotalSupply  float64 `json:"total_supply,required"`
	Decimals     float64 `json:"decimals"`
	LowLiquidity bool    `json:"low_liquidity"`
	MarketCap    float64 `json:"market_cap"`
	Name         string  `json:"name"`
	PriceUsd     float64 `json:"price_usd"`
	Symbol       string  `json:"symbol"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockNum          respjson.Field
		CirculatingSupply respjson.Field
		Contract          respjson.Field
		Datetime          respjson.Field
		Holders           respjson.Field
		Icon              respjson.Field
		NetworkID         respjson.Field
		Timestamp         respjson.Field
		TotalSupply       respjson.Field
		Decimals          respjson.Field
		LowLiquidity      respjson.Field
		MarketCap         respjson.Field
		Name              respjson.Field
		PriceUsd          respjson.Field
		Symbol            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGetMetadataResponseData) RawJSON() string { return r.JSON.raw }
func (r *TokenGetMetadataResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGetMetadataResponseDataIcon struct {
	Web3icon string `json:"web3icon,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Web3icon    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGetMetadataResponseDataIcon) RawJSON() string { return r.JSON.raw }
func (r *TokenGetMetadataResponseDataIcon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGetMetadataResponsePagination struct {
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
func (r TokenGetMetadataResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *TokenGetMetadataResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGetMetadataResponseStatistics struct {
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
func (r TokenGetMetadataResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *TokenGetMetadataResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGetMetadataParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID TokenGetMetadataParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [TokenGetMetadataParams]'s query parameters as `url.Values`.
func (r TokenGetMetadataParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type TokenGetMetadataParamsNetworkID string

const (
	TokenGetMetadataParamsNetworkIDArbitrumOne TokenGetMetadataParamsNetworkID = "arbitrum-one"
	TokenGetMetadataParamsNetworkIDAvalanche   TokenGetMetadataParamsNetworkID = "avalanche"
	TokenGetMetadataParamsNetworkIDBase        TokenGetMetadataParamsNetworkID = "base"
	TokenGetMetadataParamsNetworkIDBsc         TokenGetMetadataParamsNetworkID = "bsc"
	TokenGetMetadataParamsNetworkIDMainnet     TokenGetMetadataParamsNetworkID = "mainnet"
	TokenGetMetadataParamsNetworkIDMatic       TokenGetMetadataParamsNetworkID = "matic"
	TokenGetMetadataParamsNetworkIDOptimism    TokenGetMetadataParamsNetworkID = "optimism"
	TokenGetMetadataParamsNetworkIDUnichain    TokenGetMetadataParamsNetworkID = "unichain"
)
