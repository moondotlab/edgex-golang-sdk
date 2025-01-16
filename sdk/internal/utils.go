package internal

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/starkcurve"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

// CalcNonce calculates a nonce from source string
func CalcNonce(src string) int64 {
	h := sha256.New()
	h.Write([]byte(src))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	result, _ := big.NewInt(0).SetString(string(hash[:8]), 16)
	return result.Int64()
}

// CalcLimitOrderHash calculates the hash for a limit order
func CalcLimitOrderHash(assetIdSynthetic, assetIdCollateral, assetIdFee string, isBuyingSynthetic bool, amountSynthetic, amountCollateral, amountFee, nonce, positionID, expirationTimestamp int64) []byte {
	// Remove assetIdSynthetic, assetIdCollateral, assetIdFee 0x prefix if exists
	if len(assetIdSynthetic) > 2 && assetIdSynthetic[:2] == "0x" {
		assetIdSynthetic = assetIdSynthetic[2:]
	}
	if len(assetIdCollateral) > 2 && assetIdCollateral[:2] == "0x" {
		assetIdCollateral = assetIdCollateral[2:]
	}
	if len(assetIdFee) > 2 && assetIdFee[:2] == "0x" {
		assetIdFee = assetIdFee[2:]
	}
	var asset_id_sell *big.Int
	var asset_id_buy *big.Int
	var amount_sell, amount_buy *big.Int
	if isBuyingSynthetic {
		asset_id_sell, _ = big.NewInt(0).SetString(assetIdCollateral, 16)
		asset_id_buy, _ = big.NewInt(0).SetString(assetIdSynthetic, 16)
		amount_sell = big.NewInt(amountCollateral)
		amount_buy = big.NewInt(amountSynthetic)
	} else {
		asset_id_sell, _ = big.NewInt(0).SetString(assetIdSynthetic, 16)
		asset_id_buy, _ = big.NewInt(0).SetString(assetIdCollateral, 16)
		amount_sell = big.NewInt(amountSynthetic)
		amount_buy = big.NewInt(amountCollateral)
	}
	asset_id_fee, _ := big.NewInt(0).SetString(assetIdFee, 16)
	msg := starkcurve.CalcHash([]*big.Int{asset_id_sell, asset_id_buy})
	msgInt := big.NewInt(0).SetBytes(msg)
	msg = starkcurve.CalcHash([]*big.Int{msgInt, asset_id_fee})

	// packed_message0 = amount_sell
	// packed_message0 = packed_message0 * 2**64 + amount_buy
	// packed_message0 = packed_message0 * 2**64 + max_amount_fee
	// packed_message0 = packed_message0 * 2**32 + nonce
	packed_message0 := big.NewInt(0).Set(amount_sell)
	packed_message0 = packed_message0.Lsh(packed_message0, 64)
	packed_message0 = packed_message0.Add(packed_message0, amount_buy)
	max_amount_fee := big.NewInt(amountFee)
	packed_message0 = packed_message0.Lsh(packed_message0, 64)
	packed_message0 = packed_message0.Add(packed_message0, max_amount_fee)
	nonceInt := big.NewInt(nonce)
	packed_message0 = packed_message0.Lsh(packed_message0, 32)
	packed_message0 = packed_message0.Add(packed_message0, nonceInt)
	msgInt = big.NewInt(0).SetBytes(msg)
	msg = starkcurve.CalcHash([]*big.Int{msgInt, packed_message0})

	// packed_message1 = LIMIT_ORDER_WITH_FEES
	// packed_message1 = packed_message1 * 2**64 + position_id
	// packed_message1 = packed_message1 * 2**64 + position_id
	// packed_message1 = packed_message1 * 2**64 + position_id
	// packed_message1 = packed_message1 * 2**32 + expiration_timestamp
	// packed_message1 = packed_message1 * 2**17  # Padding.
	packed_message1 := big.NewInt(LimitOrderWithFeeType)
	packed_message1 = packed_message1.Lsh(packed_message1, 64)
	positionIDInt := big.NewInt(positionID)
	packed_message1 = packed_message1.Add(packed_message1, positionIDInt)
	packed_message1 = packed_message1.Lsh(packed_message1, 64)
	packed_message1 = packed_message1.Add(packed_message1, positionIDInt)
	packed_message1 = packed_message1.Lsh(packed_message1, 64)
	packed_message1 = packed_message1.Add(packed_message1, positionIDInt)
	expirationTimestampInt := big.NewInt(expirationTimestamp)
	packed_message1 = packed_message1.Lsh(packed_message1, 32)
	packed_message1 = packed_message1.Add(packed_message1, expirationTimestampInt)
	packed_message1 = packed_message1.Lsh(packed_message1, 17)
	msgInt = big.NewInt(0).SetBytes(msg)
	msg = starkcurve.CalcHash([]*big.Int{msgInt, packed_message1})

	return msg
}

// JoinStrings joins a slice of strings with commas
func JoinStrings(strs []string) string {
	return strings.Join(strs, ",")
}

// GetValue converts a JSON value to a string representation
func GetValue(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case []interface{}:
		if len(v) == 0 {
			return ""
		}
		values := make([]string, len(v))
		for i, item := range v {
			values[i] = GetValue(item)
		}
		return strings.Join(values, "&")
	case map[string]interface{}:
		sortedMap := make(map[string]string)
		for key, val := range v {
			sortedMap[key] = GetValue(val)
		}

		// Get sorted keys
		keys := make([]string, 0, len(sortedMap))
		for k := range sortedMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// Build key=value pairs
		pairs := make([]string, len(keys))
		for i, key := range keys {
			pairs[i] = key + "=" + sortedMap[key]
		}
		return strings.Join(pairs, "&")
	default:
		// Handle other primitive types
		return fmt.Sprint(v)
	}
}
