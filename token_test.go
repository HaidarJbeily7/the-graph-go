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

func TestTokenGetMetadata(t *testing.T) {
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
	_, err := client.Tokens.GetMetadata(
		context.TODO(),
		"0xc944e90c64b2c07662a292be6244bdf05cda44a7",
		thegraph.TokenGetMetadataParams{
			NetworkID: thegraph.TokenGetMetadataParamsNetworkIDMainnet,
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
