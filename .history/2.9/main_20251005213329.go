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
		} else if i + 1 == len(runes) {
			sb.WriteRune(c)
		} else {
			right := i + 1
			for right < len(runes) {
				if !unicode.IsDigit(runes[right]) {
					break
				}
				right++
			}
			count, err := strconv.Atoi(string(runes[i+1:right]))
			if err != nil {
				return "", errors.New("incorrect RLE string")
			}
			sb.WriteString(strings.Repeat(string(runes[i]), count))
			i = right - 1
		} 
	}
	
	return sb.String(), nil
}