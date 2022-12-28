package main

import (
	"encoding/json"
	"fmt"
	"go-factory-pattern/factory"
	"log"
)

func SetupConstructors(apiKey string, city string) factory.IWeather {
	wf := factory.GAbstractFactory("weather")
	return wf(apiKey, city).(factory.IWeather)
}

func main() {

	weatherI := SetupConstructors("606c01e9eafe84058752ab69f86d74f8", "london")

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

/*
Output

2022/12/28 13:08:28 @ provider : path : /weather?q=london&appid=606c01e9eafe84058752ab69f86d74f8&units=metric
2022/12/28 13:08:28 @ provider : completeURL : https://api.openweathermap.org/data/2.5/weather?q=london&appid=606c01e9eafe84058752ab69f86d74f8&units=metric

{
    "Temp": 10.22,
    "Pressure": 995,
    "MinTemp": 9.51,
    "MaxTemp": 10.99
}
*/
