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

// BalanceService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r BalanceService) {
	r = BalanceService{}
	r.Options = opts
	return
}

// Returns SPL token balances for Solana token accounts with mint and program data.
func (r *BalanceService) ListSvm(ctx context.Context, query BalanceListSvmParams, opts ...option.RequestOption) (res *BalanceListSvmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "balances/svm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns ERC-20 and native token balances for a wallet address with USD values.
func (r *BalanceService) GetEvm(ctx context.Context, address string, query BalanceGetEvmParams, opts ...option.RequestOption) (res *BalanceGetEvmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("balances/evm/%s", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BalanceListSvmResponse struct {
	Data         []BalanceListSvmResponseData     `json:"data,required"`
	DurationMs   float64                          `json:"duration_ms,required"`
	Pagination   BalanceListSvmResponsePagination `json:"pagination,required"`
	RequestTime  string                           `json:"request_time,required"`
	Results      float64                          `json:"results,required"`
	Statistics   BalanceListSvmResponseStatistics `json:"statistics,required"`
	TotalResults float64                          `json:"total_results,required"`
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
func (r BalanceListSvmResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceListSvmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListSvmResponseData struct {
	Amount   string  `json:"amount,required"`
	BlockNum float64 `json:"block_num,required"`
	Datetime string  `json:"datetime,required"`
	Decimals float64 `json:"decimals,required"`
	// Filter by address
	Mint string `json:"mint,required"`
	// The Graph Network ID for SVM networks https://thegraph.com/networks
	//
	// Any of "solana".
	NetworkID string `json:"network_id,required"`
	// Filter by address
	ProgramID string  `json:"program_id,required"`
	Timestamp float64 `json:"timestamp,required"`
	// Filter by address
	TokenAccount string  `json:"token_account,required"`
	Value        float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		BlockNum     respjson.Field
		Datetime     respjson.Field
		Decimals     respjson.Field
		Mint         respjson.Field
		NetworkID    respjson.Field
		ProgramID    respjson.Field
		Timestamp    respjson.Field
		TokenAccount respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BalanceListSvmResponseData) RawJSON() string { return r.JSON.raw }
func (r *BalanceListSvmResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListSvmResponsePagination struct {
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
func (r BalanceListSvmResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *BalanceListSvmResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListSvmResponseStatistics struct {
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
func (r BalanceListSvmResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *BalanceListSvmResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceGetEvmResponse struct {
	Data         []BalanceGetEvmResponseData     `json:"data,required"`
	DurationMs   float64                         `json:"duration_ms,required"`
	Pagination   BalanceGetEvmResponsePagination `json:"pagination,required"`
	RequestTime  string                          `json:"request_time,required"`
	Results      float64                         `json:"results,required"`
	Statistics   BalanceGetEvmResponseStatistics `json:"statistics,required"`
	TotalResults float64                         `json:"total_results,required"`
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
func (r BalanceGetEvmResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetEvmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceGetEvmResponseData struct {
	Amount   string  `json:"amount,required"`
	BlockNum float64 `json:"block_num,required"`
	// Filter by address
	Contract          string `json:"contract,required"`
	LastBalanceUpdate string `json:"last_balance_update,required"`
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
		Amount            respjson.Field
		BlockNum          respjson.Field
		Contract          respjson.Field
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
func (r BalanceGetEvmResponseData) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetEvmResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceGetEvmResponsePagination struct {
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
func (r BalanceGetEvmResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetEvmResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceGetEvmResponseStatistics struct {
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
func (r BalanceGetEvmResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetEvmResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListSvmParams struct {
	// The Graph Network ID for SVM networks https://thegraph.com/networks
	//
	// Any of "solana".
	NetworkID BalanceListSvmParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by mint address
	Mint param.Opt[string] `query:"mint,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by token account address
	TokenAccount param.Opt[string] `query:"token_account,omitzero" json:"-"`
	// Filter by program ID
	//
	// Any of "", "TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb",
	// "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA".
	ProgramID BalanceListSvmParamsProgramID `query:"program_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BalanceListSvmParams]'s query parameters as `url.Values`.
func (r BalanceListSvmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for SVM networks https://thegraph.com/networks
type BalanceListSvmParamsNetworkID string

const (
	BalanceListSvmParamsNetworkIDSolana BalanceListSvmParamsNetworkID = "solana"
)

// Filter by program ID
type BalanceListSvmParamsProgramID string

const (
	BalanceListSvmParamsProgramIDEmpty                                       BalanceListSvmParamsProgramID = ""
	BalanceListSvmParamsProgramIDTokenzQdBNbLqP5VEhdkAs6Epflc1PHnBqCxEpPxuEb BalanceListSvmParamsProgramID = "TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"
	BalanceListSvmParamsProgramIDTokenkegQfeZyiNwAJbNbGkpfxcWuBvf9Ss623Vq5Da BalanceListSvmParamsProgramID = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
)

type BalanceGetEvmParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID BalanceGetEvmParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by address
	Contract param.Opt[string] `query:"contract,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BalanceGetEvmParams]'s query parameters as `url.Values`.
func (r BalanceGetEvmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type BalanceGetEvmParamsNetworkID string

const (
	BalanceGetEvmParamsNetworkIDArbitrumOne BalanceGetEvmParamsNetworkID = "arbitrum-one"
	BalanceGetEvmParamsNetworkIDAvalanche   BalanceGetEvmParamsNetworkID = "avalanche"
	BalanceGetEvmParamsNetworkIDBase        BalanceGetEvmParamsNetworkID = "base"
	BalanceGetEvmParamsNetworkIDBsc         BalanceGetEvmParamsNetworkID = "bsc"
	BalanceGetEvmParamsNetworkIDMainnet     BalanceGetEvmParamsNetworkID = "mainnet"
	BalanceGetEvmParamsNetworkIDMatic       BalanceGetEvmParamsNetworkID = "matic"
	BalanceGetEvmParamsNetworkIDOptimism    BalanceGetEvmParamsNetworkID = "optimism"
	BalanceGetEvmParamsNetworkIDUnichain    BalanceGetEvmParamsNetworkID = "unichain"
)
