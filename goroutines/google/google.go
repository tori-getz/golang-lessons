package google

import (
	"fmt"
	"net/http"
)

func GetMainPage(channel chan int) {
	resp, err := http.Get("https://google.com/")

	if err != nil {
		panic(err.Error())
	}

	if channel != nil {
		channel <- resp.StatusCode
	} else {
		fmt.Println("Status code: ", resp.StatusCode)
	}
}
