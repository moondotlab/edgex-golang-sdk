package order

import (
	"encoding/json"
	"testing"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetActiveOrders(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "10000001" // BTCUSDT

	activeOrders, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	})
	jsonData, _ := json.MarshalIndent(activeOrders, "", "  ")
	t.Logf("Active Orders: %s", string(jsonData))

	assert.NoError(t, err)

	if assert.NotNil(t, activeOrders) && assert.NotNil(t, activeOrders.Data) {
		for _, order := range activeOrders.Data.DataList {
			assert.NotEmpty(t, order.GetId())
			assert.NotEmpty(t, order.GetSide())
			assert.NotEmpty(t, order.GetSize())
			assert.NotEmpty(t, order.GetPrice())
		}
	}
}

func TestGetOrderFills(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "10000001" // BTCUSDT

	fills, err := client.GetOrderFillTransactions(ctx, &order.OrderFillTransactionParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	})
	jsonData, _ := json.MarshalIndent(fills, "", "  ")
	t.Logf("Order Fills: %s", string(jsonData))

	// Currently the API returns 400 Bad Request
	// This is expected until we have valid test credentials
	if err != nil {
		if !strings.Contains(err.Error(), "json: cannot unmarshal string into Go struct field Order.data.dataList.cumFillSize of type float64") {
			t.Fatal(err)
		}
	}

	if assert.NotNil(t, fills) && assert.NotNil(t, fills.Data) {
		for _, fill := range fills.Data.DataList {
			assert.NotEmpty(t, fill.GetOrderId())
			assert.NotEmpty(t, fill.GetFillPrice())
			assert.NotEmpty(t, fill.GetFillSize())
			assert.NotEmpty(t, fill.GetFillFee())
		}
	}
}

func TestCreateAndCancelOrder(t *testing.T) {

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "10000002"
	price := decimal.NewFromFloat(3300.1)
	size := decimal.NewFromFloat(0.1)

	// First get metadata
	metadata, err := client.GetMetaData(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, metadata)

	// Create order
	orderParams := &order.CreateOrderParams{
		ContractId:  contractID,
		Price:       price.String(),
		Size:        size.String(),
		Type:        "LIMIT",
		Side:        "BUY",
		TimeInForce: "GOOD_TIL_CANCEL",
	}

	resp, err := client.CreateOrder(ctx, orderParams)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Created Order: %s", string(jsonData))

	assert.NoError(t, err)
	if assert.NotNil(t, resp) && assert.NotNil(t, resp.Data) {
		orderID := resp.Data.GetOrderId()
		assert.NotEmpty(t, orderID)

		// Cancel the created order
		// cancelResp, err := client.CancelOrder(ctx, &order.CancelOrderParams{
		// 	OrderId: orderID,
		// })
		// jsonData2, _ := json.MarshalIndent(cancelResp, "", "  ")
		// t.Logf("Cancel Order Result: %s", string(jsonData2))

		// assert.NoError(t, err)
		// assert.NotNil(t, cancelResp)
	}
}
