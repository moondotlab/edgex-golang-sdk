package internal

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"time"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"

	"github.com/edgex-Tech/edgex-golang-sdk/starkcurve"
)

// Client represents the base client with common functionality
type Client struct {
	httpClient    *http.Client
	baseURL       string
	accountID     int64
	starkPriKey   string
	openapiClient *openapi.APIClient
}

// ClientConfig holds the configuration for creating a new Client
type ClientConfig struct {
	BaseURL     string
	AccountID   int64
	StarkPriKey string
}

// NewClient creates a new base client
func NewClient(cfg *ClientConfig) (*Client, error) {
	openapiConfig := openapi.NewConfiguration()
	openapiConfig.Servers = []openapi.ServerConfiguration{
		{
			URL: cfg.BaseURL,
		},
	}
	openapiClient := openapi.NewAPIClient(openapiConfig)

	return &Client{
		httpClient:    &http.Client{Timeout: 30 * time.Second},
		baseURL:       cfg.BaseURL,
		accountID:     cfg.AccountID,
		starkPriKey:   cfg.StarkPriKey,
		openapiClient: openapiClient,
	}, nil
}

// GetAccountID returns the account ID
func (c *Client) GetAccountID() int64 {
	return c.accountID
}

// GetStarkPriKey returns the stark private key
func (c *Client) GetStarkPriKey() string {
	return c.starkPriKey
}

// Sign signs a message hash using the client's Stark private key
func (c *Client) Sign(messageHash []byte) (*L2Signature, error) {
	privateKey := c.GetStarkPriKey()
	if privateKey == "" {
		return nil, fmt.Errorf("stark private key not set")
	}

	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	starkPrivKey := big.NewInt(0).SetBytes(privKeyBytes)

	msgHashInt := big.NewInt(0).SetBytes(messageHash)
	msgHashInt = msgHashInt.Mod(msgHashInt, starkcurve.NewStarkCurve().N)

	r, s, err := starkcurve.Sign(starkPrivKey.Bytes(), msgHashInt.Bytes())
	if err != nil {
		return nil, err
	}

	rBytes := append(bytes.Repeat([]byte{0}, 32-len(r.Bytes())), r.Bytes()...)
	sBytes := append(bytes.Repeat([]byte{0}, 32-len(s.Bytes())), s.Bytes()...)

	// Convert r, s and y to hex strings
	signature := &L2Signature{
		R: hex.EncodeToString(rBytes),
		S: hex.EncodeToString(sBytes),
		V: "",
	}

	return signature, nil
}
