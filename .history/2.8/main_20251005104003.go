package main

import (
	"ntp"
	"fmt"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Pr
	}
}