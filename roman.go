// Package roman provides functions for working with Roman numerals.
package roman

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Errors returned by Encode or Decode.
var (
	ErrOutOfRange  = errors.New("Arabic number out of range. Must be 1 to 3,999")
	ErrEmptyString = errors.New("Empty string is invalid Roman numeral")
)

// error type for encountering non-Roman digits
type errInvalidDigit struct {
	roman string
	i     int
	c     rune
}

func (err *errInvalidDigit) Error() string {
	return fmt.Sprintf(
		"Invalid Roman digit %s (pos %d in \"%s\")",
		strconv.QuoteRune(err.c), err.i, err.roman,
	)
}

// strings of Roman digits and their corresponding Arabic value
type pair struct {
	roman  string
	arabic int
}

// some helpful maps
var arabicFor = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
var pairs = []pair{
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

// IsValid returns true if the argument represents a valid Roman numeral.
func IsValid(roman string) bool {
	_, err := Decode(roman)
	return err == nil
}

// Encode converts an integer into its Roman numeral representation.
// If the integer is too large or small, returns ErrOutOfRange.
func Encode(arabic int) (string, error) {
	if arabic < 1 || arabic > 3999 {
		return "", ErrOutOfRange
	}

	roman := ""
	for _, p := range pairs {
		for arabic >= p.arabic {
			arabic -= p.arabic
			roman += p.roman
		}

		if arabic == 0 {
			break
		}
	}
	return roman, nil
}

// Decode converts a Roman numeral string into the corresponding
// Arabic number.  If the string is empty, returns ErrEmptyString.
// If the string is not a valid Roman number, returns an error
// describing why.
func Decode(roman string) (int, error) {
	if len(roman) == 0 {
		return 0, ErrEmptyString
	}
	roman = strings.ToUpper(roman) // arabicFor uses upper case letters

	previousDigit := 1000
	arabic := 0
	for i, c := range roman {
		digit, ok := arabicFor[c]
		if !ok {
			return 0, &errInvalidDigit{roman, i, c}
		}
		arabic += digit

		if previousDigit < digit {
			arabic -= 2 * previousDigit
		}
		previousDigit = digit
	}

	return arabic, nil
}
