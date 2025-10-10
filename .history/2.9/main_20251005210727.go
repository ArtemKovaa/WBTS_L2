package main

import (
	"str"
	"strconv"
)

func main() {

}

func EncodeWithRLE(s string) (string, error) {
	_, err := strconv.Atoi(s)
	if err == nil {
		return "", error.New("String contains only numbers")
	}
	
	return s, nil
}