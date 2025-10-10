package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	res, err := EncodeWithRLE("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}

func UnpackRLE(s string) (string, error) {
	var sb strings.Builder
	runes := []rune(s)
	
	for i, c := range runes {
		if unicode.IsDigit(c) {
			return "", errors.New("incorrect RLE string")
		} else if i + 1 == len(runes) {
			sb.WriteRune(c)
		} else if  
	}

	if right - left > 1 {
		sb.WriteString(fmt.Sprintf("%c%d", runes[left], right - left))
	} else {
		sb.WriteRune(runes[left])
	}
	
	return sb.String(), nil
}