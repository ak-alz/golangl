package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")

	if err == nil {
		time := time.Now().Add(response.ClockOffset)
		fmt.Println(time)
	} else {
		fmt.Println("Ни тута")
	}
}
