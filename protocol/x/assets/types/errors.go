package types

// DONTCOVER

import errorsmod "cosmossdk.io/errors"

// x/assets module sentinel errors
var (
	ErrAssetDoesNotExist            = errorsmod.Register(ModuleName, 1, "Asset does not exist")
	ErrNegativeLongInterest         = errorsmod.Register(ModuleName, 2, "LongInterest cannot be negative")
	ErrNoAssetWithDenom             = errorsmod.Register(ModuleName, 3, "No asset found associated with given denom")
	ErrAssetDenomAlreadyExists      = errorsmod.Register(ModuleName, 4, "Existing asset found with the same denom")
	ErrAssetIdAlreadyExists         = errorsmod.Register(ModuleName, 5, "Existing asset found with the same asset id")
	ErrGapFoundInAssetId            = errorsmod.Register(ModuleName, 6, "Found gap in asset Id")
	ErrAssetZeroNotUsdc             = errorsmod.Register(ModuleName, 7, "First asset is not USDC")
	ErrNoAssetInGenesis             = errorsmod.Register(ModuleName, 8, "No asset found in genesis state")
	ErrInvalidMarketId              = errorsmod.Register(ModuleName, 9, "Found market id for asset without market")
	ErrInvalidAssetAtomicResolution = errorsmod.Register(ModuleName, 10, "Invalid asset atomic resolution")
	ErrInvalidDenomExponent         = errorsmod.Register(ModuleName, 11, "Invalid denom exponent")

	// Errors for Not Implemented
	ErrNotImplementedMulticollateral = errorsmod.Register(ModuleName, 401, "Not Implemented: Multi-Collateral")
	ErrNotImplementedMargin          = errorsmod.Register(ModuleName, 402, "Not Implemented: Margin-Trading of Assets")
)
