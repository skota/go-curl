package main

import (
	"flag"
	"fmt"
)

func main() {
	// if params is -h  or -H
	var url = flag.String("url", "http://www.google.com", "url to access")

	flag.Parse()
	//check params..we need atleast website

	fmt.Printf("This is the url we have %s\n", *url)
}
