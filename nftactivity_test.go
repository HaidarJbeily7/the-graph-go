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

func TestNFTActivityListWithOptionalParams(t *testing.T) {
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
	_, err := client.NFT.Activities.List(context.TODO(), thegraph.NFTActivityListParams{
		Contract:       "0xbd3531da5cf5857e7cfaa92426877b022e612cf8",
		NetworkID:      thegraph.NFTActivityListParamsNetworkIDMainnet,
		AnyAddress:     thegraph.String("anyAddress"),
		EndTime:        thegraph.Int(-9007199254740991),
		FromAddress:    thegraph.String("fromAddress"),
		Limit:          thegraph.Int(1),
		OrderBy:        thegraph.NFTActivityListParamsOrderByTimestamp,
		OrderDirection: thegraph.NFTActivityListParamsOrderDirectionAsc,
		Page:           thegraph.Int(1),
		StartTime:      thegraph.Int(-9007199254740991),
		ToAddress:      thegraph.String("toAddress"),
	})
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
