package utils

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	PickupLocation  string
	PickupDate      string
	DropOffLocation string
	DropOffDate     string
	ContractNumber  []string
	SameLocation    string
	AcrissCode      string
	TelegramChannel string
	TelegramApiKey  string
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(time.Second * 5)
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config := Config{
		PickupLocation:  viper.GetString("pickup_location"),
		PickupDate:      viper.GetString("pickup_date"),
		DropOffLocation: viper.GetString("drop_off_location"),
		DropOffDate:     viper.GetString("drop_off_date"),
		ContractNumber:  viper.GetStringSlice("contract_number"),
		SameLocation:    viper.GetString("same_location"),
		AcrissCode:      viper.GetString("acriss_code"),
		TelegramChannel: viper.GetString("telegram_channel"),
		TelegramApiKey:  viper.GetString("telegram_apikey"),
	}
	return config
}