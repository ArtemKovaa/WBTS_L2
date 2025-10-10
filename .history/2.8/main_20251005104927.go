package main

import (
	"log"
	"time"
)

func main() {
	ntpServer := "server 0.ru.pool.ntp.org"

	response, err := ntp.Query(ntpServer)

	if err != nil {
		log.Fatalf("Error getting response from %s\n", ntpServer)
	}

	currentTime := time.Now().Add(response.ClockOffset)
	

}