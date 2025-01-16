package metadata

import (
	"context"
	"fmt"
	"strconv"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the metadata client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new metadata client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// GetServerTime gets the current server time
func (c *Client) GetServerTime(ctx context.Context) (*openapi.ResultGetServerTime, error) {
	resp, _, err := c.openapiClient.Class00MetaDataPublicApiAPI.GetServerTime(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get server time: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetMetaData gets the exchange metadata
func (c *Client) GetMetaData(ctx context.Context) (*openapi.ResultMetaData, error) {
	resp, _, err := c.openapiClient.Class00MetaDataPublicApiAPI.GetMetaData(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

func mustParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}
