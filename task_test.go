package lab1

import (
	"errors"
	"testing"
)

var data = map[int]RomanNumeral{
	1:    {I},
	5:    {V},
	10:   {X},
	11:   {X, I},
	78:   {L, X, X, V, I, I, I},
	2937: {M, M, C, M, X, X, X, V, I, I},
	222:  {C, C, X, X, I, I},
	123:  {C, X, X, I, I, I},
	2999: {M, M, C, M, X, C, I, X},
}

func TestDecimalToRomanNumeral(t *testing.T) {
	for i, numeral := range data {
		res, err := DecimalToRomanNumeral(i)
		if err != nil {
			t.Fatal(err)
		}
		if err := check(res, numeral); err != nil {
			t.Fatalf("test failed. Expected %s, got %s", numeral, res)
		}
	}
}

func check(a, b RomanNumeral) error {
	if len(a) != len(b) {
		return errors.New(":)")
	}
	for i, v := range a {
		if v != b[i] {
			return errors.New(":)")
		}
	}
	return nil
}
