package lib

import (
	"math/big"
)

// BaseToQuoteQuantums converts an amount denoted in base quantums, to an equivalent amount denoted in quote
// quantums. To determine the equivalent amount, an oracle price is used.
//
//   - `priceValue * 10^priceExponent` represents the conversion rate from one full coin of base currency
//     to one full coin of quote currency.
//   - `10^baseCurrencyAtomicResolution` represents the amount of one full coin that a base quantum is equal to.
//   - `10^quoteCurrencyAtomicResolution` represents the amount of one full coin that a quote quantum is equal to.
//
// To convert from base to quote quantums, we use the following equation:
//
//	quoteQuantums =
//	  (baseQuantums * 10^baseCurrencyAtomicResolution) *
//	  (priceValue * 10^priceExponent) /
//	  (10^quoteCurrencyAtomicResolution)
//	=
//	  baseQuantums * priceValue *
//	  10^(priceExponent + baseCurrencyAtomicResolution - quoteCurrencyAtomicResolution)
//
// The result is rounded down.
func BaseToQuoteQuantums(
	bigBaseQuantums *big.Int,
	baseCurrencyAtomicResolution int32,
	priceValue uint64,
	priceExponent int32,
) (bigNotional *big.Int) {
	return multiplyByPrice(
		new(big.Rat).SetInt(bigBaseQuantums),
		baseCurrencyAtomicResolution,
		priceValue,
		priceExponent,
	)
}

// QuoteToBaseQuantums converts an amount denoted in quote quantums, to an equivalent amount denoted in base
// quantums. To determine the equivalent amount, an oracle price is used.
//
//   - `priceValue * 10^priceExponent` represents the conversion rate from one full coin of base currency
//     to one full coin of quote currency.
//   - `10^baseCurrencyAtomicResolution` represents the amount of one full coin that a base quantum is equal to.
//   - `10^quoteCurrencyAtomicResolution` represents the amount of one full coin that a quote quantum is equal to.
//
// To convert from quote to base quantums, we use the following equation:
//
//	baseQuantums =
//	  quoteQuantums / priceValue /
//	  10^(priceExponent + baseCurrencyAtomicResolution - quoteCurrencyAtomicResolution)
//
// The result is rounded down.
func QuoteToBaseQuantums(
	bigQuoteQuantums *big.Int,
	baseCurrencyAtomicResolution int32,
	priceValue uint64,
	priceExponent int32,
) (bigNotional *big.Int) {
	// Determine the non-exponent part of the equation.
	// We perform all calculations using positive rationals for consistent rounding.
	isLong := bigQuoteQuantums.Sign() >= 0
	ratAbsQuoteQuantums := new(big.Rat).Abs(
		new(big.Rat).SetInt(bigQuoteQuantums),
	)
	ratPrice := new(big.Rat).SetUint64(priceValue)
	ratQuoteQuantumsDivPrice := new(big.Rat).Quo(ratAbsQuoteQuantums, ratPrice)

	// Determine the absolute value of the return value.
	exponent := priceExponent + baseCurrencyAtomicResolution - QuoteCurrencyAtomicResolution
	ratBaseQuantums := new(big.Rat).Quo(
		ratQuoteQuantumsDivPrice,
		RatPow10(exponent),
	)

	// Round down.
	bigBaseQuantums := BigRatRound(ratBaseQuantums, false)

	// Flip the sign of the return value if necessary.
	if !isLong {
		bigBaseQuantums.Neg(bigBaseQuantums)
	}
	return bigBaseQuantums
}

// multiplyByPrice multiples a value by price, factoring in exponents of base
// and quote currencies.
// Given `value`, returns result of the following:
// `value * priceValue * 10^(priceExponent + baseCurrencyAtomicResolution - quoteCurrencyAtomicResolution)`
func multiplyByPrice(
	value *big.Rat,
	baseCurrencyAtomicResolution int32,
	priceValue uint64,
	priceExponent int32,
) (result *big.Int) {
	ratResult := new(big.Rat).SetUint64(priceValue)

	ratResult.Mul(
		ratResult,
		value,
	)

	ratResult.Mul(
		ratResult,
		RatPow10(priceExponent+baseCurrencyAtomicResolution-QuoteCurrencyAtomicResolution),
	)

	return new(big.Int).Quo(
		ratResult.Num(),
		ratResult.Denom(),
	)
}

// FundingRateToIndex converts funding rate (in ppm) to FundingIndex given the oracle price.
//
// To get funding index from funding rate, we know that:
//
//   - `fundingPaymentQuoteQuantum = fundingRatePpm / 1_000_000 * time * quoteQuantums`
//   - Divide both sides by `baseQuantums`:
//   - Left side: `fundingPaymentQuoteQuantums / baseQuantums = fundingIndexDelta / 1_000_000`
//   - right side:
//     ```
//     fundingRate * time * quoteQuantums / baseQuantums = fundingRatePpm / 1_000_000 *
//     priceValue * 10^(priceExponent + baseCurrencyAtomicResolution - quoteCurrencyAtomicResolution)
//     ```
//
// Hence, further multiplying both sides by 1_000_000, we have:
//
//	fundingIndexDelta =
//	  (fundingRatePpm * time) * priceValue *
//	  10^(priceExponent + baseCurrencyAtomicResolution - quoteCurrencyAtomicResolution)
//
// Arguments:
//
//	proratedFundingRate: prorated funding rate adjusted by time delta, in parts-per-million
//	timeSinceLastFunding: time (in seconds) since last funding index update
//	baseCurrencyAtomicResolution: atomic resolution of the base currency
//	priceValue: index price of the perpetual market according to the pricesKeeper
//	priceExponent: priceExponent of the market according to the pricesKeeper
func FundingRateToIndex(
	proratedFundingRate *big.Rat,
	baseCurrencyAtomicResolution int32,
	priceValue uint64,
	priceExponent int32,
) (fundingIndex *big.Int) {
	return multiplyByPrice(
		proratedFundingRate,
		baseCurrencyAtomicResolution,
		priceValue,
		priceExponent,
	)
}
