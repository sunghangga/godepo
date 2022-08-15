package main

import (
	"flag"

	"godepo/internal/app"
	"godepo/internal/http/deposit"
)

var (
	broker = flag.String("broker", "localhost:9092", "boostrap Kafka broker")
)

func main() {
	flag.Parse()
	deposit.Run([]string{*broker}, app.DepositStream)
}
