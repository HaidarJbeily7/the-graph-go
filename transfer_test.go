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

func TestTransferListEvmWithOptionalParams(t *testing.T) {
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
	_, err := client.Transfers.ListEvm(context.TODO(), thegraph.TransferListEvmParams{
		NetworkID:      thegraph.TransferListEvmParamsNetworkIDMainnet,
		Contract:       thegraph.String("contract"),
		EndTime:        thegraph.Int(-9007199254740991),
		From:           thegraph.String("from"),
		Limit:          thegraph.Int(1),
		OrderBy:        thegraph.TransferListEvmParamsOrderByTimestamp,
		OrderDirection: thegraph.TransferListEvmParamsOrderDirectionAsc,
		Page:           thegraph.Int(1),
		StartTime:      thegraph.Int(-9007199254740991),
		To:             thegraph.String("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"),
		TransactionID:  thegraph.String("transaction_id"),
	})
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransferListSvmWithOptionalParams(t *testing.T) {
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
	_, err := client.Transfers.ListSvm(context.TODO(), thegraph.TransferListSvmParams{
		NetworkID:      thegraph.TransferListSvmParamsNetworkIDSolana,
		Authority:      thegraph.String("authority"),
		Destination:    thegraph.String("destination"),
		EndTime:        thegraph.Int(-9007199254740991),
		Limit:          thegraph.Int(1),
		Mint:           thegraph.String("So11111111111111111111111111111111111111112"),
		OrderBy:        thegraph.TransferListSvmParamsOrderByTimestamp,
		OrderDirection: thegraph.TransferListSvmParamsOrderDirectionAsc,
		Page:           thegraph.Int(1),
		ProgramID:      thegraph.TransferListSvmParamsProgramIDEmpty,
		Source:         thegraph.String("source"),
		StartTime:      thegraph.Int(-9007199254740991),
	})
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
