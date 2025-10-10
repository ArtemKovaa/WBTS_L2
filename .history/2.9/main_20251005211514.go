package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res, err := EncodeWithRLE("45")
	if err != nil {
		fmt.Println("Error: %s", err)
	} else {
		fmt.Println("Result: %s", res)
	}
}

func EncodeWithRLE(s string) (string, error) {
	_, err := strconv.Atoi(s)
	if err == nil {
		return "", errors.New("string contains only numbers")
	}

	var sb strings.Builder
	left, right := 0, 0
	runes := []rune(s)
	
	for right < len(runes) {
		if runes[left] != runes[right] {
			sb.WriteString(fmt.Sprintf("%s%d", runes[left], right - left))
			left = i
		}
		right++
	}
	
	return sb.String(), nil
}