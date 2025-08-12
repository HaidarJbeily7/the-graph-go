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

func TestPoolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Pools.Get(context.TODO(), thegraph.PoolGetParams{
		NetworkID: thegraph.PoolGetParamsNetworkIDMainnet,
		Token:     thegraph.String("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		Factory:   thegraph.String("factory"),
		Limit:     thegraph.Int(1),
		Page:      thegraph.Int(1),
		Pool:      thegraph.String("0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"),
		Protocol:  thegraph.PoolGetParamsProtocolUniswapV3,
	})
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
