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

func TestBalanceListSvmWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.ListSvm(context.TODO(), thegraph.BalanceListSvmParams{
		NetworkID:    thegraph.BalanceListSvmParamsNetworkIDSolana,
		Limit:        thegraph.Int(1),
		Mint:         thegraph.String("So11111111111111111111111111111111111111112"),
		Page:         thegraph.Int(1),
		ProgramID:    thegraph.BalanceListSvmParamsProgramIDEmpty,
		TokenAccount: thegraph.String("token_account"),
	})
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBalanceGetEvmWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.GetEvm(
		context.TODO(),
		"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		thegraph.BalanceGetEvmParams{
			NetworkID: thegraph.BalanceGetEvmParamsNetworkIDMainnet,
			Contract:  thegraph.String("contract"),
			Limit:     thegraph.Int(1),
			Page:      thegraph.Int(1),
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
