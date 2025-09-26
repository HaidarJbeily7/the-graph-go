// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
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

// TransferService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferService] method instead.
type TransferService struct {
	Options []option.RequestOption
}

// NewTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferService(opts ...option.RequestOption) (r TransferService) {
	r = TransferService{}
	r.Options = opts
	return
}

// Returns ERC-20 and native token transfers with transaction and block data.
func (r *TransferService) ListEvm(ctx context.Context, query TransferListEvmParams, opts ...option.RequestOption) (res *TransferListEvmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transfers/evm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns SPL token transfers with program, authority, and account information.
func (r *TransferService) ListSvm(ctx context.Context, query TransferListSvmParams, opts ...option.RequestOption) (res *TransferListSvmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transfers/svm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TransferListEvmResponse struct {
	Data         []TransferListEvmResponseData     `json:"data,required"`
	DurationMs   float64                           `json:"duration_ms,required"`
	Pagination   TransferListEvmResponsePagination `json:"pagination,required"`
	RequestTime  string                            `json:"request_time,required"`
	Results      float64                           `json:"results,required"`
	Statistics   TransferListEvmResponseStatistics `json:"statistics,required"`
	TotalResults float64                           `json:"total_results,required"`
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
func (r TransferListEvmResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferListEvmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListEvmResponseData struct {
	Amount   string  `json:"amount,required"`
	BlockNum float64 `json:"block_num,required"`
	// Filter by address
	Contract string    `json:"contract,required"`
	Datetime time.Time `json:"datetime,required" format:"date-time"`
	// Filter by address
	From string `json:"from,required"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID string  `json:"network_id,required"`
	Timestamp float64 `json:"timestamp,required"`
	// Filter by address
	To            string  `json:"to,required"`
	TransactionID string  `json:"transaction_id,required"`
	Value         float64 `json:"value,required"`
	Decimals      float64 `json:"decimals"`
	Symbol        string  `json:"symbol"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		BlockNum      respjson.Field
		Contract      respjson.Field
		Datetime      respjson.Field
		From          respjson.Field
		NetworkID     respjson.Field
		Timestamp     respjson.Field
		To            respjson.Field
		TransactionID respjson.Field
		Value         respjson.Field
		Decimals      respjson.Field
		Symbol        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferListEvmResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferListEvmResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListEvmResponsePagination struct {
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
func (r TransferListEvmResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *TransferListEvmResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListEvmResponseStatistics struct {
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
func (r TransferListEvmResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *TransferListEvmResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListSvmResponse struct {
	Data         []TransferListSvmResponseData     `json:"data,required"`
	DurationMs   float64                           `json:"duration_ms,required"`
	Pagination   TransferListSvmResponsePagination `json:"pagination,required"`
	RequestTime  string                            `json:"request_time,required"`
	Results      float64                           `json:"results,required"`
	Statistics   TransferListSvmResponseStatistics `json:"statistics,required"`
	TotalResults float64                           `json:"total_results,required"`
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
func (r TransferListSvmResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferListSvmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListSvmResponseData struct {
	Amount string `json:"amount,required"`
	// Filter by address
	Authority string    `json:"authority,required"`
	BlockNum  float64   `json:"block_num,required"`
	Datetime  time.Time `json:"datetime,required" format:"date-time"`
	Decimals  int64     `json:"decimals,required"`
	// Filter by address
	Destination string `json:"destination,required"`
	// Filter by address
	Mint string `json:"mint,required"`
	// The Graph Network ID for SVM networks https://thegraph.com/networks
	//
	// Any of "solana".
	NetworkID string `json:"network_id,required"`
	// Filter by address
	ProgramID string `json:"program_id,required"`
	Signature string `json:"signature,required"`
	// Filter by address
	Source    string  `json:"source,required"`
	Timestamp float64 `json:"timestamp,required"`
	Value     float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Authority   respjson.Field
		BlockNum    respjson.Field
		Datetime    respjson.Field
		Decimals    respjson.Field
		Destination respjson.Field
		Mint        respjson.Field
		NetworkID   respjson.Field
		ProgramID   respjson.Field
		Signature   respjson.Field
		Source      respjson.Field
		Timestamp   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferListSvmResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferListSvmResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListSvmResponsePagination struct {
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
func (r TransferListSvmResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *TransferListSvmResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListSvmResponseStatistics struct {
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
func (r TransferListSvmResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *TransferListSvmResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListEvmParams struct {
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID TransferListEvmParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by address
	Contract param.Opt[string] `query:"contract,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// Filter by address
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// Filter by address
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Filter by transaction
	TransactionID param.Opt[string] `query:"transaction_id,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "timestamp".
	OrderBy TransferListEvmParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection TransferListEvmParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransferListEvmParams]'s query parameters as `url.Values`.
func (r TransferListEvmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type TransferListEvmParamsNetworkID string

const (
	TransferListEvmParamsNetworkIDArbitrumOne TransferListEvmParamsNetworkID = "arbitrum-one"
	TransferListEvmParamsNetworkIDAvalanche   TransferListEvmParamsNetworkID = "avalanche"
	TransferListEvmParamsNetworkIDBase        TransferListEvmParamsNetworkID = "base"
	TransferListEvmParamsNetworkIDBsc         TransferListEvmParamsNetworkID = "bsc"
	TransferListEvmParamsNetworkIDMainnet     TransferListEvmParamsNetworkID = "mainnet"
	TransferListEvmParamsNetworkIDMatic       TransferListEvmParamsNetworkID = "matic"
	TransferListEvmParamsNetworkIDOptimism    TransferListEvmParamsNetworkID = "optimism"
	TransferListEvmParamsNetworkIDUnichain    TransferListEvmParamsNetworkID = "unichain"
)

// The field by which to order the results.
type TransferListEvmParamsOrderBy string

const (
	TransferListEvmParamsOrderByTimestamp TransferListEvmParamsOrderBy = "timestamp"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type TransferListEvmParamsOrderDirection string

const (
	TransferListEvmParamsOrderDirectionAsc  TransferListEvmParamsOrderDirection = "asc"
	TransferListEvmParamsOrderDirectionDesc TransferListEvmParamsOrderDirection = "desc"
)

type TransferListSvmParams struct {
	// The Graph Network ID for SVM networks https://thegraph.com/networks
	//
	// Any of "solana".
	NetworkID TransferListSvmParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by authority token account address
	Authority param.Opt[string] `query:"authority,omitzero" json:"-"`
	// Filter by token account address
	Destination param.Opt[string] `query:"destination,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by mint address
	Mint param.Opt[string] `query:"mint,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by token account address
	Source param.Opt[string] `query:"source,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "timestamp".
	OrderBy TransferListSvmParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection TransferListSvmParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	// Filter by program ID
	//
	// Any of "", "TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb",
	// "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA".
	ProgramID TransferListSvmParamsProgramID `query:"program_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransferListSvmParams]'s query parameters as `url.Values`.
func (r TransferListSvmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for SVM networks https://thegraph.com/networks
type TransferListSvmParamsNetworkID string

const (
	TransferListSvmParamsNetworkIDSolana TransferListSvmParamsNetworkID = "solana"
)

// The field by which to order the results.
type TransferListSvmParamsOrderBy string

const (
	TransferListSvmParamsOrderByTimestamp TransferListSvmParamsOrderBy = "timestamp"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type TransferListSvmParamsOrderDirection string

const (
	TransferListSvmParamsOrderDirectionAsc  TransferListSvmParamsOrderDirection = "asc"
	TransferListSvmParamsOrderDirectionDesc TransferListSvmParamsOrderDirection = "desc"
)

// Filter by program ID
type TransferListSvmParamsProgramID string

const (
	TransferListSvmParamsProgramIDEmpty                                       TransferListSvmParamsProgramID = ""
	TransferListSvmParamsProgramIDTokenzQdBNbLqP5VEhdkAs6Epflc1PHnBqCxEpPxuEb TransferListSvmParamsProgramID = "TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"
	TransferListSvmParamsProgramIDTokenkegQfeZyiNwAJbNbGkpfxcWuBvf9Ss623Vq5Da TransferListSvmParamsProgramID = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
)
