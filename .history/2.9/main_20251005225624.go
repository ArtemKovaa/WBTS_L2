package main

import (
	"errors"
	"strconv"
	"unicode"
)

func UnpackRLE(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	var res []rune
	escaped := false
	var lastRune rune

	for i, r := range s {
		if escaped {
			res = append(res, r)
			lastRune = r
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if i == 0 {
				return "", errors.New("incorrect RLE string")
			}

			right := i + 1

			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", errors.New("incorrect RLE string")
			}
			if count > 0 {
				for j := 1; j < count; j++ {
					res = append(res, lastRune)
				}
			}
			continue
		}

		res = append(res, r)
		lastRune = r
	}

	if escaped {
		return "", errors.New("incorrect RLE string")
	}

	return string(res), nil
}