package main

import (
	"github.com/beevik/ntp"
	"log"
)

func main() {
	ntpServer := "server 0.ru.pool.ntp.org"

	response, err := ntp.Query(ntpServer)

	if response != nil {
		log.Fatalf("Error getting response from %s\n", ntpServer)
	}

	currentTime := time.Now().Add(response.ClockOffset)
	if err != nil {
		fmt.Println("No")
	} else{
		fmt.Println(time)
	}
}