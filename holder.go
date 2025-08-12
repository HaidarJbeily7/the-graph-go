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
	"github.com/stainless-sdks/the-graph-go/packages/param"
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// HolderService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHolderService] method instead.
type HolderService struct {
	Options []option.RequestOption
}

// NewHolderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHolderService(opts ...option.RequestOption) (r HolderService) {
	r = HolderService{}
	r.Options = opts
	return
}

// Returns token holders ranked by balance with holdings and supply percentage.
func (r *HolderService) GetByContract(ctx context.Context, contract string, query HolderGetByContractParams, opts ...option.RequestOption) (res *HolderGetByContractResponse, err error) {
	opts = append(r.Options[:], opts...)
	if contract == "" {
		err = errors.New("missing required contract parameter")
		return
	}
	path := fmt.Sprintf("holders/evm/%s", contract)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HolderGetByContractResponse struct {
	Data         []HolderGetByContractResponseData     `json:"data,required"`
	DurationMs   float64                               `json:"duration_ms,required"`
	Pagination   HolderGetByContractResponsePagination `json:"pagination,required"`
	RequestTime  string                                `json:"request_time,required"`
	Results      float64                               `json:"results,required"`
	Statistics   HolderGetByContractResponseStatistics `json:"statistics,required"`
	TotalResults float64                               `json:"total_results,required"`
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
func (r HolderGetByContractResponse) RawJSON() string { return r.JSON.raw }
func (r *HolderGetByContractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HolderGetByContractResponseData struct {
	// Filter by address
	Address           string  `json:"address,required"`
	Amount            string  `json:"amount,required"`
	BlockNum          float64 `json:"block_num,required"`
	LastBalanceUpdate string  `json:"last_balance_update,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID    string  `json:"network_id,required"`
	Value        float64 `json:"value,required"`
	Decimals     float64 `json:"decimals"`
	LowLiquidity bool    `json:"low_liquidity"`
	Name         string  `json:"name"`
	PriceUsd     float64 `json:"price_usd"`
	Symbol       string  `json:"symbol"`
	ValueUsd     float64 `json:"value_usd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address           respjson.Field
		Amount            respjson.Field
		BlockNum          respjson.Field
		LastBalanceUpdate respjson.Field
		NetworkID         respjson.Field
		Value             respjson.Field
		Decimals          respjson.Field
		LowLiquidity      respjson.Field
		Name              respjson.Field
		PriceUsd          respjson.Field
		Symbol            respjson.Field
		ValueUsd          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HolderGetByContractResponseData) RawJSON() string { return r.JSON.raw }
func (r *HolderGetByContractResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HolderGetByContractResponsePagination struct {
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
func (r HolderGetByContractResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *HolderGetByContractResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HolderGetByContractResponseStatistics struct {
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
func (r HolderGetByContractResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *HolderGetByContractResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HolderGetByContractParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID HolderGetByContractParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "value".
	OrderBy HolderGetByContractParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection HolderGetByContractParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HolderGetByContractParams]'s query parameters as
// `url.Values`.
func (r HolderGetByContractParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type HolderGetByContractParamsNetworkID string

const (
	HolderGetByContractParamsNetworkIDArbitrumOne HolderGetByContractParamsNetworkID = "arbitrum-one"
	HolderGetByContractParamsNetworkIDAvalanche   HolderGetByContractParamsNetworkID = "avalanche"
	HolderGetByContractParamsNetworkIDBase        HolderGetByContractParamsNetworkID = "base"
	HolderGetByContractParamsNetworkIDBsc         HolderGetByContractParamsNetworkID = "bsc"
	HolderGetByContractParamsNetworkIDMainnet     HolderGetByContractParamsNetworkID = "mainnet"
	HolderGetByContractParamsNetworkIDMatic       HolderGetByContractParamsNetworkID = "matic"
	HolderGetByContractParamsNetworkIDOptimism    HolderGetByContractParamsNetworkID = "optimism"
	HolderGetByContractParamsNetworkIDUnichain    HolderGetByContractParamsNetworkID = "unichain"
)

// The field by which to order the results.
type HolderGetByContractParamsOrderBy string

const (
	HolderGetByContractParamsOrderByValue HolderGetByContractParamsOrderBy = "value"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type HolderGetByContractParamsOrderDirection string

const (
	HolderGetByContractParamsOrderDirectionAsc  HolderGetByContractParamsOrderDirection = "asc"
	HolderGetByContractParamsOrderDirectionDesc HolderGetByContractParamsOrderDirection = "desc"
)
