package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	time := time.Now().Add(response.ClockOffset)

	if err == nil {
		fmt.Println(time)
	} else {
		fmt.Println("Ни тута")
	}
}
