package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/testutil/constants"
	"github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateMarketParam_GetSigners(t *testing.T) {
	msg := types.MsgUpdateMarketParam{
		Authority: constants.DaveAccAddress.String(),
	}
	require.Equal(t, []sdk.AccAddress{constants.DaveAccAddress}, msg.GetSigners())
}

func TestMsgUpdateMarketParam_ValidateBasic(t *testing.T) {
	tests := map[string]struct {
		msg         types.MsgUpdateMarketParam
		expectedErr string
	}{
		"Success": {
			msg: types.MsgUpdateMarketParam{
				Authority: "test",
				MarketParam: types.MarketParam{
					Pair:               "test",
					MinExchanges:       1,
					MinPriceChangePpm:  1_000,
					ExchangeConfigJson: "{}",
				},
			},
		},
		"Failure: Empty authority": {
			msg:         types.MsgUpdateMarketParam{},
			expectedErr: "authority cannot be empty",
		},
		"Failure: Empty pair": {
			msg: types.MsgUpdateMarketParam{
				Authority: "test",
				MarketParam: types.MarketParam{
					Pair:              "",
					MinExchanges:      1,
					MinPriceChangePpm: 1_000,
				},
			},
			expectedErr: "Pair cannot be empty",
		},
		"Failure: 0 MinExchanges": {
			msg: types.MsgUpdateMarketParam{
				Authority: "test",
				MarketParam: types.MarketParam{
					Pair:              "test",
					MinExchanges:      0,
					MinPriceChangePpm: 1_000,
				},
			},
			expectedErr: "Min exchanges must be greater than zero",
		},
		"Failure: 0 MinPriceChangePpm": {
			msg: types.MsgUpdateMarketParam{
				Authority: "test",
				MarketParam: types.MarketParam{
					Pair:              "test",
					MinExchanges:      2,
					MinPriceChangePpm: 0,
				},
			},
			expectedErr: "Min price change in parts-per-million must be greater than 0",
		},
		"Failure: 10_000 MinPriceChangePpm": {
			msg: types.MsgUpdateMarketParam{
				Authority: "test",
				MarketParam: types.MarketParam{
					Pair:              "test",
					MinExchanges:      2,
					MinPriceChangePpm: 10_000,
				},
			},
			expectedErr: "Min price change in parts-per-million must be greater than 0 and less than 10000",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.expectedErr)
			}
		})
	}
}
