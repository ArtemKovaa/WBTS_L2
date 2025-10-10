package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	res, err := UnpackRLE("abcd\\")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func UnpackRLE(s string) (string, error) {
	var sb strings.Builder
	runes := []rune(s)
	l := 0
	
	for l < len(runes) {
		if unicode.IsDigit(runes[l]) {
			return "", errors.New("incorrect RLE string")
		} else if l + 1 == len(runes) && unicode.IsLetter(runes[l] {
			sb.WriteRune(runes[l])
			break
		} else {
			if string(runes[l]) == "\\" {
				l++

				if l == len(runes) || !unicode.IsDigit(runes[l]) {
					return "", errors.New("incorrect RLE string")
				}
			}
			r := l + 1
			for r < len(runes) {
				if !unicode.IsDigit(runes[r]) {
					break
				}
				r++
			}

			if r - l > 1 {
				count, err := strconv.Atoi(string(runes[l+1:r]))
				if err != nil {
					return "", errors.New("incorrect RLE string")
				}
				sb.WriteString(strings.Repeat(string(runes[l]), count))
			} else {
				sb.WriteRune(runes[l])
			}
			l = r
		} 
	}
	
	return sb.String(), nil
}