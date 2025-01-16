package transfer

import (
	"context"
	"fmt"
	"strconv"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the transfer client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new transfer client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// GetTransferOutByIdParams represents the parameters for GetTransferOutById
type GetTransferOutByIdParams struct {
	TransferId string
}

// GetTransferOutById gets a transfer out record by ID
func (c *Client) GetTransferOutById(ctx context.Context, params GetTransferOutByIdParams) (*openapi.ResultListTransferOut, error) {
	req := c.openapiClient.Class07TransferPrivateApiAPI.GetTransferOutById(ctx)

	if params.TransferId != "" {
		req = req.TransferOutIdList(params.TransferId)
	}

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer out by id: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetTransferInByIdParams represents the parameters for GetTransferInById
type GetTransferInByIdParams struct {
	TransferId string
}

// GetTransferInById gets a transfer in record by ID
func (c *Client) GetTransferInById(ctx context.Context, params GetTransferInByIdParams) (*openapi.ResultListTransferIn, error) {
	req := c.openapiClient.Class07TransferPrivateApiAPI.GetTransferInById(ctx)

	if params.TransferId != "" {
		req = req.TransferInIdList(params.TransferId)
	}

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer in by id: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// GetWithdrawAvailableAmountParams represents the parameters for GetWithdrawAvailableAmount
type GetWithdrawAvailableAmountParams struct {
	CoinId string
}

// GetWithdrawAvailableAmount gets the available withdrawal amount
func (c *Client) GetWithdrawAvailableAmount(ctx context.Context, params GetWithdrawAvailableAmountParams) (*openapi.ResultGetTransferOutAvailableAmount, error) {
	req := c.openapiClient.Class07TransferPrivateApiAPI.GetWithdrawAvailableAmount1(ctx)

	if params.CoinId != "" {
		req = req.CoinId(params.CoinId)
	}

	// Set account ID
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get available withdrawal amount: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// CreateTransferOutParams represents the parameters for CreateTransferOut
type CreateTransferOutParams struct {
	CoinId            string
	Amount            string
	ReceiverAccountId string
	ReceiverL2Key     string
	ClientTransferId  string
	TransferReason    string
	L2Nonce           string
	L2ExpireTime      string
	L2Signature       string
	ExtraType         string
	ExtraDataJson     string
}

// CreateTransferOut creates a new transfer out order
func (c *Client) CreateTransferOut(ctx context.Context, params CreateTransferOutParams) (*openapi.ResultCreateTransferOut, error) {
	createTransferOutParam := openapi.CreateTransferOutParam{}

	// Set account ID
	createTransferOutParam.SetAccountId(strconv.FormatInt(c.GetAccountID(), 10))

	createTransferOutParam.SetCoinId(params.CoinId)
	createTransferOutParam.SetAmount(params.Amount)
	createTransferOutParam.SetReceiverAccountId(params.ReceiverAccountId)
	createTransferOutParam.SetReceiverL2Key(params.ReceiverL2Key)
	createTransferOutParam.SetClientTransferId(params.ClientTransferId)
	createTransferOutParam.SetTransferReason(params.TransferReason)
	createTransferOutParam.SetL2Nonce(params.L2Nonce)
	createTransferOutParam.SetL2ExpireTime(params.L2ExpireTime)
	createTransferOutParam.SetL2Signature(params.L2Signature)
	createTransferOutParam.SetExtraType(params.ExtraType)
	createTransferOutParam.SetExtraDataJson(params.ExtraDataJson)

	req := c.openapiClient.Class07TransferPrivateApiAPI.CreateTransferOut(ctx).
		CreateTransferOutParam(createTransferOutParam)

	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create transfer out: %w", err)
	}

	if resp.GetCode() != "SUCCESS" {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}
