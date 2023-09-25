package roman

import (
	"errors"
)

var romanNumerals = [...]struct {
	symbol string
	value  int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ToArabic(numeral string) (int, error) {
	var result int

	for len(numeral) > 0 {
		found := false
		for _, rn := range romanNumerals {
			if len(numeral) >= len(rn.symbol) && numeral[:len(rn.symbol)] == rn.symbol {
				result += rn.value
				numeral = numeral[len(rn.symbol):]
				found = true
				break
			}
		}
		if !found {
			return 0, errors.New("invalid Roman numeral")
		}
	}

	return result, nil
}

func FromArabic(numeral int) (string, error) {
	if numeral <= 0 {
		return "", errors.New("invalid Roman numeral, too small")
	}

	if numeral > 3999 {
		return "", errors.New("invalid Roman numeral, too big")
	}

	var result string

	for _, rn := range romanNumerals {
		for numeral >= rn.value {
			result += rn.symbol
			numeral -= rn.value
		}
	}

	return result, nil
}
