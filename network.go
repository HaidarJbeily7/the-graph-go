// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/the-graph-go/internal/apijson"
	"github.com/stainless-sdks/the-graph-go/internal/requestconfig"
	"github.com/stainless-sdks/the-graph-go/option"
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// NetworkService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options []option.RequestOption
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r NetworkService) {
	r = NetworkService{}
	r.Options = opts
	return
}

// Returns supported blockchain networks with identifiers and metadata.
func (r *NetworkService) List(ctx context.Context, opts ...option.RequestOption) (res *NetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "networks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NetworkListResponse struct {
	Networks []NetworkListResponseNetwork `json:"networks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Networks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponseNetwork struct {
	ID          string                         `json:"id,required"`
	Alias       []string                       `json:"alias,required"`
	Caip2ID     string                         `json:"caip2Id,required"`
	FullName    string                         `json:"fullName,required"`
	Icon        NetworkListResponseNetworkIcon `json:"icon,required"`
	NetworkType string                         `json:"networkType,required"`
	ShortName   string                         `json:"shortName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Alias       respjson.Field
		Caip2ID     respjson.Field
		FullName    respjson.Field
		Icon        respjson.Field
		NetworkType respjson.Field
		ShortName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkListResponseNetwork) RawJSON() string { return r.JSON.raw }
func (r *NetworkListResponseNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponseNetworkIcon struct {
	Web3Icons NetworkListResponseNetworkIconWeb3Icons `json:"web3Icons,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Web3Icons   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkListResponseNetworkIcon) RawJSON() string { return r.JSON.raw }
func (r *NetworkListResponseNetworkIcon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponseNetworkIconWeb3Icons struct {
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkListResponseNetworkIconWeb3Icons) RawJSON() string { return r.JSON.raw }
func (r *NetworkListResponseNetworkIconWeb3Icons) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
