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
		return "", 
	}
	
	return s, nil
}