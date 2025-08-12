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
	"github.com/stainless-sdks/the-graph-go/packages/respjson"
)

// HealthService contains methods and other services that help with interacting
// with the the-graph API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Returns API operational status and dependency health. Use `skip_endpoints=true`
// for faster database-only checks.
func (r *HealthService) Check(ctx context.Context, query HealthCheckParams, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HealthCheckResponse struct {
	Checks      HealthCheckResponseChecks `json:"checks,required"`
	DurationMs  float64                   `json:"duration_ms,required"`
	RequestTime time.Time                 `json:"request_time,required" format:"date-time"`
	// Any of "healthy", "degraded", "unhealthy".
	Status HealthCheckResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		DurationMs  respjson.Field
		RequestTime respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthCheckResponseChecks struct {
	// Any of "up", "down", "partial", "skipped".
	APIEndpoints string `json:"api_endpoints,required"`
	// Any of "up", "down", "slow".
	Database string `json:"database,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIEndpoints respjson.Field
		Database     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponseChecks) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponseChecks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthCheckResponseStatus string

const (
	HealthCheckResponseStatusHealthy   HealthCheckResponseStatus = "healthy"
	HealthCheckResponseStatusDegraded  HealthCheckResponseStatus = "degraded"
	HealthCheckResponseStatusUnhealthy HealthCheckResponseStatus = "unhealthy"
)

type HealthCheckParams struct {
	// Any of "true", "false".
	SkipEndpoints HealthCheckParamsSkipEndpoints `query:"skip_endpoints,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HealthCheckParams]'s query parameters as `url.Values`.
func (r HealthCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HealthCheckParamsSkipEndpoints string

const (
	HealthCheckParamsSkipEndpointsTrue  HealthCheckParamsSkipEndpoints = "true"
	HealthCheckParamsSkipEndpointsFalse HealthCheckParamsSkipEndpoints = "false"
)
