package main

import (
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpServer := "0.ru.pool.ntp.org"

	response, err := ntp.Time(ntpServer)

	if err != nil {
		log.Fatalf("Error getting response from %s\n", ntpServer)
	}
	
	log.Printf("NTP Server Time: %s\n", response)
}