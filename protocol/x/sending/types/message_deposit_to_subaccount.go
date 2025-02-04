package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	satypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"
)

var _ sdk.Msg = &MsgDepositToSubaccount{}

// NewMsgDepositToSubaccount constructs a `MsgDepositToSubaccount` from an
// `x/bank` account sender, an `x/subaccounts` subaccount sender, an asset ID,
// and a number of quantums.
func NewMsgDepositToSubaccount(
	sender string,
	recipient satypes.SubaccountId,
	assetId uint32,
	quantums uint64,
) *MsgDepositToSubaccount {
	return &MsgDepositToSubaccount{
		Sender:    sender,
		Recipient: recipient,
		AssetId:   assetId,
		Quantums:  quantums,
	}
}

// GetSigners specifies that the sender of the message must sign.
func (msg *MsgDepositToSubaccount) GetSigners() []sdk.AccAddress {
	// Get sender account's address.
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// ValidateBasic runs validation on the fields of a MsgDepositToSubaccount.
func (msg *MsgDepositToSubaccount) ValidateBasic() error {
	// Validate account sender.
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return ErrInvalidAccountAddress
	}

	// Validate subaccount recipient.
	if err := msg.Recipient.Validate(); err != nil {
		return err
	}

	// Validate that asset is USDC.
	if msg.AssetId != lib.UsdcAssetId {
		return ErrNonUsdcAssetTransferNotImplemented
	}

	// Validate that quantums is not zero.
	if msg.Quantums == lib.ZeroUint64 {
		return ErrInvalidTransferAmount
	}

	return nil
}
