package types

const (
	ModuleName = "liquidation"

	StoreKey = ModuleName

	RouterKey = ModuleName

    QuerierRoute = ModuleName

	MemStoreKey = "mem_liquidation"


)

func KeyPrefix(p string) []byte {
    return []byte(p)
}
