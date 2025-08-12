// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/the-graph-go/internal/apijson"
	"github.com/stainless-sdks/the-graph-go/internal/apiquery"
	"github.com/stainless-sdks/the-graph-go/internal/requestconfig"
	"github.com/stainless-sdks/the-graph-go/option"
	"github.com/stainless-sdks/the-graph-go/packages/param"
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// SwapService contains methods and other services that help with interacting with
// the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwapService] method instead.
type SwapService struct {
	Options []option.RequestOption
}

// NewSwapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSwapService(opts ...option.RequestOption) (r SwapService) {
	r = SwapService{}
	r.Options = opts
	return
}

// Returns DEX swap transactions from Uniswap protocols with token amounts and
// prices.
func (r *SwapService) GetEvm(ctx context.Context, query SwapGetEvmParams, opts ...option.RequestOption) (res *SwapGetEvmResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "swaps/evm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns AMM swap events from Solana DEXs with input/output tokens and amounts.
func (r *SwapService) GetSvm(ctx context.Context, query SwapGetSvmParams, opts ...option.RequestOption) (res *SwapGetSvmResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "swaps/svm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SwapGetEvmResponse struct {
	Data         []SwapGetEvmResponseData     `json:"data,required"`
	DurationMs   float64                      `json:"duration_ms,required"`
	Pagination   SwapGetEvmResponsePagination `json:"pagination,required"`
	RequestTime  string                       `json:"request_time,required"`
	Results      float64                      `json:"results,required"`
	Statistics   SwapGetEvmResponseStatistics `json:"statistics,required"`
	TotalResults float64                      `json:"total_results,required"`
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
func (r SwapGetEvmResponse) RawJSON() string { return r.JSON.raw }
func (r *SwapGetEvmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetEvmResponseData struct {
	Amount0  string  `json:"amount0,required"`
	Amount1  string  `json:"amount1,required"`
	BlockNum float64 `json:"block_num,required"`
	// Filter by address
	Caller   string    `json:"caller,required"`
	Datetime time.Time `json:"datetime,required" format:"date-time"`
	// Filter by address
	Factory string `json:"factory,required"`
	Fee     string `json:"fee,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID string `json:"network_id,required"`
	// Filter by pool
	Pool     string  `json:"pool,required"`
	Price0   float64 `json:"price0,required"`
	Price1   float64 `json:"price1,required"`
	Protocol string  `json:"protocol,required"`
	// Filter by address
	Recipient string `json:"recipient,required"`
	// Filter by address
	Sender        string                       `json:"sender,required"`
	Timestamp     float64                      `json:"timestamp,required"`
	Token0        SwapGetEvmResponseDataToken0 `json:"token0,required"`
	Token1        SwapGetEvmResponseDataToken1 `json:"token1,required"`
	TransactionID string                       `json:"transaction_id,required"`
	Value0        float64                      `json:"value0,required"`
	Value1        float64                      `json:"value1,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount0       respjson.Field
		Amount1       respjson.Field
		BlockNum      respjson.Field
		Caller        respjson.Field
		Datetime      respjson.Field
		Factory       respjson.Field
		Fee           respjson.Field
		NetworkID     respjson.Field
		Pool          respjson.Field
		Price0        respjson.Field
		Price1        respjson.Field
		Protocol      respjson.Field
		Recipient     respjson.Field
		Sender        respjson.Field
		Timestamp     respjson.Field
		Token0        respjson.Field
		Token1        respjson.Field
		TransactionID respjson.Field
		Value0        respjson.Field
		Value1        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwapGetEvmResponseData) RawJSON() string { return r.JSON.raw }
func (r *SwapGetEvmResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetEvmResponseDataToken0 struct {
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
func (r SwapGetEvmResponseDataToken0) RawJSON() string { return r.JSON.raw }
func (r *SwapGetEvmResponseDataToken0) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetEvmResponseDataToken1 struct {
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
func (r SwapGetEvmResponseDataToken1) RawJSON() string { return r.JSON.raw }
func (r *SwapGetEvmResponseDataToken1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetEvmResponsePagination struct {
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
func (r SwapGetEvmResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SwapGetEvmResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetEvmResponseStatistics struct {
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
func (r SwapGetEvmResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *SwapGetEvmResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetSvmResponse struct {
	Data         []SwapGetSvmResponseData     `json:"data,required"`
	DurationMs   float64                      `json:"duration_ms,required"`
	Pagination   SwapGetSvmResponsePagination `json:"pagination,required"`
	RequestTime  string                       `json:"request_time,required"`
	Results      float64                      `json:"results,required"`
	Statistics   SwapGetSvmResponseStatistics `json:"statistics,required"`
	TotalResults float64                      `json:"total_results,required"`
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
func (r SwapGetSvmResponse) RawJSON() string { return r.JSON.raw }
func (r *SwapGetSvmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetSvmResponseData struct {
	// Filter by address
	Amm              string                          `json:"amm,required"`
	AmmName          string                          `json:"amm_name,required"`
	BlockNum         float64                         `json:"block_num,required"`
	Datetime         time.Time                       `json:"datetime,required" format:"date-time"`
	InputAmount      float64                         `json:"input_amount,required"`
	InputMint        SwapGetSvmResponseDataInputMint `json:"input_mint,required"`
	InstructionIndex float64                         `json:"instruction_index,required"`
	// The Graph Network ID for SVM networks https://thegraph.com/networks
	//
	// Any of "solana".
	NetworkID    string                           `json:"network_id,required"`
	OutputAmount float64                          `json:"output_amount,required"`
	OutputMint   SwapGetSvmResponseDataOutputMint `json:"output_mint,required"`
	// Filter by address
	ProgramID        string  `json:"program_id,required"`
	ProgramName      string  `json:"program_name,required"`
	Signature        string  `json:"signature,required"`
	Timestamp        float64 `json:"timestamp,required"`
	TransactionIndex float64 `json:"transaction_index,required"`
	// Filter by address
	User string `json:"user,required"`
	// Filter by address
	AmmPool string `json:"amm_pool"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amm              respjson.Field
		AmmName          respjson.Field
		BlockNum         respjson.Field
		Datetime         respjson.Field
		InputAmount      respjson.Field
		InputMint        respjson.Field
		InstructionIndex respjson.Field
		NetworkID        respjson.Field
		OutputAmount     respjson.Field
		OutputMint       respjson.Field
		ProgramID        respjson.Field
		ProgramName      respjson.Field
		Signature        respjson.Field
		Timestamp        respjson.Field
		TransactionIndex respjson.Field
		User             respjson.Field
		AmmPool          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwapGetSvmResponseData) RawJSON() string { return r.JSON.raw }
func (r *SwapGetSvmResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetSvmResponseDataInputMint struct {
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
func (r SwapGetSvmResponseDataInputMint) RawJSON() string { return r.JSON.raw }
func (r *SwapGetSvmResponseDataInputMint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetSvmResponseDataOutputMint struct {
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
func (r SwapGetSvmResponseDataOutputMint) RawJSON() string { return r.JSON.raw }
func (r *SwapGetSvmResponseDataOutputMint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetSvmResponsePagination struct {
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
func (r SwapGetSvmResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SwapGetSvmResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetSvmResponseStatistics struct {
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
func (r SwapGetSvmResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *SwapGetSvmResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwapGetEvmParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID SwapGetEvmParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by address
	Caller param.Opt[string] `query:"caller,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by pool address
	Pool param.Opt[string] `query:"pool,omitzero" json:"-"`
	// Filter by address
	Recipient param.Opt[string] `query:"recipient,omitzero" json:"-"`
	// Filter by address
	Sender param.Opt[string] `query:"sender,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// Filter by transaction
	TransactionID param.Opt[string] `query:"transaction_id,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "timestamp".
	OrderBy SwapGetEvmParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection SwapGetEvmParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	// Protocol name
	//
	// Any of "", "uniswap_v2", "uniswap_v3", "uniswap_v4".
	Protocol SwapGetEvmParamsProtocol `query:"protocol,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SwapGetEvmParams]'s query parameters as `url.Values`.
func (r SwapGetEvmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type SwapGetEvmParamsNetworkID string

const (
	SwapGetEvmParamsNetworkIDArbitrumOne SwapGetEvmParamsNetworkID = "arbitrum-one"
	SwapGetEvmParamsNetworkIDAvalanche   SwapGetEvmParamsNetworkID = "avalanche"
	SwapGetEvmParamsNetworkIDBase        SwapGetEvmParamsNetworkID = "base"
	SwapGetEvmParamsNetworkIDBsc         SwapGetEvmParamsNetworkID = "bsc"
	SwapGetEvmParamsNetworkIDMainnet     SwapGetEvmParamsNetworkID = "mainnet"
	SwapGetEvmParamsNetworkIDMatic       SwapGetEvmParamsNetworkID = "matic"
	SwapGetEvmParamsNetworkIDOptimism    SwapGetEvmParamsNetworkID = "optimism"
	SwapGetEvmParamsNetworkIDUnichain    SwapGetEvmParamsNetworkID = "unichain"
)

// The field by which to order the results.
type SwapGetEvmParamsOrderBy string

const (
	SwapGetEvmParamsOrderByTimestamp SwapGetEvmParamsOrderBy = "timestamp"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type SwapGetEvmParamsOrderDirection string

const (
	SwapGetEvmParamsOrderDirectionAsc  SwapGetEvmParamsOrderDirection = "asc"
	SwapGetEvmParamsOrderDirectionDesc SwapGetEvmParamsOrderDirection = "desc"
)

// Protocol name
type SwapGetEvmParamsProtocol string

const (
	SwapGetEvmParamsProtocolEmpty     SwapGetEvmParamsProtocol = ""
	SwapGetEvmParamsProtocolUniswapV2 SwapGetEvmParamsProtocol = "uniswap_v2"
	SwapGetEvmParamsProtocolUniswapV3 SwapGetEvmParamsProtocol = "uniswap_v3"
	SwapGetEvmParamsProtocolUniswapV4 SwapGetEvmParamsProtocol = "uniswap_v4"
)

type SwapGetSvmParams struct {
	// The Graph Network ID for SVM networks https://thegraph.com/networks
	//
	// Any of "solana".
	NetworkID SwapGetSvmParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by amm address
	Amm param.Opt[string] `query:"amm,omitzero" json:"-"`
	// Filter by amm pool address
	AmmPool param.Opt[string] `query:"amm_pool,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// Filter by mint address
	InputMint param.Opt[string] `query:"input_mint,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by mint address
	OutputMint param.Opt[string] `query:"output_mint,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by transaction signature
	Signature param.Opt[string] `query:"signature,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// Filter by user address
	User param.Opt[string] `query:"user,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "timestamp".
	OrderBy SwapGetSvmParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection SwapGetSvmParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	// Filter by program ID
	//
	// Any of "", "675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8",
	// "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P",
	// "pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA",
	// "JUP4Fb2cqiRUcaTHdrPC8h2gNsA2ETXiPDD33WcGuJB",
	// "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4".
	ProgramID SwapGetSvmParamsProgramID `query:"program_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SwapGetSvmParams]'s query parameters as `url.Values`.
func (r SwapGetSvmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for SVM networks https://thegraph.com/networks
type SwapGetSvmParamsNetworkID string

const (
	SwapGetSvmParamsNetworkIDSolana SwapGetSvmParamsNetworkID = "solana"
)

// The field by which to order the results.
type SwapGetSvmParamsOrderBy string

const (
	SwapGetSvmParamsOrderByTimestamp SwapGetSvmParamsOrderBy = "timestamp"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type SwapGetSvmParamsOrderDirection string

const (
	SwapGetSvmParamsOrderDirectionAsc  SwapGetSvmParamsOrderDirection = "asc"
	SwapGetSvmParamsOrderDirectionDesc SwapGetSvmParamsOrderDirection = "desc"
)

// Filter by program ID
type SwapGetSvmParamsProgramID string

const (
	SwapGetSvmParamsProgramIDEmpty                                        SwapGetSvmParamsProgramID = ""
	SwapGetSvmParamsProgramID675kPx9MhTjS2zt1qfr1NyHuzeLXfQm9H24wFsUt1Mp8 SwapGetSvmParamsProgramID = "675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8"
	SwapGetSvmParamsProgramID6Ef8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P  SwapGetSvmParamsProgramID = "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P"
	SwapGetSvmParamsProgramIDPAmmBay6oceH9fJkbrhgp5D4bD4sWpmSwMn52FMfXea  SwapGetSvmParamsProgramID = "pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA"
	SwapGetSvmParamsProgramIDJup4Fb2cqiRUcaTHdrPc8h2gNsA2EtXiPdd33WcGuJb  SwapGetSvmParamsProgramID = "JUP4Fb2cqiRUcaTHdrPC8h2gNsA2ETXiPDD33WcGuJB"
	SwapGetSvmParamsProgramIDJup6LkbZbjS1jKKwapdHNy74zcZ3tLuZoi5QNyVTaV4  SwapGetSvmParamsProgramID = "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4"
)
