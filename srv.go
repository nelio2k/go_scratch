package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var options struct {
	url string
}

func argParse() {
	flag.StringVar(&options.url, "url", "", "url to test for SRV")
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage : %s [OPTIONS] \n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	argParse()

	if options.url == "" {
		usage()
		return
	}

	cname, addrs, err := net.LookupSRV("" /*service*/, "" /*proto*/, options.url)
	fmt.Printf("SRV lookup came back with cname: %v addrs: %v and err: %v\n", cname, addrs, err)
	return
}
