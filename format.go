package main

import (
	"math/big"
	"strings"
)

func format4(value *big.Rat) string { return formatDecimal(value, 4) }

func format2(value *big.Rat) string { return formatDecimal(value, 2) }

func formatDecimal(value *big.Rat, places int) string {
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(places)), nil)
	numerator := new(big.Int).Mul(new(big.Int).Set(value.Num()), scale)
	denominator := new(big.Int).Set(value.Denom())
	wholeScaled, remainder := new(big.Int), new(big.Int)
	wholeScaled.QuoRem(numerator, denominator, remainder)

	if new(big.Int).Lsh(remainder, 1).Cmp(denominator) >= 0 {
		wholeScaled.Add(wholeScaled, big.NewInt(1))
	}

	whole, fraction := new(big.Int), new(big.Int)
	whole.QuoRem(wholeScaled, scale, fraction)
	fractionText := fraction.String()
	if len(fractionText) < places {
		fractionText = strings.Repeat("0", places-len(fractionText)) + fractionText
	}
	return whole.String() + "." + fractionText
}
