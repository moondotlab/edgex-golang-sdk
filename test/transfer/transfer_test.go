package transfer

import (
	"encoding/json"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/transfer"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
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

	params := transfer.CreateTransferOutParams{
		CoinId:            "1000",
		Amount:            "1.000000",
		ReceiverAccountId: "551109015904453258",
		ReceiverL2Key:     "0x03eec711e360695bb44b1170057a25340303c1f16893a8def7450e44294405a8",
		ClientTransferId:  "3877531064364166",
		TransferReason:    "USER_TRANSFER",
		L2Nonce:           "2280110103",
		L2ExpireTime:      "1735873200000",
		L2Signature:       "0141279ec45ce1ea37b11cfa4683cfab8443bcbf8da3f066cef3e437862573f9034efe12eee1be3fc715c7b511f69e3ba32ec67a9ac89538fbb73de46fefc5e5",
		ExtraType:         "",
		ExtraDataJson:     "",
	}
	resp, err := client.CreateTransferOut(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Create Transfer Out: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetTransferOutId())
}
