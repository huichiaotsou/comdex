package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "bandoracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bandoracle"

	// Version defines the current version the IBC module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to
	PortID = "bandoracle"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("bandoracle-port-")
)
var (
	CalldataIDKey = []byte{0x02}

	CalldataKeyPrefix = []byte{0x12}
	MarketKeyPrefix   = []byte{0x13}

	MarketForAssetKeyPrefix = []byte{0x22}
	PriceForMarketKeyPrefix = []byte{0x23}
)

func CalldataKey(id uint64) []byte {
	return append(CalldataKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func MarketKey(symbol string) []byte {
	return append(MarketKeyPrefix, []byte(symbol)...)
}

func MarketForAssetKey(id uint64) []byte {
	return append(MarketForAssetKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func PriceForMarketKey(symbol string) []byte {
	return append(PriceForMarketKeyPrefix, []byte(symbol)...)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
