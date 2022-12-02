package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/myhttp"
)

func main() {
	mh := myhttp.New(time.Second)
	response, _ := mh.Get("https://jsonip.com/")
	fmt.Println("HTTP status code: ", response.StatusCode)
}
