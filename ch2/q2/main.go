package main

import (
	"flag"
	"fmt"
)

const defaultConv = "temp"

var c string

func main() {
	flag.StringVar(&c, "type", defaultConv, "convert type")
	flag.Parse()

	args := flag.NFlag()
	fmt.Println(args)
	fmt.Println(c)

}
