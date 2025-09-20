// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/HaidarJbeily7/the-graph-go/internal/apijson"
	"github.com/HaidarJbeily7/the-graph-go/internal/apiquery"
	"github.com/HaidarJbeily7/the-graph-go/internal/requestconfig"
	"github.com/HaidarJbeily7/the-graph-go/option"
	"github.com/HaidarJbeily7/the-graph-go/packages/param"
	"github.com/HaidarJbeily7/the-graph-go/packages/respjson"
)

// HistoricalBalanceService contains methods and other services that help with
// interacting with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHistoricalBalanceService] method instead.
type HistoricalBalanceService struct {
	Options []option.RequestOption
}

// NewHistoricalBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHistoricalBalanceService(opts ...option.RequestOption) (r HistoricalBalanceService) {
	r = HistoricalBalanceService{}
	r.Options = opts
	return
}

// Returns wallet token balance changes over time in OHLC format.
func (r *HistoricalBalanceService) Get(ctx context.Context, address string, query HistoricalBalanceGetParams, opts ...option.RequestOption) (res *HistoricalBalanceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("historical/balances/evm/%s", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HistoricalBalanceGetResponse struct {
	Data         []HistoricalBalanceGetResponseData     `json:"data,required"`
	DurationMs   float64                                `json:"duration_ms,required"`
	Pagination   HistoricalBalanceGetResponsePagination `json:"pagination,required"`
	RequestTime  string                                 `json:"request_time,required"`
	Results      float64                                `json:"results,required"`
	Statistics   HistoricalBalanceGetResponseStatistics `json:"statistics,required"`
	TotalResults float64                                `json:"total_results,required"`
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
func (r HistoricalBalanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *HistoricalBalanceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoricalBalanceGetResponseData struct {
	Close    float64   `json:"close,required"`
	Contract string    `json:"contract,required"`
	Datetime time.Time `json:"datetime,required" format:"date-time"`
	Decimals float64   `json:"decimals,required"`
	High     float64   `json:"high,required"`
	Low      float64   `json:"low,required"`
	Name     string    `json:"name,required"`
	Open     float64   `json:"open,required"`
	Symbol   string    `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		Contract    respjson.Field
		Datetime    respjson.Field
		Decimals    respjson.Field
		High        respjson.Field
		Low         respjson.Field
		Name        respjson.Field
		Open        respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HistoricalBalanceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *HistoricalBalanceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoricalBalanceGetResponsePagination struct {
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
func (r HistoricalBalanceGetResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *HistoricalBalanceGetResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoricalBalanceGetResponseStatistics struct {
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
func (r HistoricalBalanceGetResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *HistoricalBalanceGetResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoricalBalanceGetParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID HistoricalBalanceGetParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	Contracts []string         `query:"contracts,omitzero" json:"-"`
	// The interval for which to aggregate price data (hourly, 4-hours, daily or
	// weekly).
	Interval any `query:"interval,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HistoricalBalanceGetParams]'s query parameters as
// `url.Values`.
func (r HistoricalBalanceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type HistoricalBalanceGetParamsNetworkID string

const (
	HistoricalBalanceGetParamsNetworkIDArbitrumOne HistoricalBalanceGetParamsNetworkID = "arbitrum-one"
	HistoricalBalanceGetParamsNetworkIDAvalanche   HistoricalBalanceGetParamsNetworkID = "avalanche"
	HistoricalBalanceGetParamsNetworkIDBase        HistoricalBalanceGetParamsNetworkID = "base"
	HistoricalBalanceGetParamsNetworkIDBsc         HistoricalBalanceGetParamsNetworkID = "bsc"
	HistoricalBalanceGetParamsNetworkIDMainnet     HistoricalBalanceGetParamsNetworkID = "mainnet"
	HistoricalBalanceGetParamsNetworkIDMatic       HistoricalBalanceGetParamsNetworkID = "matic"
	HistoricalBalanceGetParamsNetworkIDOptimism    HistoricalBalanceGetParamsNetworkID = "optimism"
	HistoricalBalanceGetParamsNetworkIDUnichain    HistoricalBalanceGetParamsNetworkID = "unichain"
)
