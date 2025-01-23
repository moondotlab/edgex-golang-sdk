package ws_test

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/ws"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
)

func TestWebSocket(t *testing.T) {
	ctx := test.GetTestContext()
	baseURL := os.Getenv("TEST_WS_BASE_URL")
	if baseURL == "" {
		t.Skip("Skipping test: TEST_WS_BASE_URL environment variable is not set")
	}

	manager := ws.NewManager(baseURL, 0, "")  // No auth needed for public WebSocket

	// Connect to public WebSocket
	err := manager.ConnectPublic(ctx)
	if err != nil {
		t.Fatalf("Failed to connect to public WebSocket: %v", err)
	}

	contractID := "10000001" // BTCUSDT

	// Create channels to track message receipt
	tickerMsgCh := make(chan struct{})
	klineMsgCh := make(chan struct{})
	depthMsgCh := make(chan struct{})
	tradesMsgCh := make(chan struct{})

	// Add a debug hook to log all messages
	manager.OnPublicMessage(func(message []byte) {
		t.Logf("Raw message received: %s", string(message))
	})

	// Test cases for different subscription types
	var tickerReceived, klineReceived, depthReceived, tradesReceived bool
	testCases := []struct {
		name     string
		subFunc  func() error
		msgCh    chan struct{}
	}{
		{
			name: "Market Ticker",
			subFunc: func() error {
				return manager.SubscribeMarketTicker(contractID, func(message []byte) {
					t.Logf("Ticker message received: %s", string(message))
					if !tickerReceived {
						close(tickerMsgCh)
						tickerReceived = true
					}
				})
			},
			msgCh: tickerMsgCh,
		},
		{
			name: "KLine",
			subFunc: func() error {
				return manager.SubscribeKLine(contractID, "DAY_1", func(message []byte) {
					t.Logf("KLine message received: %s", string(message))
					if !klineReceived {
						close(klineMsgCh)
						klineReceived = true
					}
				})
			},
			msgCh: klineMsgCh,
		},
		{
			name: "Depth",
			subFunc: func() error {
				return manager.SubscribeDepth(contractID, func(message []byte) {
					t.Logf("Depth message received: %s", string(message))
					if !depthReceived {
						close(depthMsgCh)
						depthReceived = true
					}
				})
			},
			msgCh: depthMsgCh,
		},
		{
			name: "Trades",
			subFunc: func() error {
				return manager.SubscribeTrades(contractID, func(message []byte) {
					t.Logf("Trades message received: %s", string(message))
					if !tradesReceived {
						close(tradesMsgCh)
						tradesReceived = true
					}
				})
			},
			msgCh: tradesMsgCh,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Subscribe to the channel
			err := tc.subFunc()
			if err != nil {
				t.Fatalf("Failed to subscribe to %s: %v", tc.name, err)
			}

			// Wait for message or timeout
			select {
			case <-tc.msgCh:
				t.Logf("%s message received successfully", tc.name)
			case <-time.After(5 * time.Second):
				t.Errorf("Timeout waiting for %s message", tc.name)
			}
		})
	}

	// Clean up
	manager.Close()
}

func TestPrivateWebSocket(t *testing.T) {
	// Get test credentials from environment variables
	baseURL := os.Getenv("TEST_WS_BASE_URL")
	if baseURL == "" {
		t.Skip("Skipping test: TEST_WS_BASE_URL environment variable is not set")
	}

	accountIDStr := os.Getenv("TEST_ACCOUNT_ID")
	if accountIDStr == "" {
		t.Skip("Skipping test: TEST_ACCOUNT_ID environment variable is not set")
	}

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		t.Fatalf("Invalid TEST_ACCOUNT_ID: %v", err)
	}

	starkPrivateKey := os.Getenv("TEST_STARK_PRIVATE_KEY")
	if starkPrivateKey == "" {
		t.Skip("Skipping test: TEST_STARK_PRIVATE_KEY environment variable is not set")
	}

	ctx := test.GetTestContext()
	manager := ws.NewManager(
		baseURL,
		accountID,
		starkPrivateKey,
	)

	// Connect to private WebSocket
	err = manager.ConnectPrivate(ctx)
	if err != nil {
		t.Logf("Failed to connect to private WebSocket: %v", err)
	}

	// Listen for account updates
	done := make(chan struct{})
	err = manager.OnPrivateMessage("ACCOUNT_UPDATE", func(message []byte) {
		var msg map[string]interface{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			t.Logf("Failed to unmarshal message: %v", err)
			return
		}
		t.Logf("Received account update: %v", msg)
		close(done)
	})
	if err != nil {
		t.Fatalf("Failed to register account update handler: %v", err)
	}

	// Wait for message or timeout
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Log("Timeout waiting for account update")
	}

	// Clean up
	manager.Close()
}
