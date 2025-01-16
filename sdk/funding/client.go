package funding

import (
	"context"
	"fmt"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the funding client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new funding client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// GetFundingRateParams represents the parameters for GetFundingRate
type GetFundingRateParams struct {
	ContractID string
	From       *int64
	To         *int64
	Size       *int32
	Offset     *string
}

// GetFundingRate gets the funding rate for a contract
func (c *Client) GetFundingRate(ctx context.Context, params GetFundingRateParams) (*openapi.ResultPageDataFundingRate, error) {
	req := c.openapiClient.Class01FundingPublicApiAPI.GetFundingRatePage(ctx).
		ContractId(params.ContractID).
		FilterSettlementFundingRate("true") // Only get settlement funding rates

	if params.Size != nil {
		req = req.Size(fmt.Sprintf("%d", *params.Size))
	}
	if params.Offset != nil {
		req = req.OffsetData(*params.Offset)
	}
	if params.From != nil {
		req = req.FilterBeginTimeInclusive(fmt.Sprintf("%d", *params.From))
	}
	if params.To != nil {
		req = req.FilterEndTimeExclusive(fmt.Sprintf("%d", *params.To))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get funding rate: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetLatestFundingRateParams represents the parameters for GetLatestFundingRate
type GetLatestFundingRateParams struct {
	ContractID string
}

// GetLatestFundingRate gets the latest funding rate for a contract
func (c *Client) GetLatestFundingRate(ctx context.Context, params GetLatestFundingRateParams) (*openapi.ResultListFundingRate, error) {
	req := c.openapiClient.Class01FundingPublicApiAPI.GetLatestFundingRate(ctx).
		ContractId(params.ContractID)

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest funding rate: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}
