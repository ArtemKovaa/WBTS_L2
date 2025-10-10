package main

import (
	"errors"
	"strconv"
)

func main() {
	res, err := EncodeWithRLE("45")
	if err != nil {
		fmt.
	}
}

func EncodeWithRLE(s string) (string, error) {
	_, err := strconv.Atoi(s)
	if err == nil {
		return "", errors.New("string contains only numbers")
	}
	
	return s, nil
}