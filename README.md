# EdgeX Golang SDK

A Go SDK for interacting with the EdgeX Exchange API.

## Installation

```bash
go get github.com/edgex-Tech/edgex-golang-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/edgex-Tech/edgex-golang-sdk/sdk"
)

func main() {
    // Create a new client
    client, err := sdk.NewClient(
        sdk.WithBaseURL("https://testnet.edgex.exchange"),
        sdk.WithAccountID(12345),
        sdk.WithStarkPrivateKey("your-stark-private-key"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Create context
    ctx := context.Background()

    // Get account assets
    assets, err := client.Asset.GetAccountAsset(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Print asset information
    fmt.Printf("Account Assets: %+v\n", assets)
}
```

## Available APIs

The SDK currently supports the following API modules:

- **Asset API**: Manage assets, withdrawals, and get coin rates
- **Transfer API**: Handle transfers between accounts

For detailed examples of each API endpoint, please refer to the test files:
- Asset API examples: [asset_test.go](test/asset/asset_test.go)
- Transfer API examples: [transfer_test.go](test/transfer/transfer_test.go)

## Environment Variables

For running tests, the following environment variables are required:

```bash
export TEST_BASE_URL=https://testnet.edgex.exchange
export TEST_ACCOUNT_ID=your_account_id
export TEST_STARK_PRIVATE_KEY=your_stark_private_key
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Last Updated

2025-01-16
