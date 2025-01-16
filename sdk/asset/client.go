package asset

import (
	"context"
	"fmt"
	"strconv"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the asset client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new asset client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// GetAllOrdersPageParams represents the parameters for GetAllOrdersPage
type GetAllOrdersPageParams struct {
	StartTime  string
	EndTime    string
	ChainId    string
	TypeList   string
	Size       string
	OffsetData string
}

// GetAllOrdersPage gets all asset orders with pagination
func (c *Client) GetAllOrdersPage(ctx context.Context, params GetAllOrdersPageParams) (*openapi.ResultPageDataAssetOrder, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetAllOrdersPage(ctx)

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	// Set optional parameters
	if params.StartTime != "" {
		req = req.StartTime(params.StartTime)
	}
	if params.EndTime != "" {
		req = req.EndTime(params.EndTime)
	}
	if params.ChainId != "" {
		req = req.ChainId(params.ChainId)
	}
	if params.TypeList != "" {
		req = req.TypeList(params.TypeList)
	}
	if params.Size != "" {
		req = req.Size(params.Size)
	}
	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get asset orders: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetCoinRateParams represents the parameters for GetCoinRate
type GetCoinRateParams struct {
	ChainId string
	Coin    string
}

// GetCoinRate gets the coin rate
func (c *Client) GetCoinRate(ctx context.Context, params GetCoinRateParams) (*openapi.ResultGetCoinRate, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetCoinRate(ctx)

	if params.ChainId != "" {
		req = req.ChainId(params.ChainId)
	}
	if params.Coin != "" {
		req = req.Coin(params.Coin)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get coin rate: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetCrossWithdrawByIdParams represents the parameters for GetCrossWithdrawById
type GetCrossWithdrawByIdParams struct {
	CrossWithdrawIdList string
}

// GetCrossWithdrawById gets cross withdraw records by ID
func (c *Client) GetCrossWithdrawById(ctx context.Context, params GetCrossWithdrawByIdParams) (*openapi.ResultListCrossWithdraw, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetCrossWithdrawById(ctx)

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	if params.CrossWithdrawIdList != "" {
		req = req.CrossWithdrawIdList(params.CrossWithdrawIdList)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get cross withdraw by id: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetCrossWithdrawSignInfoParams represents the parameters for GetCrossWithdrawSignInfo
type GetCrossWithdrawSignInfoParams struct {
	ChainId string
	Amount  string
}

// GetCrossWithdrawSignInfo gets cross withdraw sign info
func (c *Client) GetCrossWithdrawSignInfo(ctx context.Context, params GetCrossWithdrawSignInfoParams) (*openapi.ResultGetCrossWithdrawSignInfo, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetCrossWithdrawSignInfo(ctx)

	if params.ChainId != "" {
		req = req.ChainId(params.ChainId)
	}
	if params.Amount != "" {
		req = req.Amount(params.Amount)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get cross withdraw sign info: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetFastWithdrawByIdParams represents the parameters for GetFastWithdrawById
type GetFastWithdrawByIdParams struct {
	FastWithdrawIdList string
}

// GetFastWithdrawById gets fast withdraw records by ID
func (c *Client) GetFastWithdrawById(ctx context.Context, params GetFastWithdrawByIdParams) (*openapi.ResultListFastWithdraw, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetFastWithdrawById(ctx)

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	if params.FastWithdrawIdList != "" {
		req = req.FastWithdrawIdList(params.FastWithdrawIdList)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get fast withdraw by id: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetFastWithdrawSignInfoParams represents the parameters for GetFastWithdrawSignInfo
type GetFastWithdrawSignInfoParams struct {
	ChainId string
	Amount  string
}

// GetFastWithdrawSignInfo gets fast withdraw sign info
func (c *Client) GetFastWithdrawSignInfo(ctx context.Context, params GetFastWithdrawSignInfoParams) (*openapi.ResultGetFastWithdrawSignInfo, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetFastWithdrawSignInfo(ctx)

	if params.ChainId != "" {
		req = req.ChainId(params.ChainId)
	}
	if params.Amount != "" {
		req = req.Amount(params.Amount)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get fast withdraw sign info: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetNormalWithdrawByIdParams represents the parameters for GetNormalWithdrawById
type GetNormalWithdrawByIdParams struct {
	NormalWithdrawIdList string
}

// GetNormalWithdrawById gets normal withdraw records by ID
func (c *Client) GetNormalWithdrawById(ctx context.Context, params GetNormalWithdrawByIdParams) (*openapi.ResultListNormalWithdraw, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetNormalWithdrawById(ctx)

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	if params.NormalWithdrawIdList != "" {
		req = req.NormalWithdrawIdList(params.NormalWithdrawIdList)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get normal withdraw by id: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetNormalWithdrawableAmountParams represents the parameters for GetNormalWithdrawableAmount
type GetNormalWithdrawableAmountParams struct {
	Address string
}

// GetNormalWithdrawableAmount gets normal withdrawable amount
func (c *Client) GetNormalWithdrawableAmount(ctx context.Context, params GetNormalWithdrawableAmountParams) (*openapi.ResultGetNormalWithdrawableAmount, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.GetNormalWithdrawableAmount(ctx)

	if params.Address != "" {
		req = req.Address(params.Address)
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get normal withdrawable amount: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// CreateNormalWithdrawParams represents the parameters for CreateNormalWithdraw
type CreateNormalWithdrawParams struct {
	CoinId           string
	Amount           string
	EthAddress       string
	ClientWithdrawId string
	ExpireTime       string
	L2Signature      string
}

// CreateNormalWithdraw creates a normal withdrawal order
func (c *Client) CreateNormalWithdraw(ctx context.Context, params CreateNormalWithdrawParams) (*openapi.ResultCreateNormalWithdraw, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.CreateNormalWithdraw(ctx)

	// Convert account ID to string
	accountId := strconv.FormatInt(c.GetAccountID(), 10)

	// Create request body
	body := openapi.CreateNormalWithdrawParam{
		AccountId:        &accountId,
		CoinId:           &params.CoinId,
		Amount:           &params.Amount,
		EthAddress:       &params.EthAddress,
		ClientWithdrawId: &params.ClientWithdrawId,
		ExpireTime:       &params.ExpireTime,
		L2Signature:      &params.L2Signature,
	}

	resp, _, err := req.CreateNormalWithdrawParam(body).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create normal withdraw: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// CreateCrossWithdrawParams represents the parameters for CreateCrossWithdraw
type CreateCrossWithdrawParams struct {
	CoinId                string
	Amount                string
	EthAddress            string
	Erc20Address          string
	LpAccountId           string
	ClientCrossWithdrawId string
	ExpireTime            string
	L2Signature           string
	Fee                   string
	ChainId               string
	MpcAddress            string
	MpcSignature          string
	MpcSignTime           string
}

// CreateCrossWithdraw creates a cross-chain withdrawal order
func (c *Client) CreateCrossWithdraw(ctx context.Context, params CreateCrossWithdrawParams) (*openapi.ResultCreateCrossWithdraw, error) {
	req := c.openapiClient.Class09AssetsPrivateApiAPI.CreateCrossWithdraw(ctx)

	// Convert account ID to string
	accountId := strconv.FormatInt(c.GetAccountID(), 10)

	// Create request body
	body := openapi.CreateCrossWithdrawParam{
		AccountId:             &accountId,
		CoinId:                &params.CoinId,
		Amount:                &params.Amount,
		EthAddress:            &params.EthAddress,
		Erc20Address:          &params.Erc20Address,
		LpAccountId:           &params.LpAccountId,
		ClientCrossWithdrawId: &params.ClientCrossWithdrawId,
		ExpireTime:            &params.ExpireTime,
		L2Signature:           &params.L2Signature,
		Fee:                   &params.Fee,
		ChainId:               &params.ChainId,
		MpcAddress:            &params.MpcAddress,
		MpcSignature:          &params.MpcSignature,
		MpcSignTime:           &params.MpcSignTime,
	}

	resp, _, err := req.CreateCrossWithdrawParam(body).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create cross withdraw: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}
