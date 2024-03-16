package main

import (
	"fmt"
	"os"

	"github.com/stevenwilkin/fx/currencybeacon"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cb := currencybeacon.NewCurrencyBeaconFromEnv()

	rates, err := cb.GetRates()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(rates))
}
