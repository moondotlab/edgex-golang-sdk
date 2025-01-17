package transfer

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"math/big"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/shopspring/decimal"
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
func (c *Client) CreateTransferOut(ctx context.Context, params CreateTransferOutParams, metadata openapi.MetaData) (*openapi.ResultCreateTransferOut, error) {
	createTransferOutParam := openapi.CreateTransferOutParam{}

	// Set account ID
	accountId := strconv.FormatInt(c.GetAccountID(), 10)
	createTransferOutParam.SetAccountId(accountId)

	// Generate client transfer ID if not provided
	if params.ClientTransferId == "" {
		params.ClientTransferId = internal.GenerateUUID()
	}

	// Set expiration time if not provided (default to 1 hour from now)
	if params.L2ExpireTime == "" {
		expireTime := time.Now().Add(14 * 24 * time.Hour).UnixMilli()
		params.L2ExpireTime = strconv.FormatInt(expireTime, 10)
	}

	// Set nonce if not provided
	if params.L2Nonce == "" {
		nonce := internal.CalcNonce(params.ClientTransferId)
		params.L2Nonce = strconv.FormatInt(nonce, 10)
	}

	// Convert parameters to appropriate types for hash calculation
	amountDm, _ := decimal.NewFromString(params.Amount)
	amount := amountDm.Shift(6).IntPart()
	nonce, _ := strconv.ParseInt(params.L2Nonce, 10, 64)
	expireTime, _ := strconv.ParseInt(params.L2ExpireTime, 10, 64)
	expireTimeUnix := expireTime / (60 * 60 * 1000) // Convert to hours

	// Remove 0x prefix from receiver L2 key if present
	receiverL2Key := strings.TrimPrefix(params.ReceiverL2Key, "0x")

	// Get asset IDs from metadata
	global := metadata.GetGlobal()
	collateralCoin := global.GetStarkExCollateralCoin()
	assetIDStr := collateralCoin.GetStarkExAssetId()
	assetID, ok := new(big.Int).SetString(assetIDStr, 0)
	if !ok {
		return nil, fmt.Errorf("invalid asset ID format: %s", assetIDStr)
	}

	// Convert receiver L2 key to big.Int
	receiverPublicKey, ok := new(big.Int).SetString(receiverL2Key, 16)
	if !ok {
		return nil, fmt.Errorf("invalid receiver L2 key format: %s", receiverL2Key)
	}

	// Get position IDs (same as account IDs)
	senderPositionId := c.GetAccountID()
	receiverPositionId, err := strconv.ParseInt(params.ReceiverAccountId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver account ID: %w", err)
	}
	feePositionId := senderPositionId // Fee position is same as sender for now
	maxAmountFee := int64(0)

	// Calculate transfer hash and sign it
	msgHash := internal.CalcTransferHash(
		assetID,
		big.NewInt(0),
		receiverPublicKey,
		senderPositionId,
		receiverPositionId,
		feePositionId,
		nonce,
		amount,
		maxAmountFee,
		expireTimeUnix,
	)
	signature, err := c.Sign(msgHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transfer hash: %w", err)
	}

	// Set all parameters
	createTransferOutParam.SetCoinId(params.CoinId)
	createTransferOutParam.SetAmount(amountDm.String())
	createTransferOutParam.SetReceiverAccountId(params.ReceiverAccountId)
	createTransferOutParam.SetReceiverL2Key(params.ReceiverL2Key)
	createTransferOutParam.SetClientTransferId(params.ClientTransferId)
	createTransferOutParam.SetTransferReason(params.TransferReason)
	createTransferOutParam.SetL2Nonce(params.L2Nonce)
	createTransferOutParam.SetL2ExpireTime(params.L2ExpireTime)
	createTransferOutParam.SetL2Signature(fmt.Sprintf("%s%s%s", signature.R, signature.S, signature.V))
	createTransferOutParam.SetExtraType(params.ExtraType)
	createTransferOutParam.SetExtraDataJson(params.ExtraDataJson)

	// Execute the request
	req := c.openapiClient.Class07TransferPrivateApiAPI.CreateTransferOut(ctx)
	req = req.CreateTransferOutParam(createTransferOutParam)

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
