package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	var prev rune

	for _, r := range s {
		if unicode.IsDigit(r) {
			if prev == 0 {
				return "", ErrInvalidString
			}

			num, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrInvalidString
			}

			repeat := strings.Repeat(string(prev), num)
			res.WriteString(repeat)
			prev = 0
		} else {
			if prev != 0 {
				res.WriteString(string(prev))
			}
			prev = r
		}
	}

	if prev != 0 {
		res.WriteString(string(prev))
	}

	return res.String(), nil
}
