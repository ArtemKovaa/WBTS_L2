package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	res, err := UnpackRLE("a4bc2d5e")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}

func UnpackRLE(s string) (string, error) {
	var sb strings.Builder
	runes := []rune(s)
	l := 0
	
	for l < len(runes) {
		if unicode.IsDigit(runes[l]) {
			return "", errors.New("incorrect RLE string")
		} else if l + 1 == len(runes) {
			sb.WriteRune(runes[l])
		} else {
			r := l + 1
			for r < len(runes) {
				if !unicode.IsDigit(runes[r]) {
					break
				}
				r++
			}
			count, err := strconv.Atoi(string(runes[l+1:right]))
			if err != nil {
				return "", errors.New("incorrect RLE string")
			}
			sb.WriteString(strings.Repeat(string(runes[i]), count))
			l = r - 1
		} 
	}
	
	return sb.String(), nil
}