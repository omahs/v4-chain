package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/testutil/constants"
	types "github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateOracleMarket_GetSigners(t *testing.T) {
	msg := types.MsgCreateOracleMarket{
		Authority: constants.AliceAccAddress.String(),
	}
	require.Equal(t, []sdk.AccAddress{constants.AliceAccAddress}, msg.GetSigners())
}

func TestMsgCreateOracleMarket_ValidateBasic(t *testing.T) {
	validExchangeConfigJson := `{"exchanges":[{"exchangeName":"Binance","ticker":"BTCUSDT"}]}`
	tests := []struct {
		desc        string
		msg         types.MsgCreateOracleMarket
		expectedErr string
	}{
		{
			desc:        "Empty authority",
			msg:         types.MsgCreateOracleMarket{},
			expectedErr: "authority cannot be empty",
		},
		{
			desc: "Valid MsgCreateOracleMarket",
			msg: types.MsgCreateOracleMarket{
				Authority: "test",
				Params: types.MarketParam{
					Pair:               "BTC-USD",
					MinExchanges:       1,
					MinPriceChangePpm:  1_000,
					ExchangeConfigJson: validExchangeConfigJson,
				},
			},
			expectedErr: "",
		},
		{
			desc: "Empty pair",
			msg: types.MsgCreateOracleMarket{
				Authority: "test",
				Params: types.MarketParam{
					Pair:               "",
					MinExchanges:       1,
					MinPriceChangePpm:  1_000,
					ExchangeConfigJson: validExchangeConfigJson,
				},
			},
			expectedErr: "Pair cannot be empty",
		},
		{
			desc: "Invalid MinPriceChangePpm",
			msg: types.MsgCreateOracleMarket{
				Authority: "test",
				Params: types.MarketParam{
					Pair:               "BTC-USD",
					MinExchanges:       1,
					MinPriceChangePpm:  0,
					ExchangeConfigJson: validExchangeConfigJson,
				},
			},
			expectedErr: "Min price change in parts-per-million must be greater than 0",
		},
		{
			desc: "Empty ExchangeConfigJson",
			msg: types.MsgCreateOracleMarket{
				Authority: "test",
				Params: types.MarketParam{
					Pair:               "BTC-USD",
					MinExchanges:       1,
					MinPriceChangePpm:  1_000,
					ExchangeConfigJson: "",
				},
			},
			expectedErr: "ExchangeConfigJson string is not valid",
		},
		{
			desc: "Typo in ExchangeConfigJson",
			msg: types.MsgCreateOracleMarket{
				Authority: "test",
				Params: types.MarketParam{
					Pair:               "BTC-USD",
					MinExchanges:       1,
					MinPriceChangePpm:  1_000,
					ExchangeConfigJson: `{"exchanges":[]`, // missing a bracket
				},
			},
			expectedErr: "ExchangeConfigJson string is not valid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.expectedErr)
			}
		})
	}
}
