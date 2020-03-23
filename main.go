package main

import (
	"enterprice-rental-car/app"
	"enterprice-rental-car/utils"
	"fmt"
	"strconv"
	"time"
)

func main() {
	config := utils.ReadConfig()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	locationId := app.GetLocationId(config, timestamp)
	for _, contractId := range config.ContractNumber {
		app.CreateSession(config, locationId, contractId)
		price := app.GetPrice(config)
		fmt.Println(price)
	}
}
