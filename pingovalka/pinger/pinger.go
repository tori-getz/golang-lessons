package pinger

import "net/http"

func Ping(url string, statusCodeCh chan int, errorCh chan error) {
	resp, err := http.Get(url)

	if err != nil {
		errorCh <- err
		return
	}

	statusCodeCh <- resp.StatusCode
}
