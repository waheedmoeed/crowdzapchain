package types

const (
	// ModuleName is the name of the module
	ModuleName = "smartcontracts"

	// StoreKey to be used when creating the KVStore
	StoreKey  = ModuleName
	StoreKeyB = ModuleName + "2"

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)
