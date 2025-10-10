package main

import (
	"github.com/beevik/ntp"
	"fmt"
)

func main() {
	ntpServer := "server 0.ru.pool.ntp.org"

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("No")
	} else{
		fmt.Println(time)
	}
}