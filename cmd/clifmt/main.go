package main

import (
	"flag"
	"git.xdrm.io/go/clifmt"
	"time"
)

func main() {

	start := time.Now()

	// custom --help output
	flag.CommandLine.Usage = help
	flag.Parse()

	clifmt.Printf("*executed in* %s\n", time.Now().Sub(start))

}
