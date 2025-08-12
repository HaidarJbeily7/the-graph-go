// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/the-graph-go"
	"github.com/stainless-sdks/the-graph-go/internal/testutil"
	"github.com/stainless-sdks/the-graph-go/option"
)

func TestOhlcPriceGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := thegraph.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Ohlc.Prices.Get(
		context.TODO(),
		"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		thegraph.OhlcPriceGetParams{
			NetworkID: thegraph.OhlcPriceGetParamsNetworkIDMainnet,
			EndTime:   thegraph.Int(-9007199254740991),
			Interval:  map[string]interface{}{},
			Limit:     thegraph.Int(1),
			Page:      thegraph.Int(1),
			StartTime: thegraph.Int(-9007199254740991),
		},
	)
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
