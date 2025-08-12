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

func TestNFTOwnershipGetByAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.NFT.Ownerships.GetByAddress(
		context.TODO(),
		"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		thegraph.NFTOwnershipGetByAddressParams{
			NetworkID:     thegraph.NFTOwnershipGetByAddressParamsNetworkIDMainnet,
			Contract:      thegraph.String("contract"),
			Limit:         thegraph.Int(1),
			Page:          thegraph.Int(1),
			TokenStandard: thegraph.NFTOwnershipGetByAddressParamsTokenStandardEmpty,
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
