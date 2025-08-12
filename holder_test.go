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

func TestHolderGetByContractWithOptionalParams(t *testing.T) {
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
	_, err := client.Holders.GetByContract(
		context.TODO(),
		"0xc944e90c64b2c07662a292be6244bdf05cda44a7",
		thegraph.HolderGetByContractParams{
			NetworkID:      thegraph.HolderGetByContractParamsNetworkIDMainnet,
			Limit:          thegraph.Int(1),
			OrderBy:        thegraph.HolderGetByContractParamsOrderByValue,
			OrderDirection: thegraph.HolderGetByContractParamsOrderDirectionAsc,
			Page:           thegraph.Int(1),
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
