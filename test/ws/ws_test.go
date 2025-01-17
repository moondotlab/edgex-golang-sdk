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
		t.Logf("Failed to connect to public WebSocket: %v", err)
	}

	// Subscribe to market ticker
	done := make(chan struct{})
	err = manager.SubscribeMarketTicker("BTC-USDT", func(message []byte) {
		var msg map[string]interface{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			t.Logf("Failed to unmarshal message: %v", err)
			return
		}
		t.Logf("Received ticker message: %v", msg)
		close(done)
	})
	if err != nil {
		t.Fatalf("Failed to subscribe to market ticker: %v", err)
	}

	// Wait for message or timeout
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Log("Timeout waiting for ticker message")
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
