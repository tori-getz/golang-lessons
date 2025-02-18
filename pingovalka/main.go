package main

import (
	"flag"
	"fmt"
	"pingovalka/app/files"
	"pingovalka/app/log"
	"pingovalka/app/pinger"

	"github.com/fatih/color"
)

func main() {
	color.Magenta("__ pingovalka app __")

	filePath := flag.String("file", "url.txt", "Path to file with URLs")

	flag.Parse()

	urls, err := files.Read(*filePath)

	statusCodeCh := make(chan int)
	errorCh := make(chan error)

	if err != nil {
		log.Error("Unable read file!")
		return
	}

	for _, url := range urls {
		go pinger.Ping(url, statusCodeCh, errorCh)
	}

	for _, url := range urls {
		select {
		case err := <-errorCh:
			log.Error(err.Error())
		case statusCode := <-statusCodeCh:
			codeStr := fmt.Sprintf("%v", statusCode)
			fmt.Printf("GET %v - Status code %v\n", color.MagentaString(url), color.YellowString(codeStr))
		}
	}
}
