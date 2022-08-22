package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev"

func main() {
	fmt.Println("HTT Respose handeler")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)

}
