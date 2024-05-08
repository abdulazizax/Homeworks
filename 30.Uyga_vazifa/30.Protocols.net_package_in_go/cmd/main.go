package main

import (
	"fmt"
	funk "homework/30.Protocols.net_package_in_go/functions"
	"log"

	"github.com/k0kubun/pp"
)

func main() {
	var url_str string
	fmt.Printf("Enter the url: ")
	fmt.Scanln(&url_str)

	// url_str := "https://www.example.com/path/to/resource?param1=value1&param2=value2#section1"

	isValid := funk.IsValidURL(url_str)
	if isValid {
		fmt.Println("Valid URL")
	} else {
		fmt.Println("Invalid URL")
	}

	url_pr, err := funk.GetURL_Parametrs(url_str)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(url_pr)

}
