// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"net/http"
	"net/url"

	"github.com/HaidarJbeily7/the-graph-go/internal/apijson"
	"github.com/HaidarJbeily7/the-graph-go/internal/apiquery"
	"github.com/HaidarJbeily7/the-graph-go/internal/requestconfig"
	"github.com/HaidarJbeily7/the-graph-go/option"
	"github.com/HaidarJbeily7/the-graph-go/packages/param"
	"github.com/HaidarJbeily7/the-graph-go/packages/respjson"
)

// NFTSaleService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNFTSaleService] method instead.
type NFTSaleService struct {
	Options []option.RequestOption
}

// NewNFTSaleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNFTSaleService(opts ...option.RequestOption) (r NFTSaleService) {
	r = NFTSaleService{}
	r.Options = opts
	return
}

// Returns NFT marketplace sales with price, buyer, seller, and transaction data.
func (r *NFTSaleService) List(ctx context.Context, query NFTSaleListParams, opts ...option.RequestOption) (res *NFTSaleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "nft/sales/evm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NFTSaleListResponse struct {
	Data         []NFTSaleListResponseData     `json:"data,required"`
	DurationMs   float64                       `json:"duration_ms,required"`
	Pagination   NFTSaleListResponsePagination `json:"pagination,required"`
	RequestTime  string                        `json:"request_time,required"`
	Results      float64                       `json:"results,required"`
	Statistics   NFTSaleListResponseStatistics `json:"statistics,required"`
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
func (r NFTSaleListResponse) RawJSON() string { return r.JSON.raw }
func (r *NFTSaleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTSaleListResponseData struct {
	Token        string  `json:"token,required"`
	BlockNum     float64 `json:"block_num,required"`
	Name         string  `json:"name,required"`
	Offerer      string  `json:"offerer,required"`
	Recipient    string  `json:"recipient,required"`
	SaleAmount   float64 `json:"sale_amount,required"`
	SaleCurrency string  `json:"sale_currency,required"`
	Symbol       string  `json:"symbol,required"`
	Timestamp    string  `json:"timestamp,required"`
	TokenID      string  `json:"token_id,required"`
	TxHash       string  `json:"tx_hash,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token        respjson.Field
		BlockNum     respjson.Field
		Name         respjson.Field
		Offerer      respjson.Field
		Recipient    respjson.Field
		SaleAmount   respjson.Field
		SaleCurrency respjson.Field
		Symbol       respjson.Field
		Timestamp    respjson.Field
		TokenID      respjson.Field
		TxHash       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTSaleListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NFTSaleListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTSaleListResponsePagination struct {
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
func (r NFTSaleListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *NFTSaleListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTSaleListResponseStatistics struct {
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
func (r NFTSaleListResponseStatistics) RawJSON() string { return r.JSON.raw }
func (r *NFTSaleListResponseStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NFTSaleListParams struct {
	// Filter by NFT contract address
	Contract string `query:"contract,required" json:"-"`
	// The Graph Network ID for EVM networks https://thegraph.com/networks
	//
	// Any of "arbitrum-one", "avalanche", "base", "bsc", "mainnet", "matic",
	// "optimism", "unichain".
	NetworkID NFTSaleListParamsNetworkID `query:"network_id,omitzero,required" json:"-"`
	// Filter by address
	AnyAddress param.Opt[string] `query:"anyAddress,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	EndTime param.Opt[int64] `query:"endTime,omitzero" json:"-"`
	// The maximum number of items returned in a single request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by address
	OffererAddress param.Opt[string] `query:"offererAddress,omitzero" json:"-"`
	// The page number of the results to return.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by address
	RecipientAddress param.Opt[string] `query:"recipientAddress,omitzero" json:"-"`
	// UNIX timestamp in seconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	// NFT token ID
	TokenID param.Opt[string] `query:"token_id,omitzero" json:"-"`
	// The field by which to order the results.
	//
	// Any of "timestamp".
	OrderBy NFTSaleListParamsOrderBy `query:"orderBy,omitzero" json:"-"`
	// The order in which to return the results: Ascending (asc) or Descending (desc).
	//
	// Any of "asc", "desc".
	OrderDirection NFTSaleListParamsOrderDirection `query:"orderDirection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NFTSaleListParams]'s query parameters as `url.Values`.
func (r NFTSaleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Graph Network ID for EVM networks https://thegraph.com/networks
type NFTSaleListParamsNetworkID string

const (
	NFTSaleListParamsNetworkIDArbitrumOne NFTSaleListParamsNetworkID = "arbitrum-one"
	NFTSaleListParamsNetworkIDAvalanche   NFTSaleListParamsNetworkID = "avalanche"
	NFTSaleListParamsNetworkIDBase        NFTSaleListParamsNetworkID = "base"
	NFTSaleListParamsNetworkIDBsc         NFTSaleListParamsNetworkID = "bsc"
	NFTSaleListParamsNetworkIDMainnet     NFTSaleListParamsNetworkID = "mainnet"
	NFTSaleListParamsNetworkIDMatic       NFTSaleListParamsNetworkID = "matic"
	NFTSaleListParamsNetworkIDOptimism    NFTSaleListParamsNetworkID = "optimism"
	NFTSaleListParamsNetworkIDUnichain    NFTSaleListParamsNetworkID = "unichain"
)

// The field by which to order the results.
type NFTSaleListParamsOrderBy string

const (
	NFTSaleListParamsOrderByTimestamp NFTSaleListParamsOrderBy = "timestamp"
)

// The order in which to return the results: Ascending (asc) or Descending (desc).
type NFTSaleListParamsOrderDirection string

const (
	NFTSaleListParamsOrderDirectionAsc  NFTSaleListParamsOrderDirection = "asc"
	NFTSaleListParamsOrderDirectionDesc NFTSaleListParamsOrderDirection = "desc"
)
