package main

import (
	"fmt"
	"github.com/dsparling/go-apache-log-parser"
	"log"
)

func main() {
	lines, err := apachelogparser.Parse("logs/access_log")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		//fmt.Println(line)
		fmt.Printf("remote host: %s\n", line.RemoteHost)
		fmt.Printf("time: %s\n", line.Time)
		fmt.Printf("request: %s\n", line.Request)
		fmt.Printf("status: %d\n", line.Status)
		fmt.Printf("bytes: %d\n", line.Bytes)
		fmt.Printf("referer: %s\n", line.Referer)
		fmt.Printf("user agent: %s\n", line.UserAgent)
		fmt.Printf("url: %s\n", line.URL)
	}
}
