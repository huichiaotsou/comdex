package types

import (
	"encoding/binary"
	"fmt"
)

type (
	// OracleScriptID is the type-safe unique identifier type for oracle scripts.
	OracleScriptID uint64

	// OracleRequestID is the type-safe unique identifier type for data requests.
	OracleRequestID int64
)

const (
	MaxMarketSymbolLength = 8
	MaxAssetNameLength = 16
)

func (m *Market) Validate() error {
	if m.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if len(m.Symbol) > MaxMarketSymbolLength {
		return fmt.Errorf("symbol length cannot be greater than %d", MaxMarketSymbolLength)
	}
	if m.ScriptID == 0 {
		return fmt.Errorf("script_id cannot be zero")
	}

	return nil
}

// int64ToBytes convert int64 to a byte slice
func int64ToBytes(num int64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(num))
	return result
}
