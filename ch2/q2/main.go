package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const defaultConv = "temp"

var c string

func main() {
	flag.StringVar(&c, "type", defaultConv, "convert type")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		args = append(args, input.Text())
	}
	for _, arg := range args {
		intInput, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Println(intInput)
	}
}
