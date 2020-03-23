package app

import (
	"encoding/json"
	"enterprice-rental-car/utils"
	"github.com/labstack/gommon/log"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

var request = gorequest.New()

func GetLocationId(config utils.Config, timestamp string) string {
	_, body, _ := request.Get("https://prd.location.enterprise.com/enterprise-sls/search/location/dotcom/text/" + config.PickupLocation + "?countryCode=ES&brand=ENTERPRISE&locale=es_ES&now=" + timestamp).End()
	return gjson.Get(body, "*.0.peopleSoftId").String()
}

func CreateSession(config utils.Config, locationId string, contractId string) {
	request.Post("https://prd-west.webapi.enterprise.es/enterprise-ewt/enterprise/reservations/initiate?locale=es_ES").
		Send(`{"pickupLocation":{"id":"` + locationId + `","type":"BRANCH","dateTime":"` + config.PickupDate + `","countryCode":"ES"},"returnLocation":{"id":"` + locationId + `","type":"BRANCH","dateTime":"` + config.DropOffDate + `","countryCode":"ES"},"contract_number":"` + contractId + `","renter_age":25,"country_of_residence_code":"ES","enable_north_american_prepay_rates":false,"sameLocation":true,"view_currency_code":"EUR","additional_information":[]}`).End()
}

func GetPrice(config utils.Config) string {
	_, body, _ := request.Get("https://prd-west.webapi.enterprise.es/enterprise-ewt//enterprise/reservations/vehicles/availability").End()
	var cars Document
	if err := json.Unmarshal([]byte(body), &cars); err != nil {
		log.Error(err)
	}
	for _, car := range cars.Availablecars {
		if car.Code == config.AcrissCode {
			return car.Charges.PAYLATER.TotalPricePayment.Amount
		}
	}
	return "-1"
}