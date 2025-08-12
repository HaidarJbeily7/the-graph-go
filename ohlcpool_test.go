// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package thegraph_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/HaidarJbeily7/the-graph-go"
	"github.com/HaidarJbeily7/the-graph-go/internal/testutil"
	"github.com/HaidarJbeily7/the-graph-go/option"
)

func TestOhlcPoolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Ohlc.Pools.Get(
		context.TODO(),
		"0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640",
		thegraph.OhlcPoolGetParams{
			NetworkID: thegraph.OhlcPoolGetParamsNetworkIDMainnet,
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
