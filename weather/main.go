package main

import (
	"flag"
	"fmt"
	"weather/app/geo"
	"weather/app/log"
	"weather/app/weather"

	"github.com/fatih/color"
)

func main() {
	color.Magenta("__ weather app __")

	city := flag.String("city", "", "Your city")
	format := flag.Int("format", 1, "Format")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)

	if err != nil {
		log.Error(err.Error())
		return
	}

	weatherResponse, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		log.Error(err.Error())
		return
	}

	fmt.Println(weatherResponse)
}
