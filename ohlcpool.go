// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"errors"
	"fmt"
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

// OhlcPoolService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOhlcPoolService] method instead.
type OhlcPoolService struct {
	Options []option.RequestOption
}

// NewOhlcPoolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOhlcPoolService(opts ...option.RequestOption) (r OhlcPoolService) {
	r = OhlcPoolService{}
	r.Options = opts
	return
}

// Returns candlestick price data for liquidity pools across time intervals.
func (r *OhlcPoolService) Get(ctx context.Context, pool string, query OhlcPoolGetParams, opts ...option.RequestOption) (res *OhlcPoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if pool == "" {
		err = errors.New("missing required pool parameter")
		return
	}
	path := fmt.Sprintf("ohlc/pools/evm/%s", pool)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OhlcPoolGetResponse struct {
	Data         []OhlcPoolGetResponseData     `json:"data,required"`
	DurationMs   float64                       `json:"duration_ms,required"`
	Pagination   OhlcPoolGetResponsePagination `json:"pagination,required"`
	RequestTime  string                        `json:"request_time,required"`
	Results      float64                       `json:"results,required"`
	Statistics   OhlcPoolGetResponseStatistics `json:"statistics,required"`
	TotalResults float64                       `json:"total_results,required"`
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
func (r OhlcPoolGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OhlcPoolGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OhlcPoolGetResponseData struct {
	Close        float64   `json:"close,required"`
	Datetime     time.Time `json:"datetime,required" format:"date-time"`
	High         float64   `json:"high,required"`
	Low          float64   `json:"low,required"`
	Open         float64   `json:"open,required"`
	Ticker       string    `json:"ticker,required"`
	Transactions float64   `json:"transactions,required"`
	Uaw          float64   `json:"uaw,required"`
	Volume       float64   `json:"volume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close        respjson.Field
		Datetime     respjson.Field
		High         respjson.Field
		Low          respjson.Field
		Open         respjson.Field
		Ticker       respjson.Field
		Transactions respjson.Field
		Uaw          respjson.Field
		Volume       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OhlcPoolGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *OhlcPoolGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OhlcPoolGetResponsePagination struct {
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
func (r OhlcPoolGetResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *OhlcPoolGetResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OhlcPoolGetResponseStatistics struct {
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
func (r OhlcPoolGetResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *OhlcPoolGetResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OhlcPoolGetParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID OhlcPoolGetParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// The interval for which to aggregate price data (hourly, 4-hours, daily or
	// weekly).
	Interval any `query:"interval,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OhlcPoolGetParams]'s query parameters as `url.Values`.
func (r OhlcPoolGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type OhlcPoolGetParamsNetworkID string

const (
	OhlcPoolGetParamsNetworkIDArbitrumOne OhlcPoolGetParamsNetworkID = "arbitrum-one"
	OhlcPoolGetParamsNetworkIDAvalanche   OhlcPoolGetParamsNetworkID = "avalanche"
	OhlcPoolGetParamsNetworkIDBase        OhlcPoolGetParamsNetworkID = "base"
	OhlcPoolGetParamsNetworkIDBsc         OhlcPoolGetParamsNetworkID = "bsc"
	OhlcPoolGetParamsNetworkIDMainnet     OhlcPoolGetParamsNetworkID = "mainnet"
	OhlcPoolGetParamsNetworkIDMatic       OhlcPoolGetParamsNetworkID = "matic"
	OhlcPoolGetParamsNetworkIDOptimism    OhlcPoolGetParamsNetworkID = "optimism"
	OhlcPoolGetParamsNetworkIDUnichain    OhlcPoolGetParamsNetworkID = "unichain"
)
