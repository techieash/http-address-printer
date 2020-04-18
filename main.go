package main

import (
	"flag"
	"fmt"
	"http-address-printer/util"
)


func main() {

	var concurrencyLimit int
	flag.IntVar(&concurrencyLimit, "parallel", 10, "Number of concurrent Request")
	flag.Parse()

	commandLineArgs := flag.Args()

	urls := make([]string, len(commandLineArgs))
	for index, input := range commandLineArgs {
		if !util.IsUrl(input) {
			urls[index] = util.ConvertToURL(input)
		} else {
			urls[index] = input
		}
	}
	results := util.BoundedParallelGet(urls, concurrencyLimit);
	for _, value := range results {
		fmt.Println(value)
	}

}
