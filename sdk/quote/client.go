package quote

import (
	"context"
	"fmt"
	"strings"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the quote client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new quote client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// GetQuoteSummary gets the quote summary for a given contract
func (c *Client) GetQuoteSummary(ctx context.Context, contractID string) (*openapi.ResultGetTickerSummaryModel, error) {
	resp, _, err := c.openapiClient.Class01QuotePublicApiAPI.GetTicketSummary(ctx).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get quote summary: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// Get24HourQuotes gets the 24-hour quotes for given contracts
func (c *Client) Get24HourQuote(ctx context.Context, contractId string) (*openapi.ResultListTicker, error) {
	resp, _, err := c.openapiClient.Class01QuotePublicApiAPI.GetTicker(ctx).
		ContractId(contractId). // API only supports one contract ID
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get 24-hour quotes: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetKLineParams represents the parameters for GetKLine
type GetKLineParams struct {
	ContractID string
	Interval   string
	Size       int32
	From       *int64
	To         *int64
	PriceType  string
}

// GetKLine gets the K-line data for a contract
func (c *Client) GetKLine(ctx context.Context, params GetKLineParams) (*openapi.ResultPageDataKline, error) {
	req := c.openapiClient.Class01QuotePublicApiAPI.GetKline(ctx).
		ContractId(params.ContractID).
		KlineType(params.Interval).
		Size(fmt.Sprintf("%d", params.Size))

	if params.PriceType != "" {
		req = req.PriceType(params.PriceType)
	}
	if params.From != nil {
		req = req.FilterBeginKlineTimeInclusive(fmt.Sprintf("%d", *params.From))
	}
	if params.To != nil {
		req = req.FilterEndKlineTimeExclusive(fmt.Sprintf("%d", *params.To))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get k-line data: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetOrderBookDepthParams represents the parameters for GetOrderBookDepth
type GetOrderBookDepthParams struct {
	ContractID string
	Size       int32
	Precision  *string
}

// GetOrderBookDepth gets the order book depth for a contract
func (c *Client) GetOrderBookDepth(ctx context.Context, params GetOrderBookDepthParams) (*openapi.ResultListDepth, error) {
	req := c.openapiClient.Class01QuotePublicApiAPI.GetDepth(ctx).
		ContractId(params.ContractID).
		Level(fmt.Sprintf("%d", params.Size))

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get order book depth: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetMultiContractKLineParams represents the parameters for GetMultiContractKLine
type GetMultiContractKLineParams struct {
	ContractIDs []string
	Interval    string
	Size        int32
	From        *int64
	To          *int64
	PriceType   string
}

// GetMultiContractKLine gets the K-line data for multiple contracts
func (c *Client) GetMultiContractKLine(ctx context.Context, params GetMultiContractKLineParams) (*openapi.ResultListContractKline, error) {
	req := c.openapiClient.Class01QuotePublicApiAPI.GetMultiContractKline(ctx).
		ContractIdList(strings.Join(params.ContractIDs, ",")).
		KlineType(params.Interval).
		Size(fmt.Sprintf("%d", params.Size))

	if params.PriceType != "" {
		req = req.PriceType(params.PriceType)
	}
	if params.From != nil {
		req = req.FilterBeginKlineTimeInclusive(fmt.Sprintf("%d", *params.From))
	}
	if params.To != nil {
		req = req.FilterEndKlineTimeExclusive(fmt.Sprintf("%d", *params.To))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get multi-contract k-line data: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}
