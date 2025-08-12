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

func TestNFTItemEvmContractGetByToken(t *testing.T) {
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
	_, err := client.NFT.Items.Evm.Contract.GetByToken(
		context.TODO(),
		"5712",
		thegraph.NFTItemEvmContractGetByTokenParams{
			Contract:  "0xbd3531da5cf5857e7cfaa92426877b022e612cf8",
			NetworkID: thegraph.NFTItemEvmContractGetByTokenParamsNetworkIDMainnet,
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
