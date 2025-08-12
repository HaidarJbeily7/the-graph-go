// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/HaidarJbeily7/the-graph-go/internal/apijson"
	"github.com/HaidarJbeily7/the-graph-go/internal/apiquery"
	"github.com/HaidarJbeily7/the-graph-go/internal/requestconfig"
	"github.com/HaidarJbeily7/the-graph-go/option"
	"github.com/HaidarJbeily7/the-graph-go/packages/param"
	"github.com/HaidarJbeily7/the-graph-go/packages/respjson"
)

// PoolService contains methods and other services that help with interacting with
// the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPoolService] method instead.
type PoolService struct {
	Options []option.RequestOption
}

// NewPoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPoolService(opts ...option.RequestOption) (r PoolService) {
	r = PoolService{}
	r.Options = opts
	return
}

// Returns Uniswap liquidity pool metadata including token pairs, fees, and
// protocol versions.
func (r *PoolService) Get(ctx context.Context, query PoolGetParams, opts ...option.RequestOption) (res *PoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pools/evm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PoolGetResponse struct {
	Data         []PoolGetResponseData     `json:"data,required"`
	DurationMs   float64                   `json:"duration_ms,required"`
	Pagination   PoolGetResponsePagination `json:"pagination,required"`
	RequestTime  string                    `json:"request_time,required"`
	Results      float64                   `json:"results,required"`
	Statistics   PoolGetResponseStatistics `json:"statistics,required"`
	TotalResults float64                   `json:"total_results,required"`
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
func (r PoolGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PoolGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolGetResponseData struct {
	BlockNum float64   `json:"block_num,required"`
	Datetime time.Time `json:"datetime,required" format:"date-time"`
	// Filter by address
	Factory string  `json:"factory,required"`
	Fee     float64 `json:"fee,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID string `json:"network_id,required"`
	// Filter by pool
	Pool          string                    `json:"pool,required"`
	Protocol      string                    `json:"protocol,required"`
	Token0        PoolGetResponseDataToken0 `json:"token0,required"`
	Token1        PoolGetResponseDataToken1 `json:"token1,required"`
	TransactionID string                    `json:"transaction_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockNum      respjson.Field
		Datetime      respjson.Field
		Factory       respjson.Field
		Fee           respjson.Field
		NetworkID     respjson.Field
		Pool          respjson.Field
		Protocol      respjson.Field
		Token0        respjson.Field
		Token1        respjson.Field
		TransactionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoolGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PoolGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolGetResponseDataToken0 struct {
	// Filter by address
	Address  string  `json:"address,required"`
	Decimals float64 `json:"decimals,required"`
	Symbol   string  `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Decimals    respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoolGetResponseDataToken0) RawJSON() string { return r.JSON.raw }
func (r *PoolGetResponseDataToken0) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolGetResponseDataToken1 struct {
	// Filter by address
	Address  string  `json:"address,required"`
	Decimals float64 `json:"decimals,required"`
	Symbol   string  `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Decimals    respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoolGetResponseDataToken1) RawJSON() string { return r.JSON.raw }
func (r *PoolGetResponseDataToken1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolGetResponsePagination struct {
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
func (r PoolGetResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *PoolGetResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolGetResponseStatistics struct {
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
func (r PoolGetResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *PoolGetResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolGetParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID PoolGetParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by contract address
	Token param.Opt[string] `query:"token,omitzero" json:"-"`
	// Filter by address
	Factory param.Opt[string] `query:"factory,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by pool address
	Pool param.Opt[string] `query:"pool,omitzero" json:"-"`
	// Protocol name
	//
	// Any of "", "uniswap_v2", "uniswap_v3", "uniswap_v4".
	Protocol PoolGetParamsProtocol `query:"protocol,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PoolGetParams]'s query parameters as `url.Values`.
func (r PoolGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type PoolGetParamsNetworkID string

const (
	PoolGetParamsNetworkIDArbitrumOne PoolGetParamsNetworkID = "arbitrum-one"
	PoolGetParamsNetworkIDAvalanche   PoolGetParamsNetworkID = "avalanche"
	PoolGetParamsNetworkIDBase        PoolGetParamsNetworkID = "base"
	PoolGetParamsNetworkIDBsc         PoolGetParamsNetworkID = "bsc"
	PoolGetParamsNetworkIDMainnet     PoolGetParamsNetworkID = "mainnet"
	PoolGetParamsNetworkIDMatic       PoolGetParamsNetworkID = "matic"
	PoolGetParamsNetworkIDOptimism    PoolGetParamsNetworkID = "optimism"
	PoolGetParamsNetworkIDUnichain    PoolGetParamsNetworkID = "unichain"
)

// Protocol name
type PoolGetParamsProtocol string

const (
	PoolGetParamsProtocolEmpty     PoolGetParamsProtocol = ""
	PoolGetParamsProtocolUniswapV2 PoolGetParamsProtocol = "uniswap_v2"
	PoolGetParamsProtocolUniswapV3 PoolGetParamsProtocol = "uniswap_v3"
	PoolGetParamsProtocolUniswapV4 PoolGetParamsProtocol = "uniswap_v4"
)
