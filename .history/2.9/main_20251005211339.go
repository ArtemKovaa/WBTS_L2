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
	
	for i, right := len(s) {
		if s[left] != s[right] {
			sb.WriteString(fmt.Sprintf("%s%d", s[left], right - left))
			left = right
		}
		right++
	}
	
	return sb.String(), nil
}