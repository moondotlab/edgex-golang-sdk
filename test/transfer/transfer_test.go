package transfer

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/transfer"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetTransferOutById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := transfer.GetTransferOutByIdParams{
		TransferId: "123",
	}
	resp, err := client.GetTransferOutById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Transfer Out: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	if len(data) > 0 {
		record := data[0]
		assert.Equal(t, "123", record.GetId())
		assert.NotEmpty(t, record.GetCoinId())
		assert.NotEmpty(t, record.GetAmount())
	}
}

func TestGetTransferInById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := transfer.GetTransferInByIdParams{
		TransferId: "123",
	}
	resp, err := client.GetTransferInById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Transfer In: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	if len(data) > 0 {
		record := data[0]
		assert.Equal(t, "123", record.GetId())
		assert.NotEmpty(t, record.GetCoinId())
		assert.NotEmpty(t, record.GetAmount())
	}
}

func TestGetWithdrawAvailableAmount(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := transfer.GetWithdrawAvailableAmountParams{
		CoinId: "1000",
	}
	resp, err := client.GetWithdrawAvailableAmount(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Withdraw Available Amount: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetAvailableAmount())
}

func TestCreateTransferOut(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Test parameters
	params := transfer.CreateTransferOutParams{
		CoinId:            "1000", // Asset ID
		Amount:            "1", // 1 unit
		ReceiverAccountId: "542103805685137746",
		ReceiverL2Key:     "0x046bcf2e07c20550c49986aca69f405ae4672507fae2568640d3f1d2dcf1bfeb",
		TransferReason:    "USER_TRANSFER",
		ClientTransferId:  "test_transfer_" + time.Now().Format("20060102150405"),
		ExtraType:         "",
		ExtraDataJson:     "",
	}

	// Create transfer out - should auto-generate nonce, expiry, and signature
	resp, err := client.CreateTransferOut(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Log response for debugging
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Create Transfer Out Response: %s", string(jsonData))

	// Verify response
	assert.Equal(t, "SUCCESS", resp.GetCode())
	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetTransferOutId())

	// Get the transfer details to verify
	getParams := transfer.GetTransferOutByIdParams{
		TransferId: data.GetTransferOutId(),
	}
	getResp, err := client.GetTransferOutById(ctx, getParams)
	assert.NoError(t, err)
	assert.NotNil(t, getResp)

	// Log transfer details for debugging
	jsonData, _ = json.MarshalIndent(getResp, "", "  ")
	t.Logf("Transfer Details: %s", string(jsonData))

	// Verify transfer details
	transferData := getResp.GetData()
	assert.NotEmpty(t, transferData)
	if len(transferData) > 0 {
		transfer := transferData[0]
		assert.Equal(t, params.CoinId, transfer.GetCoinId())
		
		// Compare amounts using decimal to handle precision correctly
		expectedAmount, _ := decimal.NewFromString(params.Amount)
		actualAmount, _ := decimal.NewFromString(transfer.GetAmount())
		assert.True(t, expectedAmount.Equal(actualAmount), "Amount mismatch: expected %s, got %s", expectedAmount, actualAmount)
		
		assert.Equal(t, params.ReceiverAccountId, transfer.GetReceiverAccountId())
		assert.Equal(t, "USER_TRANSFER", transfer.GetTransferReason())
		assert.NotEmpty(t, transfer.GetL2Signature())
	}
}
