// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/the-graph-go/internal/apijson"
	"github.com/stainless-sdks/the-graph-go/internal/apiquery"
	"github.com/stainless-sdks/the-graph-go/internal/requestconfig"
	"github.com/stainless-sdks/the-graph-go/option"
	"github.com/stainless-sdks/the-graph-go/packages/param"
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// NFTActivityService contains methods and other services that help with
// interacting with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTActivityService] method instead.
type NFTActivityService struct {
	Options []option.RequestOption
}

// NewNFTActivityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNFTActivityService(opts ...option.RequestOption) (r NFTActivityService) {
	r = NFTActivityService{}
	r.Options = opts
	return
}

// Returns NFT transfer events including mints, burns, and ownership changes.
func (r *NFTActivityService) List(ctx context.Context, query NFTActivityListParams, opts ...option.RequestOption) (res *NFTActivityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "nft/activities/evm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NFTActivityListResponse struct {
	Data         []NFTActivityListResponseData     `json:"data,required"`
	DurationMs   float64                           `json:"duration_ms,required"`
	Pagination   NFTActivityListResponsePagination `json:"pagination,required"`
	RequestTime  string                            `json:"request_time,required"`
	Results      float64                           `json:"results,required"`
	Statistics   NFTActivityListResponseStatistics `json:"statistics,required"`
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
func (r NFTActivityListResponse) RawJSON() string { return r.JSON.raw }
func (r *NFTActivityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTActivityListResponseData struct {
	// Any of "TRANSFER", "MINT", "BURN".
	Type          string  `json:"@type,required"`
	Amount        float64 `json:"amount,required"`
	BlockHash     string  `json:"block_hash,required"`
	BlockNum      float64 `json:"block_num,required"`
	Contract      string  `json:"contract,required"`
	From          string  `json:"from,required"`
	Timestamp     string  `json:"timestamp,required"`
	To            string  `json:"to,required"`
	TokenID       string  `json:"token_id,required"`
	TxHash        string  `json:"tx_hash,required"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	TokenStandard string  `json:"token_standard"`
	TransferType  string  `json:"transfer_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		Amount        respjson.Field
		BlockHash     respjson.Field
		BlockNum      respjson.Field
		Contract      respjson.Field
		From          respjson.Field
		Timestamp     respjson.Field
		To            respjson.Field
		TokenID       respjson.Field
		TxHash        respjson.Field
		Name          respjson.Field
		Symbol        respjson.Field
		TokenStandard respjson.Field
		TransferType  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTActivityListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NFTActivityListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTActivityListResponsePagination struct {
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
func (r NFTActivityListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *NFTActivityListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTActivityListResponseStatistics struct {
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
func (r NFTActivityListResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *NFTActivityListResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTActivityListParams struct {
	// Filter by NFT contract address
	Contract string `query:"contract,required" json:"-"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID NFTActivityListParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by address
	AnyAddress param.Opt[string] `query:"anyAddress,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// Filter by address
	FromAddress param.Opt[string] `query:"fromAddress,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// Filter by address
	ToAddress param.Opt[string] `query:"toAddress,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "timestamp".
	OrderBy NFTActivityListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection NFTActivityListParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NFTActivityListParams]'s query parameters as `url.Values`.
func (r NFTActivityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type NFTActivityListParamsNetworkID string

const (
	NFTActivityListParamsNetworkIDArbitrumOne NFTActivityListParamsNetworkID = "arbitrum-one"
	NFTActivityListParamsNetworkIDAvalanche   NFTActivityListParamsNetworkID = "avalanche"
	NFTActivityListParamsNetworkIDBase        NFTActivityListParamsNetworkID = "base"
	NFTActivityListParamsNetworkIDBsc         NFTActivityListParamsNetworkID = "bsc"
	NFTActivityListParamsNetworkIDMainnet     NFTActivityListParamsNetworkID = "mainnet"
	NFTActivityListParamsNetworkIDMatic       NFTActivityListParamsNetworkID = "matic"
	NFTActivityListParamsNetworkIDOptimism    NFTActivityListParamsNetworkID = "optimism"
	NFTActivityListParamsNetworkIDUnichain    NFTActivityListParamsNetworkID = "unichain"
)

// The field by which to order the results.
type NFTActivityListParamsOrderBy string

const (
	NFTActivityListParamsOrderByTimestamp NFTActivityListParamsOrderBy = "timestamp"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type NFTActivityListParamsOrderDirection string

const (
	NFTActivityListParamsOrderDirectionAsc  NFTActivityListParamsOrderDirection = "asc"
	NFTActivityListParamsOrderDirectionDesc NFTActivityListParamsOrderDirection = "desc"
)
