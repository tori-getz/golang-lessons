package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/app/geo"
)

func GetWeather(geoData geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geoData.City)

	if err != nil {
		panic(err.Error())
	}

	params := url.Values{}
	params.Add("format", fmt.Sprintf("%v", format))

	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", errors.New(response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
