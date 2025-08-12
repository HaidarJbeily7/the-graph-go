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

func TestSwapGetEvmWithOptionalParams(t *testing.T) {
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
	_, err := client.Swaps.GetEvm(context.TODO(), thegraph.SwapGetEvmParams{
		NetworkID:      thegraph.SwapGetEvmParamsNetworkIDMainnet,
		Caller:         thegraph.String("caller"),
		EndTime:        thegraph.Int(-9007199254740991),
		Limit:          thegraph.Int(1),
		OrderBy:        thegraph.SwapGetEvmParamsOrderByTimestamp,
		OrderDirection: thegraph.SwapGetEvmParamsOrderDirectionAsc,
		Page:           thegraph.Int(1),
		Pool:           thegraph.String("0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"),
		Protocol:       thegraph.SwapGetEvmParamsProtocolUniswapV3,
		Recipient:      thegraph.String("recipient"),
		Sender:         thegraph.String("sender"),
		StartTime:      thegraph.Int(-9007199254740991),
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

func TestSwapGetSvmWithOptionalParams(t *testing.T) {
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
	_, err := client.Swaps.GetSvm(context.TODO(), thegraph.SwapGetSvmParams{
		NetworkID:      thegraph.SwapGetSvmParamsNetworkIDSolana,
		Amm:            thegraph.String("amm"),
		AmmPool:        thegraph.String("amm_pool"),
		EndTime:        thegraph.Int(-9007199254740991),
		InputMint:      thegraph.String("input_mint"),
		Limit:          thegraph.Int(1),
		OrderBy:        thegraph.SwapGetSvmParamsOrderByTimestamp,
		OrderDirection: thegraph.SwapGetSvmParamsOrderDirectionAsc,
		OutputMint:     thegraph.String("output_mint"),
		Page:           thegraph.Int(1),
		ProgramID:      thegraph.SwapGetSvmParamsProgramIDPAmmBay6oceH9fJkbrhgp5D4bD4sWpmSwMn52FMfXea,
		Signature:      thegraph.String("signature"),
		StartTime:      thegraph.Int(-9007199254740991),
		User:           thegraph.String("user"),
	})
	if err != nil {
		var apierr *thegraph.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
