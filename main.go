package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var fancy bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s n (where n is an integer)\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&fancy, "fancy", false, "fancy output")
	flag.Parse()
}

func main() {
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	input, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected argument %q\n", flag.Arg(0))
		flag.Usage()
		os.Exit(1)
	}

	result := Process(input)

	if fancy {
		fmt.Printf("The result of Process(%d) is: %d\n", input, result)
	} else {
		fmt.Println(result)
	}
}

func Process(x int) int {
	return 2 * x
}
