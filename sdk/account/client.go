package account

import (
	"context"
	"fmt"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the account client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new account client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// GetAccountAsset gets the account asset information
func (c *Client) GetAccountAsset(ctx context.Context) (*openapi.ResultGetAccountAsset, error) {
	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetAccountAsset(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get account asset: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetAccountPositions gets the account positions
func (c *Client) GetAccountPositions(ctx context.Context) (*openapi.ResultListPosition, error) {
	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetAccountAsset(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get account positions: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	data := resp.GetData()
	if !data.HasPositionList() {
		return &openapi.ResultListPosition{
			Code:       resp.Code,
			Data:       []openapi.Position{},
			ErrorParam: resp.ErrorParam,
		}, nil
	}

	// Convert ResultGetAccountAsset to ResultListPosition
	result := &openapi.ResultListPosition{
		Code:       resp.Code,
		Data:       data.GetPositionList(),
		ErrorParam: resp.ErrorParam,
	}

	return result, nil
}

// GetPositionTransactionPageParams represents the parameters for GetPositionTransactionPage
type GetPositionTransactionPageParams struct {
	Size                   int32
	OffsetData             string
	FilterCoinIDList       []string
	FilterContractIDList   []string
	FilterTypeList         []string
	FilterStartCreatedTime int64
	FilterEndCreatedTime   int64
	FilterCloseOnly        *bool
	FilterOpenOnly         *bool
}

// GetPositionTransactionPage gets the position transactions with pagination
func (c *Client) GetPositionTransactionPage(ctx context.Context, params GetPositionTransactionPageParams) (*openapi.ResultPageDataPositionTransaction, error) {
	req := c.openapiClient.Class03AccountPrivateApiAPI.GetPositionTransactionPage(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Size(fmt.Sprintf("%d", params.Size))

	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	if len(params.FilterCoinIDList) > 0 {
		req = req.FilterCoinIdList(internal.JoinStrings(params.FilterCoinIDList))
	}

	if len(params.FilterContractIDList) > 0 {
		req = req.FilterContractIdList(internal.JoinStrings(params.FilterContractIDList))
	}

	if len(params.FilterTypeList) > 0 {
		req = req.FilterTypeList(internal.JoinStrings(params.FilterTypeList))
	}

	if params.FilterStartCreatedTime > 0 {
		req = req.FilterStartCreatedTimeInclusive(fmt.Sprintf("%d", params.FilterStartCreatedTime))
	}

	if params.FilterEndCreatedTime > 0 {
		req = req.FilterEndCreatedTimeExclusive(fmt.Sprintf("%d", params.FilterEndCreatedTime))
	}

	if params.FilterCloseOnly != nil {
		req = req.FilterCloseOnly(fmt.Sprintf("%v", *params.FilterCloseOnly))
	}

	if params.FilterOpenOnly != nil {
		req = req.FilterOpenOnly(fmt.Sprintf("%v", *params.FilterOpenOnly))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get position transaction page: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetCollateralTransactionPageParams represents the parameters for GetCollateralTransactionPage
type GetCollateralTransactionPageParams struct {
	Size                   int32
	OffsetData             string
	FilterCoinIDList       []string
	FilterTypeList         []string
	FilterStartCreatedTime int64
	FilterEndCreatedTime   int64
}

// GetCollateralTransactionPage gets the collateral transactions with pagination
func (c *Client) GetCollateralTransactionPage(ctx context.Context, params GetCollateralTransactionPageParams) (*openapi.ResultPageDataCollateralTransaction, error) {
	req := c.openapiClient.Class03AccountPrivateApiAPI.GetCollateralTransactionPage(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Size(fmt.Sprintf("%d", params.Size))

	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	if len(params.FilterCoinIDList) > 0 {
		req = req.FilterCoinIdList(internal.JoinStrings(params.FilterCoinIDList))
	}

	if len(params.FilterTypeList) > 0 {
		req = req.FilterTypeList(internal.JoinStrings(params.FilterTypeList))
	}

	if params.FilterStartCreatedTime > 0 {
		req = req.FilterStartCreatedTimeInclusive(fmt.Sprintf("%d", params.FilterStartCreatedTime))
	}

	if params.FilterEndCreatedTime > 0 {
		req = req.FilterEndCreatedTimeExclusive(fmt.Sprintf("%d", params.FilterEndCreatedTime))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral transaction page: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetPositionByContractID gets position information for specific contracts
func (c *Client) GetPositionByContractID(ctx context.Context, contractIDs []string) (*openapi.ResultListPosition, error) {
	if len(contractIDs) == 0 {
		return nil, fmt.Errorf("at least one contractId is required")
	}

	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetCollateralByCoinId(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		ContractIdList(internal.JoinStrings(contractIDs)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get position by contract ID: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetPositionTermPageParams represents the parameters for GetPositionTermPage
type GetPositionTermPageParams struct {
	Size                   int32
	OffsetData             string
	FilterCoinIDList       []string
	FilterContractIDList   []string
	FilterIsLongPosition   *bool
	FilterStartCreatedTime int64
	FilterEndCreatedTime   int64
}

// GetPositionTermPage gets position terms with pagination
func (c *Client) GetPositionTermPage(ctx context.Context, params GetPositionTermPageParams) (*openapi.ResultPageDataPositionTerm, error) {
	req := c.openapiClient.Class03AccountPrivateApiAPI.GetPositionTermPage(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Size(fmt.Sprintf("%d", params.Size))

	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	if len(params.FilterCoinIDList) > 0 {
		req = req.FilterCoinIdList(internal.JoinStrings(params.FilterCoinIDList))
	}

	if len(params.FilterContractIDList) > 0 {
		req = req.FilterContractIdList(internal.JoinStrings(params.FilterContractIDList))
	}

	if params.FilterIsLongPosition != nil {
		req = req.FilterIsLongPosition(fmt.Sprintf("%v", *params.FilterIsLongPosition))
	}

	if params.FilterStartCreatedTime > 0 {
		req = req.FilterStartCreatedTimeInclusive(fmt.Sprintf("%d", params.FilterStartCreatedTime))
	}

	if params.FilterEndCreatedTime > 0 {
		req = req.FilterEndCreatedTimeExclusive(fmt.Sprintf("%d", params.FilterEndCreatedTime))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get position term page: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetCollateralByCoinID gets collateral information for specific coins
func (c *Client) GetCollateralByCoinID(ctx context.Context, coinIDs []string) (*openapi.ResultListCollateral, error) {
	req := c.openapiClient.Class03AccountPrivateApiAPI.GetCollateralByCoinId1(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID()))

	if len(coinIDs) > 0 {
		req = req.CoinIdList(internal.JoinStrings(coinIDs))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral by coin ID: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetAccountByID gets account information by ID
func (c *Client) GetAccountByID(ctx context.Context) (*openapi.ResultAccount, error) {
	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetAccountById(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get account by ID: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetAccountAssetSnapshotPageParams represents the parameters for GetAccountAssetSnapshotPage
type GetAccountAssetSnapshotPageParams struct {
	Size            int32
	OffsetData      string
	CoinID          string
	FilterTimeTag   *int32
	FilterStartTime int64
	FilterEndTime   int64
}

// GetAccountAssetSnapshotPage gets account asset snapshots with pagination
func (c *Client) GetAccountAssetSnapshotPage(ctx context.Context, params GetAccountAssetSnapshotPageParams) (*openapi.ResultPageDataAccountAssetSnapshot, error) {
	if params.CoinID == "" {
		return nil, fmt.Errorf("coinId is required")
	}

	req := c.openapiClient.Class03AccountPrivateApiAPI.GetAccountAssetSnapshotPage(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Size(fmt.Sprintf("%d", params.Size)).
		CoinId(params.CoinID)

	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	if params.FilterTimeTag != nil {
		req = req.FilterTimeTag(fmt.Sprintf("%d", *params.FilterTimeTag))
	}

	if params.FilterStartTime > 0 {
		req = req.FilterStartTimeInclusive(fmt.Sprintf("%d", params.FilterStartTime))
	}

	if params.FilterEndTime > 0 {
		req = req.FilterEndTimeExclusive(fmt.Sprintf("%d", params.FilterEndTime))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get account asset snapshot page: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetPositionTransactionByID gets specific position transactions by IDs
func (c *Client) GetPositionTransactionByID(ctx context.Context, transactionIDs []string) (*openapi.ResultListPositionTransaction, error) {
	if len(transactionIDs) == 0 {
		return nil, fmt.Errorf("at least one transactionId is required")
	}

	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetPositionTransactionById(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		PositionTransactionIdList(internal.JoinStrings(transactionIDs)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get position transaction by ID: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetCollateralTransactionByID gets specific collateral transactions by IDs
func (c *Client) GetCollateralTransactionByID(ctx context.Context, transactionIDs []string) (*openapi.ResultListCollateralTransaction, error) {
	if len(transactionIDs) == 0 {
		return nil, fmt.Errorf("at least one transactionId is required")
	}

	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetCollateralTransactionById(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		CollateralTransactionIdList(internal.JoinStrings(transactionIDs)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral transaction by ID: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetAccountDeleverageLight gets account deleverage light information
func (c *Client) GetAccountDeleverageLight(ctx context.Context) (*openapi.ResultGetAccountDeleverageLight, error) {
	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.GetAccountDeleverageLight(ctx).
		AccountId(fmt.Sprintf("%d", c.GetAccountID())).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get account deleverage light: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// UpdateLeverageSetting updates the account leverage settings
func (c *Client) UpdateLeverageSetting(ctx context.Context, contractID string, leverage string) error {
	param := openapi.NewUpdateLeverageSettingParam()
	param.SetAccountId(fmt.Sprintf("%d", c.GetAccountID()))
	param.SetContractId(contractID)
	param.SetLeverage(leverage)

	resp, _, err := c.openapiClient.Class03AccountPrivateApiAPI.UpdateLeverageSetting(ctx).
		UpdateLeverageSettingParam(*param).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update leverage setting: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return nil
}
