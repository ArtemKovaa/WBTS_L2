package main

import (
	"errors"
	"strconv"
)

func main() {

}

func EncodeWithRLE(s string) (string, error) {
	_, err := strconv.Atoi(s)
	if err == nil {
		return "", errors.New("String contains only numbers")
	}
	
	return s, nil
}