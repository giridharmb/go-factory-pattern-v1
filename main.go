package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-factory-pattern/factory"
	"log"
)

func SetupConstructors(apiKey string, city string) factory.IWeather {
	wf := factory.GAbstractFactory("weather")
	return wf(apiKey, city).(factory.IWeather)
}

func main() {

	apiKeyPtr := flag.String("apikey", "<apikey>", "open-weather api key")
	cityPtr := flag.String("city", "<city>", "name of the city")

	var apiKey string
	var city string

	flag.Parse()

	city = *cityPtr
	apiKey = *apiKeyPtr

	if city == "" {
		log.Printf("provide valid -city <city>")
		return
	}

	if city == "" {
		log.Printf("provide valid -apikey <apikey>")
		return
	}

	weatherI := SetupConstructors(apiKey, city)

	weather, err := weatherI.GetWeather()
	if err != nil {
		log.Printf("%v", err.Error())
		return
	}

	PrettyPrintData(weather)

}

func PrettyPrintData(data interface{}) {
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("error : could not MarshalIndent json : %v", err.Error())
		return
	}
	fmt.Printf("\n%v\n\n", string(dataBytes))
}
