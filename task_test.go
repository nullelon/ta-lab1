package lab1

import (
	"errors"
	"testing"
)

func TestTask1(t *testing.T) {
	data := map[int]RomanNumeral{
		1:  {I},
		5:  {V},
		10: {X},
		11: {X, I},
	}

	for i, numeral := range data {
		res := Task1(i)
		if err := check(res, numeral); err != nil {
			t.Fatalf("test failed. Expected %s, got %s", numeral, res)
		}
	}
}

func check(a, b RomanNumeral) error {
	if len(a) != len(b) {
		return errors.New(":(")
	}
	for i, v := range a {
		if v != b[i] {
			return errors.New(":(")
		}
	}
	return nil
}
